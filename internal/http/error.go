package http

import (
	"net/http"

	"gorm.io/gorm"
)

type Error interface {
	GetStatusCodeFromError(err error) int
}

func NewHttpError() Error {
	return &httpError{
		errorStatusMap: buildStatusFromErrorMap(),
	}
}

type httpError struct {
	errorStatusMap  map[error]int
	errorMessageMap map[int]string
}

func (h *httpError) GetStatusCodeFromError(err error) int {
	status, ok := h.errorStatusMap[err]
	if !ok {
		return http.StatusInternalServerError
	}
	return status
}

func buildStatusFromErrorMap() map[error]int {
	return map[error]int{
		gorm.ErrRecordNotFound:     http.StatusNotFound,
		gorm.ErrPrimaryKeyRequired: http.StatusBadRequest,
	}
}
