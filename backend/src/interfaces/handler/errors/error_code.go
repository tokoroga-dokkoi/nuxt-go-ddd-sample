package handler

import "net/http"

// ErrorCode is error type to bu used in this service api.
type ErrorCode string

const (
	// 40x ErrorCode
	BadRequest          ErrorCode = "400"
	Forbidden           ErrorCode = "403"
	Unauthorized        ErrorCode = "401"
	NotFound            ErrorCode = "404"
	UnprocessableEntity ErrorCode = "422"
	// 50x ErrorCode
	InternalServerError ErrorCode = "500"
)

var codeStatusMap = map[ErrorCode]int{
	BadRequest:          http.StatusBadRequest,
	Forbidden:           http.StatusForbidden,
	Unauthorized:        http.StatusUnauthorized,
	NotFound:            http.StatusNotFound,
	UnprocessableEntity: http.StatusUnprocessableEntity,
	InternalServerError: http.StatusInternalServerError,
}

func GetHTTPStatus(code ErrorCode) int {
	return codeStatusMap[code]
}
