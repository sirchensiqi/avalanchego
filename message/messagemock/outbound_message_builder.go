// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ava-labs/avalanchego/message (interfaces: OutboundMsgBuilder)
//
// Generated by this command:
//
//	mockgen -package=messagemock -destination=message/messagemock/outbound_message_builder.go -mock_names=OutboundMsgBuilder=OutboundMsgBuilder github.com/ava-labs/avalanchego/message OutboundMsgBuilder
//

// Package messagemock is a generated GoMock package.
package messagemock

import (
	netip "net/netip"
	reflect "reflect"
	time "time"

	ids "github.com/ava-labs/avalanchego/ids"
	message "github.com/ava-labs/avalanchego/message"
	p2p "github.com/ava-labs/avalanchego/proto/pb/p2p"
	ips "github.com/ava-labs/avalanchego/utils/ips"
	gomock "go.uber.org/mock/gomock"
)

// OutboundMsgBuilder is a mock of OutboundMsgBuilder interface.
type OutboundMsgBuilder struct {
	ctrl     *gomock.Controller
	recorder *OutboundMsgBuilderMockRecorder
}

// OutboundMsgBuilderMockRecorder is the mock recorder for OutboundMsgBuilder.
type OutboundMsgBuilderMockRecorder struct {
	mock *OutboundMsgBuilder
}

// NewOutboundMsgBuilder creates a new mock instance.
func NewOutboundMsgBuilder(ctrl *gomock.Controller) *OutboundMsgBuilder {
	mock := &OutboundMsgBuilder{ctrl: ctrl}
	mock.recorder = &OutboundMsgBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *OutboundMsgBuilder) EXPECT() *OutboundMsgBuilderMockRecorder {
	return m.recorder
}

// Accepted mocks base method.
func (m *OutboundMsgBuilder) Accepted(arg0 ids.ID, arg1 uint32, arg2 []ids.ID) (message.OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accepted", arg0, arg1, arg2)
	ret0, _ := ret[0].(message.OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Accepted indicates an expected call of Accepted.
func (mr *OutboundMsgBuilderMockRecorder) Accepted(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accepted", reflect.TypeOf((*OutboundMsgBuilder)(nil).Accepted), arg0, arg1, arg2)
}

// AcceptedFrontier mocks base method.
func (m *OutboundMsgBuilder) AcceptedFrontier(arg0 ids.ID, arg1 uint32, arg2 ids.ID) (message.OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptedFrontier", arg0, arg1, arg2)
	ret0, _ := ret[0].(message.OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptedFrontier indicates an expected call of AcceptedFrontier.
func (mr *OutboundMsgBuilderMockRecorder) AcceptedFrontier(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptedFrontier", reflect.TypeOf((*OutboundMsgBuilder)(nil).AcceptedFrontier), arg0, arg1, arg2)
}

// AcceptedStateSummary mocks base method.
func (m *OutboundMsgBuilder) AcceptedStateSummary(arg0 ids.ID, arg1 uint32, arg2 []ids.ID) (message.OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptedStateSummary", arg0, arg1, arg2)
	ret0, _ := ret[0].(message.OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptedStateSummary indicates an expected call of AcceptedStateSummary.
func (mr *OutboundMsgBuilderMockRecorder) AcceptedStateSummary(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptedStateSummary", reflect.TypeOf((*OutboundMsgBuilder)(nil).AcceptedStateSummary), arg0, arg1, arg2)
}

// Ancestors mocks base method.
func (m *OutboundMsgBuilder) Ancestors(arg0 ids.ID, arg1 uint32, arg2 [][]byte) (message.OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ancestors", arg0, arg1, arg2)
	ret0, _ := ret[0].(message.OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ancestors indicates an expected call of Ancestors.
func (mr *OutboundMsgBuilderMockRecorder) Ancestors(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ancestors", reflect.TypeOf((*OutboundMsgBuilder)(nil).Ancestors), arg0, arg1, arg2)
}

// AppError mocks base method.
func (m *OutboundMsgBuilder) AppError(arg0 ids.ID, arg1 uint32, arg2 int32, arg3 string) (message.OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppError", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(message.OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppError indicates an expected call of AppError.
func (mr *OutboundMsgBuilderMockRecorder) AppError(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppError", reflect.TypeOf((*OutboundMsgBuilder)(nil).AppError), arg0, arg1, arg2, arg3)
}

// AppGossip mocks base method.
func (m *OutboundMsgBuilder) AppGossip(arg0 ids.ID, arg1 []byte) (message.OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppGossip", arg0, arg1)
	ret0, _ := ret[0].(message.OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppGossip indicates an expected call of AppGossip.
func (mr *OutboundMsgBuilderMockRecorder) AppGossip(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppGossip", reflect.TypeOf((*OutboundMsgBuilder)(nil).AppGossip), arg0, arg1)
}

// AppRequest mocks base method.
func (m *OutboundMsgBuilder) AppRequest(arg0 ids.ID, arg1 uint32, arg2 time.Duration, arg3 []byte) (message.OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppRequest", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(message.OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppRequest indicates an expected call of AppRequest.
func (mr *OutboundMsgBuilderMockRecorder) AppRequest(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppRequest", reflect.TypeOf((*OutboundMsgBuilder)(nil).AppRequest), arg0, arg1, arg2, arg3)
}

// AppResponse mocks base method.
func (m *OutboundMsgBuilder) AppResponse(arg0 ids.ID, arg1 uint32, arg2 []byte) (message.OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppResponse", arg0, arg1, arg2)
	ret0, _ := ret[0].(message.OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppResponse indicates an expected call of AppResponse.
func (mr *OutboundMsgBuilderMockRecorder) AppResponse(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppResponse", reflect.TypeOf((*OutboundMsgBuilder)(nil).AppResponse), arg0, arg1, arg2)
}

// Chits mocks base method.
func (m *OutboundMsgBuilder) Chits(arg0 ids.ID, arg1 uint32, arg2, arg3, arg4 ids.ID) (message.OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Chits", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(message.OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Chits indicates an expected call of Chits.
func (mr *OutboundMsgBuilderMockRecorder) Chits(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chits", reflect.TypeOf((*OutboundMsgBuilder)(nil).Chits), arg0, arg1, arg2, arg3, arg4)
}

// Get mocks base method.
func (m *OutboundMsgBuilder) Get(arg0 ids.ID, arg1 uint32, arg2 time.Duration, arg3 ids.ID) (message.OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(message.OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *OutboundMsgBuilderMockRecorder) Get(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*OutboundMsgBuilder)(nil).Get), arg0, arg1, arg2, arg3)
}

// GetAccepted mocks base method.
func (m *OutboundMsgBuilder) GetAccepted(arg0 ids.ID, arg1 uint32, arg2 time.Duration, arg3 []ids.ID) (message.OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccepted", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(message.OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccepted indicates an expected call of GetAccepted.
func (mr *OutboundMsgBuilderMockRecorder) GetAccepted(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccepted", reflect.TypeOf((*OutboundMsgBuilder)(nil).GetAccepted), arg0, arg1, arg2, arg3)
}

// GetAcceptedFrontier mocks base method.
func (m *OutboundMsgBuilder) GetAcceptedFrontier(arg0 ids.ID, arg1 uint32, arg2 time.Duration) (message.OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAcceptedFrontier", arg0, arg1, arg2)
	ret0, _ := ret[0].(message.OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAcceptedFrontier indicates an expected call of GetAcceptedFrontier.
func (mr *OutboundMsgBuilderMockRecorder) GetAcceptedFrontier(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAcceptedFrontier", reflect.TypeOf((*OutboundMsgBuilder)(nil).GetAcceptedFrontier), arg0, arg1, arg2)
}

// GetAcceptedStateSummary mocks base method.
func (m *OutboundMsgBuilder) GetAcceptedStateSummary(arg0 ids.ID, arg1 uint32, arg2 time.Duration, arg3 []uint64) (message.OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAcceptedStateSummary", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(message.OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAcceptedStateSummary indicates an expected call of GetAcceptedStateSummary.
func (mr *OutboundMsgBuilderMockRecorder) GetAcceptedStateSummary(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAcceptedStateSummary", reflect.TypeOf((*OutboundMsgBuilder)(nil).GetAcceptedStateSummary), arg0, arg1, arg2, arg3)
}

// GetAncestors mocks base method.
func (m *OutboundMsgBuilder) GetAncestors(arg0 ids.ID, arg1 uint32, arg2 time.Duration, arg3 ids.ID, arg4 p2p.EngineType) (message.OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAncestors", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(message.OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAncestors indicates an expected call of GetAncestors.
func (mr *OutboundMsgBuilderMockRecorder) GetAncestors(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAncestors", reflect.TypeOf((*OutboundMsgBuilder)(nil).GetAncestors), arg0, arg1, arg2, arg3, arg4)
}

// GetPeerList mocks base method.
func (m *OutboundMsgBuilder) GetPeerList(arg0, arg1 []byte, arg2 bool) (message.OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPeerList", arg0, arg1, arg2)
	ret0, _ := ret[0].(message.OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPeerList indicates an expected call of GetPeerList.
func (mr *OutboundMsgBuilderMockRecorder) GetPeerList(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeerList", reflect.TypeOf((*OutboundMsgBuilder)(nil).GetPeerList), arg0, arg1, arg2)
}

// GetStateSummaryFrontier mocks base method.
func (m *OutboundMsgBuilder) GetStateSummaryFrontier(arg0 ids.ID, arg1 uint32, arg2 time.Duration) (message.OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStateSummaryFrontier", arg0, arg1, arg2)
	ret0, _ := ret[0].(message.OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStateSummaryFrontier indicates an expected call of GetStateSummaryFrontier.
func (mr *OutboundMsgBuilderMockRecorder) GetStateSummaryFrontier(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStateSummaryFrontier", reflect.TypeOf((*OutboundMsgBuilder)(nil).GetStateSummaryFrontier), arg0, arg1, arg2)
}

// Handshake mocks base method.
func (m *OutboundMsgBuilder) Handshake(arg0 uint32, arg1 uint64, arg2 netip.AddrPort, arg3 string, arg4, arg5, arg6 uint32, arg7 uint64, arg8, arg9 []byte, arg10 []ids.ID, arg11, arg12 []uint32, arg13, arg14 []byte, arg15 bool) (message.OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handshake", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15)
	ret0, _ := ret[0].(message.OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Handshake indicates an expected call of Handshake.
func (mr *OutboundMsgBuilderMockRecorder) Handshake(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handshake", reflect.TypeOf((*OutboundMsgBuilder)(nil).Handshake), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15)
}

// PeerList mocks base method.
func (m *OutboundMsgBuilder) PeerList(arg0 []*ips.ClaimedIPPort, arg1 bool) (message.OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeerList", arg0, arg1)
	ret0, _ := ret[0].(message.OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PeerList indicates an expected call of PeerList.
func (mr *OutboundMsgBuilderMockRecorder) PeerList(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeerList", reflect.TypeOf((*OutboundMsgBuilder)(nil).PeerList), arg0, arg1)
}

// Ping mocks base method.
func (m *OutboundMsgBuilder) Ping(arg0 uint32, arg1 []*p2p.SubnetUptime) (message.OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", arg0, arg1)
	ret0, _ := ret[0].(message.OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ping indicates an expected call of Ping.
func (mr *OutboundMsgBuilderMockRecorder) Ping(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*OutboundMsgBuilder)(nil).Ping), arg0, arg1)
}

// Pong mocks base method.
func (m *OutboundMsgBuilder) Pong() (message.OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pong")
	ret0, _ := ret[0].(message.OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pong indicates an expected call of Pong.
func (mr *OutboundMsgBuilderMockRecorder) Pong() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pong", reflect.TypeOf((*OutboundMsgBuilder)(nil).Pong))
}

// PullQuery mocks base method.
func (m *OutboundMsgBuilder) PullQuery(arg0 ids.ID, arg1 uint32, arg2 time.Duration, arg3 ids.ID, arg4 uint64) (message.OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PullQuery", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(message.OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PullQuery indicates an expected call of PullQuery.
func (mr *OutboundMsgBuilderMockRecorder) PullQuery(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullQuery", reflect.TypeOf((*OutboundMsgBuilder)(nil).PullQuery), arg0, arg1, arg2, arg3, arg4)
}

// PushQuery mocks base method.
func (m *OutboundMsgBuilder) PushQuery(arg0 ids.ID, arg1 uint32, arg2 time.Duration, arg3 []byte, arg4 uint64) (message.OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushQuery", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(message.OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PushQuery indicates an expected call of PushQuery.
func (mr *OutboundMsgBuilderMockRecorder) PushQuery(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushQuery", reflect.TypeOf((*OutboundMsgBuilder)(nil).PushQuery), arg0, arg1, arg2, arg3, arg4)
}

// Put mocks base method.
func (m *OutboundMsgBuilder) Put(arg0 ids.ID, arg1 uint32, arg2 []byte) (message.OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", arg0, arg1, arg2)
	ret0, _ := ret[0].(message.OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put.
func (mr *OutboundMsgBuilderMockRecorder) Put(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*OutboundMsgBuilder)(nil).Put), arg0, arg1, arg2)
}

// StateSummaryFrontier mocks base method.
func (m *OutboundMsgBuilder) StateSummaryFrontier(arg0 ids.ID, arg1 uint32, arg2 []byte) (message.OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateSummaryFrontier", arg0, arg1, arg2)
	ret0, _ := ret[0].(message.OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateSummaryFrontier indicates an expected call of StateSummaryFrontier.
func (mr *OutboundMsgBuilderMockRecorder) StateSummaryFrontier(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateSummaryFrontier", reflect.TypeOf((*OutboundMsgBuilder)(nil).StateSummaryFrontier), arg0, arg1, arg2)
}