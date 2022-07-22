package cache

import "github.com/stretchr/testify/mock"

type Cache struct {
	mock.Mock
}

func (c *Cache) Set(key string, value interface{}) error {
	called := c.Called(key, value)
	return called.Error(0)
}

func (c *Cache) Get(key string) ([]byte, error) {
	called := c.Called(key)
	return called.Get(0).([]byte), called.Error(1)
}

func (c *Cache) Remove(key string) error {
	called := c.Called(key)
	return called.Error(0)
}
