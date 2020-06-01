package book

import (
	"gf-empty/library/errors"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

type CreateForm struct {
	Name   string  `v:"required|min-length:1#invalid name"`
	Author string  `v:"required|min-length:1#invalid author"`
	Price  float64 `v:"min:0#invalid price"`
}

type GetForm struct {
	Id int64 `v:"min:1#invalid id"`
}

type HTTP struct {
	Create *ghttp.Request
	Get    *ghttp.Request
}

func (s *HTTP) CreateRequest() (*CreateRequest, errors.Error) {
	req := s.Create

	var data *CreateForm

	if err := req.GetFormStruct(&data); err != nil {
		return nil, errors.NewServiceError(
			errors.INVALID_ARGS,
			errors.GetMessage(errors.INVALID_ARGS),
		)
	}

	if err := gvalid.CheckStruct(data, nil); err != nil {
		return nil, errors.NewServiceError(
			errors.INVALID_ARGS,
			err.FirstString(),
		)
	}

	var r *CreateRequest

	if err := gconv.Struct(data, &r); err != nil {
		return nil, errors.NewInternalError(
			errors.INTERNAL_ERROR,
			err.Error(),
		)
	}

	return r, nil
}

func (s *HTTP) GetRequest() (*GetRequest, errors.Error) {
	req := s.Get

	var data *GetForm

	if err := req.GetQueryStruct(&data); err != nil {
		return nil, errors.NewServiceError(
			errors.INVALID_ARGS,
			errors.GetMessage(errors.INVALID_ARGS),
		)
	}

	if err := gvalid.CheckStruct(data, nil); err != nil {
		return nil, errors.NewServiceError(
			errors.INVALID_ARGS,
			err.FirstString(),
		)
	}

	return &GetRequest{
		Id: data.Id,
	}, nil
}
