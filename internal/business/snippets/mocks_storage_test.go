// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package snippets_test is a generated GoMock package.
package snippets_test

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	snippets "github.com/titusjaka/go-sample/internal/business/snippets"
	service "github.com/titusjaka/go-sample/internal/infrastructure/service"
	reflect "reflect"
)

// MockStorage is a mock of Storage interface
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockStorage) Get(ctx context.Context, id uint) (snippets.Snippet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(snippets.Snippet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockStorageMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStorage)(nil).Get), ctx, id)
}

// Create mocks base method
func (m *MockStorage) Create(ctx context.Context, snippet snippets.Snippet) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, snippet)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockStorageMockRecorder) Create(ctx, snippet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStorage)(nil).Create), ctx, snippet)
}

// List mocks base method
func (m *MockStorage) List(ctx context.Context, pagination service.Pagination) ([]snippets.Snippet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, pagination)
	ret0, _ := ret[0].([]snippets.Snippet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockStorageMockRecorder) List(ctx, pagination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStorage)(nil).List), ctx, pagination)
}

// SoftDelete mocks base method
func (m *MockStorage) SoftDelete(ctx context.Context, id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SoftDelete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// SoftDelete indicates an expected call of SoftDelete
func (mr *MockStorageMockRecorder) SoftDelete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SoftDelete", reflect.TypeOf((*MockStorage)(nil).SoftDelete), ctx, id)
}

// Total mocks base method
func (m *MockStorage) Total(ctx context.Context) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Total", ctx)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Total indicates an expected call of Total
func (mr *MockStorageMockRecorder) Total(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Total", reflect.TypeOf((*MockStorage)(nil).Total), ctx)
}
