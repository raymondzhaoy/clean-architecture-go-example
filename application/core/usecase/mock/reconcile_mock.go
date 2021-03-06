// Code generated by MockGen. DO NOT EDIT.
// Source: reconcile.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockOnSuccess is a mock of OnSuccess interface
type MockOnSuccess struct {
	ctrl     *gomock.Controller
	recorder *MockOnSuccessMockRecorder
}

// MockOnSuccessMockRecorder is the mock recorder for MockOnSuccess
type MockOnSuccessMockRecorder struct {
	mock *MockOnSuccess
}

// NewMockOnSuccess creates a new mock instance
func NewMockOnSuccess(ctrl *gomock.Controller) *MockOnSuccess {
	mock := &MockOnSuccess{ctrl: ctrl}
	mock.recorder = &MockOnSuccessMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOnSuccess) EXPECT() *MockOnSuccessMockRecorder {
	return m.recorder
}

// AuditSuccess mocks base method
func (m *MockOnSuccess) AuditSuccess() {
	m.ctrl.Call(m, "AuditSuccess")
}

// AuditSuccess indicates an expected call of AuditSuccess
func (mr *MockOnSuccessMockRecorder) AuditSuccess() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuditSuccess", reflect.TypeOf((*MockOnSuccess)(nil).AuditSuccess))
}

// MockOnFailure is a mock of OnFailure interface
type MockOnFailure struct {
	ctrl     *gomock.Controller
	recorder *MockOnFailureMockRecorder
}

// MockOnFailureMockRecorder is the mock recorder for MockOnFailure
type MockOnFailureMockRecorder struct {
	mock *MockOnFailure
}

// NewMockOnFailure creates a new mock instance
func NewMockOnFailure(ctrl *gomock.Controller) *MockOnFailure {
	mock := &MockOnFailure{ctrl: ctrl}
	mock.recorder = &MockOnFailureMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOnFailure) EXPECT() *MockOnFailureMockRecorder {
	return m.recorder
}

// AuditFailure mocks base method
func (m *MockOnFailure) AuditFailure() {
	m.ctrl.Call(m, "AuditFailure")
}

// AuditFailure indicates an expected call of AuditFailure
func (mr *MockOnFailureMockRecorder) AuditFailure() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuditFailure", reflect.TypeOf((*MockOnFailure)(nil).AuditFailure))
}

// MockIGetSerialNumberFromReality is a mock of IGetSerialNumberFromReality interface
type MockIGetSerialNumberFromReality struct {
	ctrl     *gomock.Controller
	recorder *MockIGetSerialNumberFromRealityMockRecorder
}

// MockIGetSerialNumberFromRealityMockRecorder is the mock recorder for MockIGetSerialNumberFromReality
type MockIGetSerialNumberFromRealityMockRecorder struct {
	mock *MockIGetSerialNumberFromReality
}

// NewMockIGetSerialNumberFromReality creates a new mock instance
func NewMockIGetSerialNumberFromReality(ctrl *gomock.Controller) *MockIGetSerialNumberFromReality {
	mock := &MockIGetSerialNumberFromReality{ctrl: ctrl}
	mock.recorder = &MockIGetSerialNumberFromRealityMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIGetSerialNumberFromReality) EXPECT() *MockIGetSerialNumberFromRealityMockRecorder {
	return m.recorder
}

// GetSerialNumber mocks base method
func (m *MockIGetSerialNumberFromReality) GetSerialNumber(hostname string) string {
	ret := m.ctrl.Call(m, "GetSerialNumber", hostname)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSerialNumber indicates an expected call of GetSerialNumber
func (mr *MockIGetSerialNumberFromRealityMockRecorder) GetSerialNumber(hostname interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSerialNumber", reflect.TypeOf((*MockIGetSerialNumberFromReality)(nil).GetSerialNumber), hostname)
}

// MockIGetSerialNumberFromModel is a mock of IGetSerialNumberFromModel interface
type MockIGetSerialNumberFromModel struct {
	ctrl     *gomock.Controller
	recorder *MockIGetSerialNumberFromModelMockRecorder
}

// MockIGetSerialNumberFromModelMockRecorder is the mock recorder for MockIGetSerialNumberFromModel
type MockIGetSerialNumberFromModelMockRecorder struct {
	mock *MockIGetSerialNumberFromModel
}

// NewMockIGetSerialNumberFromModel creates a new mock instance
func NewMockIGetSerialNumberFromModel(ctrl *gomock.Controller) *MockIGetSerialNumberFromModel {
	mock := &MockIGetSerialNumberFromModel{ctrl: ctrl}
	mock.recorder = &MockIGetSerialNumberFromModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIGetSerialNumberFromModel) EXPECT() *MockIGetSerialNumberFromModelMockRecorder {
	return m.recorder
}

// GetSerialNumber mocks base method
func (m *MockIGetSerialNumberFromModel) GetSerialNumber(hostname string) string {
	ret := m.ctrl.Call(m, "GetSerialNumber", hostname)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSerialNumber indicates an expected call of GetSerialNumber
func (mr *MockIGetSerialNumberFromModelMockRecorder) GetSerialNumber(hostname interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSerialNumber", reflect.TypeOf((*MockIGetSerialNumberFromModel)(nil).GetSerialNumber), hostname)
}

// MockIUpdateSerialNumberInModel is a mock of IUpdateSerialNumberInModel interface
type MockIUpdateSerialNumberInModel struct {
	ctrl     *gomock.Controller
	recorder *MockIUpdateSerialNumberInModelMockRecorder
}

// MockIUpdateSerialNumberInModelMockRecorder is the mock recorder for MockIUpdateSerialNumberInModel
type MockIUpdateSerialNumberInModelMockRecorder struct {
	mock *MockIUpdateSerialNumberInModel
}

// NewMockIUpdateSerialNumberInModel creates a new mock instance
func NewMockIUpdateSerialNumberInModel(ctrl *gomock.Controller) *MockIUpdateSerialNumberInModel {
	mock := &MockIUpdateSerialNumberInModel{ctrl: ctrl}
	mock.recorder = &MockIUpdateSerialNumberInModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIUpdateSerialNumberInModel) EXPECT() *MockIUpdateSerialNumberInModelMockRecorder {
	return m.recorder
}

// UpdateSerialNumber mocks base method
func (m *MockIUpdateSerialNumberInModel) UpdateSerialNumber(hostname, serialNumber string) error {
	ret := m.ctrl.Call(m, "UpdateSerialNumber", hostname, serialNumber)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSerialNumber indicates an expected call of UpdateSerialNumber
func (mr *MockIUpdateSerialNumberInModelMockRecorder) UpdateSerialNumber(hostname, serialNumber interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSerialNumber", reflect.TypeOf((*MockIUpdateSerialNumberInModel)(nil).UpdateSerialNumber), hostname, serialNumber)
}

// MockIGetAllDevicesHostname is a mock of IGetAllDevicesHostname interface
type MockIGetAllDevicesHostname struct {
	ctrl     *gomock.Controller
	recorder *MockIGetAllDevicesHostnameMockRecorder
}

// MockIGetAllDevicesHostnameMockRecorder is the mock recorder for MockIGetAllDevicesHostname
type MockIGetAllDevicesHostnameMockRecorder struct {
	mock *MockIGetAllDevicesHostname
}

// NewMockIGetAllDevicesHostname creates a new mock instance
func NewMockIGetAllDevicesHostname(ctrl *gomock.Controller) *MockIGetAllDevicesHostname {
	mock := &MockIGetAllDevicesHostname{ctrl: ctrl}
	mock.recorder = &MockIGetAllDevicesHostnameMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIGetAllDevicesHostname) EXPECT() *MockIGetAllDevicesHostnameMockRecorder {
	return m.recorder
}

// GetAllDevicesHostnames mocks base method
func (m *MockIGetAllDevicesHostname) GetAllDevicesHostnames() []string {
	ret := m.ctrl.Call(m, "GetAllDevicesHostnames")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetAllDevicesHostnames indicates an expected call of GetAllDevicesHostnames
func (mr *MockIGetAllDevicesHostnameMockRecorder) GetAllDevicesHostnames() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllDevicesHostnames", reflect.TypeOf((*MockIGetAllDevicesHostname)(nil).GetAllDevicesHostnames))
}
