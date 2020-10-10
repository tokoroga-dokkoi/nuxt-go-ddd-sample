package handler

// errorResponse define struct hold data
// response error for client
type errorResponse struct {
	Code   ErrorCode `json:"code"`
	Errors []string  `json:"errors"`
}

// generalError interface should be
type generalError interface {
	Code() ErrorCode
	Messages() []string
	Error() string
}

type internalError interface {
	Internal() bool
	Error() string
}
