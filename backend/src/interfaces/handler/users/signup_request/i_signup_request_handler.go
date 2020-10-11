package users_signup_handler

import "github.com/labstack/echo"

// ISignUpRequestHandler は仮登録処理を行う
type ISignUpRequestHandler interface {
	SignUpRequest() echo.HandlerFunc
}
