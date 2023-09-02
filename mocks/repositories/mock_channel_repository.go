// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sugar-cat7/vspo-common-api/domain/repositories (interfaces: ChannelRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/sugar-cat7/vspo-common-api/domain/entities"
)

// MockChannelRepository is a mock of ChannelRepository interface.
type MockChannelRepository struct {
	ctrl     *gomock.Controller
	recorder *MockChannelRepositoryMockRecorder
}

// MockChannelRepositoryMockRecorder is the mock recorder for MockChannelRepository.
type MockChannelRepositoryMockRecorder struct {
	mock *MockChannelRepository
}

// NewMockChannelRepository creates a new mock instance.
func NewMockChannelRepository(ctrl *gomock.Controller) *MockChannelRepository {
	mock := &MockChannelRepository{ctrl: ctrl}
	mock.recorder = &MockChannelRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChannelRepository) EXPECT() *MockChannelRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockChannelRepository) Create(arg0 *entities.Channel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockChannelRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockChannelRepository)(nil).Create), arg0)
}

// CreateInBatch mocks base method.
func (m *MockChannelRepository) CreateInBatch(arg0 []*entities.Channel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInBatch", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateInBatch indicates an expected call of CreateInBatch.
func (mr *MockChannelRepositoryMockRecorder) CreateInBatch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInBatch", reflect.TypeOf((*MockChannelRepository)(nil).CreateInBatch), arg0)
}

// Delete mocks base method.
func (m *MockChannelRepository) Delete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockChannelRepositoryMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockChannelRepository)(nil).Delete), arg0)
}

// GetAll mocks base method.
func (m *MockChannelRepository) GetAll() ([]*entities.Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*entities.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockChannelRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockChannelRepository)(nil).GetAll))
}

// GetByID mocks base method.
func (m *MockChannelRepository) GetByID(arg0 string) (*entities.Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(*entities.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockChannelRepositoryMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockChannelRepository)(nil).GetByID), arg0)
}

// GetInBatch mocks base method.
func (m *MockChannelRepository) GetInBatch(arg0 []string) ([]*entities.Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInBatch", arg0)
	ret0, _ := ret[0].([]*entities.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInBatch indicates an expected call of GetInBatch.
func (mr *MockChannelRepositoryMockRecorder) GetInBatch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInBatch", reflect.TypeOf((*MockChannelRepository)(nil).GetInBatch), arg0)
}

// Update mocks base method.
func (m *MockChannelRepository) Update(arg0 *entities.Channel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockChannelRepositoryMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockChannelRepository)(nil).Update), arg0)
}

// UpdateInBatch mocks base method.
func (m *MockChannelRepository) UpdateInBatch(arg0 []*entities.Channel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInBatch", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInBatch indicates an expected call of UpdateInBatch.
func (mr *MockChannelRepositoryMockRecorder) UpdateInBatch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInBatch", reflect.TypeOf((*MockChannelRepository)(nil).UpdateInBatch), arg0)
}
