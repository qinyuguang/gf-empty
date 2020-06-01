package recovery

import (
	"github.com/gogf/gf/frame/g"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Logger is a log handler for recovering from a panic.
func Logger(p interface{}) error {
	g.Log().Error(p)
	return status.Errorf(codes.Unknown, "recovered from panic %#v", p)
}
