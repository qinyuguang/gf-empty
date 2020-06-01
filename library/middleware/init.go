package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
)

func Init(r *ghttp.Request) {
	// start time
	r.SetCtxVar("request_time", gtime.TimestampMilli())

	r.Middleware.Next()
}
