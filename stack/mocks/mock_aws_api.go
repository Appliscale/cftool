// Code generated by MockGen. DO NOT EDIT.
// Source: ./awsapi/cloudformation.go

// Package mocks is a generated GoMock package.
package mocks

import (
	cloudformation "github.com/aws/aws-sdk-go/service/cloudformation"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCloudFormationAPI is a mock of CloudFormationAPI interface
type MockCloudFormationAPI struct {
	ctrl     *gomock.Controller
	recorder *MockCloudFormationAPIMockRecorder
}

// MockCloudFormationAPIMockRecorder is the mock recorder for MockCloudFormationAPI
type MockCloudFormationAPIMockRecorder struct {
	mock *MockCloudFormationAPI
}

// NewMockCloudFormationAPI creates a new mock instance
func NewMockCloudFormationAPI(ctrl *gomock.Controller) *MockCloudFormationAPI {
	mock := &MockCloudFormationAPI{ctrl: ctrl}
	mock.recorder = &MockCloudFormationAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCloudFormationAPI) EXPECT() *MockCloudFormationAPIMockRecorder {
	return m.recorder
}

// CreateStack mocks base method
func (m *MockCloudFormationAPI) CreateStack(input *cloudformation.CreateStackInput) (*cloudformation.CreateStackOutput, error) {
	ret := m.ctrl.Call(m, "CreateStack", input)
	ret0, _ := ret[0].(*cloudformation.CreateStackOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStack indicates an expected call of CreateStack
func (mr *MockCloudFormationAPIMockRecorder) CreateStack(input interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStack", reflect.TypeOf((*MockCloudFormationAPI)(nil).CreateStack), input)
}

// DeleteStack mocks base method
func (m *MockCloudFormationAPI) DeleteStack(input *cloudformation.DeleteStackInput) (*cloudformation.DeleteStackOutput, error) {
	ret := m.ctrl.Call(m, "DeleteStack", input)
	ret0, _ := ret[0].(*cloudformation.DeleteStackOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteStack indicates an expected call of DeleteStack
func (mr *MockCloudFormationAPIMockRecorder) DeleteStack(input interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStack", reflect.TypeOf((*MockCloudFormationAPI)(nil).DeleteStack), input)
}

// UpdateStack mocks base method
func (m *MockCloudFormationAPI) UpdateStack(input *cloudformation.UpdateStackInput) (*cloudformation.UpdateStackOutput, error) {
	ret := m.ctrl.Call(m, "UpdateStack", input)
	ret0, _ := ret[0].(*cloudformation.UpdateStackOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStack indicates an expected call of UpdateStack
func (mr *MockCloudFormationAPIMockRecorder) UpdateStack(input interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStack", reflect.TypeOf((*MockCloudFormationAPI)(nil).UpdateStack), input)
}

// SetStackPolicy mocks base method
func (m *MockCloudFormationAPI) SetStackPolicy(input *cloudformation.SetStackPolicyInput) (*cloudformation.SetStackPolicyOutput, error) {
	ret := m.ctrl.Call(m, "SetStackPolicy", input)
	ret0, _ := ret[0].(*cloudformation.SetStackPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetStackPolicy indicates an expected call of SetStackPolicy
func (mr *MockCloudFormationAPIMockRecorder) SetStackPolicy(input interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStackPolicy", reflect.TypeOf((*MockCloudFormationAPI)(nil).SetStackPolicy), input)
}

// UpdateTerminationProtection mocks base method
func (m *MockCloudFormationAPI) UpdateTerminationProtection(input *cloudformation.UpdateTerminationProtectionInput) (*cloudformation.UpdateTerminationProtectionOutput, error) {
	ret := m.ctrl.Call(m, "UpdateTerminationProtection", input)
	ret0, _ := ret[0].(*cloudformation.UpdateTerminationProtectionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTerminationProtection indicates an expected call of UpdateTerminationProtection
func (mr *MockCloudFormationAPIMockRecorder) UpdateTerminationProtection(input interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTerminationProtection", reflect.TypeOf((*MockCloudFormationAPI)(nil).UpdateTerminationProtection), input)
}

// EstimateTemplateCost mocks base method
func (m *MockCloudFormationAPI) EstimateTemplateCost(input *cloudformation.EstimateTemplateCostInput) (*cloudformation.EstimateTemplateCostOutput, error) {
	ret := m.ctrl.Call(m, "EstimateTemplateCost", input)
	ret0, _ := ret[0].(*cloudformation.EstimateTemplateCostOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EstimateTemplateCost indicates an expected call of EstimateTemplateCost
func (mr *MockCloudFormationAPIMockRecorder) EstimateTemplateCost(input interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimateTemplateCost", reflect.TypeOf((*MockCloudFormationAPI)(nil).EstimateTemplateCost), input)
}

// ValidateTemplate mocks base method
func (m *MockCloudFormationAPI) ValidateTemplate(input *cloudformation.ValidateTemplateInput) (*cloudformation.ValidateTemplateOutput, error) {
	ret := m.ctrl.Call(m, "ValidateTemplate", input)
	ret0, _ := ret[0].(*cloudformation.ValidateTemplateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateTemplate indicates an expected call of ValidateTemplate
func (mr *MockCloudFormationAPIMockRecorder) ValidateTemplate(input interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateTemplate", reflect.TypeOf((*MockCloudFormationAPI)(nil).ValidateTemplate), input)
}

// CreateChangeSet mocks base method
func (m *MockCloudFormationAPI) CreateChangeSet(input *cloudformation.CreateChangeSetInput) (*cloudformation.CreateChangeSetOutput, error) {
	ret := m.ctrl.Call(m, "CreateChangeSet", input)
	ret0, _ := ret[0].(*cloudformation.CreateChangeSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChangeSet indicates an expected call of CreateChangeSet
func (mr *MockCloudFormationAPIMockRecorder) CreateChangeSet(input interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChangeSet", reflect.TypeOf((*MockCloudFormationAPI)(nil).CreateChangeSet), input)
}

// DescribeChangeSet mocks base method
func (m *MockCloudFormationAPI) DescribeChangeSet(input *cloudformation.DescribeChangeSetInput) (*cloudformation.DescribeChangeSetOutput, error) {
	ret := m.ctrl.Call(m, "DescribeChangeSet", input)
	ret0, _ := ret[0].(*cloudformation.DescribeChangeSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeChangeSet indicates an expected call of DescribeChangeSet
func (mr *MockCloudFormationAPIMockRecorder) DescribeChangeSet(input interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeChangeSet", reflect.TypeOf((*MockCloudFormationAPI)(nil).DescribeChangeSet), input)
}

// WaitUntilChangeSetCreateComplete mocks base method
func (m *MockCloudFormationAPI) WaitUntilChangeSetCreateComplete(input *cloudformation.DescribeChangeSetInput) error {
	ret := m.ctrl.Call(m, "WaitUntilChangeSetCreateComplete", input)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitUntilChangeSetCreateComplete indicates an expected call of WaitUntilChangeSetCreateComplete
func (mr *MockCloudFormationAPIMockRecorder) WaitUntilChangeSetCreateComplete(input interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitUntilChangeSetCreateComplete", reflect.TypeOf((*MockCloudFormationAPI)(nil).WaitUntilChangeSetCreateComplete), input)
}
