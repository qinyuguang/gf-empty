package errors

import (
	"encoding/json"

	"github.com/gogf/gf/frame/g"
)

type Error interface {
	error
	Code() ResponseCode
	Message() string
}

type ServiceError struct {
	error
	code    ResponseCode
	message string
}

func NewServiceError(code ResponseCode, message string) Error {
	err := &ServiceError{
		code:    code,
		message: message,
	}

	g.Log().Infof("ServiceError: code:%d message:%s", code, message)

	return err
}

func (err *ServiceError) Error() string {
	result, _ := json.Marshal(err)

	return string(result)
}

func (err *ServiceError) Code() ResponseCode {
	return err.code
}

func (err *ServiceError) Message() string {
	return err.message
}

type InternalError struct {
	error
	code    ResponseCode
	message string
}

func NewInternalError(code ResponseCode, message string) Error {
	err := &InternalError{
		code:    code,
		message: message,
	}

	g.Log().Errorf("InternalError: code:%d message:%s", code, message)

	return err
}

func (err *InternalError) Error() string {
	result, _ := json.Marshal(err)

	return string(result)
}

func (err *InternalError) Code() ResponseCode {
	return err.code
}

func (err *InternalError) Message() string {
	return err.message
}
