package handler

import (
	"net/http"

	usecase_auth "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/application/usecase/auth/signup"
	"github.com/labstack/echo"
)

type IAuthHandler interface {
	SignUp() echo.HandlerFunc
}

type AuthHandler struct {
	signUpUsecase usecase_auth.IUserAuthUsecase
}

func NewAuthHandler(authUsecase usecase_auth.IUserAuthUsecase) IAuthHandler {
	authHandler := AuthHandler{signUpUsecase: authUsecase}

	return &authHandler
}

// POST /api/v1/auth/signup
func (handler *AuthHandler) SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := handler.signUpUsecase.SignUp()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, result)
	}
}
