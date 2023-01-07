package redis

import (
	"fmt"
	"log"

	redigo "github.com/gomodule/redigo/redis"

	"app/internal/storage"
)

type redis struct {
	addr string
	port string
}

func New(host, port string) storage.Cache {
	return &redis{
		addr: host,
		port: port,
	}
}

func (r *redis) connect() redigo.Conn {
	conn, err := redigo.Dial("tcp", r.getHost())
	if err != nil {
		log.Fatalf("failed to connect to redis server: %v\n", err)
	}
	return conn
}

func (r *redis) Set(key string, value interface{}) error {
	conn := r.connect()
	defer r.closeConnection(conn)

	_, err := conn.Do("SET", key, value)
	if err != nil {
		log.Printf("failed to set key %s, value %s: %v\n", key, value, err)
	}
	return err
}

func (r *redis) Get(key string) ([]byte, error) {
	conn := r.connect()
	defer r.closeConnection(conn)

	data, err := redigo.Bytes(conn.Do("GET", key))
	if err != nil {
		log.Printf("failed to get key %s: %v\n", key, err)
	}
	return data, err
}

func (r *redis) Remove(key string) error {
	conn := r.connect()
	defer r.closeConnection(conn)

	_, err := conn.Do("DEL", key)
	if err != nil {
		log.Printf("failed to remove key %s: %v\n", key, err)
	}
	return err
}

func (r *redis) getHost() string {
	return fmt.Sprintf("%s:%s", r.addr, r.port)
}

func (r *redis) closeConnection(conn redigo.Conn) {
	err := conn.Close()
	log.Println("failed to close connection:", err)
}
