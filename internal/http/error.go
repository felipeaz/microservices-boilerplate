package http

import (
	"net/http"

	"gorm.io/gorm"
)

type ResponseError struct {
	StatusCode int
	Error      error
}

type Error interface {
	GetHttpResponseError(err error) ResponseError
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

func (h *httpError) GetHttpResponseError(err error) ResponseError {
	return ResponseError{
		StatusCode: h.getStatusCodeFromError(err),
		Error:      err,
	}
}

func (h *httpError) getStatusCodeFromError(err error) int {
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
