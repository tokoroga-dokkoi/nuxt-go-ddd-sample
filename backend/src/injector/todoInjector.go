package injector

import (
  "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/repository"
  "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/handler"
  "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/infra/mysql"
  "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/application/usecase"
)

// Create Sql handler
func InjectDb()  mysql.SQLHandler {
  sqlHandler := mysql.NewSqlHandler()
  return *sqlHandler
}
// Create Repository
func InjectTodoRepository() repository.TodoRepository {
  sqlHandler := InjectDb()
  return mysql.NewTodoRepository(sqlHandler)
}

// Create Usecase
func InjectTodoUsecase() usecase.TodoUsecase {
  todoRepo := InjectTodoRepository()
  return usecase.NewTodoUsecase(todoRepo)
}

// Create Handler
func InjectTodoHandler() handler.TodoHandler {
  todoUsecase := InjectTodoUsecase()
  return handler.NewTodoHandler(todoUsecase)
}
