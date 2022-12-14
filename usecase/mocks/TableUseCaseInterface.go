// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "wmb-rest-api/model/entity"

	mock "github.com/stretchr/testify/mock"
)

// TableUseCaseInterface is an autogenerated mock type for the TableUseCaseInterface type
type TableUseCaseInterface struct {
	mock.Mock
}

// CreateTable provides a mock function with given fields: t
func (_m *TableUseCaseInterface) CreateTable(t *entity.Table) error {
	ret := _m.Called(t)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Table) error); ok {
		r0 = rf(t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTable provides a mock function with given fields: t
func (_m *TableUseCaseInterface) DeleteTable(t *entity.Table) error {
	ret := _m.Called(t)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Table) error); ok {
		r0 = rf(t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTable provides a mock function with given fields: t
func (_m *TableUseCaseInterface) GetTable(t *entity.Table) error {
	ret := _m.Called(t)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Table) error); ok {
		r0 = rf(t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTable provides a mock function with given fields: t
func (_m *TableUseCaseInterface) UpdateTable(t *entity.Table) error {
	ret := _m.Called(t)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Table) error); ok {
		r0 = rf(t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTableAvailability provides a mock function with given fields: t, isAvailable
func (_m *TableUseCaseInterface) UpdateTableAvailability(t *entity.Table, isAvailable bool) error {
	ret := _m.Called(t, isAvailable)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Table, bool) error); ok {
		r0 = rf(t, isAvailable)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTableUseCaseInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewTableUseCaseInterface creates a new instance of TableUseCaseInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTableUseCaseInterface(t mockConstructorTestingTNewTableUseCaseInterface) *TableUseCaseInterface {
	mock := &TableUseCaseInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
