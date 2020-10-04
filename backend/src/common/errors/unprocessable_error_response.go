package common

// UnprocessableErrorResponse は422エラーを示す
type UnprocessableErrorResponse struct {
	code     uint
	messages string
}

// NewUnProcessableErrorResponse はコンストラクタ
func NewUnProcessableErrorResponse(err error) IErrorResponse {
	errObj := UnprocessableErrorResponse{
		code:     422,
		messages: err.Error(),
	}

	return &errObj
}

func (err UnprocessableErrorResponse) getMessage() string {
	return err.messages
}

func (err UnprocessableErrorResponse) getCode() uint {
	return err.code
}
