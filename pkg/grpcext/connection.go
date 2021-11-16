package grpcext

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"sync"
)

func NewConnection(target string) *Connection {
	return &Connection{
		target: target,
		locker: new(sync.Mutex),
	}
}

type Connection struct {
	target     string
	locker     *sync.Mutex
	clientConn *grpc.ClientConn
}

func (c *Connection) conn() (*grpc.ClientConn, error) {
	if c.clientConn != nil && c.clientConn.GetState() == connectivity.Ready {
		return c.clientConn, nil
	}

	c.locker.Lock()
	defer c.locker.Unlock()
	if c.clientConn != nil && c.clientConn.GetState() == connectivity.Ready {
		return c.clientConn, nil
	}
	conn, err := grpc.Dial(c.target, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c.clientConn = conn
	return conn, nil
}

func (c *Connection) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	conn, err := c.conn()
	if err != nil {
		return err
	}
	return conn.Invoke(ctx, method, args, reply, opts...)
}

func (c *Connection) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	conn, err := c.conn()
	if err != nil {
		return nil, err
	}
	return conn.NewStream(ctx, desc, method, opts...)
}
