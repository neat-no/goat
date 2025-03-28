package server

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"

	grpcMocks "github.com/neat-no/goat/gen/grpc/mocks"
	"github.com/neat-no/goat/gen/testproto"
)

func TestUnaryTransportStream(t *testing.T) {
	t.Run("Method", func(t *testing.T) {
		ts := NewUnaryServerTransportStream("foo")
		require.Equal(t, "foo", ts.Method())
	})

	t.Run("Headers", func(t *testing.T) {
		is := require.New(t)

		ts := NewUnaryServerTransportStream("")
		md1 := metadata.New(map[string]string{"foo": "1"})
		md2 := metadata.New(map[string]string{"foo": "2"})
		md3 := metadata.New(map[string]string{"foo": "2"})

		is.NoError(ts.SetHeader(md1))
		is.NoError(ts.SetHeader(md2))
		is.NoError(ts.SendHeader(md3))

		got := ts.GetHeaders()
		is.Equal(metadata.Join(md1, md2, md3), got)

		// Can't "send" again
		is.Error(ts.SetHeader(md1))
		is.Error(ts.SendHeader(md1))
	})

	t.Run("Trailers", func(t *testing.T) {
		is := require.New(t)

		ts := NewUnaryServerTransportStream("")
		md1 := metadata.New(map[string]string{"foo": "1"})
		md2 := metadata.New(map[string]string{"foo": "2"})

		is.NoError(ts.SetTrailer(md1))
		is.NoError(ts.SetTrailer(md2))

		got := ts.GetTrailers()
		is.Equal(metadata.Join(md1, md2), got)
	})
}

func TestServerTransportStream(t *testing.T) {
	t.Run("Method", func(t *testing.T) {
		ts := NewServerTransportStream("foo",
			grpcMocks.NewMockBidiStreamingServer[testproto.Msg, testproto.Msg](t))
		require.Equal(t, "foo", ts.Method())
	})

	t.Run("Headers and trailers", func(t *testing.T) {
		is := require.New(t)

		md := metadata.New(map[string]string{"foo": "1"})

		m := grpcMocks.NewMockBidiStreamingServer[testproto.Msg, testproto.Msg](t)
		ts := NewServerTransportStream("", m)

		m.EXPECT().SetHeader(md).Return(nil).Once()
		is.NoError(ts.SetHeader(md))

		m.EXPECT().SetHeader(md).Return(errTest).Once()
		is.Error(ts.SetHeader(md))

		m.EXPECT().SendHeader(md).Return(nil).Once()
		is.NoError(ts.SendHeader(md))

		m.EXPECT().SendHeader(md).Return(errTest).Once()
		is.Error(ts.SendHeader(md))

		m.EXPECT().SetTrailer(md).Return().Once()
		is.NoError(ts.SetTrailer(md))
	})

}
