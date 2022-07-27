package storage

import uuid "github.com/satori/go.uuid"

type Database interface {
	Create(obj interface{}) error
	Update(id uuid.UUID, obj interface{}) error
	Set(obj interface{}, field string, value interface{}) error
	Select(obj interface{}) error
	Raw(query string, obj interface{}) error
	Delete(id uuid.UUID, obj interface{}) error
}
