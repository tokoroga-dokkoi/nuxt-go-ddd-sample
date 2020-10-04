package application_common

// InternalServerErrorResponse は500エラーを示す
type InternalServerErrorResponse struct {
	code    uint
	message string
}

// NewInternalServerErrorResponse はコンストラクタ
func NewInternalServerErrorResponse(err error) IErrorResponse {
	errObj := InternalServerErrorResponse{
		code:    500,
		message: err.Error(),
	}
	// [Todo] 初期化時にSentryなどのツールに送る
	return &errObj
}

func (err InternalServerErrorResponse) getMessage() string {
	return err.message
}

func (err InternalServerErrorResponse) getCode() uint {
	return err.code
}
