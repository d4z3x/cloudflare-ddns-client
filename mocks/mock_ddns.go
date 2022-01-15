// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/matt/cloudflare-ddns-client/ddns/ddns.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	ddns "github.com/mattolenik/cloudflare-ddns-client/ddns"
)

// MockDDNSProvider is a mock of DDNSProvider interface.
type MockDDNSProvider struct {
	ctrl     *gomock.Controller
	recorder *MockDDNSProviderMockRecorder
}

// MockDDNSProviderMockRecorder is the mock recorder for MockDDNSProvider.
type MockDDNSProviderMockRecorder struct {
	mock *MockDDNSProvider
}

// NewMockDDNSProvider creates a new mock instance.
func NewMockDDNSProvider(ctrl *gomock.Controller) *MockDDNSProvider {
	mock := &MockDDNSProvider{ctrl: ctrl}
	mock.recorder = &MockDDNSProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDDNSProvider) EXPECT() *MockDDNSProviderMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockDDNSProvider) Get(domain, record string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", domain, record)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDDNSProviderMockRecorder) Get(domain, record interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDDNSProvider)(nil).Get), domain, record)
}

// Update mocks base method.
func (m *MockDDNSProvider) Update(domain, record, ip string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", domain, record, ip)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockDDNSProviderMockRecorder) Update(domain, record, ip interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDDNSProvider)(nil).Update), domain, record, ip)
}

// MockDaemon is a mock of Daemon interface.
type MockDaemon struct {
	ctrl     *gomock.Controller
	recorder *MockDaemonMockRecorder
}

// MockDaemonMockRecorder is the mock recorder for MockDaemon.
type MockDaemonMockRecorder struct {
	mock *MockDaemon
}

// NewMockDaemon creates a new mock instance.
func NewMockDaemon(ctrl *gomock.Controller) *MockDaemon {
	mock := &MockDaemon{ctrl: ctrl}
	mock.recorder = &MockDaemonMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDaemon) EXPECT() *MockDaemonMockRecorder {
	return m.recorder
}

// Start mocks base method.
func (m *MockDaemon) Start(provider ddns.DDNSProvider, updatePeriod, failureRetryDelay time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", provider, updatePeriod, failureRetryDelay)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockDaemonMockRecorder) Start(provider, updatePeriod, failureRetryDelay interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockDaemon)(nil).Start), provider, updatePeriod, failureRetryDelay)
}

// Stop mocks base method.
func (m *MockDaemon) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockDaemonMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockDaemon)(nil).Stop))
}

// Update mocks base method.
func (m *MockDaemon) Update(provider ddns.DDNSProvider) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", provider)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockDaemonMockRecorder) Update(provider interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDaemon)(nil).Update), provider)
}
