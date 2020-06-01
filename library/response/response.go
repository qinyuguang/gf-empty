package response

import (
	"gf-empty/library/errors"

	"github.com/gogf/gf/net/ghttp"
)

type JSONResponse struct {
	Code errors.ResponseCode `json:"code"` // 错误码 0成功
	Msg  string              `json:"msg"`  // 提示信息
	Data interface{}         `json:"data"` // 返回数据
}

func Success(r *ghttp.Request, data interface{}) {
	r.Response.WriteJson(JSONResponse{
		Code: 0,
		Msg:  "",
		Data: data,
	})
}

func Fail(r *ghttp.Request, code errors.ResponseCode, msg string, data interface{}) {
	r.Response.WriteJson(JSONResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	})
	r.Exit()
}
