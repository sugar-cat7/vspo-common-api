// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sugar-cat7/vspo-common-api/domain/repositories (interfaces: SongRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/sugar-cat7/vspo-common-api/domain/entities"
)

// MockSongRepository is a mock of SongRepository interface.
type MockSongRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSongRepositoryMockRecorder
}

// MockSongRepositoryMockRecorder is the mock recorder for MockSongRepository.
type MockSongRepositoryMockRecorder struct {
	mock *MockSongRepository
}

// NewMockSongRepository creates a new mock instance.
func NewMockSongRepository(ctrl *gomock.Controller) *MockSongRepository {
	mock := &MockSongRepository{ctrl: ctrl}
	mock.recorder = &MockSongRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSongRepository) EXPECT() *MockSongRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSongRepository) Create(arg0 *entities.Video) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockSongRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSongRepository)(nil).Create), arg0)
}

// CreateInBatch mocks base method.
func (m *MockSongRepository) CreateInBatch(arg0 []*entities.Video) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInBatch", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateInBatch indicates an expected call of CreateInBatch.
func (mr *MockSongRepositoryMockRecorder) CreateInBatch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInBatch", reflect.TypeOf((*MockSongRepository)(nil).CreateInBatch), arg0)
}

// Delete mocks base method.
func (m *MockSongRepository) Delete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockSongRepositoryMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSongRepository)(nil).Delete), arg0)
}

// GetAll mocks base method.
func (m *MockSongRepository) GetAll() ([]*entities.Video, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*entities.Video)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockSongRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockSongRepository)(nil).GetAll))
}

// GetByID mocks base method.
func (m *MockSongRepository) GetByID(arg0 string) (*entities.Video, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(*entities.Video)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockSongRepositoryMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockSongRepository)(nil).GetByID), arg0)
}

// Update mocks base method.
func (m *MockSongRepository) Update(arg0 *entities.Video) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockSongRepositoryMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSongRepository)(nil).Update), arg0)
}

// UpdateInBatch mocks base method.
func (m *MockSongRepository) UpdateInBatch(arg0 []*entities.Video) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInBatch", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInBatch indicates an expected call of UpdateInBatch.
func (mr *MockSongRepositoryMockRecorder) UpdateInBatch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInBatch", reflect.TypeOf((*MockSongRepository)(nil).UpdateInBatch), arg0)
}
