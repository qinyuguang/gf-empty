package book

import (
	"context"

	dto_book "gf-empty/app/dto/book"
	service_book "gf-empty/app/service/book"
	bookpb "gf-empty/proto/message/book"
)

type Book struct{}

func (s *Book) Create(ctx context.Context, req *bookpb.CreateRequest) (*bookpb.CreateResponse, error) {
	if err := service_book.Ins.Create(ctx, &dto_book.RPC{Create: req}); err != nil {
		return CreateFail(err.Code(), err.Message()), nil
	}

	return CreateSuccess(), nil
}

func (s *Book) Get(ctx context.Context, req *bookpb.GetRequest) (*bookpb.GetResponse, error) {
	res, err := service_book.Ins.Get(ctx, &dto_book.RPC{Get: req})
	if err != nil {
		return GetFail(err.Code(), err.Message()), nil
	}

	return GetSuccess(res), nil
}
