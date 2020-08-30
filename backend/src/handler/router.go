package handler

import (
  "github.com/labstack/echo"
)

func InitRouting(e *echo.Echo, todoHandler TodoHandler) {
  // Todo
  e.GET("/api/v1/todos", todoHandler.Index())
  e.POST("/api/v1/todos", todoHandler.Create())
  e.PUT("/api/v1/todos/:id", todoHandler.Update())
  e.DELETE("/api/v1/todos/:id", todoHandler.Delete())
}
