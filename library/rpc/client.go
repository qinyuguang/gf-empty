package rpc

import (
	"context"
	"errors"
	"time"

	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/frame/g"
	"google.golang.org/grpc"
)

var clientIns = gmap.NewStrAnyMap(true)

// GRPCClient wraps the grpc.ClientConn and provides more feature.
type GRPCClient struct {
	conn           *grpc.ClientConn
	addr           string
	opts           []grpc.DialOption
	connectTimeout time.Duration
	invokeTimeout  time.Duration
}

// Client returns an instance of GRPCClient with specified name.
func Client(service string) *GRPCClient {
	client := clientIns.GetVarOrSetFuncLock(service, func() interface{} {
		config := g.Cfg().GetJson("rpc.service." + service)

		client := &GRPCClient{
			addr:           config.GetString("address"),
			connectTimeout: config.GetDuration("connect_timeout"),
			invokeTimeout:  config.GetDuration("invoke_timeout"),
			opts:           []grpc.DialOption{grpc.WithBlock()},
		}

		return client
	})

	if client != nil {
		return client.Val().(*GRPCClient)
	}

	return nil
}

// Options adds grpc.DialOption for dial.
func (c *GRPCClient) Options(opts ...grpc.DialOption) *GRPCClient {
	c.opts = append(c.opts, opts...)
	return c
}

// Dial connects service with connect timeout and grpc.DialOption.
func (c *GRPCClient) Dial() error {
	ctx, _ := context.WithTimeout(context.Background(), c.connectTimeout*time.Second)

	conn, err := grpc.DialContext(ctx, c.addr, c.opts...)
	if err != nil {
		return err
	}

	c.conn = conn

	return nil
}

// Conn returns an instance of grpc.ClientConn
// and reconnect when connection is nil.
func (c *GRPCClient) Conn() (*grpc.ClientConn, error) {
	if c.conn == nil {
		if err := c.Dial(); err != nil {
			return nil, errors.New("service doesn't dial")
		}
	}

	return c.conn, nil
}
