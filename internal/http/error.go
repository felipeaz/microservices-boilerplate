package http

import (
	"gorm.io/gorm"
	"net/http"
)

type Error interface {
	GetStatusCodeFromError(err error) int
}

func NewHttpError() Error {
	return &httpError{
		errorStatusMap: buildErrorMap(),
	}
}

type httpError struct {
	errorStatusMap map[error]int
}

func (h *httpError) GetStatusCodeFromError(err error) int {
	status, ok := h.errorStatusMap[err]
	if !ok {
		return http.StatusInternalServerError
	}
	return status
}

func buildErrorMap() map[error]int {
	return map[error]int{
		gorm.ErrRecordNotFound:     http.StatusNotFound,
		gorm.ErrPrimaryKeyRequired: http.StatusBadRequest,
	}
}
