// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/services/services.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	core "github.com/arthurshafikov/banner-rotation/internal/core"
)

// MockBanners is a mock of Banners interface.
type MockBanners struct {
	ctrl     *gomock.Controller
	recorder *MockBannersMockRecorder
}

// MockBannersMockRecorder is the mock recorder for MockBanners.
type MockBannersMockRecorder struct {
	mock *MockBanners
}

// NewMockBanners creates a new mock instance.
func NewMockBanners(ctrl *gomock.Controller) *MockBanners {
	mock := &MockBanners{ctrl: ctrl}
	mock.recorder = &MockBannersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBanners) EXPECT() *MockBannersMockRecorder {
	return m.recorder
}

// AddBanner mocks base method.
func (m *MockBanners) AddBanner(ctx context.Context, description string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBanner", ctx, description)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddBanner indicates an expected call of AddBanner.
func (mr *MockBannersMockRecorder) AddBanner(ctx, description interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBanner", reflect.TypeOf((*MockBanners)(nil).AddBanner), ctx, description)
}

// DeleteBanner mocks base method.
func (m *MockBanners) DeleteBanner(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBanner", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBanner indicates an expected call of DeleteBanner.
func (mr *MockBannersMockRecorder) DeleteBanner(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBanner", reflect.TypeOf((*MockBanners)(nil).DeleteBanner), ctx, id)
}

// GetBanner mocks base method.
func (m *MockBanners) GetBanner(ctx context.Context, id int64) (*core.Banner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBanner", ctx, id)
	ret0, _ := ret[0].(*core.Banner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBanner indicates an expected call of GetBanner.
func (mr *MockBannersMockRecorder) GetBanner(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBanner", reflect.TypeOf((*MockBanners)(nil).GetBanner), ctx, id)
}

// MockSlots is a mock of Slots interface.
type MockSlots struct {
	ctrl     *gomock.Controller
	recorder *MockSlotsMockRecorder
}

// MockSlotsMockRecorder is the mock recorder for MockSlots.
type MockSlotsMockRecorder struct {
	mock *MockSlots
}

// NewMockSlots creates a new mock instance.
func NewMockSlots(ctrl *gomock.Controller) *MockSlots {
	mock := &MockSlots{ctrl: ctrl}
	mock.recorder = &MockSlotsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSlots) EXPECT() *MockSlotsMockRecorder {
	return m.recorder
}

// AddSlot mocks base method.
func (m *MockSlots) AddSlot(ctx context.Context, description string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSlot", ctx, description)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSlot indicates an expected call of AddSlot.
func (mr *MockSlotsMockRecorder) AddSlot(ctx, description interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSlot", reflect.TypeOf((*MockSlots)(nil).AddSlot), ctx, description)
}

// DeleteSlot mocks base method.
func (m *MockSlots) DeleteSlot(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSlot", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSlot indicates an expected call of DeleteSlot.
func (mr *MockSlotsMockRecorder) DeleteSlot(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSlot", reflect.TypeOf((*MockSlots)(nil).DeleteSlot), ctx, id)
}

// GetSlot mocks base method.
func (m *MockSlots) GetSlot(ctx context.Context, id int64) (*core.Slot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSlot", ctx, id)
	ret0, _ := ret[0].(*core.Slot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSlot indicates an expected call of GetSlot.
func (mr *MockSlotsMockRecorder) GetSlot(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSlot", reflect.TypeOf((*MockSlots)(nil).GetSlot), ctx, id)
}

// MockBannerSlots is a mock of BannerSlots interface.
type MockBannerSlots struct {
	ctrl     *gomock.Controller
	recorder *MockBannerSlotsMockRecorder
}

// MockBannerSlotsMockRecorder is the mock recorder for MockBannerSlots.
type MockBannerSlotsMockRecorder struct {
	mock *MockBannerSlots
}

// NewMockBannerSlots creates a new mock instance.
func NewMockBannerSlots(ctrl *gomock.Controller) *MockBannerSlots {
	mock := &MockBannerSlots{ctrl: ctrl}
	mock.recorder = &MockBannerSlotsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBannerSlots) EXPECT() *MockBannerSlotsMockRecorder {
	return m.recorder
}

// AssociateBannerToSlot mocks base method.
func (m *MockBannerSlots) AssociateBannerToSlot(ctx context.Context, bannerId, slotId int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssociateBannerToSlot", ctx, bannerId, slotId)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssociateBannerToSlot indicates an expected call of AssociateBannerToSlot.
func (mr *MockBannerSlotsMockRecorder) AssociateBannerToSlot(ctx, bannerId, slotId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssociateBannerToSlot", reflect.TypeOf((*MockBannerSlots)(nil).AssociateBannerToSlot), ctx, bannerId, slotId)
}

// DissociateBannerFromSlot mocks base method.
func (m *MockBannerSlots) DissociateBannerFromSlot(ctx context.Context, bannerId, slotId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DissociateBannerFromSlot", ctx, bannerId, slotId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DissociateBannerFromSlot indicates an expected call of DissociateBannerFromSlot.
func (mr *MockBannerSlotsMockRecorder) DissociateBannerFromSlot(ctx, bannerId, slotId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DissociateBannerFromSlot", reflect.TypeOf((*MockBannerSlots)(nil).DissociateBannerFromSlot), ctx, bannerId, slotId)
}

// GetByBannerAndSlotIds mocks base method.
func (m *MockBannerSlots) GetByBannerAndSlotIds(ctx context.Context, bannerId, slotId int64) (*core.BannerSlot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByBannerAndSlotIds", ctx, bannerId, slotId)
	ret0, _ := ret[0].(*core.BannerSlot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByBannerAndSlotIds indicates an expected call of GetByBannerAndSlotIds.
func (mr *MockBannerSlotsMockRecorder) GetByBannerAndSlotIds(ctx, bannerId, slotId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByBannerAndSlotIds", reflect.TypeOf((*MockBannerSlots)(nil).GetByBannerAndSlotIds), ctx, bannerId, slotId)
}

// GetRandomBannerIdExceptExcluded mocks base method.
func (m *MockBannerSlots) GetRandomBannerIdExceptExcluded(ctx context.Context, slotId, excludedBannerId int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRandomBannerIdExceptExcluded", ctx, slotId, excludedBannerId)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRandomBannerIdExceptExcluded indicates an expected call of GetRandomBannerIdExceptExcluded.
func (mr *MockBannerSlotsMockRecorder) GetRandomBannerIdExceptExcluded(ctx, slotId, excludedBannerId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRandomBannerIdExceptExcluded", reflect.TypeOf((*MockBannerSlots)(nil).GetRandomBannerIdExceptExcluded), ctx, slotId, excludedBannerId)
}

// MockSocialGroups is a mock of SocialGroups interface.
type MockSocialGroups struct {
	ctrl     *gomock.Controller
	recorder *MockSocialGroupsMockRecorder
}

// MockSocialGroupsMockRecorder is the mock recorder for MockSocialGroups.
type MockSocialGroupsMockRecorder struct {
	mock *MockSocialGroups
}

// NewMockSocialGroups creates a new mock instance.
func NewMockSocialGroups(ctrl *gomock.Controller) *MockSocialGroups {
	mock := &MockSocialGroups{ctrl: ctrl}
	mock.recorder = &MockSocialGroupsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSocialGroups) EXPECT() *MockSocialGroupsMockRecorder {
	return m.recorder
}

// AddSocialGroup mocks base method.
func (m *MockSocialGroups) AddSocialGroup(ctx context.Context, description string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSocialGroup", ctx, description)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSocialGroup indicates an expected call of AddSocialGroup.
func (mr *MockSocialGroupsMockRecorder) AddSocialGroup(ctx, description interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSocialGroup", reflect.TypeOf((*MockSocialGroups)(nil).AddSocialGroup), ctx, description)
}

// DeleteSocialGroup mocks base method.
func (m *MockSocialGroups) DeleteSocialGroup(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSocialGroup", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSocialGroup indicates an expected call of DeleteSocialGroup.
func (mr *MockSocialGroupsMockRecorder) DeleteSocialGroup(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSocialGroup", reflect.TypeOf((*MockSocialGroups)(nil).DeleteSocialGroup), ctx, id)
}

// GetSocialGroup mocks base method.
func (m *MockSocialGroups) GetSocialGroup(ctx context.Context, id int64) (*core.SocialGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSocialGroup", ctx, id)
	ret0, _ := ret[0].(*core.SocialGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSocialGroup indicates an expected call of GetSocialGroup.
func (mr *MockSocialGroupsMockRecorder) GetSocialGroup(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSocialGroup", reflect.TypeOf((*MockSocialGroups)(nil).GetSocialGroup), ctx, id)
}

// MockBannerSlotSocialGroups is a mock of BannerSlotSocialGroups interface.
type MockBannerSlotSocialGroups struct {
	ctrl     *gomock.Controller
	recorder *MockBannerSlotSocialGroupsMockRecorder
}

// MockBannerSlotSocialGroupsMockRecorder is the mock recorder for MockBannerSlotSocialGroups.
type MockBannerSlotSocialGroupsMockRecorder struct {
	mock *MockBannerSlotSocialGroups
}

// NewMockBannerSlotSocialGroups creates a new mock instance.
func NewMockBannerSlotSocialGroups(ctrl *gomock.Controller) *MockBannerSlotSocialGroups {
	mock := &MockBannerSlotSocialGroups{ctrl: ctrl}
	mock.recorder = &MockBannerSlotSocialGroupsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBannerSlotSocialGroups) EXPECT() *MockBannerSlotSocialGroupsMockRecorder {
	return m.recorder
}

// GetBannerIdToShow mocks base method.
func (m *MockBannerSlotSocialGroups) GetBannerIdToShow(ctx context.Context, inp core.GetBannerRequest) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBannerIdToShow", ctx, inp)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBannerIdToShow indicates an expected call of GetBannerIdToShow.
func (mr *MockBannerSlotSocialGroupsMockRecorder) GetBannerIdToShow(ctx, inp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBannerIdToShow", reflect.TypeOf((*MockBannerSlotSocialGroups)(nil).GetBannerIdToShow), ctx, inp)
}

// IncrementClick mocks base method.
func (m *MockBannerSlotSocialGroups) IncrementClick(ctx context.Context, inp core.IncrementClickInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrementClick", ctx, inp)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncrementClick indicates an expected call of IncrementClick.
func (mr *MockBannerSlotSocialGroupsMockRecorder) IncrementClick(ctx, inp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrementClick", reflect.TypeOf((*MockBannerSlotSocialGroups)(nil).IncrementClick), ctx, inp)
}

// MockQueue is a mock of Queue interface.
type MockQueue struct {
	ctrl     *gomock.Controller
	recorder *MockQueueMockRecorder
}

// MockQueueMockRecorder is the mock recorder for MockQueue.
type MockQueueMockRecorder struct {
	mock *MockQueue
}

// NewMockQueue creates a new mock instance.
func NewMockQueue(ctrl *gomock.Controller) *MockQueue {
	mock := &MockQueue{ctrl: ctrl}
	mock.recorder = &MockQueueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueue) EXPECT() *MockQueueMockRecorder {
	return m.recorder
}

// AddToQueue mocks base method.
func (m *MockQueue) AddToQueue(topic string, value interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToQueue", topic, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToQueue indicates an expected call of AddToQueue.
func (mr *MockQueueMockRecorder) AddToQueue(topic, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToQueue", reflect.TypeOf((*MockQueue)(nil).AddToQueue), topic, value)
}
