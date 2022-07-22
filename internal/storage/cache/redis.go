package cache

import (
	"fmt"

	redigo "github.com/gomodule/redigo/redis"

	"microservices-boilerplate/internal/storage"
)

type redis struct {
	addr string
}

func NewRedis(host string) storage.Cache {
	return redis{
		addr: host,
	}
}

func (r redis) connect() (redigo.Conn, error) {
	conn, err := redigo.Dial("tcp", r.addr)
	if err != nil {
		return conn, fmt.Errorf("failed to connect to dial redis: %v", err)
	}

	return conn, nil
}

func (r redis) Set(key string, value interface{}) error {
	conn, err := r.connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Do("SET", key, value)
	if err != nil {
		err = fmt.Errorf("error setting key %s with value %s: %v", key, value, err)
	}
	return err
}

func (r redis) Get(key string) ([]byte, error) {
	conn, err := r.connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	data, err := redigo.Bytes(conn.Do("GET", key))
	if err != nil {
		err = fmt.Errorf("error getting key %s: %v", key, err)
	}
	return data, err
}

func (r redis) Remove(key string) error {
	conn, err := r.connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Do("DEL", key)
	if err != nil {
		return fmt.Errorf("failed to remove key %s: %v", key, err)
	}
	return nil
}
