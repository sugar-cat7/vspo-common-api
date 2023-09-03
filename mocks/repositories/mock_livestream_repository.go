// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sugar-cat7/vspo-common-api/domain/repositories (interfaces: LiveStreamRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/sugar-cat7/vspo-common-api/domain/entities/legacy"
)

// MockLiveStreamRepository is a mock of LiveStreamRepository interface.
type MockLiveStreamRepository struct {
	ctrl     *gomock.Controller
	recorder *MockLiveStreamRepositoryMockRecorder
}

// MockLiveStreamRepositoryMockRecorder is the mock recorder for MockLiveStreamRepository.
type MockLiveStreamRepositoryMockRecorder struct {
	mock *MockLiveStreamRepository
}

// NewMockLiveStreamRepository creates a new mock instance.
func NewMockLiveStreamRepository(ctrl *gomock.Controller) *MockLiveStreamRepository {
	mock := &MockLiveStreamRepository{ctrl: ctrl}
	mock.recorder = &MockLiveStreamRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLiveStreamRepository) EXPECT() *MockLiveStreamRepositoryMockRecorder {
	return m.recorder
}

// CreateInBatch mocks base method.
func (m *MockLiveStreamRepository) CreateInBatch(arg0 []*entities.OldVideo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInBatch", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateInBatch indicates an expected call of CreateInBatch.
func (mr *MockLiveStreamRepositoryMockRecorder) CreateInBatch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInBatch", reflect.TypeOf((*MockLiveStreamRepository)(nil).CreateInBatch), arg0)
}

// FindAllByPeriod mocks base method.
func (m *MockLiveStreamRepository) FindAllByPeriod(arg0, arg1 string) ([]*entities.OldVideo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllByPeriod", arg0, arg1)
	ret0, _ := ret[0].([]*entities.OldVideo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllByPeriod indicates an expected call of FindAllByPeriod.
func (mr *MockLiveStreamRepositoryMockRecorder) FindAllByPeriod(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllByPeriod", reflect.TypeOf((*MockLiveStreamRepository)(nil).FindAllByPeriod), arg0, arg1)
}

// UpdateInBatch mocks base method.
func (m *MockLiveStreamRepository) UpdateInBatch(arg0 []*entities.OldVideo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInBatch", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInBatch indicates an expected call of UpdateInBatch.
func (mr *MockLiveStreamRepositoryMockRecorder) UpdateInBatch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInBatch", reflect.TypeOf((*MockLiveStreamRepository)(nil).UpdateInBatch), arg0)
}