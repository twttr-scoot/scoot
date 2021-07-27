// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/twitter/scoot/bazel/cas/connection (interfaces: GRPCDialer,ClientConnPtr)

// Package mock_connection is a generated GoMock package.
package mock_connection

import (
	gomock "github.com/golang/mock/gomock"
	connection "github.com/twitter/scoot/bazel/cas/connection"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockGRPCDialer is a mock of GRPCDialer interface
type MockGRPCDialer struct {
	ctrl     *gomock.Controller
	recorder *MockGRPCDialerMockRecorder
}

// MockGRPCDialerMockRecorder is the mock recorder for MockGRPCDialer
type MockGRPCDialerMockRecorder struct {
	mock *MockGRPCDialer
}

// NewMockGRPCDialer creates a new mock instance
func NewMockGRPCDialer(ctrl *gomock.Controller) *MockGRPCDialer {
	mock := &MockGRPCDialer{ctrl: ctrl}
	mock.recorder = &MockGRPCDialerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGRPCDialer) EXPECT() *MockGRPCDialerMockRecorder {
	return m.recorder
}

// Dial mocks base method
func (m *MockGRPCDialer) Dial(arg0 string, arg1 ...grpc.DialOption) (connection.ClientConnPtr, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Dial", varargs...)
	ret0, _ := ret[0].(connection.ClientConnPtr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Dial indicates an expected call of Dial
func (mr *MockGRPCDialerMockRecorder) Dial(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dial", reflect.TypeOf((*MockGRPCDialer)(nil).Dial), varargs...)
}

// MockClientConnPtr is a mock of ClientConnPtr interface
type MockClientConnPtr struct {
	ctrl     *gomock.Controller
	recorder *MockClientConnPtrMockRecorder
}

// MockClientConnPtrMockRecorder is the mock recorder for MockClientConnPtr
type MockClientConnPtrMockRecorder struct {
	mock *MockClientConnPtr
}

// NewMockClientConnPtr creates a new mock instance
func NewMockClientConnPtr(ctrl *gomock.Controller) *MockClientConnPtr {
	mock := &MockClientConnPtr{ctrl: ctrl}
	mock.recorder = &MockClientConnPtrMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientConnPtr) EXPECT() *MockClientConnPtrMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockClientConnPtr) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockClientConnPtrMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClientConnPtr)(nil).Close))
}
