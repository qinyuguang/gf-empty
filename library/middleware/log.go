package middleware

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
)

func Log(r *ghttp.Request) {
	r.Middleware.Next()

	g.Log().Cat("_request").Infof(
		`%s %s%s "%s" %s %d`,
		r.Method, r.Host, r.URL.String(), r.GetBody(), r.GetClientIp(), gtime.TimestampMilli()-r.GetCtxVar("request_time").Int64(),
	)

	g.Log().Debug("Header:", r.Header)
	g.Log().Debug("Cookie:", r.Cookie)
}
