package errors

// errorResponse define struct hold data
// response error for client
type errorResponse struct {
	Code   ErrorCode `json:"code"`
	Errors []string  `json:"errors"`
}

// generalError interface should be
type GeneralError interface {
	Code() ErrorCode
	Messages() []string
}

type internalError interface {
	Internal() bool
}
