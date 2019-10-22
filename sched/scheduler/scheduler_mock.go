// Code generated by MockGen. DO NOT EDIT.
// Source: scheduler.go

// Package scheduler is a generated GoMock package.
package scheduler

import (
	gomock "github.com/golang/mock/gomock"
	stats "github.com/twitter/scoot/common/stats"
	saga "github.com/twitter/scoot/saga"
	sched "github.com/twitter/scoot/sched"
	reflect "reflect"
)

// MockScheduler is a mock of Scheduler interface
type MockScheduler struct {
	ctrl     *gomock.Controller
	recorder *MockSchedulerMockRecorder
}

// MockSchedulerMockRecorder is the mock recorder for MockScheduler
type MockSchedulerMockRecorder struct {
	mock *MockScheduler
}

// NewMockScheduler creates a new mock instance
func NewMockScheduler(ctrl *gomock.Controller) *MockScheduler {
	mock := &MockScheduler{ctrl: ctrl}
	mock.recorder = &MockSchedulerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockScheduler) EXPECT() *MockSchedulerMockRecorder {
	return m.recorder
}

// ScheduleJob mocks base method
func (m *MockScheduler) ScheduleJob(jobDef sched.JobDefinition) (string, error) {
	ret := m.ctrl.Call(m, "ScheduleJob", jobDef)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScheduleJob indicates an expected call of ScheduleJob
func (mr *MockSchedulerMockRecorder) ScheduleJob(jobDef interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleJob", reflect.TypeOf((*MockScheduler)(nil).ScheduleJob), jobDef)
}

// KillJob mocks base method
func (m *MockScheduler) KillJob(jobId string) error {
	ret := m.ctrl.Call(m, "KillJob", jobId)
	ret0, _ := ret[0].(error)
	return ret0
}

// KillJob indicates an expected call of KillJob
func (mr *MockSchedulerMockRecorder) KillJob(jobId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KillJob", reflect.TypeOf((*MockScheduler)(nil).KillJob), jobId)
}

// GetSagaCoord mocks base method
func (m *MockScheduler) GetSagaCoord() saga.SagaCoordinator {
	ret := m.ctrl.Call(m, "GetSagaCoord")
	ret0, _ := ret[0].(saga.SagaCoordinator)
	return ret0
}

// GetSagaCoord indicates an expected call of GetSagaCoord
func (mr *MockSchedulerMockRecorder) GetSagaCoord() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSagaCoord", reflect.TypeOf((*MockScheduler)(nil).GetSagaCoord))
}

// OfflineWorker mocks base method
func (m *MockScheduler) OfflineWorker(req sched.OfflineWorkerReq) error {
	ret := m.ctrl.Call(m, "OfflineWorker", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// OfflineWorker indicates an expected call of OfflineWorker
func (mr *MockSchedulerMockRecorder) OfflineWorker(req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OfflineWorker", reflect.TypeOf((*MockScheduler)(nil).OfflineWorker), req)
}

// ReinstateWorker mocks base method
func (m *MockScheduler) ReinstateWorker(req sched.ReinstateWorkerReq) error {
	ret := m.ctrl.Call(m, "ReinstateWorker", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReinstateWorker indicates an expected call of ReinstateWorker
func (mr *MockSchedulerMockRecorder) ReinstateWorker(req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReinstateWorker", reflect.TypeOf((*MockScheduler)(nil).ReinstateWorker), req)
}

// SetSchedulerStatus mocks base method
func (m *MockScheduler) SetSchedulerStatus(maxTasks int) error {
	ret := m.ctrl.Call(m, "SetSchedulerStatus", maxTasks)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSchedulerStatus indicates an expected call of SetSchedulerStatus
func (mr *MockSchedulerMockRecorder) SetSchedulerStatus(maxTasks interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSchedulerStatus", reflect.TypeOf((*MockScheduler)(nil).SetSchedulerStatus), maxTasks)
}

// GetSchedulerStatus mocks base method
func (m *MockScheduler) GetSchedulerStatus() (int, int) {
	ret := m.ctrl.Call(m, "GetSchedulerStatus")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// GetSchedulerStatus indicates an expected call of GetSchedulerStatus
func (mr *MockSchedulerMockRecorder) GetSchedulerStatus() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchedulerStatus", reflect.TypeOf((*MockScheduler)(nil).GetSchedulerStatus))
}

// MockSchedulingAlgorithm is a mock of SchedulingAlgorithm interface
type MockSchedulingAlgorithm struct {
	ctrl     *gomock.Controller
	recorder *MockSchedulingAlgorithmMockRecorder
}

// MockSchedulingAlgorithmMockRecorder is the mock recorder for MockSchedulingAlgorithm
type MockSchedulingAlgorithmMockRecorder struct {
	mock *MockSchedulingAlgorithm
}

// NewMockSchedulingAlgorithm creates a new mock instance
func NewMockSchedulingAlgorithm(ctrl *gomock.Controller) *MockSchedulingAlgorithm {
	mock := &MockSchedulingAlgorithm{ctrl: ctrl}
	mock.recorder = &MockSchedulingAlgorithmMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSchedulingAlgorithm) EXPECT() *MockSchedulingAlgorithmMockRecorder {
	return m.recorder
}

// GetTasksToBeAssigned mocks base method
func (m *MockSchedulingAlgorithm) GetTasksToBeAssigned(jobs []*jobState, stat stats.StatsReceiver, cs *clusterState, requestors map[string][]*jobState, cfg SchedulerConfig) []*taskState {
	ret := m.ctrl.Call(m, "GetTasksToBeAssigned", jobs, stat, cs, requestors, cfg)
	ret0, _ := ret[0].([]*taskState)
	return ret0
}

// GetTasksToBeAssigned indicates an expected call of GetTasksToBeAssigned
func (mr *MockSchedulingAlgorithmMockRecorder) GetTasksToBeAssigned(jobs, stat, cs, requestors, cfg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasksToBeAssigned", reflect.TypeOf((*MockSchedulingAlgorithm)(nil).GetTasksToBeAssigned), jobs, stat, cs, requestors, cfg)
}
