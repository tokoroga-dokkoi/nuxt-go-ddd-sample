package handler

import (
	"net/http"
	"strconv"

	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/application/command"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/application/response"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/application/usecase"
	"github.com/labstack/echo"
)

// Create Handler
type TodoHandler struct {
	todoUsecase usecase.TodoUsecase
}

func NewTodoHandler(todoUsecase usecase.TodoUsecase) TodoHandler {
	todoHandler := TodoHandler{todoUsecase: todoUsecase}

	return todoHandler
}

// CRUD Operation

// GET /api/v1/todos
func (handler *TodoHandler) Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		var responseTodos []response.TodoResponse
		result, err := handler.todoUsecase.FindAll()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		for _, data := range result {
			todo := response.TodoResponse{
				ID:          data.GetId(),
				Name:        data.GetName(),
				Priority:    data.GetPriority(),
				Description: data.GetDescription(),
				CreatedAt:   data.CreatedAt,
				UpdatedAt:   data.UpdatedAt,
			}

			responseTodos = append(responseTodos, todo)
		}
		return c.JSON(http.StatusOK, responseTodos)
	}
}

// POST /api/v1/todos
func (handler *TodoHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var todo command.TodoCreateCommand
		c.Bind(&todo)
		result, err := handler.todoUsecase.Add(&todo)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusCreated, result)
	}
}

// PUT /api/v1/todos/{:id}
func (handler *TodoHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var todo command.TodoUpdateCommand
		todo.Id, _ = strconv.Atoi(c.Param("id"))
		c.Bind(&todo)
		result, err := handler.todoUsecase.Update(&todo)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, result)
	}
}

// DELETE /api/v1/todos/:id
func (handler *TodoHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		todoId, _ := strconv.Atoi(c.Param("id"))
		err := handler.todoUsecase.Delete(todoId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.NoContent(http.StatusNoContent)
	}
}
