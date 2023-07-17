// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/usecase/category_usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	model "github.com/Ikhlashmulya/golang-clean-architecture-project-structure/pkg/model"
	gomock "go.uber.org/mock/gomock"
)

// MockCategoryUsecase is a mock of CategoryUsecase interface.
type MockCategoryUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryUsecaseMockRecorder
}

// MockCategoryUsecaseMockRecorder is the mock recorder for MockCategoryUsecase.
type MockCategoryUsecaseMockRecorder struct {
	mock *MockCategoryUsecase
}

// NewMockCategoryUsecase creates a new mock instance.
func NewMockCategoryUsecase(ctrl *gomock.Controller) *MockCategoryUsecase {
	mock := &MockCategoryUsecase{ctrl: ctrl}
	mock.recorder = &MockCategoryUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategoryUsecase) EXPECT() *MockCategoryUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCategoryUsecase) Create(ctx context.Context, request model.CreateCategoryRequest) model.CategoryResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, request)
	ret0, _ := ret[0].(model.CategoryResponse)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockCategoryUsecaseMockRecorder) Create(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCategoryUsecase)(nil).Create), ctx, request)
}

// Delete mocks base method.
func (m *MockCategoryUsecase) Delete(ctx context.Context, categoryId int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", ctx, categoryId)
}

// Delete indicates an expected call of Delete.
func (mr *MockCategoryUsecaseMockRecorder) Delete(ctx, categoryId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCategoryUsecase)(nil).Delete), ctx, categoryId)
}

// FindAll mocks base method.
func (m *MockCategoryUsecase) FindAll(ctx context.Context) []model.CategoryResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx)
	ret0, _ := ret[0].([]model.CategoryResponse)
	return ret0
}

// FindAll indicates an expected call of FindAll.
func (mr *MockCategoryUsecaseMockRecorder) FindAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockCategoryUsecase)(nil).FindAll), ctx)
}

// FindById mocks base method.
func (m *MockCategoryUsecase) FindById(ctx context.Context, categoryId int) model.CategoryResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", ctx, categoryId)
	ret0, _ := ret[0].(model.CategoryResponse)
	return ret0
}

// FindById indicates an expected call of FindById.
func (mr *MockCategoryUsecaseMockRecorder) FindById(ctx, categoryId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockCategoryUsecase)(nil).FindById), ctx, categoryId)
}

// Update mocks base method.
func (m *MockCategoryUsecase) Update(ctx context.Context, request model.UpdateCategoryRequest) model.CategoryResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, request)
	ret0, _ := ret[0].(model.CategoryResponse)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockCategoryUsecaseMockRecorder) Update(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCategoryUsecase)(nil).Update), ctx, request)
}