package redis

import (
	"fmt"
	"log"

	redigo "github.com/gomodule/redigo/redis"

	"microservices-boilerplate/internal/storage"
)

type redis struct {
	addr string
	port string
}

func New(host, port string) storage.Cache {
	return redis{
		addr: host,
		port: port,
	}
}

func (r redis) connect() (redigo.Conn, error) {
	conn, err := redigo.Dial("tcp", r.getHost())
	if err != nil {
		log.Printf("failed to connect to redis server: %v\n", err)
	}
	return conn, err
}

func (r redis) Set(key string, value interface{}) error {
	conn, err := r.connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Do("SET", key, value)
	if err != nil {
		log.Printf("failed to set key %s, value %s: %v\n", key, value, err)
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
		log.Printf("failed to get key %s: %v\n", key, err)
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
		log.Printf("failed to remove key %s: %v\n", key, err)
	}
	return err
}

func (r redis) getHost() string {
	return fmt.Sprintf("%s:%s", r.addr, r.port)
}
