package goat

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/encoding/proto"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/avos-io/goat/gen/goatorepo"
	"github.com/avos-io/goat/internal"
	"github.com/avos-io/goat/internal/server"
)

// ServerOption is an option used when constructing a NewServer.
type ServerOption interface {
	apply(*Server)
}

type serverOptFunc func(*Server)

func (fn serverOptFunc) apply(s *Server) {
	fn(s)
}

// UnaryInterceptor returns a ServerOption that sets the UnaryServerInterceptor
// for the server. Only one unary interceptor can be installed. The construction
// of multiple interceptors (e.g., chaining) can be implemented at the caller.
func UnaryInterceptor(i grpc.UnaryServerInterceptor) ServerOption {
	return serverOptFunc(func(s *Server) {
		s.unaryInterceptor = i
	})
}

// StreamInterceptor returns a ServerOption that sets the StreamServerInterceptor
// for the server. Only one stream interceptor can be installed.
func StreamInterceptor(i grpc.StreamServerInterceptor) ServerOption {
	return serverOptFunc(func(s *Server) {
		s.streamInterceptor = i
	})
}

func NewServer(id string, opts ...ServerOption) *Server {
	ctx, cancel := context.WithCancel(context.Background())

	srv := Server{
		id:       id,
		ctx:      ctx,
		cancel:   cancel,
		services: make(map[string]*serviceInfo),
	}

	for _, opt := range opts {
		opt.apply(&srv)
	}

	return &srv
}

type serviceInfo struct {
	name        string
	serviceImpl interface{}
	methods     map[string]*grpc.MethodDesc
	streams     map[string]*grpc.StreamDesc
	mdata       interface{}
}

type Server struct {
	ctx    context.Context
	cancel context.CancelFunc
	id     string

	services map[string]*serviceInfo // service name -> service info

	unaryInterceptor  grpc.UnaryServerInterceptor
	streamInterceptor grpc.StreamServerInterceptor
}

func (s *Server) Stop() {
	log.Info().Msg("Server stop")
	s.cancel()
}

// grpc.ServiceRegistrar
func (s *Server) RegisterService(sd *grpc.ServiceDesc, ss interface{}) {
	if ss != nil {
		ht := reflect.TypeOf(sd.HandlerType).Elem()
		st := reflect.TypeOf(ss)
		if !st.Implements(ht) {
			log.Panic().Msgf("RegisterService handler of type %v does not satisfy %v", st, ht)
		}
	}

	if _, ok := s.services[sd.ServiceName]; ok {
		log.Panic().Msgf("RegisterService duplicate service registration for %q", sd.ServiceName)
	}
	info := &serviceInfo{
		name:        sd.ServiceName,
		serviceImpl: ss,
		methods:     make(map[string]*grpc.MethodDesc),
		streams:     make(map[string]*grpc.StreamDesc),
		mdata:       sd.Metadata,
	}
	for i := range sd.Methods {
		d := &sd.Methods[i]
		info.methods[d.MethodName] = d
	}
	for i := range sd.Streams {
		d := &sd.Streams[i]
		info.streams[d.StreamName] = d
	}
	s.services[sd.ServiceName] = info
}

func (s *Server) Serve(ctx context.Context, rw RpcReadWriter) error {
	h := newHandler(s.ctx, s, rw)
	return h.serve(ctx)
}

type unaryRpcArgs struct {
	info *serviceInfo
	md   *grpc.MethodDesc
	rpc  *goatorepo.Rpc
}

// handler for a specific goat.RpcReadWriter
type handler struct {
	ctx    context.Context
	cancel context.CancelFunc

	srv *Server

	rw    RpcReadWriter
	codec encoding.Codec

	mu      sync.Mutex // protects streams
	streams map[uint64]chan *goatorepo.Rpc

	writeChan    chan *goatorepo.Rpc
	unaryRpcChan chan unaryRpcArgs
}

func newHandler(ctx context.Context, srv *Server, rw RpcReadWriter) *handler {
	ctx, cancel := context.WithCancel(ctx)

	return &handler{
		ctx:          ctx,
		cancel:       cancel,
		srv:          srv,
		rw:           rw,
		codec:        encoding.GetCodec(proto.Name),
		streams:      map[uint64]chan *goatorepo.Rpc{},
		writeChan:    make(chan *goatorepo.Rpc),
		unaryRpcChan: make(chan unaryRpcArgs),
	}
}

func (h *handler) serve(clientCtx context.Context) error {
	defer h.cancel()

	writeCtx, cancel := context.WithCancel(h.ctx)
	defer cancel()

	go func() {
		for {
			select {
			case rpc := <-h.writeChan:
				h.rw.Write(writeCtx, rpc)
			case <-writeCtx.Done():
				return
			}
		}
	}()

	unaryRpcCtx, unaryRpcCtxCancel := context.WithCancel(h.ctx)
	defer unaryRpcCtxCancel()

	const numRpcWorkers = 8

	for i := 0; i < numRpcWorkers; i++ {
		go func() {
			for {
				select {
				case args := <-h.unaryRpcChan:
					h.writeChan <- h.processUnaryRpc(clientCtx, args.info, args.md, args.rpc)
				case <-unaryRpcCtx.Done():
					return
				}
			}
		}()
	}

	for {
		rpc, err := h.rw.Read(h.ctx)
		if err != nil {
			return errors.Wrap(err, "read error")
		}

		if rpc.GetHeader() == nil {
			return fmt.Errorf("no header")
		}

		rawMethod := rpc.GetHeader().Method
		service, method, err := parseRawMethod(rawMethod)
		if err != nil {
			return errors.Wrapf(err, "failed to parse %s", rawMethod)
		}

		if rpc.GetHeader().Destination != h.srv.id {
			log.Warn().
				Msgf("Server %s: invalid destination %s", h.srv.id, rpc.GetHeader().Destination)
			continue
		}

		si, known := h.srv.services[service]
		if !known {
			log.Warn().Msgf("Server: unknown service, %s", service)
			continue
		}
		if md, ok := si.methods[method]; ok {
			h.unaryRpcChan <- unaryRpcArgs{si, md, rpc}
			continue
		}
		if sd, ok := si.streams[method]; ok {
			h.processStreamingRpc(clientCtx, si, sd, rpc)
			continue
		}
		log.Warn().Msgf("Server: unhandled method, %s %s", service, method)
	}
}

func (h *handler) processUnaryRpc(
	clientCtx context.Context,
	info *serviceInfo,
	md *grpc.MethodDesc,
	rpc *goatorepo.Rpc,
) *goatorepo.Rpc {
	ctx, cancel, err := contextFromHeaders(clientCtx, rpc.GetHeader())
	if err != nil {
		log.Panic().Err(err).Msg("Server: failed to get context from headers")
	}
	defer cancel()

	fullMethod := fmt.Sprintf("/%s/%s", info.name, md.MethodName)
	sts := server.NewUnaryServerTransportStream(fullMethod)
	ctx = grpc.NewContextWithServerTransportStream(ctx, sts)

	body := rpc.GetBody()

	dec := func(msg interface{}) error {
		if body.GetData() == nil {
			return nil
		}

		if err := h.codec.Unmarshal(body.GetData(), msg); err != nil {
			return status.Error(codes.InvalidArgument, err.Error())
		}
		return nil
	}

	resp, appErr := md.Handler(info.serviceImpl, ctx, dec, h.srv.unaryInterceptor)

	respH := internal.ToKeyValue(sts.GetHeaders())
	respHeader := &goatorepo.RequestHeader{
		Method:      fullMethod,
		Headers:     respH,
		Source:      rpc.Header.Destination,
		Destination: rpc.Header.Source,
	}
	if len(rpc.Header.ProxyRecord) > 1 {
		respHeader.ProxyNext = rpc.Header.ProxyRecord[0 : len(rpc.Header.ProxyRecord)-1]
	}
	respTrailer := &goatorepo.Trailer{
		Metadata: internal.ToKeyValue(sts.GetTrailers()),
	}

	var respStatus *goatorepo.ResponseStatus

	if appErr != nil {
		st, ok := status.FromError(appErr)
		if !ok {
			st = status.FromContextError(appErr)
		}
		respStatus = &goatorepo.ResponseStatus{
			Code:    st.Proto().GetCode(),
			Message: st.Proto().GetMessage(),
			Details: st.Proto().GetDetails(),
		}
	}

	var respBody *goatorepo.Body

	if resp != nil {
		data, err := h.codec.Marshal(resp)

		if err == nil {
			respBody = &goatorepo.Body{
				Data: data,
			}
		}
	}

	return &goatorepo.Rpc{
		Id:      rpc.GetId(),
		Header:  respHeader,
		Status:  respStatus,
		Body:    respBody,
		Trailer: respTrailer,
	}
}

func (h *handler) processStreamingRpc(
	clientCtx context.Context,
	info *serviceInfo,
	sd *grpc.StreamDesc,
	rpc *goatorepo.Rpc,
) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if ch, ok := h.streams[rpc.Id]; ok {
		ch <- rpc
		return nil
	}

	streamId := rpc.Id
	ch := make(chan *goatorepo.Rpc, 1)
	h.streams[streamId] = ch

	go h.runStream(clientCtx, info, sd, rpc, streamId, ch)
	return nil
}

func (h *handler) runStream(
	clientCtx context.Context,
	info *serviceInfo,
	sd *grpc.StreamDesc,
	rpc *goatorepo.Rpc,
	streamId uint64,
	rCh chan *goatorepo.Rpc,
) error {
	defer h.unregisterStream(streamId)

	if rpc.GetBody() != nil {
		// The first Rpc in a stream is used to tell the server to start the stream;
		// it must have an empty body. If this isn't the case, it must be because
		// we've missed the first Rpc in the stream.
		log.Info().Msgf("did not expect body: calling RST stream %d", streamId)
		h.resetStream(rpc)
		return nil
	}

	if rpc.GetTrailer() != nil {
		// The client may send a trailer to end a stream after we've already ended
		// it, in which case we don't want to lazily create a new stream here.
		return nil
	}

	r := func(ctx context.Context) (*goatorepo.Rpc, error) {
		select {
		case msg, ok := <-rCh:
			if !ok {
				return nil, fmt.Errorf("rCh closed")
			}
			return msg, nil
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
	rw := internal.NewFnReadWriter(r, func(ctx context.Context, r *goatorepo.Rpc) error {
		h.writeChan <- r
		return nil
	})

	ctx, cancel, err := contextFromHeaders(clientCtx, rpc.GetHeader())
	if err != nil {
		log.Panic().Err(err).Msg("Server: failed to get context from headers")
	}
	defer cancel()

	si := &grpc.StreamServerInfo{
		FullMethod:     fmt.Sprintf("/%s/%s", info.name, sd.StreamName),
		IsClientStream: sd.ClientStreams,
		IsServerStream: sd.ServerStreams,
	}

	stream, err := server.NewServerStream(
		ctx,
		streamId,
		si.FullMethod,
		rpc.Header.Destination,
		rpc.Header.Source,
		rw,
	)
	if err != nil {
		log.Error().Err(err).Msg("Server: newServerStream failed")
		return err
	}

	sts := server.NewServerTransportStream(si.FullMethod, stream)
	stream.SetContext(grpc.NewContextWithServerTransportStream(ctx, sts))

	var appErr error
	if h.srv.streamInterceptor != nil {
		appErr = h.srv.streamInterceptor(info.serviceImpl, stream, si, sd.Handler)
	} else {
		appErr = sd.Handler(info.serviceImpl, stream)
	}

	err = stream.SendTrailer(appErr)
	if err != nil {
		log.Error().Err(err).Msg("Server: SendTrailer")
		return err
	}

	return nil
}

func (h *handler) unregisterStream(id uint64) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if ch, ok := h.streams[id]; ok {
		close(ch)
	}

	delete(h.streams, id)
}

// resetStream instructs the caller to tear down and restart the stream. We call
// this if something has gone irrecoverably wrong in the stream.
func (h *handler) resetStream(rpc *goatorepo.Rpc) {
	reset := &goatorepo.Rpc{
		Id: rpc.GetId(),
		Header: &goatorepo.RequestHeader{
			Method:      rpc.GetHeader().GetMethod(),
			Source:      rpc.GetHeader().GetDestination(),
			Destination: rpc.GetHeader().GetSource(),
		},
		Status: &goatorepo.ResponseStatus{
			Code:    int32(codes.Aborted),
			Message: "RST stream",
		},
		Trailer: &goatorepo.Trailer{},
	}

	if len(rpc.Header.ProxyRecord) > 1 {
		reset.Header.ProxyNext = rpc.Header.ProxyRecord[0 : len(rpc.Header.ProxyRecord)-1]
	}

	h.rw.Write(h.ctx, reset)
}

// contextFromHeaders returns a new incoming context with metadata populated
// by the given request headers. If the given headers contain a GRPC-Timeout, it
// is used to set the deadline on the returned context.
func contextFromHeaders(
	parent context.Context,
	h *goatorepo.RequestHeader,
) (context.Context, context.CancelFunc, error) {
	cancel := func() {}
	md, err := internal.ToMetadata(h.Headers)
	if err != nil {
		return parent, cancel, err
	}
	ctx := metadata.NewIncomingContext(parent, md)

	for _, hdr := range h.Headers {
		if strings.ToLower(hdr.Key) == "grpc-timeout" {
			if timeout, ok := parseGrpcTimeout(hdr.Value); ok {
				ctx, cancel = context.WithTimeout(ctx, timeout)
				break
			}
		}
	}
	return ctx, cancel, nil
}

// See https://grpc.io/docs/guides/wire.html#requests
func parseGrpcTimeout(timeout string) (time.Duration, bool) {
	if timeout == "" {
		return 0, false
	}
	suffix := timeout[len(timeout)-1]

	val, err := strconv.ParseInt(timeout[:len(timeout)-1], 10, 64)
	if err != nil {
		return 0, false
	}
	getUnit := func(suffix byte) time.Duration {
		switch suffix {
		case 'H':
			return time.Hour
		case 'M':
			return time.Minute
		case 'S':
			return time.Second
		case 'm':
			return time.Millisecond
		case 'u':
			return time.Microsecond
		case 'n':
			return time.Nanosecond
		default:
			return 0
		}
	}
	unit := getUnit(suffix)
	if unit == 0 {
		return 0, false
	}

	return time.Duration(val) * unit, true
}

func parseRawMethod(sm string) (string, string, error) {
	if sm != "" && sm[0] == '/' {
		sm = sm[1:]
	}
	pos := strings.LastIndex(sm, "/")
	if pos == -1 {
		return "", "", fmt.Errorf("invalid method name %s", sm)
	}
	service := sm[:pos]
	method := sm[pos+1:]
	return service, method, nil
}
