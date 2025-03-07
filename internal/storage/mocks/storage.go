// Code generated by MockGen. DO NOT EDIT.
// Source: internal/storage/storage.go

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

// FoundShortUrl mocks base method.
func (m *MockStorage) FoundShortUrl(url string) (string, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FoundShortUrl", url)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FoundShortUrl indicates an expected call of FoundShortUrl.
func (mr *MockStorageMockRecorder) FoundShortUrl(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FoundShortUrl", reflect.TypeOf((*MockStorage)(nil).FoundShortUrl), url)
}

// FoundUrl mocks base method.
func (m *MockStorage) FoundUrl(shortUrl string) (string, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FoundUrl", shortUrl)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FoundUrl indicates an expected call of FoundUrl.
func (mr *MockStorageMockRecorder) FoundUrl(shortUrl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FoundUrl", reflect.TypeOf((*MockStorage)(nil).FoundUrl), shortUrl)
}

// SaveUrl mocks base method.
func (m *MockStorage) SaveUrl(url, shortUrl string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUrl", url, shortUrl)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveUrl indicates an expected call of SaveUrl.
func (mr *MockStorageMockRecorder) SaveUrl(url, shortUrl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUrl", reflect.TypeOf((*MockStorage)(nil).SaveUrl), url, shortUrl)
}
