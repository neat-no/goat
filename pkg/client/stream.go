package client

import (
	"context"
	"errors"
	"io"
	"sync"

	"github.com/rs/zerolog/log"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/encoding/proto"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/avos-io/goat"
	wrapped "github.com/avos-io/goat/gen"
	"github.com/avos-io/goat/internal"
)

type clientStream struct {
	ctx context.Context

	id     uint64
	method string
	codec  encoding.Codec

	ready  sync.WaitGroup
	header metadata.MD

	rw goat.RpcReadWriter

	mu        sync.Mutex // protects: done, headerErr, rErr, trailer
	done      bool
	headerErr error
	rErr      error
	trailer   *wrapped.Trailer

	teardown func()

	rCh chan *wrapped.Body
}

var _ grpc.ClientStream = (*clientStream)(nil)

func newClientStream(
	ctx context.Context,
	id uint64,
	method string,
	rw goat.RpcReadWriter,
	teardown func(),
) grpc.ClientStream {

	ctx, cancel := context.WithCancel(ctx)

	csteardown := func() {
		teardown()
		cancel()
	}

	cs := &clientStream{
		ctx:      ctx,
		id:       id,
		method:   method,
		codec:    encoding.GetCodec(proto.Name),
		rw:       rw,
		teardown: csteardown,
		rCh:      make(chan *wrapped.Body),
	}

	cs.ready.Add(1)
	go cs.readLoop()

	return cs
}

// Header returns the header metadata received from the server if there
// is any. It blocks if the metadata is not ready to read.
func (cs *clientStream) Header() (metadata.MD, error) {
	cs.ready.Wait()
	return cs.header, cs.headerErr
}

// Trailer returns the trailer metadata from the server, if there is any.
// It must only be called after stream.CloseAndRecv has returned, or
// stream.Recv has returned a non-nil error (including io.EOF).
func (cs *clientStream) Trailer() metadata.MD {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	if cs.trailer.GetMetadata() == nil {
		return nil
	}
	md, err := internal.ToMetadata(cs.trailer.Metadata)
	if err != nil {
		log.Panic().Err(err).Msg("Trailer err")
	}
	return md
}

// CloseSend closes the send direction of the stream. It closes the stream
// when non-nil error is met. It is also not safe to call CloseSend
// concurrently with SendMsg.
func (cs *clientStream) CloseSend() error {
	tr := wrapped.Rpc{
		Id: cs.id,
		Header: &wrapped.RequestHeader{
			Method: cs.method,
		},
		Status: &wrapped.ResponseStatus{
			Code:    int32(codes.OK),
			Message: codes.OK.String(),
		},
		Trailer: &wrapped.Trailer{},
	}

	return cs.rw.Write(cs.ctx, &tr)
}

// Context returns the context for this stream.
func (wcs *clientStream) Context() context.Context {
	return wcs.ctx
}

// SendMsg is generally called by generated code. On error, SendMsg aborts
// the stream. If the error was generated by the client, the status is
// returned directly; otherwise, io.EOF is returned and the status of
// the stream may be discovered using RecvMsg.
//
// SendMsg blocks until:
//   - There is sufficient flow control to schedule m with the transport, or
//   - The stream is done, or
//   - The stream breaks.
//
// SendMsg does not wait until the message is received by the server. An
// untimely stream closure may result in lost messages. To ensure delivery,
// users should ensure the RPC completed successfully using RecvMsg.
//
// It is safe to have a goroutine calling SendMsg and another goroutine
// calling RecvMsg on the same stream at the same time, but it is not safe
// to call SendMsg on the same stream in different goroutines. It is also
// not safe to call CloseSend concurrently with SendMsg.
func (cs *clientStream) SendMsg(m interface{}) error {
	if done, err := cs.readErrorIfDone(); done {
		return err
	}

	body, err := cs.codec.Marshal(m)
	if err != nil {
		cs.teardown()
		return err
	}
	rpc := wrapped.Rpc{
		Id: cs.id,
		Header: &wrapped.RequestHeader{
			Method: cs.method,
		},
		Body: &wrapped.Body{
			Data: body,
		},
	}
	err = cs.rw.Write(cs.ctx, &rpc)
	if err != nil {
		cs.teardown()
		return err
	}

	return nil
}

// RecvMsg blocks until it receives a message into m or the stream is
// done. It returns io.EOF when the stream completes successfully. On
// any other error, the stream is aborted and the error contains the RPC
// status.
//
// It is safe to have a goroutine calling SendMsg and another goroutine
// calling RecvMsg on the same stream at the same time, but it is not
// safe to call RecvMsg on the same stream in different goroutines.
func (cs *clientStream) RecvMsg(m interface{}) error {
	if done, err := cs.readErrorIfDone(); done {
		return err
	}

	select {
	case <-cs.ctx.Done():
		return toStatusError(cs.ctx.Err())
	case body, ok := <-cs.rCh:
		if !ok {
			done, err := cs.readErrorIfDone()
			if !done {
				panic("cs.rCh was closed but cs.done == false!")
			}
			return err
		}
		return cs.codec.Unmarshal(body.Data, m)
	}
}

// readLoop reads from the given RpcReadWriter, filling out headers, trailers,
// and recevied errors, and enqueing messages onto the internal rCh channel.
//
// On receiving an error or the end of the stream, the stream is closed down
// for reading. Errors are held internally and will be yielded on the next
// RecvMsg or SendMsg call, as per the gRPC semantics.
func (cs *clientStream) readLoop() error {
	log.Info().Uint64("id", cs.id).Msg("ClientStream started")

	var rErr error
	var trailer *wrapped.Trailer

	defer func() {
		cs.mu.Lock()
		defer cs.mu.Unlock()

		close(cs.rCh)
		cs.teardown()

		cs.done = true
		cs.rErr = rErr
		cs.trailer = trailer

		log.Info().Uint64("id", cs.id).Msg("ClientStream done")
	}()

	onReady := func(err error, headers metadata.MD) {
		cs.mu.Lock()
		defer cs.mu.Unlock()

		if cs.header == nil {
			cs.headerErr = err
			cs.header = headers
			cs.ready.Done()
		}
	}

	for {
		rpc, err := cs.rw.Read(cs.ctx)
		if err != nil {
			rErr = toStatusError(err)
			onReady(rErr, nil)
			return err
		}

		if cs.header == nil {
			md, err := internal.ToMetadata(rpc.GetHeader().GetHeaders())
			if err != nil {
				return err
			}
			onReady(nil, md)
		}

		if done, err := errorIfDone(rpc); done {
			trailer = rpc.GetTrailer()
			rErr = err
			return err
		}

		if rpc.Body == nil {
			continue
		}

		select {
		case <-cs.ctx.Done():
			rErr = toStatusError(cs.ctx.Err())
			return rErr
		case cs.rCh <- rpc.Body:
			//
		}
	}
}

func (cs *clientStream) readErrorIfDone() (bool, error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	if !cs.done {
		return false, nil
	}

	return true, cs.rErr
}

// toStatusError converts an error to one compatible with the grpc.status pkg.
func toStatusError(err error) error {
	if errors.Is(err, context.DeadlineExceeded) ||
		errors.Is(err, context.Canceled) {
		return status.FromContextError(err).Err()
	}
	stErr, _ := status.FromError(err)
	return stErr.Err()
}

// errorIfDone returns a grpc/status compatible error if the given RPC is the
// last RPC in the stream.
func errorIfDone(rpc *wrapped.Rpc) (bool, error) {
	if rpc.Trailer == nil {
		return false, nil
	}
	st := rpc.GetStatus()
	if st.GetCode() == int32(codes.OK) {
		return true, io.EOF
	}
	sp := spb.Status{
		Code:    st.GetCode(),
		Message: st.GetMessage(),
		Details: st.GetDetails(),
	}
	return true, status.FromProto(&sp).Err()
}
