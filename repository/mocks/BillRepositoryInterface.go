// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "wmb-rest-api/model/entity"

	mock "github.com/stretchr/testify/mock"
)

// BillRepositoryInterface is an autogenerated mock type for the BillRepositoryInterface type
type BillRepositoryInterface struct {
	mock.Mock
}

// CreateBillPayment provides a mock function with given fields: bp
func (_m *BillRepositoryInterface) CreateBillPayment(bp *entity.BillPayment) error {
	ret := _m.Called(bp)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.BillPayment) error); ok {
		r0 = rf(bp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateTransaction provides a mock function with given fields: c, t, tt, details
func (_m *BillRepositoryInterface) CreateTransaction(c *entity.Customer, t *entity.Table, tt *entity.TransactionType, details *[]entity.BillDetail) (int, error) {
	ret := _m.Called(c, t, tt, details)

	var r0 int
	if rf, ok := ret.Get(0).(func(*entity.Customer, *entity.Table, *entity.TransactionType, *[]entity.BillDetail) int); ok {
		r0 = rf(c, t, tt, details)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.Customer, *entity.Table, *entity.TransactionType, *[]entity.BillDetail) error); ok {
		r1 = rf(c, t, tt, details)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: b
func (_m *BillRepositoryInterface) Delete(b *entity.Bill) error {
	ret := _m.Called(b)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Bill) error); ok {
		r0 = rf(b)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAllBillDetail provides a mock function with given fields: by
func (_m *BillRepositoryInterface) FindAllBillDetail(by map[string]interface{}) ([]entity.BillDetail, error) {
	ret := _m.Called(by)

	var r0 []entity.BillDetail
	if rf, ok := ret.Get(0).(func(map[string]interface{}) []entity.BillDetail); ok {
		r0 = rf(by)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.BillDetail)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]interface{}) error); ok {
		r1 = rf(by)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllByDate provides a mock function with given fields: date
func (_m *BillRepositoryInterface) FindAllByDate(date string) ([]entity.Bill, error) {
	ret := _m.Called(date)

	var r0 []entity.Bill
	if rf, ok := ret.Get(0).(func(string) []entity.Bill); ok {
		r0 = rf(date)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Bill)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(date)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: b
func (_m *BillRepositoryInterface) FindById(b *entity.Bill) error {
	ret := _m.Called(b)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Bill) error); ok {
		r0 = rf(b)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewBillRepositoryInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewBillRepositoryInterface creates a new instance of BillRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBillRepositoryInterface(t mockConstructorTestingTNewBillRepositoryInterface) *BillRepositoryInterface {
	mock := &BillRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
