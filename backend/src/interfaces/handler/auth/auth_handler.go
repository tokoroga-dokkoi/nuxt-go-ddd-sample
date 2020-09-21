package handler

import (
	"log"
	"net/http"

	usecase_auth "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/application/usecase/auth/signup"
	session "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/interfaces/session/firebase"
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
		authenticator := session.NewFirebaseAuth()
		// fetch command
		signUpCommand := new(usecase_auth.UserSignUpInputCommand)

		if err := c.Bind(signUpCommand); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		// validate id token
		result := authenticator.ValidateIdToken(signUpCommand.IDToken)

		if !result {
			return c.JSON(http.StatusForbidden, "invalid id token")
		}

		log.Printf("params: %v\n", signUpCommand)
		user, err := handler.signUpUsecase.SignUp(signUpCommand)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, user)
	}
}
