package book

import (
	dto_book "gf-empty/app/dto/book"
	service_book "gf-empty/app/service/book"
	"gf-empty/library/response"

	"github.com/gogf/gf/net/ghttp"
)

// @summary Create is a demonstration route handler
// @tags    book
// @produce json
// @param   name	formData   string    true    "book name"
// @param   author  formData   string    true    "book author"
// @param   price	formData   float64   true    "book price"
// @router  /book/create [POST]
// @success 200 {object} response.JSONResponse "response"
func Create(r *ghttp.Request) {
	if err := service_book.Ins.Create(r.Context(), &dto_book.HTTP{Create: r}); err != nil {
		response.Fail(r, err.Code(), err.Message(), nil)
	}

	response.Success(r, nil)
}

// @summary Get is a demonstration route handler
// @tags    book
// @produce json
// @param   id  query   int     true     "book id"
// @router  /book/get [GET]
// @success 200 {object} response.JSONResponse "response"
func Get(r *ghttp.Request) {
	res, err := service_book.Ins.Get(r.Context(), &dto_book.HTTP{Get: r})
	if err != nil {
		response.Fail(r, err.Code(), err.Message(), nil)
	}

	response.Success(r, res)
}
