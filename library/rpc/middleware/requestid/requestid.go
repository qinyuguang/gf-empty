package requestid

import (
	"context"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/guuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// MetadataKey is metadata key of requestid.
const MetadataKey = "x-request-id"

// UnaryServerInterceptor returns a new requestid unary server interceptor.
func UnaryServerInterceptor(opts ...OptionFunc) grpc.UnaryServerInterceptor {
	o := newOptions(opts...)
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.MD{}
		}

		id := metadataValue(md, MetadataKey)
		if id != "" && existRequestID(o, id) {
			return nil, status.Errorf(codes.AlreadyExists, "duplicated request. requestid=%s", id)
		}

		ctx = context.WithValue(ctx, RequestIDKey, id)

		return handler(ctx, req)
	}
}

// UnaryClientInterceptor returns a new requestid unary client interceptor.
func UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.MD{}
		}

		id := gconv.String(ctx.Value(RequestIDKey))
		if id == "" {
			id = guuid.New().String()
		}

		md.Set(MetadataKey, id)

		ctx = metadata.NewOutgoingContext(ctx, md)

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func metadataValue(md metadata.MD, key string) string {
	if vals := md.Get(key); len(vals) > 0 {
		return vals[0]
	}

	return ""
}

func existRequestID(o *options, id string) bool {
	res, err := g.Redis("requestid").DoVar("SET", o.scope+id, 1, "NX", "EX", o.ttl)
	if err != nil {
		g.Log().Error("requestid err: ", err)
		return false
	}

	return !res.Bool()
}
