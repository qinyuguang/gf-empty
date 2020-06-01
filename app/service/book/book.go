package book

import (
	"context"

	// for "go test ./..."
	_ "gf-empty/boot"

	dto_book "gf-empty/app/dto/book"
	model_book "gf-empty/app/model/book"
	"gf-empty/library/errors"
	"gf-empty/library/snowflake"

	"github.com/gogf/gf/util/gconv"
)

var Ins Service = new(service)

type Service interface {
	Create(context.Context, dto_book.Request) errors.Error
	Get(context.Context, dto_book.Request) (*model_book.Entity, errors.Error)
}

type service struct{}

func (s *service) Create(ctx context.Context, request dto_book.Request) errors.Error {
	req, err := request.CreateRequest()
	if err != nil {
		return err
	}

	var entity *model_book.Entity
	gconv.Struct(req, &entity)
	entity.Id = snowflake.NextIDInt64()

	dbErr := model_book.Ins.Create(entity)
	if dbErr != nil {
		return errors.NewInternalError(
			errors.UNAVAILABLE,
			errors.GetMessage(errors.UNAVAILABLE),
		)
	}

	return nil
}

func (s *service) Get(ctx context.Context, request dto_book.Request) (*model_book.Entity, errors.Error) {
	req, err := request.GetRequest()
	if err != nil {
		return nil, err
	}

	res, dbErr := model_book.Ins.GetByID(req.Id)
	if dbErr != nil {
		return nil, errors.NewInternalError(
			errors.UNAVAILABLE,
			errors.GetMessage(errors.UNAVAILABLE),
		)
	}

	return res, nil
}
