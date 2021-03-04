// Code generated by MockGen. DO NOT EDIT.
// Source: azure.go

// Package mock_azure is a generated GoMock package.
package mock_azure

import (
	context "context"
	reflect "reflect"
	time "time"

	azqueue "github.com/Azure/azure-storage-queue-go/azqueue"
	gomock "github.com/golang/mock/gomock"
)

// MockAzure is a mock of Azure interface.
type MockAzure struct {
	ctrl     *gomock.Controller
	recorder *MockAzureMockRecorder
}

// MockAzureMockRecorder is the mock recorder for MockAzure.
type MockAzureMockRecorder struct {
	mock *MockAzure
}

// NewMockAzure creates a new mock instance.
func NewMockAzure(ctrl *gomock.Controller) *MockAzure {
	mock := &MockAzure{ctrl: ctrl}
	mock.recorder = &MockAzureMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAzure) EXPECT() *MockAzureMockRecorder {
	return m.recorder
}

// Enqueue mocks base method.
func (m *MockAzure) Enqueue() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Enqueue")
}

// Enqueue indicates an expected call of Enqueue.
func (mr *MockAzureMockRecorder) Enqueue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enqueue", reflect.TypeOf((*MockAzure)(nil).Enqueue))
}

// MockmessagesURLEnqueue is a mock of messagesURLEnqueue interface.
type MockmessagesURLEnqueue struct {
	ctrl     *gomock.Controller
	recorder *MockmessagesURLEnqueueMockRecorder
}

// MockmessagesURLEnqueueMockRecorder is the mock recorder for MockmessagesURLEnqueue.
type MockmessagesURLEnqueueMockRecorder struct {
	mock *MockmessagesURLEnqueue
}

// NewMockmessagesURLEnqueue creates a new mock instance.
func NewMockmessagesURLEnqueue(ctrl *gomock.Controller) *MockmessagesURLEnqueue {
	mock := &MockmessagesURLEnqueue{ctrl: ctrl}
	mock.recorder = &MockmessagesURLEnqueueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockmessagesURLEnqueue) EXPECT() *MockmessagesURLEnqueueMockRecorder {
	return m.recorder
}

// Enqueue mocks base method.
func (m *MockmessagesURLEnqueue) Enqueue(ctx context.Context, messageText string, visibilityTimeout, timeToLive time.Duration) (*azqueue.EnqueueMessageResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enqueue", ctx, messageText, visibilityTimeout, timeToLive)
	ret0, _ := ret[0].(*azqueue.EnqueueMessageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Enqueue indicates an expected call of Enqueue.
func (mr *MockmessagesURLEnqueueMockRecorder) Enqueue(ctx, messageText, visibilityTimeout, timeToLive interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enqueue", reflect.TypeOf((*MockmessagesURLEnqueue)(nil).Enqueue), ctx, messageText, visibilityTimeout, timeToLive)
}

// MockenqueueMessageResponse is a mock of enqueueMessageResponse interface.
type MockenqueueMessageResponse struct {
	ctrl     *gomock.Controller
	recorder *MockenqueueMessageResponseMockRecorder
}

// MockenqueueMessageResponseMockRecorder is the mock recorder for MockenqueueMessageResponse.
type MockenqueueMessageResponseMockRecorder struct {
	mock *MockenqueueMessageResponse
}

// NewMockenqueueMessageResponse creates a new mock instance.
func NewMockenqueueMessageResponse(ctrl *gomock.Controller) *MockenqueueMessageResponse {
	mock := &MockenqueueMessageResponse{ctrl: ctrl}
	mock.recorder = &MockenqueueMessageResponseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockenqueueMessageResponse) EXPECT() *MockenqueueMessageResponseMockRecorder {
	return m.recorder
}

// Response mocks base method.
func (m *MockenqueueMessageResponse) Response(arg0 *azqueue.EnqueueMessageResponse) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Response", arg0)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Response indicates an expected call of Response.
func (mr *MockenqueueMessageResponseMockRecorder) Response(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Response", reflect.TypeOf((*MockenqueueMessageResponse)(nil).Response), arg0)
}

// StatusCode mocks base method.
func (m *MockenqueueMessageResponse) StatusCode(arg0 *azqueue.EnqueueMessageResponse) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatusCode", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// StatusCode indicates an expected call of StatusCode.
func (mr *MockenqueueMessageResponseMockRecorder) StatusCode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusCode", reflect.TypeOf((*MockenqueueMessageResponse)(nil).StatusCode), arg0)
}