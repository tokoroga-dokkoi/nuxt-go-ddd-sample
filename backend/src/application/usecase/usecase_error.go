package usecase

import (
	"strings"

	handler "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/interfaces/handler/errors"
)

// UsecaseError はユースケース層で発生したエラーをハンドリングするクラス
type UsecaseError struct {
	code          handler.ErrorCode
	message       []string
	originalError error
}

// NewUsecaseError はユースケース層で発生したエラーのコンストラクタ
func NewUsecaseError(code string, message []string, originalError error) *UsecaseError {
	errCode := handler.ErrorCode("400")
	errObj := UsecaseError{
		code:          errCode,
		message:       message,
		originalError: originalError,
	}

	return &errObj
}

// Code はErrorコードを返却する
func (e *UsecaseError) Code() handler.ErrorCode {
	return e.code
}

// Message はユーザに出力するエラーメッセージを返却する
func (e *UsecaseError) Message() []string {
	return e.message
}

// Error は、ユーザに出力するエラーメッセージとトレースバックを返却する
func (e *UsecaseError) Error() string {
	joinedMessage := strings.Join(e.message, "\n")

	return joinedMessage + ":" + e.originalError.Error()
}
