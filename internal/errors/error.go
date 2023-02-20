package errors

import (
	"net/http"

	"gorm.io/gorm"
)

var errorStatusMap = map[error]int{
	gorm.ErrRecordNotFound:     http.StatusNotFound,
	gorm.ErrPrimaryKeyRequired: http.StatusBadRequest,
}

func GetStatus(err error) int {
	status, ok := errorStatusMap[err]
	if !ok {
		return http.StatusInternalServerError
	}
	return status
}
