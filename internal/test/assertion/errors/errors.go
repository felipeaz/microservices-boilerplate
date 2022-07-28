package errors

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrGeneric      = errors.New("generic error")
	ErrNotFound     = gorm.ErrRecordNotFound
	ErrCreatingUUID = errors.New("failed to create UUID from string")
)
