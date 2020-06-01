package router

import (
	api_book "gf-empty/app/api/book"
	api_health "gf-empty/app/api/health"
	rpc_book "gf-empty/app/rpc/book"
	rpc_health "gf-empty/app/rpc/health"
	"gf-empty/library/rpc"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	// gRPC handlers
	rpc.Server().RegisterService(
		new(rpc_book.Service),
		new(rpc_health.Service),
	)

	s := g.Server()

	s.Group("/health", func(group *ghttp.RouterGroup) {
		group.GET("/check", api_health.Check)
	})

	s.Group("/book", func(group *ghttp.RouterGroup) {
		group.POST("/create", api_book.Create)
		group.GET("/get", api_book.Get)
	})
}
