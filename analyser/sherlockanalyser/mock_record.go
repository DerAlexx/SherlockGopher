// Code generated by MockGen. DO NOT EDIT.
// Source: /home/devuser/snap/go/pkg/mod/github.com/neo4j/neo4j-go-driver@v1.7.4/neo4j/record.go

// Package mock_neo4j is a generated GoMock package.
package sherlockanalyser

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRecord is a mock of Record interface
type MockRecord struct {
	ctrl     *gomock.Controller
	recorder *MockRecordMockRecorder
}

// MockRecordMockRecorder is the mock recorder for MockRecord
type MockRecordMockRecorder struct {
	mock *MockRecord
}

// NewMockRecord creates a new mock instance
func NewMockRecord(ctrl *gomock.Controller) *MockRecord {
	mock := &MockRecord{ctrl: ctrl}
	mock.recorder = &MockRecordMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRecord) EXPECT() *MockRecordMockRecorder {
	return m.recorder
}

// Keys mocks base method
func (m *MockRecord) Keys() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Keys indicates an expected call of Keys
func (mr *MockRecordMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockRecord)(nil).Keys))
}

// Values mocks base method
func (m *MockRecord) Values() []interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Values")
	ret0, _ := ret[0].([]interface{})
	return ret0
}

// Values indicates an expected call of Values
func (mr *MockRecordMockRecorder) Values() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Values", reflect.TypeOf((*MockRecord)(nil).Values))
}

// Get mocks base method
func (m *MockRecord) Get(key string) (interface{}, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRecordMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRecord)(nil).Get), key)
}

// GetByIndex mocks base method
func (m *MockRecord) GetByIndex(index int) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIndex", index)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// GetByIndex indicates an expected call of GetByIndex
func (mr *MockRecordMockRecorder) GetByIndex(index interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIndex", reflect.TypeOf((*MockRecord)(nil).GetByIndex), index)
}