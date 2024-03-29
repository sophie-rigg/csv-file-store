// Code generated by MockGen. DO NOT EDIT.
// Source: storage.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockStorage) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockStorageMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStorage)(nil).Close))
}

// CreateFile mocks base method.
func (m *MockStorage) CreateFile(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFile", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFile indicates an expected call of CreateFile.
func (mr *MockStorageMockRecorder) CreateFile(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFile", reflect.TypeOf((*MockStorage)(nil).CreateFile), id)
}

// ListFiles mocks base method.
func (m *MockStorage) ListFiles() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFiles")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFiles indicates an expected call of ListFiles.
func (mr *MockStorageMockRecorder) ListFiles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFiles", reflect.TypeOf((*MockStorage)(nil).ListFiles))
}

// OpenFile mocks base method.
func (m *MockStorage) OpenFile(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenFile", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// OpenFile indicates an expected call of OpenFile.
func (mr *MockStorageMockRecorder) OpenFile(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenFile", reflect.TypeOf((*MockStorage)(nil).OpenFile), id)
}

// Read mocks base method.
func (m *MockStorage) Read(p []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockStorageMockRecorder) Read(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockStorage)(nil).Read), p)
}

// RemoveFile mocks base method.
func (m *MockStorage) RemoveFile(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFile", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFile indicates an expected call of RemoveFile.
func (mr *MockStorageMockRecorder) RemoveFile(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFile", reflect.TypeOf((*MockStorage)(nil).RemoveFile), id)
}

// Write mocks base method.
func (m *MockStorage) Write(p []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockStorageMockRecorder) Write(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockStorage)(nil).Write), p)
}
