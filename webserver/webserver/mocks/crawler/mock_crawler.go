// Code generated by MockGen. DO NOT EDIT.
// Source: /home/markus/Projects/SherlockGopher/sherlockcrawler/proto/crawlerproto/crawlerproto.pb.micro.go

// Package mock_crawlerwebserverproto is a generated GoMock package.
package crawler

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	client "github.com/micro/go-micro/client"
	crawlerwebserverproto "github.com/ob-algdatii-20ss/SherlockGopher/sherlockcrawler/proto"
)

// MockCrawlerService is a mock of CrawlerService interface
type MockCrawlerService struct {
	ctrl     *gomock.Controller
	recorder *MockCrawlerServiceMockRecorder
}

// MockCrawlerServiceMockRecorder is the mock recorder for MockCrawlerService
type MockCrawlerServiceMockRecorder struct {
	mock *MockCrawlerService
}

// NewMockCrawlerService creates a new mock instance
func NewMockCrawlerService(ctrl *gomock.Controller) *MockCrawlerService {
	mock := &MockCrawlerService{ctrl: ctrl}
	mock.recorder = &MockCrawlerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCrawlerService) EXPECT() *MockCrawlerServiceMockRecorder {
	return m.recorder
}

// ReceiveURL mocks base method
func (m *MockCrawlerService) ReceiveURL(ctx context.Context, in *crawlerwebserverproto.SubmitURLRequest, opts ...client.CallOption) (*crawlerwebserverproto.SubmitURLResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReceiveURL", varargs...)
	ret0, _ := ret[0].(*crawlerwebserverproto.SubmitURLResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReceiveURL indicates an expected call of ReceiveURL
func (mr *MockCrawlerServiceMockRecorder) ReceiveURL(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiveURL", reflect.TypeOf((*MockCrawlerService)(nil).ReceiveURL), varargs...)
}

// StatusOfTaskQueue mocks base method
func (m *MockCrawlerService) StatusOfTaskQueue(ctx context.Context, in *crawlerwebserverproto.TaskStatusRequest, opts ...client.CallOption) (*crawlerwebserverproto.TaskStatusResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StatusOfTaskQueue", varargs...)
	ret0, _ := ret[0].(*crawlerwebserverproto.TaskStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatusOfTaskQueue indicates an expected call of StatusOfTaskQueue
func (mr *MockCrawlerServiceMockRecorder) StatusOfTaskQueue(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusOfTaskQueue", reflect.TypeOf((*MockCrawlerService)(nil).StatusOfTaskQueue), varargs...)
}

// SetState mocks base method
func (m *MockCrawlerService) SetState(ctx context.Context, in *crawlerwebserverproto.StateRequest, opts ...client.CallOption) (*crawlerwebserverproto.StateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetState", varargs...)
	ret0, _ := ret[0].(*crawlerwebserverproto.StateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetState indicates an expected call of SetState
func (mr *MockCrawlerServiceMockRecorder) SetState(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetState", reflect.TypeOf((*MockCrawlerService)(nil).SetState), varargs...)
}

// GetState mocks base method
func (m *MockCrawlerService) GetState(ctx context.Context, in *crawlerwebserverproto.StateGetRequest, opts ...client.CallOption) (*crawlerwebserverproto.StateGetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetState", varargs...)
	ret0, _ := ret[0].(*crawlerwebserverproto.StateGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetState indicates an expected call of GetState
func (mr *MockCrawlerServiceMockRecorder) GetState(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetState", reflect.TypeOf((*MockCrawlerService)(nil).GetState), varargs...)
}

// CreateTask mocks base method
func (m *MockCrawlerService) CreateTask(ctx context.Context, in *crawlerwebserverproto.CrawlTaskCreateRequest, opts ...client.CallOption) (*crawlerwebserverproto.CrawlTaskCreateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTask", varargs...)
	ret0, _ := ret[0].(*crawlerwebserverproto.CrawlTaskCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTask indicates an expected call of CreateTask
func (mr *MockCrawlerServiceMockRecorder) CreateTask(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockCrawlerService)(nil).CreateTask), varargs...)
}

// MockCrawlerHandler is a mock of CrawlerHandler interface
type MockCrawlerHandler struct {
	ctrl     *gomock.Controller
	recorder *MockCrawlerHandlerMockRecorder
}

// MockCrawlerHandlerMockRecorder is the mock recorder for MockCrawlerHandler
type MockCrawlerHandlerMockRecorder struct {
	mock *MockCrawlerHandler
}

// NewMockCrawlerHandler creates a new mock instance
func NewMockCrawlerHandler(ctrl *gomock.Controller) *MockCrawlerHandler {
	mock := &MockCrawlerHandler{ctrl: ctrl}
	mock.recorder = &MockCrawlerHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCrawlerHandler) EXPECT() *MockCrawlerHandlerMockRecorder {
	return m.recorder
}

// ReceiveURL mocks base method
func (m *MockCrawlerHandler) ReceiveURL(arg0 context.Context, arg1 *crawlerwebserverproto.SubmitURLRequest, arg2 *crawlerwebserverproto.SubmitURLResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceiveURL", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReceiveURL indicates an expected call of ReceiveURL
func (mr *MockCrawlerHandlerMockRecorder) ReceiveURL(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiveURL", reflect.TypeOf((*MockCrawlerHandler)(nil).ReceiveURL), arg0, arg1, arg2)
}

// StatusOfTaskQueue mocks base method
func (m *MockCrawlerHandler) StatusOfTaskQueue(arg0 context.Context, arg1 *crawlerwebserverproto.TaskStatusRequest, arg2 *crawlerwebserverproto.TaskStatusResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatusOfTaskQueue", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// StatusOfTaskQueue indicates an expected call of StatusOfTaskQueue
func (mr *MockCrawlerHandlerMockRecorder) StatusOfTaskQueue(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusOfTaskQueue", reflect.TypeOf((*MockCrawlerHandler)(nil).StatusOfTaskQueue), arg0, arg1, arg2)
}

// SetState mocks base method
func (m *MockCrawlerHandler) SetState(arg0 context.Context, arg1 *crawlerwebserverproto.StateRequest, arg2 *crawlerwebserverproto.StateResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetState", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetState indicates an expected call of SetState
func (mr *MockCrawlerHandlerMockRecorder) SetState(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetState", reflect.TypeOf((*MockCrawlerHandler)(nil).SetState), arg0, arg1, arg2)
}

// GetState mocks base method
func (m *MockCrawlerHandler) GetState(arg0 context.Context, arg1 *crawlerwebserverproto.StateGetRequest, arg2 *crawlerwebserverproto.StateGetResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetState", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetState indicates an expected call of GetState
func (mr *MockCrawlerHandlerMockRecorder) GetState(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetState", reflect.TypeOf((*MockCrawlerHandler)(nil).GetState), arg0, arg1, arg2)
}

// CreateTask mocks base method
func (m *MockCrawlerHandler) CreateTask(arg0 context.Context, arg1 *crawlerwebserverproto.CrawlTaskCreateRequest, arg2 *crawlerwebserverproto.CrawlTaskCreateResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTask indicates an expected call of CreateTask
func (mr *MockCrawlerHandlerMockRecorder) CreateTask(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockCrawlerHandler)(nil).CreateTask), arg0, arg1, arg2)
}
