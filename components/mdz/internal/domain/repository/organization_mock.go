// Code generated by MockGen. DO NOT EDIT.
// Source: /home/max/Workspace/midaz/components/mdz/internal/domain/repository/organization.go
//
// Generated by this command:
//
//	mockgen -source=/home/max/Workspace/midaz/components/mdz/internal/domain/repository/organization.go -destination=/home/max/Workspace/midaz/components/mdz/internal/domain/repository/organization_mock.go -package repository
//

// Package repository is a generated GoMock package.
package repository

import (
	reflect "reflect"

	mmodel "github.com/LerianStudio/midaz/pkg/mmodel"
	gomock "go.uber.org/mock/gomock"
)

// MockOrganization is a mock of Organization interface.
type MockOrganization struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationMockRecorder
	isgomock struct{}
}

// MockOrganizationMockRecorder is the mock recorder for MockOrganization.
type MockOrganizationMockRecorder struct {
	mock *MockOrganization
}

// NewMockOrganization creates a new mock instance.
func NewMockOrganization(ctrl *gomock.Controller) *MockOrganization {
	mock := &MockOrganization{ctrl: ctrl}
	mock.recorder = &MockOrganizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganization) EXPECT() *MockOrganizationMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockOrganization) Create(org mmodel.CreateOrganizationInput) (*mmodel.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", org)
	ret0, _ := ret[0].(*mmodel.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockOrganizationMockRecorder) Create(org any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrganization)(nil).Create), org)
}

// Delete mocks base method.
func (m *MockOrganization) Delete(organizationID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", organizationID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockOrganizationMockRecorder) Delete(organizationID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOrganization)(nil).Delete), organizationID)
}

// Get mocks base method.
func (m *MockOrganization) Get(limit, page int, SortOrder, StartDate, EndDate string) (*mmodel.Organizations, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", limit, page, SortOrder, StartDate, EndDate)
	ret0, _ := ret[0].(*mmodel.Organizations)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockOrganizationMockRecorder) Get(limit, page, SortOrder, StartDate, EndDate any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockOrganization)(nil).Get), limit, page, SortOrder, StartDate, EndDate)
}

// GetByID mocks base method.
func (m *MockOrganization) GetByID(organizationID string) (*mmodel.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", organizationID)
	ret0, _ := ret[0].(*mmodel.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockOrganizationMockRecorder) GetByID(organizationID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockOrganization)(nil).GetByID), organizationID)
}

// Update mocks base method.
func (m *MockOrganization) Update(organizationID string, orgInput mmodel.UpdateOrganizationInput) (*mmodel.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", organizationID, orgInput)
	ret0, _ := ret[0].(*mmodel.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockOrganizationMockRecorder) Update(organizationID, orgInput any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOrganization)(nil).Update), organizationID, orgInput)
}