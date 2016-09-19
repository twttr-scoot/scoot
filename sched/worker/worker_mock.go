// Automatically generated by MockGen. DO NOT EDIT!
// Source: worker.go

package worker

import (
	gomock "github.com/golang/mock/gomock"
	runner "github.com/scootdev/scoot/runner"
	sched "github.com/scootdev/scoot/sched"
)

// Mock of Worker interface
type MockWorker struct {
	ctrl     *gomock.Controller
	recorder *_MockWorkerRecorder
}

// Recorder for MockWorker (not exported)
type _MockWorkerRecorder struct {
	mock *MockWorker
}

func NewMockWorker(ctrl *gomock.Controller) *MockWorker {
	mock := &MockWorker{ctrl: ctrl}
	mock.recorder = &_MockWorkerRecorder{mock}
	return mock
}

func (_m *MockWorker) EXPECT() *_MockWorkerRecorder {
	return _m.recorder
}

func (_m *MockWorker) RunAndWait(task sched.TaskDefinition) (runner.ProcessStatus, error) {
	ret := _m.ctrl.Call(_m, "RunAndWait", task)
	ret0, _ := ret[0].(runner.ProcessStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockWorkerRecorder) RunAndWait(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RunAndWait", arg0)
}
