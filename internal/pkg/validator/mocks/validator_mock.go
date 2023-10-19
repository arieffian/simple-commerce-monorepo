// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pkg/validator/validator.go

// Package mock_validator is a generated GoMock package.
package mock_validator

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockValidatorService is a mock of ValidatorService interface.
type MockValidatorService struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorServiceMockRecorder
}

// MockValidatorServiceMockRecorder is the mock recorder for MockValidatorService.
type MockValidatorServiceMockRecorder struct {
	mock *MockValidatorService
}

// NewMockValidatorService creates a new mock instance.
func NewMockValidatorService(ctrl *gomock.Controller) *MockValidatorService {
	mock := &MockValidatorService{ctrl: ctrl}
	mock.recorder = &MockValidatorServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidatorService) EXPECT() *MockValidatorServiceMockRecorder {
	return m.recorder
}

// Validate mocks base method.
func (m *MockValidatorService) Validate(ctx context.Context, i interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", ctx, i)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockValidatorServiceMockRecorder) Validate(ctx, i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockValidatorService)(nil).Validate), ctx, i)
}