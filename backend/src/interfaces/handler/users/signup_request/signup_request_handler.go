package handler

import (
	"net/http"

	signup_request "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/application/usecase/signup_request"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/interfaces/handler/errors"
	"github.com/labstack/echo"
)

// SignUpRequestHandler は仮登録のハンドラー
type SignUpRequestHandler struct {
	// ユースケース
	signUpRequestUsecase signup_request.IUserAuthSignUpRequestUsecase
}

// NewSignUpRequestHandler はコンストラクタ
func NewSignUpRequestHandler(usecase signup_request.IUserAuthSignUpRequestUsecase) ISignUpRequestHandler {
	handler := SignUpRequestHandler{
		signUpRequestUsecase: usecase,
	}

	return &handler
}

// SignUpRequest : POST /api/v1/users/signup_request
func (handler *SignUpRequestHandler) SignUpRequest() echo.HandlerFunc {
	return func(ec echo.Context) error {
		c := ec.(*errors.AppErrorContext)
		// input command
		params := new(signup_request.UserSignUpRequestInputCommand)
		c.Bind(params)

		// call usecase
		dto, err := handler.signUpRequestUsecase.SignUpRequest(params)

		// error response
		if err != nil {
			return c.ErrorResponse(err)
		}

		// json response
		resJson := emailVerificationJsonResponse{
			id:        dto.Id,
			email:     dto.Email,
			createdAt: dto.CreatedAt,
		}

		return c.JSON(http.StatusOK, resJson)
	}
}
