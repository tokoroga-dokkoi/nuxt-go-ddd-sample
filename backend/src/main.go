package main

import (
  "fmt"
  "os"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/handler"
  "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/injector"
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
        return true
      }
      return false
    },
    Handler: dumpHandler,
  }))

  // Todos Routing
  todoHandler := injector.InjectTodoHandler()
  handler.InitRouting(e, todoHandler)
  e.Logger.Fatal(e.Start(":8080"))
}
