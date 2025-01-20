package client

import (
	"context"
	"errors"
	"io"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/encoding/proto"
	"google.golang.org/grpc/mem"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/stats"
	"google.golang.org/grpc/status"

	goatorepo "github.com/avos-io/goat/gen/goatorepo"
	"github.com/avos-io/goat/internal"
	"github.com/avos-io/goat/types"
)

type clientStream struct {
	ctx context.Context

	id     uint64
	method string
	codec  encoding.CodecV2

	sourceAddress, destAddress string

	ready  sync.WaitGroup
	header metadata.MD

	rw types.RpcReadWriter

	statsHandlers []stats.Handler
	beginTime     time.Time

	protected struct {
		sync.Mutex

		done      bool
		headerErr error
		eErr      error
		rErr      error
		trailer   *goatorepo.Trailer
	}

	teardown func(bool)

	rCh chan *goatorepo.Body
}

var _ grpc.ClientStream = (*clientStream)(nil)

func NewStream(
	ctx context.Context,
	id uint64,
	method string,
	rw types.RpcReadWriter,
	teardown func(),
	sourceAddress, destAddress string,
	statsHandlers []stats.Handler,
	beginTime time.Time,
) grpc.ClientStream {

	ctx, cancel := context.WithCancel(ctx)

	cs := &clientStream{
		ctx:           ctx,
		id:            id,
		method:        method,
		codec:         encoding.GetCodecV2(proto.Name),
		rw:            rw,
		rCh:           make(chan *goatorepo.Body),
		sourceAddress: sourceAddress,
		destAddress:   destAddress,
		statsHandlers: statsHandlers,
		beginTime:     beginTime,
	}

	csteardown := func(sendRst bool) {
		teardown()

		if sendRst {
			rpc := goatorepo.Rpc{
				Id: id,
				Header: &goatorepo.RequestHeader{
					Method:      method,
					Source:      sourceAddress,
					Destination: destAddress,
				},
				Reset_: &goatorepo.Reset{
					Type: "RST_STREAM",
				},
			}

			writeCtx, cancelWrite := context.WithDeadline(context.TODO(), time.Now().Add(30*time.Second))
			defer cancelWrite()
			err := rw.Write(writeCtx, &rpc)
			if err != nil {
				log.Err(err).Msg("NewStream: failed to teardown")
			}
		}

		cancel()
	}
	cs.teardown = csteardown

	cs.ready.Add(1)
	go cs.readLoop()

	return cs
}

// Header returns the header metadata received from the server if there
// is any. It blocks if the metadata is not ready to read.
func (cs *clientStream) Header() (metadata.MD, error) {
	// unblocked once we receive our first headers
	cs.ready.Wait()

	cs.protected.Lock()
	defer cs.protected.Unlock()

	return cs.header, cs.protected.headerErr
}

// Trailer returns the trailer metadata from the server, if there is any.
// It must only be called after stream.CloseAndRecv has returned, or
// stream.Recv has returned a non-nil error (including io.EOF).
func (cs *clientStream) Trailer() metadata.MD {
	cs.protected.Lock()
	defer cs.protected.Unlock()

	if cs.protected.trailer.GetMetadata() == nil {
		return nil
	}
	md, err := internal.ToMetadata(cs.protected.trailer.Metadata)
	if err != nil {
		log.Panic().Err(err).Msg("Trailer err")
	}
	return md
}

// CloseSend closes the send direction of the stream. It closes the stream
// when non-nil error is met. It is also not safe to call CloseSend
// concurrently with SendMsg.
func (cs *clientStream) CloseSend() error {
	tr := goatorepo.Rpc{
		Id: cs.id,
		Header: &goatorepo.RequestHeader{
			Method:      cs.method,
			Source:      cs.sourceAddress,
			Destination: cs.destAddress,
		},
		Status: &goatorepo.ResponseStatus{
			Code:    int32(codes.OK),
			Message: codes.OK.String(),
		},
		Trailer: &goatorepo.Trailer{},
	}

	for _, sh := range cs.statsHandlers {
		md, _ := internal.ToMetadata(tr.GetTrailer().GetMetadata())
		sh.HandleRPC(cs.ctx, &stats.OutTrailer{
			Client:  true,
			Trailer: md,
		})
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
		cs.teardown(false)
		return err
	}
	rpc := goatorepo.Rpc{
		Id: cs.id,
		Header: &goatorepo.RequestHeader{
			Method:      cs.method,
			Source:      cs.sourceAddress,
			Destination: cs.destAddress,
		},
		Body: &goatorepo.Body{
			Data: body.Materialize(),
		},
	}
	err = cs.rw.Write(cs.ctx, &rpc)
	if err != nil {
		cs.teardown(false)
		return err
	}

	for _, sh := range cs.statsHandlers {
		sh.HandleRPC(cs.ctx, &stats.OutPayload{
			Client:   true,
			Payload:  m,
			Length:   len(body),
			SentTime: time.Now(),
		})
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

		buf := mem.NewBuffer(&body.Data, nil)
		bs := mem.BufferSlice{buf}

		err := cs.codec.Unmarshal(bs, m)
		if err != nil {
			return err
		}
		for _, sh := range cs.statsHandlers {
			sh.HandleRPC(cs.ctx, &stats.InPayload{
				RecvTime: time.Now(),
				Payload:  m,
				Length:   len(body.GetData()),
			})
		}
		return nil
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
	var trailer *goatorepo.Trailer

	defer func() {
		cs.protected.Lock()
		defer cs.protected.Unlock()

		close(cs.rCh)
		sendRst := trailer == nil
		cs.teardown(sendRst)

		cs.protected.done = true
		cs.protected.rErr = rErr
		cs.protected.trailer = trailer

		now := time.Now()
		for _, sh := range cs.statsHandlers {
			end := &stats.End{
				BeginTime: cs.beginTime,
				EndTime:   now,
			}
			if rErr != nil && rErr != io.EOF {
				end.Error = rErr
			}
			sh.HandleRPC(cs.ctx, end)
		}

		log.Info().Uint64("id", cs.id).Msg("ClientStream done")
	}()

	onReady := func(err error, headers metadata.MD) {
		cs.protected.Lock()
		defer cs.protected.Unlock()

		if cs.header == nil {
			cs.protected.headerErr = err
			cs.header = headers
			cs.ready.Done()

			if err == nil {
				for _, sh := range cs.statsHandlers {
					sh.HandleRPC(cs.ctx, &stats.InHeader{
						Client:     true,
						FullMethod: cs.method,
						Header:     headers,
					})
				}
			}
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
			// ok
		}
	}
}

func (cs *clientStream) readErrorIfDone() (bool, error) {
	cs.protected.Lock()
	defer cs.protected.Unlock()

	if !cs.protected.done {
		return false, nil
	}

	return true, cs.protected.rErr
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
func errorIfDone(rpc *goatorepo.Rpc) (bool, error) {
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
