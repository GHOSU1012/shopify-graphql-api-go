// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/r0busta/go-shopify-graphql/v5 (interfaces: LocationService)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/r0busta/go-shopify-graphql-model/v2/graph/model"
)

// MockLocationService is a mock of LocationService interface.
type MockLocationService struct {
	ctrl     *gomock.Controller
	recorder *MockLocationServiceMockRecorder
}

// MockLocationServiceMockRecorder is the mock recorder for MockLocationService.
type MockLocationServiceMockRecorder struct {
	mock *MockLocationService
}

// NewMockLocationService creates a new mock instance.
func NewMockLocationService(ctrl *gomock.Controller) *MockLocationService {
	mock := &MockLocationService{ctrl: ctrl}
	mock.recorder = &MockLocationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocationService) EXPECT() *MockLocationServiceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockLocationService) Get(arg0 string) (*model.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*model.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockLocationServiceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockLocationService)(nil).Get), arg0)
}