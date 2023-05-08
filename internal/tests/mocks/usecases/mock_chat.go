// Code generated by MockGen. DO NOT EDIT.
// Source: chat.go

// Package mock_usecases is a generated GoMock package.
package mock_usecases

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/swallowarc/simple-line-ai-bot/internal/domain"
	usecases "github.com/swallowarc/simple-line-ai-bot/internal/usecases"
)

// MockChat is a mock of Chat interface.
type MockChat struct {
	ctrl     *gomock.Controller
	recorder *MockChatMockRecorder
}

// MockChatMockRecorder is the mock recorder for MockChat.
type MockChatMockRecorder struct {
	mock *MockChat
}

// NewMockChat creates a new mock instance.
func NewMockChat(ctrl *gomock.Controller) *MockChat {
	mock := &MockChat{ctrl: ctrl}
	mock.recorder = &MockChatMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChat) EXPECT() *MockChatMockRecorder {
	return m.recorder
}

// Chat mocks base method.
func (m *MockChat) Chat(ctx context.Context, es domain.EventSource, req string, callback usecases.Callback) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Chat", ctx, es, req, callback)
	ret0, _ := ret[0].(error)
	return ret0
}

// Chat indicates an expected call of Chat.
func (mr *MockChatMockRecorder) Chat(ctx, es, req, callback interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chat", reflect.TypeOf((*MockChat)(nil).Chat), ctx, es, req, callback)
}

// ClearChatHistory mocks base method.
func (m *MockChat) ClearChatHistory(ctx context.Context, es domain.EventSource, callback usecases.Callback) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearChatHistory", ctx, es, callback)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearChatHistory indicates an expected call of ClearChatHistory.
func (mr *MockChatMockRecorder) ClearChatHistory(ctx, es, callback interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearChatHistory", reflect.TypeOf((*MockChat)(nil).ClearChatHistory), ctx, es, callback)
}

// Help mocks base method.
func (m *MockChat) Help(ctx context.Context, callback usecases.Callback) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Help", ctx, callback)
	ret0, _ := ret[0].(error)
	return ret0
}

// Help indicates an expected call of Help.
func (mr *MockChatMockRecorder) Help(ctx, callback interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Help", reflect.TypeOf((*MockChat)(nil).Help), ctx, callback)
}
