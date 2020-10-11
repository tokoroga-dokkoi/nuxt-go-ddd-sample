package router

import (
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/interfaces/handler"
	"github.com/labstack/echo"
)

// InitRouting is assign route
func InitRouting(e *echo.Echo, handler handler.AppHandler) {
	/*
	   ユーザの仮登録
	   /api/v1/users/signup_request
	*/
	e.POST("/api/v1/users/signup_request", handler.UsersHandler.SignUpRequestHandler.SignUpRequest())
}
