package storage

import "github.com/stretchr/testify/mock"

type Database struct {
	mock.Mock
}

func (d *Database) Create(obj interface{}) error {
	called := d.Called(obj)
	return called.Error(0)
}

func (d *Database) Update(obj interface{}) error {
	called := d.Called(obj)
	return called.Error(0)
}

func (d *Database) Set(obj interface{}, field string, value interface{}) error {
	called := d.Called(obj, field, value)
	return called.Error(0)
}

func (d *Database) Select(obj interface{}) error {
	called := d.Called(obj)
	return called.Error(0)
}

func (d *Database) Raw(query string, obj interface{}) error {
	called := d.Called(query, obj)
	return called.Error(0)
}

func (d *Database) Delete(obj interface{}) error {
	called := d.Called(obj)
	return called.Error(0)
}
