// Code generated by MockGen. DO NOT EDIT.
// Source: get_details.go

// Package mock is a generated GoMock package.
package mock

import (
	entity "github.com/basilhe/tdd/application/core/entity"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIGetDeviceDetails is a mock of IGetDeviceDetails interface
type MockIGetDeviceDetails struct {
	ctrl     *gomock.Controller
	recorder *MockIGetDeviceDetailsMockRecorder
}

// MockIGetDeviceDetailsMockRecorder is the mock recorder for MockIGetDeviceDetails
type MockIGetDeviceDetailsMockRecorder struct {
	mock *MockIGetDeviceDetails
}

// NewMockIGetDeviceDetails creates a new mock instance
func NewMockIGetDeviceDetails(ctrl *gomock.Controller) *MockIGetDeviceDetails {
	mock := &MockIGetDeviceDetails{ctrl: ctrl}
	mock.recorder = &MockIGetDeviceDetailsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIGetDeviceDetails) EXPECT() *MockIGetDeviceDetailsMockRecorder {
	return m.recorder
}

// GetDetails mocks base method
func (m *MockIGetDeviceDetails) GetDetails(hostname string) (*entity.BroadbandAccessDevice, error) {
	ret := m.ctrl.Call(m, "GetDetails", hostname)
	ret0, _ := ret[0].(*entity.BroadbandAccessDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetails indicates an expected call of GetDetails
func (mr *MockIGetDeviceDetailsMockRecorder) GetDetails(hostname interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetails", reflect.TypeOf((*MockIGetDeviceDetails)(nil).GetDetails), hostname)
}