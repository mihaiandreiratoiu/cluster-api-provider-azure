/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: ../resourcehealth.go

// Package mock_resourcehealth is a generated GoMock package.
package mock_resourcehealth

import (
	reflect "reflect"

	azcore "github.com/Azure/azure-sdk-for-go/sdk/azcore"
	autorest "github.com/Azure/go-autorest/autorest"
	gomock "go.uber.org/mock/gomock"
	v1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
	conditions "sigs.k8s.io/cluster-api/util/conditions"
)

// MockResourceHealthScope is a mock of ResourceHealthScope interface.
type MockResourceHealthScope struct {
	ctrl     *gomock.Controller
	recorder *MockResourceHealthScopeMockRecorder
}

// MockResourceHealthScopeMockRecorder is the mock recorder for MockResourceHealthScope.
type MockResourceHealthScopeMockRecorder struct {
	mock *MockResourceHealthScope
}

// NewMockResourceHealthScope creates a new mock instance.
func NewMockResourceHealthScope(ctrl *gomock.Controller) *MockResourceHealthScope {
	mock := &MockResourceHealthScope{ctrl: ctrl}
	mock.recorder = &MockResourceHealthScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceHealthScope) EXPECT() *MockResourceHealthScopeMockRecorder {
	return m.recorder
}

// Authorizer mocks base method.
func (m *MockResourceHealthScope) Authorizer() autorest.Authorizer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorizer")
	ret0, _ := ret[0].(autorest.Authorizer)
	return ret0
}

// Authorizer indicates an expected call of Authorizer.
func (mr *MockResourceHealthScopeMockRecorder) Authorizer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorizer", reflect.TypeOf((*MockResourceHealthScope)(nil).Authorizer))
}

// AvailabilityStatusResource mocks base method.
func (m *MockResourceHealthScope) AvailabilityStatusResource() conditions.Setter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailabilityStatusResource")
	ret0, _ := ret[0].(conditions.Setter)
	return ret0
}

// AvailabilityStatusResource indicates an expected call of AvailabilityStatusResource.
func (mr *MockResourceHealthScopeMockRecorder) AvailabilityStatusResource() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailabilityStatusResource", reflect.TypeOf((*MockResourceHealthScope)(nil).AvailabilityStatusResource))
}

// AvailabilityStatusResourceURI mocks base method.
func (m *MockResourceHealthScope) AvailabilityStatusResourceURI() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailabilityStatusResourceURI")
	ret0, _ := ret[0].(string)
	return ret0
}

// AvailabilityStatusResourceURI indicates an expected call of AvailabilityStatusResourceURI.
func (mr *MockResourceHealthScopeMockRecorder) AvailabilityStatusResourceURI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailabilityStatusResourceURI", reflect.TypeOf((*MockResourceHealthScope)(nil).AvailabilityStatusResourceURI))
}

// BaseURI mocks base method.
func (m *MockResourceHealthScope) BaseURI() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseURI")
	ret0, _ := ret[0].(string)
	return ret0
}

// BaseURI indicates an expected call of BaseURI.
func (mr *MockResourceHealthScopeMockRecorder) BaseURI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseURI", reflect.TypeOf((*MockResourceHealthScope)(nil).BaseURI))
}

// ClientID mocks base method.
func (m *MockResourceHealthScope) ClientID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientID indicates an expected call of ClientID.
func (mr *MockResourceHealthScopeMockRecorder) ClientID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientID", reflect.TypeOf((*MockResourceHealthScope)(nil).ClientID))
}

// ClientSecret mocks base method.
func (m *MockResourceHealthScope) ClientSecret() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientSecret")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientSecret indicates an expected call of ClientSecret.
func (mr *MockResourceHealthScopeMockRecorder) ClientSecret() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientSecret", reflect.TypeOf((*MockResourceHealthScope)(nil).ClientSecret))
}

// CloudEnvironment mocks base method.
func (m *MockResourceHealthScope) CloudEnvironment() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudEnvironment")
	ret0, _ := ret[0].(string)
	return ret0
}

// CloudEnvironment indicates an expected call of CloudEnvironment.
func (mr *MockResourceHealthScopeMockRecorder) CloudEnvironment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudEnvironment", reflect.TypeOf((*MockResourceHealthScope)(nil).CloudEnvironment))
}

// HashKey mocks base method.
func (m *MockResourceHealthScope) HashKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// HashKey indicates an expected call of HashKey.
func (mr *MockResourceHealthScopeMockRecorder) HashKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashKey", reflect.TypeOf((*MockResourceHealthScope)(nil).HashKey))
}

// SubscriptionID mocks base method.
func (m *MockResourceHealthScope) SubscriptionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SubscriptionID indicates an expected call of SubscriptionID.
func (mr *MockResourceHealthScopeMockRecorder) SubscriptionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionID", reflect.TypeOf((*MockResourceHealthScope)(nil).SubscriptionID))
}

// TenantID mocks base method.
func (m *MockResourceHealthScope) TenantID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantID")
	ret0, _ := ret[0].(string)
	return ret0
}

// TenantID indicates an expected call of TenantID.
func (mr *MockResourceHealthScopeMockRecorder) TenantID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantID", reflect.TypeOf((*MockResourceHealthScope)(nil).TenantID))
}

// Token mocks base method.
func (m *MockResourceHealthScope) Token() azcore.TokenCredential {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Token")
	ret0, _ := ret[0].(azcore.TokenCredential)
	return ret0
}

// Token indicates an expected call of Token.
func (mr *MockResourceHealthScopeMockRecorder) Token() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Token", reflect.TypeOf((*MockResourceHealthScope)(nil).Token))
}

// MockAvailabilityStatusFilterer is a mock of AvailabilityStatusFilterer interface.
type MockAvailabilityStatusFilterer struct {
	ctrl     *gomock.Controller
	recorder *MockAvailabilityStatusFiltererMockRecorder
}

// MockAvailabilityStatusFiltererMockRecorder is the mock recorder for MockAvailabilityStatusFilterer.
type MockAvailabilityStatusFiltererMockRecorder struct {
	mock *MockAvailabilityStatusFilterer
}

// NewMockAvailabilityStatusFilterer creates a new mock instance.
func NewMockAvailabilityStatusFilterer(ctrl *gomock.Controller) *MockAvailabilityStatusFilterer {
	mock := &MockAvailabilityStatusFilterer{ctrl: ctrl}
	mock.recorder = &MockAvailabilityStatusFiltererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAvailabilityStatusFilterer) EXPECT() *MockAvailabilityStatusFiltererMockRecorder {
	return m.recorder
}

// AvailabilityStatusFilter mocks base method.
func (m *MockAvailabilityStatusFilterer) AvailabilityStatusFilter(cond *v1beta1.Condition) *v1beta1.Condition {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailabilityStatusFilter", cond)
	ret0, _ := ret[0].(*v1beta1.Condition)
	return ret0
}

// AvailabilityStatusFilter indicates an expected call of AvailabilityStatusFilter.
func (mr *MockAvailabilityStatusFiltererMockRecorder) AvailabilityStatusFilter(cond interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailabilityStatusFilter", reflect.TypeOf((*MockAvailabilityStatusFilterer)(nil).AvailabilityStatusFilter), cond)
}
