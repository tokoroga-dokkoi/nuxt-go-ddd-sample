package router

import (
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/interfaces/handler"
	"github.com/labstack/echo"
)

// InitRouting is assign route
func InitRouting(e *echo.Echo, handler handler.IAppHandler) {
	// /api/v1/auth/signup
	e.POST("/api/v1/auth/signup", handler.SignUp())
}
