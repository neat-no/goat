package testutil

import (
	"context"
	"encoding/binary"
	"io"
	"net"
	"sync"

	goatorepo "github.com/neat-no/goat/gen/goatorepo"
	"google.golang.org/protobuf/proto"
)

type goatOverPipe struct {
	conn       net.Conn
	readMutex  sync.Mutex
	writeMutex sync.Mutex
}

func NewGoatOverPipe(c net.Conn) *goatOverPipe {
	return &goatOverPipe{conn: c}
}

func (g *goatOverPipe) Read(ctx context.Context) (*goatorepo.Rpc, error) {
	g.readMutex.Lock()
	defer g.readMutex.Unlock()

	var msgSize uint32
	err := binary.Read(g.conn, binary.BigEndian, &msgSize)
	if err != nil {
		return nil, err
	}

	data := make([]byte, msgSize)
	_, err = io.ReadFull(g.conn, data)
	if err != nil {
		return nil, err
	}

	var rpc goatorepo.Rpc
	err = proto.Unmarshal(data, &rpc)
	if err != nil {
		panic(err)
	}

	return &rpc, nil
}

func (g *goatOverPipe) Write(ctx context.Context, pkt *goatorepo.Rpc) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	g.writeMutex.Lock()
	defer g.writeMutex.Unlock()

	data, err := proto.Marshal(pkt)
	if err != nil {
		panic(err)
	}

	var msgSize uint32 = uint32(len(data))
	err = binary.Write(g.conn, binary.BigEndian, &msgSize)
	if err != nil {
		return err
	}

	_, err = g.conn.Write(data)
	if err != nil {
		return err
	}

	return nil
}
