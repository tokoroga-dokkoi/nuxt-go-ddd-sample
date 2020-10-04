package common

// BadRequestErrorResponse は400エラーを示す
type BadRequestErrorResponse struct {
	code     uint
	messages string
}

// NewBadRequestErrorResponse はコンストラクタ
func NewBadRequestErrorResponse(err error) IErrorResponse {
	errObj := BadRequestErrorResponse{
		code:     400,
		messages: err.Error(),
	}
	return &errObj
}

func (err BadRequestErrorResponse) getMessage() string {
	return err.messages
}

func (err BadRequestErrorResponse) getCode() uint {
	return err.code
}
