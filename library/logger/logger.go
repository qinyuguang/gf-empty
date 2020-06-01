package logger

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

type Writer struct {
	gf *glog.Logger
}

func NewWriter() *Writer {
	return &Writer{
		gf: glog.New(),
	}
}

func (w *Writer) Write(p []byte) (n int, err error) {
	return w.gf.Write(p)
}

func Middleware(r *ghttp.Request) {
	// new logger
	// g.Log().SetWriter(NewWriter())

	// db logger
	g.DB().SetLogger(g.Log())

	// set logger
	g.Log().SetAsync(true)

	g.Log().Cat("_request").Infof(
		`%s %s%s "%s" %s`,
		r.Method, r.Host, r.URL.String(), r.GetBody(), r.GetClientIp(),
	)

	g.Log().Debug("Header:", r.Header)
	g.Log().Debug("Cookie:", r.Cookie)

	r.Middleware.Next()
}
