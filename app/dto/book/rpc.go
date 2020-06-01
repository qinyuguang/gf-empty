package book

import (
	"gf-empty/library/errors"

	bookpb "gf-empty/proto/message/book"

	"github.com/gogf/gf/util/gconv"
)

type RPC struct {
	Create *bookpb.CreateRequest
	Get    *bookpb.GetRequest
}

func (s *RPC) CreateRequest() (*CreateRequest, errors.Error) {
	req := s.Create

	var r *CreateRequest

	if err := gconv.Struct(req.GetBook(), &r); err != nil {
		return nil, errors.NewInternalError(
			errors.INTERNAL_ERROR,
			err.Error(),
		)
	}

	return r, nil
}

func (s *RPC) GetRequest() (*GetRequest, errors.Error) {
	return &GetRequest{
		Id: s.Get.GetId(),
	}, nil
}
