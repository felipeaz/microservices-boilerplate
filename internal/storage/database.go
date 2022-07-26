package storage

type Database interface {
	Create(obj interface{}) error
	Update(obj interface{}) error
	Set(obj interface{}, field string, value interface{}) error
	Select(obj interface{}) error
	Raw(query string, obj interface{}) error
	Delete(obj interface{}) error
}
