// Code generated by MockGen. DO NOT EDIT.
// Source: ./resolver.go

// Package testdata is a generated GoMock package.
package testdata

import (
	reflect "reflect"

	core "github.com/d34dl0ck/coupler/internal/core"
	gomock "github.com/golang/mock/gomock"
)

// MockResolver is a mock of Resolver interface.
type MockResolver struct {
	ctrl     *gomock.Controller
	recorder *MockResolverMockRecorder
}

// MockResolverMockRecorder is the mock recorder for MockResolver.
type MockResolverMockRecorder struct {
	mock *MockResolver
}

// NewMockResolver creates a new mock instance.
func NewMockResolver(ctrl *gomock.Controller) *MockResolver {
	mock := &MockResolver{ctrl: ctrl}
	mock.recorder = &MockResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResolver) EXPECT() *MockResolverMockRecorder {
	return m.recorder
}

// Resolve mocks base method.
func (m *MockResolver) Resolve(k core.DependencyKey) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resolve", k)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Resolve indicates an expected call of Resolve.
func (mr *MockResolverMockRecorder) Resolve(k interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resolve", reflect.TypeOf((*MockResolver)(nil).Resolve), k)
}
