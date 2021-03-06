// Code generated by MockGen. DO NOT EDIT.
// Source: app/model/book/book.go

// Package mock_book is a generated GoMock package.
package mock_book

import (
	book "gf-empty/app/model/book"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIModel is a mock of IModel interface
type MockIModel struct {
	ctrl     *gomock.Controller
	recorder *MockIModelMockRecorder
}

// MockIModelMockRecorder is the mock recorder for MockIModel
type MockIModelMockRecorder struct {
	mock *MockIModel
}

// NewMockIModel creates a new mock instance
func NewMockIModel(ctrl *gomock.Controller) *MockIModel {
	mock := &MockIModel{ctrl: ctrl}
	mock.recorder = &MockIModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIModel) EXPECT() *MockIModelMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockIModel) Create(arg0 *book.Entity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockIModelMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIModel)(nil).Create), arg0)
}

// GetByID mocks base method
func (m *MockIModel) GetByID(arg0 int64) (*book.Entity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(*book.Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockIModelMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIModel)(nil).GetByID), arg0)
}
