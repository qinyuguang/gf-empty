package health

import (
	"github.com/gogf/gf/net/ghttp"
)

func Check(r *ghttp.Request) {
	r.Response.Write("OK")
}
