package main

import (
	"fmt"
	"os"

	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/injector"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/interfaces/handler/errors"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/interfaces/router"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func dumpHandler(c echo.Context, reqBody, resBody []byte) {
	fmt.Fprint(os.Stdout, "Request:", string(reqBody))
}

func main() {
	// Create Echo Instance
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Skipper: func(c echo.Context) bool {
			if c.Request().Header.Get("X-Debug") == "" {
				return false
			}
			return false
		},
		Handler: dumpHandler,
	}))
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return h(&errors.AppErrorContext{c})
		}
	})
	// e.Use(MyMiddleware.FirebaseJwtValidator)

	// Todos Routing
	appHandler := injector.InjectHandlers()
	router.InitRouting(e, appHandler)
	e.Logger.Fatal(e.Start(":3001"))
}
