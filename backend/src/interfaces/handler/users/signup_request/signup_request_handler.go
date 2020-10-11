package users_signup_handler

import (
	"fmt"
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
	fmt.Println("[Call] ハンドラーが呼び出されました")
	return func(ec echo.Context) error {
		c := ec.(*errors.AppErrorContext)
		// input command
		params := new(signup_request.UserSignUpRequestInputCommand)
		if err := c.Bind(params); err != nil {
			fmt.Println(err.Error())
			fmt.Println("bind failed")
			return c.ErrorResponse(err)
		}

		dto, err := handler.signUpRequestUsecase.SignUpRequest(params)

		// error response
		if err != nil {
			return c.ErrorResponse(err)
		}

		// json response
		resJson := emailVerificationJsonResponse{
			Id:        dto.Id,
			Email:     dto.Email,
			CreatedAt: dto.CreatedAt,
		}

		return c.JSON(http.StatusCreated, resJson)
	}
}
