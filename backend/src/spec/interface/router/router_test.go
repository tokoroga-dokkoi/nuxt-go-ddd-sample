package router_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	usecaseAuth "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/application/usecase/auth/signup"
	handler "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/interfaces/handler"
	authHandler "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/interfaces/handler/auth"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/interfaces/router"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestRouterSignUpSucceeds(t *testing.T) {
	testAuthUsecase := usecaseAuth.NewTestAuthUsecase()
	authHandler := authHandler.NewAuthHandler(testAuthUsecase)
	appHandler := handler.NewAppHandler(authHandler)

	e := echo.New()
	router.InitRouting(e, appHandler)

	signUpJson := `{"name": "hogehoge", "email": "hogehoge@example.com"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/signup", strings.NewReader(signUpJson))
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
