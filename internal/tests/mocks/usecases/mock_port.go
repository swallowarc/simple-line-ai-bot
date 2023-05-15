// Code generated by MockGen. DO NOT EDIT.
// Source: port.go

// Package mock_usecases is a generated GoMock package.
package mock_usecases

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/swallowarc/simple-line-ai-bot/internal/domain"
)

// MockChatRepository is a mock of ChatRepository interface.
type MockChatRepository struct {
	ctrl     *gomock.Controller
	recorder *MockChatRepositoryMockRecorder
}

// MockChatRepositoryMockRecorder is the mock recorder for MockChatRepository.
type MockChatRepositoryMockRecorder struct {
	mock *MockChatRepository
}

// NewMockChatRepository creates a new mock instance.
func NewMockChatRepository(ctrl *gomock.Controller) *MockChatRepository {
	mock := &MockChatRepository{ctrl: ctrl}
	mock.recorder = &MockChatRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatRepository) EXPECT() *MockChatRepositoryMockRecorder {
	return m.recorder
}

// Chat mocks base method.
func (m *MockChatRepository) Chat(ctx context.Context, messages domain.ChatMessages) (domain.ChatMessages, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Chat", ctx, messages)
	ret0, _ := ret[0].(domain.ChatMessages)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Chat indicates an expected call of Chat.
func (mr *MockChatRepositoryMockRecorder) Chat(ctx, messages interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chat", reflect.TypeOf((*MockChatRepository)(nil).Chat), ctx, messages)
}

// DeleteCacheMessages mocks base method.
func (m *MockChatRepository) DeleteCacheMessages(ctx context.Context, es domain.EventSource) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCacheMessages", ctx, es)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCacheMessages indicates an expected call of DeleteCacheMessages.
func (mr *MockChatRepositoryMockRecorder) DeleteCacheMessages(ctx, es interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCacheMessages", reflect.TypeOf((*MockChatRepository)(nil).DeleteCacheMessages), ctx, es)
}

// ListCacheMessages mocks base method.
func (m *MockChatRepository) ListCacheMessages(ctx context.Context, es domain.EventSource) (domain.ChatMessages, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCacheMessages", ctx, es)
	ret0, _ := ret[0].(domain.ChatMessages)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCacheMessages indicates an expected call of ListCacheMessages.
func (mr *MockChatRepositoryMockRecorder) ListCacheMessages(ctx, es interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCacheMessages", reflect.TypeOf((*MockChatRepository)(nil).ListCacheMessages), ctx, es)
}

// UpsertCacheMessages mocks base method.
func (m *MockChatRepository) UpsertCacheMessages(ctx context.Context, es domain.EventSource, cms domain.ChatMessages) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertCacheMessages", ctx, es, cms)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertCacheMessages indicates an expected call of UpsertCacheMessages.
func (mr *MockChatRepositoryMockRecorder) UpsertCacheMessages(ctx, es, cms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertCacheMessages", reflect.TypeOf((*MockChatRepository)(nil).UpsertCacheMessages), ctx, es, cms)
}

// MockMessagingRepository is a mock of MessagingRepository interface.
type MockMessagingRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMessagingRepositoryMockRecorder
}

// MockMessagingRepositoryMockRecorder is the mock recorder for MockMessagingRepository.
type MockMessagingRepositoryMockRecorder struct {
	mock *MockMessagingRepository
}

// NewMockMessagingRepository creates a new mock instance.
func NewMockMessagingRepository(ctrl *gomock.Controller) *MockMessagingRepository {
	mock := &MockMessagingRepository{ctrl: ctrl}
	mock.recorder = &MockMessagingRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessagingRepository) EXPECT() *MockMessagingRepositoryMockRecorder {
	return m.recorder
}

// GetGroupName mocks base method.
func (m *MockMessagingRepository) GetGroupName(arg0 context.Context, groupID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupName", arg0, groupID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupName indicates an expected call of GetGroupName.
func (mr *MockMessagingRepositoryMockRecorder) GetGroupName(arg0, groupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupName", reflect.TypeOf((*MockMessagingRepository)(nil).GetGroupName), arg0, groupID)
}

// GetUserName mocks base method.
func (m *MockMessagingRepository) GetUserName(arg0 context.Context, userID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserName", arg0, userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserName indicates an expected call of GetUserName.
func (mr *MockMessagingRepositoryMockRecorder) GetUserName(arg0, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserName", reflect.TypeOf((*MockMessagingRepository)(nil).GetUserName), arg0, userID)
}

// ListRoomMemberNames mocks base method.
func (m *MockMessagingRepository) ListRoomMemberNames(arg0 context.Context, roomID string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoomMemberNames", arg0, roomID)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoomMemberNames indicates an expected call of ListRoomMemberNames.
func (mr *MockMessagingRepositoryMockRecorder) ListRoomMemberNames(arg0, roomID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoomMemberNames", reflect.TypeOf((*MockMessagingRepository)(nil).ListRoomMemberNames), arg0, roomID)
}

// PushMessage mocks base method.
func (m *MockMessagingRepository) PushMessage(ctx context.Context, eventSourceID, message string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushMessage", ctx, eventSourceID, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// PushMessage indicates an expected call of PushMessage.
func (mr *MockMessagingRepositoryMockRecorder) PushMessage(ctx, eventSourceID, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushMessage", reflect.TypeOf((*MockMessagingRepository)(nil).PushMessage), ctx, eventSourceID, message)
}

// ReplyMessage mocks base method.
func (m *MockMessagingRepository) ReplyMessage(ctx context.Context, replyToken, message string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplyMessage", ctx, replyToken, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReplyMessage indicates an expected call of ReplyMessage.
func (mr *MockMessagingRepositoryMockRecorder) ReplyMessage(ctx, replyToken, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplyMessage", reflect.TypeOf((*MockMessagingRepository)(nil).ReplyMessage), ctx, replyToken, message)
}

// MockLicenseRepository is a mock of LicenseRepository interface.
type MockLicenseRepository struct {
	ctrl     *gomock.Controller
	recorder *MockLicenseRepositoryMockRecorder
}

// MockLicenseRepositoryMockRecorder is the mock recorder for MockLicenseRepository.
type MockLicenseRepositoryMockRecorder struct {
	mock *MockLicenseRepository
}

// NewMockLicenseRepository creates a new mock instance.
func NewMockLicenseRepository(ctrl *gomock.Controller) *MockLicenseRepository {
	mock := &MockLicenseRepository{ctrl: ctrl}
	mock.recorder = &MockLicenseRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLicenseRepository) EXPECT() *MockLicenseRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockLicenseRepository) Delete(ctx context.Context, es domain.EventSource) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, es)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockLicenseRepositoryMockRecorder) Delete(ctx, es interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockLicenseRepository)(nil).Delete), ctx, es)
}

// Get mocks base method.
func (m *MockLicenseRepository) Get(ctx context.Context, es domain.EventSource) (domain.License, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, es)
	ret0, _ := ret[0].(domain.License)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockLicenseRepositoryMockRecorder) Get(ctx, es interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockLicenseRepository)(nil).Get), ctx, es)
}

// Update mocks base method.
func (m *MockLicenseRepository) Update(ctx context.Context, lc domain.License, lt time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, lc, lt)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockLicenseRepositoryMockRecorder) Update(ctx, lc, lt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockLicenseRepository)(nil).Update), ctx, lc, lt)
}

// Upsert mocks base method.
func (m *MockLicenseRepository) Upsert(ctx context.Context, lc domain.License, lt time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", ctx, lc, lt)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockLicenseRepositoryMockRecorder) Upsert(ctx, lc, lt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockLicenseRepository)(nil).Upsert), ctx, lc, lt)
}
