// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sjanota/budget/backend/pkg/resolver (interfaces: BudgetResolverStorage)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/sjanota/budget/backend/pkg/models"
)

// MockBudgetResolverStorage is a mock of BudgetResolverStorage interface
type MockBudgetResolverStorage struct {
	ctrl     *gomock.Controller
	recorder *MockBudgetResolverStorageMockRecorder
}

// MockBudgetResolverStorageMockRecorder is the mock recorder for MockBudgetResolverStorage
type MockBudgetResolverStorageMockRecorder struct {
	mock *MockBudgetResolverStorage
}

// NewMockBudgetResolverStorage creates a new mock instance
func NewMockBudgetResolverStorage(ctrl *gomock.Controller) *MockBudgetResolverStorage {
	mock := &MockBudgetResolverStorage{ctrl: ctrl}
	mock.recorder = &MockBudgetResolverStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBudgetResolverStorage) EXPECT() *MockBudgetResolverStorageMockRecorder {
	return m.recorder
}

// GetMonthlyReport mocks base method
func (m *MockBudgetResolverStorage) GetMonthlyReport(arg0 context.Context, arg1 models.MonthlyReportID) (*models.MonthlyReport, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMonthlyReport", arg0, arg1)
	ret0, _ := ret[0].(*models.MonthlyReport)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMonthlyReport indicates an expected call of GetMonthlyReport
func (mr *MockBudgetResolverStorageMockRecorder) GetMonthlyReport(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMonthlyReport", reflect.TypeOf((*MockBudgetResolverStorage)(nil).GetMonthlyReport), arg0, arg1)
}
