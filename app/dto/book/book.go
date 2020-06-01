package book

import (
	"gf-empty/library/errors"
)

type Request interface {
	CreateRequest() (*CreateRequest, errors.Error)
	GetRequest() (*GetRequest, errors.Error)
}

type CreateRequest struct {
	Name   string
	Author string
	Price  float64
}

type GetRequest struct {
	Id int64
}
