package book

import (
	model_book "gf-empty/app/model/book"
	"gf-empty/library/errors"

	bookpb "gf-empty/proto/message/book"

	"github.com/gogf/gf/util/gconv"
)

func CreateSuccess() *bookpb.CreateResponse {
	return &bookpb.CreateResponse{
		Code: 0,
		Msg:  "",
	}
}

func CreateFail(code errors.ResponseCode, msg string) *bookpb.CreateResponse {
	return &bookpb.CreateResponse{
		Code: int64(code),
		Msg:  msg,
	}
}

func GetSuccess(data *model_book.Entity) *bookpb.GetResponse {
	var book *bookpb.Book

	gconv.Struct(data, &book)

	return &bookpb.GetResponse{
		Code: 0,
		Msg:  "",
		Data: book,
	}
}

func GetFail(code errors.ResponseCode, msg string) *bookpb.GetResponse {
	return &bookpb.GetResponse{
		Code: int64(code),
		Msg:  msg,
		Data: nil,
	}
}
