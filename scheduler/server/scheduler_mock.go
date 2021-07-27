// Code generated by MockGen. DO NOT EDIT.
// Source: scheduler.go

// Package server is a generated GoMock package.
package server

import (
	gomock "github.com/golang/mock/gomock"
	stats "github.com/twitter/scoot/common/stats"
	saga "github.com/twitter/scoot/saga"
	domain "github.com/twitter/scoot/scheduler/domain"
	reflect "reflect"
	time "time"
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
func (m *MockScheduler) ScheduleJob(jobDef domain.JobDefinition) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScheduleJob", jobDef)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScheduleJob indicates an expected call of ScheduleJob
func (mr *MockSchedulerMockRecorder) ScheduleJob(jobDef interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleJob", reflect.TypeOf((*MockScheduler)(nil).ScheduleJob), jobDef)
}

// KillJob mocks base method
func (m *MockScheduler) KillJob(jobId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KillJob", jobId)
	ret0, _ := ret[0].(error)
	return ret0
}

// KillJob indicates an expected call of KillJob
func (mr *MockSchedulerMockRecorder) KillJob(jobId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KillJob", reflect.TypeOf((*MockScheduler)(nil).KillJob), jobId)
}

// GetSagaCoord mocks base method
func (m *MockScheduler) GetSagaCoord() saga.SagaCoordinator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSagaCoord")
	ret0, _ := ret[0].(saga.SagaCoordinator)
	return ret0
}

// GetSagaCoord indicates an expected call of GetSagaCoord
func (mr *MockSchedulerMockRecorder) GetSagaCoord() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSagaCoord", reflect.TypeOf((*MockScheduler)(nil).GetSagaCoord))
}

// OfflineWorker mocks base method
func (m *MockScheduler) OfflineWorker(req domain.OfflineWorkerReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OfflineWorker", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// OfflineWorker indicates an expected call of OfflineWorker
func (mr *MockSchedulerMockRecorder) OfflineWorker(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OfflineWorker", reflect.TypeOf((*MockScheduler)(nil).OfflineWorker), req)
}

// ReinstateWorker mocks base method
func (m *MockScheduler) ReinstateWorker(req domain.ReinstateWorkerReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReinstateWorker", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReinstateWorker indicates an expected call of ReinstateWorker
func (mr *MockSchedulerMockRecorder) ReinstateWorker(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReinstateWorker", reflect.TypeOf((*MockScheduler)(nil).ReinstateWorker), req)
}

// SetSchedulerStatus mocks base method
func (m *MockScheduler) SetSchedulerStatus(maxTasks int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSchedulerStatus", maxTasks)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSchedulerStatus indicates an expected call of SetSchedulerStatus
func (mr *MockSchedulerMockRecorder) SetSchedulerStatus(maxTasks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSchedulerStatus", reflect.TypeOf((*MockScheduler)(nil).SetSchedulerStatus), maxTasks)
}

// GetSchedulerStatus mocks base method
func (m *MockScheduler) GetSchedulerStatus() (int, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSchedulerStatus")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// GetSchedulerStatus indicates an expected call of GetSchedulerStatus
func (mr *MockSchedulerMockRecorder) GetSchedulerStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchedulerStatus", reflect.TypeOf((*MockScheduler)(nil).GetSchedulerStatus))
}

// GetClassLoadPercents mocks base method
func (m *MockScheduler) GetClassLoadPercents() (map[string]int32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClassLoadPercents")
	ret0, _ := ret[0].(map[string]int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClassLoadPercents indicates an expected call of GetClassLoadPercents
func (mr *MockSchedulerMockRecorder) GetClassLoadPercents() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClassLoadPercents", reflect.TypeOf((*MockScheduler)(nil).GetClassLoadPercents))
}

// SetClassLoadPercents mocks base method
func (m *MockScheduler) SetClassLoadPercents(classLoads map[string]int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetClassLoadPercents", classLoads)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetClassLoadPercents indicates an expected call of SetClassLoadPercents
func (mr *MockSchedulerMockRecorder) SetClassLoadPercents(classLoads interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClassLoadPercents", reflect.TypeOf((*MockScheduler)(nil).SetClassLoadPercents), classLoads)
}

// GetRequestorToClassMap mocks base method
func (m *MockScheduler) GetRequestorToClassMap() (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestorToClassMap")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequestorToClassMap indicates an expected call of GetRequestorToClassMap
func (mr *MockSchedulerMockRecorder) GetRequestorToClassMap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestorToClassMap", reflect.TypeOf((*MockScheduler)(nil).GetRequestorToClassMap))
}

// SetRequestorToClassMap mocks base method
func (m *MockScheduler) SetRequestorToClassMap(requestorToClassMap map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetRequestorToClassMap", requestorToClassMap)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetRequestorToClassMap indicates an expected call of SetRequestorToClassMap
func (mr *MockSchedulerMockRecorder) SetRequestorToClassMap(requestorToClassMap interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRequestorToClassMap", reflect.TypeOf((*MockScheduler)(nil).SetRequestorToClassMap), requestorToClassMap)
}

// GetRebalanceMinimumDuration mocks base method
func (m *MockScheduler) GetRebalanceMinimumDuration() (time.Duration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRebalanceMinimumDuration")
	ret0, _ := ret[0].(time.Duration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRebalanceMinimumDuration indicates an expected call of GetRebalanceMinimumDuration
func (mr *MockSchedulerMockRecorder) GetRebalanceMinimumDuration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRebalanceMinimumDuration", reflect.TypeOf((*MockScheduler)(nil).GetRebalanceMinimumDuration))
}

// SetRebalanceMinimumDuration mocks base method
func (m *MockScheduler) SetRebalanceMinimumDuration(durationMin time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetRebalanceMinimumDuration", durationMin)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetRebalanceMinimumDuration indicates an expected call of SetRebalanceMinimumDuration
func (mr *MockSchedulerMockRecorder) SetRebalanceMinimumDuration(durationMin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRebalanceMinimumDuration", reflect.TypeOf((*MockScheduler)(nil).SetRebalanceMinimumDuration), durationMin)
}

// GetRebalanceThreshold mocks base method
func (m *MockScheduler) GetRebalanceThreshold() (int32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRebalanceThreshold")
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRebalanceThreshold indicates an expected call of GetRebalanceThreshold
func (mr *MockSchedulerMockRecorder) GetRebalanceThreshold() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRebalanceThreshold", reflect.TypeOf((*MockScheduler)(nil).GetRebalanceThreshold))
}

// SetRebalanceThreshold mocks base method
func (m *MockScheduler) SetRebalanceThreshold(durationMin int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetRebalanceThreshold", durationMin)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetRebalanceThreshold indicates an expected call of SetRebalanceThreshold
func (mr *MockSchedulerMockRecorder) SetRebalanceThreshold(durationMin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRebalanceThreshold", reflect.TypeOf((*MockScheduler)(nil).SetRebalanceThreshold), durationMin)
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
func (m *MockSchedulingAlgorithm) GetTasksToBeAssigned(jobs []*jobState, stat stats.StatsReceiver, cs *clusterState, requestors map[string][]*jobState) ([]*taskState, []*taskState) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTasksToBeAssigned", jobs, stat, cs, requestors)
	ret0, _ := ret[0].([]*taskState)
	ret1, _ := ret[1].([]*taskState)
	return ret0, ret1
}

// GetTasksToBeAssigned indicates an expected call of GetTasksToBeAssigned
func (mr *MockSchedulingAlgorithmMockRecorder) GetTasksToBeAssigned(jobs, stat, cs, requestors interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasksToBeAssigned", reflect.TypeOf((*MockSchedulingAlgorithm)(nil).GetTasksToBeAssigned), jobs, stat, cs, requestors)
}
