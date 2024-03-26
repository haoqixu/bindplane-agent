// Code generated by MockGen. DO NOT EDIT.
// Source: ./exporter/chronicleexporter/protos/generated/ingestion_grpc.pb.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	generated "github.com/observiq/bindplane-agent/exporter/chronicleexporter/protos/api"
	grpc "google.golang.org/grpc"
)

// MockIngestionServiceV2Client is a mock of IngestionServiceV2Client interface.
type MockIngestionServiceV2Client struct {
	ctrl     *gomock.Controller
	recorder *MockIngestionServiceV2ClientMockRecorder
}

// MockIngestionServiceV2ClientMockRecorder is the mock recorder for MockIngestionServiceV2Client.
type MockIngestionServiceV2ClientMockRecorder struct {
	mock *MockIngestionServiceV2Client
}

// NewMockIngestionServiceV2Client creates a new mock instance.
func NewMockIngestionServiceV2Client(ctrl *gomock.Controller) *MockIngestionServiceV2Client {
	mock := &MockIngestionServiceV2Client{ctrl: ctrl}
	mock.recorder = &MockIngestionServiceV2ClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIngestionServiceV2Client) EXPECT() *MockIngestionServiceV2ClientMockRecorder {
	return m.recorder
}

// BatchCreateEvents mocks base method.
func (m *MockIngestionServiceV2Client) BatchCreateEvents(ctx context.Context, in *generated.BatchCreateEventsRequest, opts ...grpc.CallOption) (*generated.BatchCreateEventsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchCreateEvents", varargs...)
	ret0, _ := ret[0].(*generated.BatchCreateEventsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchCreateEvents indicates an expected call of BatchCreateEvents.
func (mr *MockIngestionServiceV2ClientMockRecorder) BatchCreateEvents(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchCreateEvents", reflect.TypeOf((*MockIngestionServiceV2Client)(nil).BatchCreateEvents), varargs...)
}

// BatchCreateLogs mocks base method.
func (m *MockIngestionServiceV2Client) BatchCreateLogs(ctx context.Context, in *generated.BatchCreateLogsRequest, opts ...grpc.CallOption) (*generated.BatchCreateLogsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchCreateLogs", varargs...)
	ret0, _ := ret[0].(*generated.BatchCreateLogsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchCreateLogs indicates an expected call of BatchCreateLogs.
func (mr *MockIngestionServiceV2ClientMockRecorder) BatchCreateLogs(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchCreateLogs", reflect.TypeOf((*MockIngestionServiceV2Client)(nil).BatchCreateLogs), varargs...)
}

// MockIngestionServiceV2Server is a mock of IngestionServiceV2Server interface.
type MockIngestionServiceV2Server struct {
	ctrl     *gomock.Controller
	recorder *MockIngestionServiceV2ServerMockRecorder
}

// MockIngestionServiceV2ServerMockRecorder is the mock recorder for MockIngestionServiceV2Server.
type MockIngestionServiceV2ServerMockRecorder struct {
	mock *MockIngestionServiceV2Server
}

// NewMockIngestionServiceV2Server creates a new mock instance.
func NewMockIngestionServiceV2Server(ctrl *gomock.Controller) *MockIngestionServiceV2Server {
	mock := &MockIngestionServiceV2Server{ctrl: ctrl}
	mock.recorder = &MockIngestionServiceV2ServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIngestionServiceV2Server) EXPECT() *MockIngestionServiceV2ServerMockRecorder {
	return m.recorder
}

// BatchCreateEvents mocks base method.
func (m *MockIngestionServiceV2Server) BatchCreateEvents(arg0 context.Context, arg1 *generated.BatchCreateEventsRequest) (*generated.BatchCreateEventsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchCreateEvents", arg0, arg1)
	ret0, _ := ret[0].(*generated.BatchCreateEventsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchCreateEvents indicates an expected call of BatchCreateEvents.
func (mr *MockIngestionServiceV2ServerMockRecorder) BatchCreateEvents(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchCreateEvents", reflect.TypeOf((*MockIngestionServiceV2Server)(nil).BatchCreateEvents), arg0, arg1)
}

// BatchCreateLogs mocks base method.
func (m *MockIngestionServiceV2Server) BatchCreateLogs(arg0 context.Context, arg1 *generated.BatchCreateLogsRequest) (*generated.BatchCreateLogsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchCreateLogs", arg0, arg1)
	ret0, _ := ret[0].(*generated.BatchCreateLogsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchCreateLogs indicates an expected call of BatchCreateLogs.
func (mr *MockIngestionServiceV2ServerMockRecorder) BatchCreateLogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchCreateLogs", reflect.TypeOf((*MockIngestionServiceV2Server)(nil).BatchCreateLogs), arg0, arg1)
}

// mustEmbedUnimplementedIngestionServiceV2Server mocks base method.
func (m *MockIngestionServiceV2Server) mustEmbedUnimplementedIngestionServiceV2Server() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedIngestionServiceV2Server")
}

// mustEmbedUnimplementedIngestionServiceV2Server indicates an expected call of mustEmbedUnimplementedIngestionServiceV2Server.
func (mr *MockIngestionServiceV2ServerMockRecorder) mustEmbedUnimplementedIngestionServiceV2Server() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedIngestionServiceV2Server", reflect.TypeOf((*MockIngestionServiceV2Server)(nil).mustEmbedUnimplementedIngestionServiceV2Server))
}

// MockUnsafeIngestionServiceV2Server is a mock of UnsafeIngestionServiceV2Server interface.
type MockUnsafeIngestionServiceV2Server struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeIngestionServiceV2ServerMockRecorder
}

// MockUnsafeIngestionServiceV2ServerMockRecorder is the mock recorder for MockUnsafeIngestionServiceV2Server.
type MockUnsafeIngestionServiceV2ServerMockRecorder struct {
	mock *MockUnsafeIngestionServiceV2Server
}

// NewMockUnsafeIngestionServiceV2Server creates a new mock instance.
func NewMockUnsafeIngestionServiceV2Server(ctrl *gomock.Controller) *MockUnsafeIngestionServiceV2Server {
	mock := &MockUnsafeIngestionServiceV2Server{ctrl: ctrl}
	mock.recorder = &MockUnsafeIngestionServiceV2ServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeIngestionServiceV2Server) EXPECT() *MockUnsafeIngestionServiceV2ServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedIngestionServiceV2Server mocks base method.
func (m *MockUnsafeIngestionServiceV2Server) mustEmbedUnimplementedIngestionServiceV2Server() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedIngestionServiceV2Server")
}

// mustEmbedUnimplementedIngestionServiceV2Server indicates an expected call of mustEmbedUnimplementedIngestionServiceV2Server.
func (mr *MockUnsafeIngestionServiceV2ServerMockRecorder) mustEmbedUnimplementedIngestionServiceV2Server() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedIngestionServiceV2Server", reflect.TypeOf((*MockUnsafeIngestionServiceV2Server)(nil).mustEmbedUnimplementedIngestionServiceV2Server))
}
