// Code generated by mockery v2.14.0. DO NOT EDIT.

package storage

import mock "github.com/stretchr/testify/mock"

// Database is an autogenerated mock type for the Database type
type Database struct {
	mock.Mock
}

// Create provides a mock function with given fields: obj
func (_m *Database) Create(obj interface{}) error {
	ret := _m.Called(obj)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: obj
func (_m *Database) Delete(obj interface{}) error {
	ret := _m.Called(obj)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Raw provides a mock function with given fields: query, obj
func (_m *Database) Raw(query string, obj interface{}) error {
	ret := _m.Called(query, obj)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(query, obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Select provides a mock function with given fields: obj
func (_m *Database) Select(obj interface{}) error {
	ret := _m.Called(obj)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Set provides a mock function with given fields: obj, field, value
func (_m *Database) Set(obj interface{}, field string, value interface{}) error {
	ret := _m.Called(obj, field, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, interface{}) error); ok {
		r0 = rf(obj, field, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: obj
func (_m *Database) Update(obj interface{}) error {
	ret := _m.Called(obj)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewDatabase interface {
	mock.TestingT
	Cleanup(func())
}

// NewDatabase creates a new instance of Database. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDatabase(t mockConstructorTestingTNewDatabase) *Database {
	mock := &Database{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}