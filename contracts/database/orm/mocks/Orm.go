// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	orm "github.com/goravel/framework/contracts/database/orm"
	mock "github.com/stretchr/testify/mock"

	sql "database/sql"
)

// Orm is an autogenerated mock type for the Orm type
type Orm struct {
	mock.Mock
}

// Connection provides a mock function with given fields: name
func (_m *Orm) Connection(name string) orm.Orm {
	ret := _m.Called(name)

	var r0 orm.Orm
	if rf, ok := ret.Get(0).(func(string) orm.Orm); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Orm)
		}
	}

	return r0
}

// DB provides a mock function with given fields:
func (_m *Orm) DB() (*sql.DB, error) {
	ret := _m.Called()

	var r0 *sql.DB
	var r1 error
	if rf, ok := ret.Get(0).(func() (*sql.DB, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *sql.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.DB)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Factory provides a mock function with given fields:
func (_m *Orm) Factory() orm.Factory {
	ret := _m.Called()

	var r0 orm.Factory
	if rf, ok := ret.Get(0).(func() orm.Factory); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Factory)
		}
	}

	return r0
}

// Observe provides a mock function with given fields: model, observer
func (_m *Orm) Observe(model interface{}, observer orm.Observer) {
	_m.Called(model, observer)
}

// Query provides a mock function with given fields:
func (_m *Orm) Query() orm.Query {
	ret := _m.Called()

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func() orm.Query); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Transaction provides a mock function with given fields: txFunc
func (_m *Orm) Transaction(txFunc func(orm.Transaction) error) error {
	ret := _m.Called(txFunc)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(orm.Transaction) error) error); ok {
		r0 = rf(txFunc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WithContext provides a mock function with given fields: ctx
func (_m *Orm) WithContext(ctx context.Context) orm.Orm {
	ret := _m.Called(ctx)

	var r0 orm.Orm
	if rf, ok := ret.Get(0).(func(context.Context) orm.Orm); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Orm)
		}
	}

	return r0
}

// NewOrm creates a new instance of Orm. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrm(t interface {
	mock.TestingT
	Cleanup(func())
}) *Orm {
	mock := &Orm{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
