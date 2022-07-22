package cache

type Cache interface {
	Set(key string, value interface{}) error
	Get(key string) ([]byte, error)
	Remove(key string) error
}
