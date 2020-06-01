package debug

import (
	"context"

	"github.com/gogf/gf/frame/g"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// UnaryServerInterceptor returns a new debug unary server interceptor.
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if info.FullMethod == "/grpc.health.v1.Health/Check" {
			return handler(ctx, req)
		}

		g.Log().Debug("ServerMethod:", info.FullMethod)

		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			g.Log().Debug("ServerMetadata:", md)
		}

		g.Log().Debug("ServerRequest:", req)

		resp, err = handler(ctx, req)

		g.Log().Debug("ServerResponse:", resp)

		return
	}
}

// UnaryClientInterceptor returns a new debug unary client interceptor.
func UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
		g.Log().Debug("ClientMethod:", method)

		md, ok := metadata.FromOutgoingContext(ctx)
		if ok {
			g.Log().Debug("ClientMetadata:", md)
		}

		g.Log().Debug("ClientRequest:", req)

		err = invoker(ctx, method, req, reply, cc, opts...)

		g.Log().Debug("ClientReply:", reply)

		return
	}
}
