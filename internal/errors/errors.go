package errors

import "errors"

var (
	ErrCreatingUUIDFromString = errors.New("failed to create UUID from string")
)
