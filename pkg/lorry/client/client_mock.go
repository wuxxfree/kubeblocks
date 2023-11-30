// /*
// Copyright (C) 2022-2023 ApeCloud Co., Ltd
//
// This file is part of KubeBlocks project
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
// */
//
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/apecloud/kubeblocks/pkg/lorry/client (interfaces: Client)

// Package client is a generated GoMock package.
package client

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockClient) CreateUser(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockClientMockRecorder) CreateUser(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockClient)(nil).CreateUser), arg0, arg1, arg2, arg3)
}

// DeleteUser mocks base method.
func (m *MockClient) DeleteUser(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockClientMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockClient)(nil).DeleteUser), arg0, arg1)
}

// DescribeUser mocks base method.
func (m *MockClient) DescribeUser(arg0 context.Context, arg1 string) (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeUser", arg0, arg1)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeUser indicates an expected call of DescribeUser.
func (mr *MockClientMockRecorder) DescribeUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeUser", reflect.TypeOf((*MockClient)(nil).DescribeUser), arg0, arg1)
}

// GetRole mocks base method.
func (m *MockClient) GetRole(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRole", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRole indicates an expected call of GetRole.
func (mr *MockClientMockRecorder) GetRole(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRole", reflect.TypeOf((*MockClient)(nil).GetRole), arg0)
}

// GrantUserRole mocks base method.
func (m *MockClient) GrantUserRole(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GrantUserRole", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GrantUserRole indicates an expected call of GrantUserRole.
func (mr *MockClientMockRecorder) GrantUserRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantUserRole", reflect.TypeOf((*MockClient)(nil).GrantUserRole), arg0, arg1, arg2)
}

// JoinMember mocks base method.
func (m *MockClient) JoinMember(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JoinMember", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// JoinMember indicates an expected call of JoinMember.
func (mr *MockClientMockRecorder) JoinMember(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JoinMember", reflect.TypeOf((*MockClient)(nil).JoinMember), arg0)
}

// LeaveMember mocks base method.
func (m *MockClient) LeaveMember(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LeaveMember", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// LeaveMember indicates an expected call of LeaveMember.
func (mr *MockClientMockRecorder) LeaveMember(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeaveMember", reflect.TypeOf((*MockClient)(nil).LeaveMember), arg0)
}

// ListSystemAccounts mocks base method.
func (m *MockClient) ListSystemAccounts(arg0 context.Context) ([]map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSystemAccounts", arg0)
	ret0, _ := ret[0].([]map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSystemAccounts indicates an expected call of ListSystemAccounts.
func (mr *MockClientMockRecorder) ListSystemAccounts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSystemAccounts", reflect.TypeOf((*MockClient)(nil).ListSystemAccounts), arg0)
}

// ListUsers mocks base method.
func (m *MockClient) ListUsers(arg0 context.Context) ([]map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", arg0)
	ret0, _ := ret[0].([]map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockClientMockRecorder) ListUsers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockClient)(nil).ListUsers), arg0)
}

// RevokeUserRole mocks base method.
func (m *MockClient) RevokeUserRole(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeUserRole", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeUserRole indicates an expected call of RevokeUserRole.
func (mr *MockClientMockRecorder) RevokeUserRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeUserRole", reflect.TypeOf((*MockClient)(nil).RevokeUserRole), arg0, arg1, arg2)
}
