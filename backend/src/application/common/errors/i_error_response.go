package application_common

// IErrorResponse はエラーレスポンスのインタフェース
type IErrorResponse interface {
	getMessage() string
	getCode() uint
}
