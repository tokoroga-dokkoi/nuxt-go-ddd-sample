package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

// FirebaseJwtValidator is verifing id token on middleware
func FirebaseJwtValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()
		app, err := firebase.NewApp(ctx, nil)

		if err != nil {
			log.Errorf("error initializing firebase app %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		firebaseAuthClient, err := app.Auth(ctx)

		if err != nil {
			log.Errorf("error getting firebase auth client: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		authHeader := c.Request().Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		_, err = firebaseAuthClient.VerifyIDToken(ctx, idToken)

		if err != nil {
			unAuthMes := fmt.Sprintf("error verifying ID Token: %v \n", idToken)
			return echo.NewHTTPError(http.StatusBadRequest, unAuthMes)
		}

		return next(c)
	}
}
