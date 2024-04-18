// Code generated by MockGen. DO NOT EDIT.
// Source: version.go
//
// Generated by this command:
//
//	mockgen -source=version.go -package=version -destination=./version_mock.go
//

// Package version is a generated GoMock package.
package version

import (
	reflect "reflect"

	version "github.com/hashicorp/go-version"
	gomock "go.uber.org/mock/gomock"
)

// MockRosaVersion is a mock of RosaVersion interface.
type MockRosaVersion struct {
	ctrl     *gomock.Controller
	recorder *MockRosaVersionMockRecorder
}

// MockRosaVersionMockRecorder is the mock recorder for MockRosaVersion.
type MockRosaVersionMockRecorder struct {
	mock *MockRosaVersion
}

// NewMockRosaVersion creates a new mock instance.
func NewMockRosaVersion(ctrl *gomock.Controller) *MockRosaVersion {
	mock := &MockRosaVersion{ctrl: ctrl}
	mock.recorder = &MockRosaVersionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRosaVersion) EXPECT() *MockRosaVersionMockRecorder {
	return m.recorder
}

// IsLatest mocks base method.
func (m *MockRosaVersion) IsLatest(latestVersion string) (*version.Version, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLatest", latestVersion)
	ret0, _ := ret[0].(*version.Version)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// IsLatest indicates an expected call of IsLatest.
func (mr *MockRosaVersionMockRecorder) IsLatest(latestVersion any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLatest", reflect.TypeOf((*MockRosaVersion)(nil).IsLatest), latestVersion)
}
