// Code generated by MockGen. DO NOT EDIT.
// Source: /home/markus/Projects/SherlockGopher/analyser/proto/analysertowebserver/analysertowebserver.pb.micro.go

// Package mock_analysertowebserver is a generated GoMock package.
package mock_analysertowebserver

import (
	context "context"
	analysertowebserver "github.com/DerAlexx/SherlockGopher/analyser/proto"
	gomock "github.com/golang/mock/gomock"
	client "github.com/micro/go-micro/client"
	reflect "reflect"
)

// MockAnalyserService is a mock of AnalyserService interface
type MockAnalyserService struct {
	ctrl     *gomock.Controller
	recorder *MockAnalyserServiceMockRecorder
}

// MockAnalyserServiceMockRecorder is the mock recorder for MockAnalyserService
type MockAnalyserServiceMockRecorder struct {
	mock *MockAnalyserService
}

// NewMockAnalyserService creates a new mock instance
func NewMockAnalyserService(ctrl *gomock.Controller) *MockAnalyserService {
	mock := &MockAnalyserService{ctrl: ctrl}
	mock.recorder = &MockAnalyserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAnalyserService) EXPECT() *MockAnalyserServiceMockRecorder {
	return m.recorder
}

// WorkloadRPC mocks base method
func (m *MockAnalyserService) WorkloadRPC(ctx context.Context, in *analysertowebserver.WorkloadRequest, opts ...client.CallOption) (*analysertowebserver.WorkloadResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WorkloadRPC", varargs...)
	ret0, _ := ret[0].(*analysertowebserver.WorkloadResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WorkloadRPC indicates an expected call of WorkloadRPC
func (mr *MockAnalyserServiceMockRecorder) WorkloadRPC(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WorkloadRPC", reflect.TypeOf((*MockAnalyserService)(nil).WorkloadRPC), varargs...)
}

// ChangeStateRPC mocks base method
func (m *MockAnalyserService) ChangeStateRPC(ctx context.Context, in *analysertowebserver.ChangeStateRequest, opts ...client.CallOption) (*analysertowebserver.ChangeStateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChangeStateRPC", varargs...)
	ret0, _ := ret[0].(*analysertowebserver.ChangeStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeStateRPC indicates an expected call of ChangeStateRPC
func (mr *MockAnalyserServiceMockRecorder) ChangeStateRPC(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeStateRPC", reflect.TypeOf((*MockAnalyserService)(nil).ChangeStateRPC), varargs...)
}

// StateRPC mocks base method
func (m *MockAnalyserService) StateRPC(ctx context.Context, in *analysertowebserver.StateRequest, opts ...client.CallOption) (*analysertowebserver.StateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StateRPC", varargs...)
	ret0, _ := ret[0].(*analysertowebserver.StateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateRPC indicates an expected call of StateRPC
func (mr *MockAnalyserServiceMockRecorder) StateRPC(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateRPC", reflect.TypeOf((*MockAnalyserService)(nil).StateRPC), varargs...)
}

// WebsiteData mocks base method
func (m *MockAnalyserService) WebsiteData(ctx context.Context, opts ...client.CallOption) (analysertowebserver.Analyser_WebsiteDataService, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WebsiteData", varargs...)
	ret0, _ := ret[0].(analysertowebserver.Analyser_WebsiteDataService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WebsiteData indicates an expected call of WebsiteData
func (mr *MockAnalyserServiceMockRecorder) WebsiteData(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WebsiteData", reflect.TypeOf((*MockAnalyserService)(nil).WebsiteData), varargs...)
}

// MockAnalyser_WebsiteDataService is a mock of Analyser_WebsiteDataService interface
type MockAnalyser_WebsiteDataService struct {
	ctrl     *gomock.Controller
	recorder *MockAnalyser_WebsiteDataServiceMockRecorder
}

// MockAnalyser_WebsiteDataServiceMockRecorder is the mock recorder for MockAnalyser_WebsiteDataService
type MockAnalyser_WebsiteDataServiceMockRecorder struct {
	mock *MockAnalyser_WebsiteDataService
}

// NewMockAnalyser_WebsiteDataService creates a new mock instance
func NewMockAnalyser_WebsiteDataService(ctrl *gomock.Controller) *MockAnalyser_WebsiteDataService {
	mock := &MockAnalyser_WebsiteDataService{ctrl: ctrl}
	mock.recorder = &MockAnalyser_WebsiteDataServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAnalyser_WebsiteDataService) EXPECT() *MockAnalyser_WebsiteDataServiceMockRecorder {
	return m.recorder
}

// SendMsg mocks base method
func (m *MockAnalyser_WebsiteDataService) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockAnalyser_WebsiteDataServiceMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockAnalyser_WebsiteDataService)(nil).SendMsg), arg0)
}

// RecvMsg mocks base method
func (m *MockAnalyser_WebsiteDataService) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockAnalyser_WebsiteDataServiceMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockAnalyser_WebsiteDataService)(nil).RecvMsg), arg0)
}

// Close mocks base method
func (m *MockAnalyser_WebsiteDataService) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockAnalyser_WebsiteDataServiceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAnalyser_WebsiteDataService)(nil).Close))
}

// Send mocks base method
func (m *MockAnalyser_WebsiteDataService) Send(arg0 *analysertowebserver.CrawlerPackage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockAnalyser_WebsiteDataServiceMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockAnalyser_WebsiteDataService)(nil).Send), arg0)
}

// Recv mocks base method
func (m *MockAnalyser_WebsiteDataService) Recv() (*analysertowebserver.CrawlerAck, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*analysertowebserver.CrawlerAck)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockAnalyser_WebsiteDataServiceMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockAnalyser_WebsiteDataService)(nil).Recv))
}

// MockAnalyserHandler is a mock of AnalyserHandler interface
type MockAnalyserHandler struct {
	ctrl     *gomock.Controller
	recorder *MockAnalyserHandlerMockRecorder
}

// MockAnalyserHandlerMockRecorder is the mock recorder for MockAnalyserHandler
type MockAnalyserHandlerMockRecorder struct {
	mock *MockAnalyserHandler
}

// NewMockAnalyserHandler creates a new mock instance
func NewMockAnalyserHandler(ctrl *gomock.Controller) *MockAnalyserHandler {
	mock := &MockAnalyserHandler{ctrl: ctrl}
	mock.recorder = &MockAnalyserHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAnalyserHandler) EXPECT() *MockAnalyserHandlerMockRecorder {
	return m.recorder
}

// WorkloadRPC mocks base method
func (m *MockAnalyserHandler) WorkloadRPC(arg0 context.Context, arg1 *analysertowebserver.WorkloadRequest, arg2 *analysertowebserver.WorkloadResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WorkloadRPC", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WorkloadRPC indicates an expected call of WorkloadRPC
func (mr *MockAnalyserHandlerMockRecorder) WorkloadRPC(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WorkloadRPC", reflect.TypeOf((*MockAnalyserHandler)(nil).WorkloadRPC), arg0, arg1, arg2)
}

// ChangeStateRPC mocks base method
func (m *MockAnalyserHandler) ChangeStateRPC(arg0 context.Context, arg1 *analysertowebserver.ChangeStateRequest, arg2 *analysertowebserver.ChangeStateResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeStateRPC", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeStateRPC indicates an expected call of ChangeStateRPC
func (mr *MockAnalyserHandlerMockRecorder) ChangeStateRPC(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeStateRPC", reflect.TypeOf((*MockAnalyserHandler)(nil).ChangeStateRPC), arg0, arg1, arg2)
}

// StateRPC mocks base method
func (m *MockAnalyserHandler) StateRPC(arg0 context.Context, arg1 *analysertowebserver.StateRequest, arg2 *analysertowebserver.StateResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateRPC", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// StateRPC indicates an expected call of StateRPC
func (mr *MockAnalyserHandlerMockRecorder) StateRPC(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateRPC", reflect.TypeOf((*MockAnalyserHandler)(nil).StateRPC), arg0, arg1, arg2)
}

// WebsiteData mocks base method
func (m *MockAnalyserHandler) WebsiteData(arg0 context.Context, arg1 analysertowebserver.Analyser_WebsiteDataStream) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WebsiteData", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WebsiteData indicates an expected call of WebsiteData
func (mr *MockAnalyserHandlerMockRecorder) WebsiteData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WebsiteData", reflect.TypeOf((*MockAnalyserHandler)(nil).WebsiteData), arg0, arg1)
}

// MockAnalyser_WebsiteDataStream is a mock of Analyser_WebsiteDataStream interface
type MockAnalyser_WebsiteDataStream struct {
	ctrl     *gomock.Controller
	recorder *MockAnalyser_WebsiteDataStreamMockRecorder
}

// MockAnalyser_WebsiteDataStreamMockRecorder is the mock recorder for MockAnalyser_WebsiteDataStream
type MockAnalyser_WebsiteDataStreamMockRecorder struct {
	mock *MockAnalyser_WebsiteDataStream
}

// NewMockAnalyser_WebsiteDataStream creates a new mock instance
func NewMockAnalyser_WebsiteDataStream(ctrl *gomock.Controller) *MockAnalyser_WebsiteDataStream {
	mock := &MockAnalyser_WebsiteDataStream{ctrl: ctrl}
	mock.recorder = &MockAnalyser_WebsiteDataStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAnalyser_WebsiteDataStream) EXPECT() *MockAnalyser_WebsiteDataStreamMockRecorder {
	return m.recorder
}

// SendMsg mocks base method
func (m *MockAnalyser_WebsiteDataStream) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockAnalyser_WebsiteDataStreamMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockAnalyser_WebsiteDataStream)(nil).SendMsg), arg0)
}

// RecvMsg mocks base method
func (m *MockAnalyser_WebsiteDataStream) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockAnalyser_WebsiteDataStreamMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockAnalyser_WebsiteDataStream)(nil).RecvMsg), arg0)
}

// Close mocks base method
func (m *MockAnalyser_WebsiteDataStream) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockAnalyser_WebsiteDataStreamMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAnalyser_WebsiteDataStream)(nil).Close))
}

// Send mocks base method
func (m *MockAnalyser_WebsiteDataStream) Send(arg0 *analysertowebserver.CrawlerAck) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockAnalyser_WebsiteDataStreamMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockAnalyser_WebsiteDataStream)(nil).Send), arg0)
}

// Recv mocks base method
func (m *MockAnalyser_WebsiteDataStream) Recv() (*analysertowebserver.CrawlerPackage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*analysertowebserver.CrawlerPackage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockAnalyser_WebsiteDataStreamMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockAnalyser_WebsiteDataStream)(nil).Recv))
}
