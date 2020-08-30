package mysql

import (
  "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model"
  "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/repository"
  "errors"
)

// Implements Repository
type TodoRepository struct {
  sqlHandler SQLHandler
}

func NewTodoRepository(sqlHandler SQLHandler) repository.TodoRepository {
  todoRepository := TodoRepository{sqlHandler}
  return &todoRepository
}

// SQL Operation
func (r *TodoRepository) Find(id int) (*model.Todo, error) {
  var todo model.Todo
  result := r.sqlHandler.Conn.Where("id = ?", id).First(&todo)
  if result.RecordNotFound() {
    return nil, errors.New("Record Is Not Found")
  }
  return &todo, result.Error
}
func (r *TodoRepository) FindAll(limit int) ([]*model.Todo, error) {
  var todoList []*model.Todo
  result := r.sqlHandler.Conn.Limit(10).Find(&todoList)
  return todoList, result.Error
}
func (r *TodoRepository) Create(todo *model.Todo) error {
  result := r.sqlHandler.Conn.Create(&todo)
  return result.Error
}
func (r *TodoRepository) Save(todo *model.Todo) error {
  result := r.sqlHandler.Conn.Where("id = ?", todo.GetId()).
                  Assign(model.Todo{
                    Name: todo.Name,
                    Priority: todo.Priority,
                    Description: todo.Description,
                  }).
                  FirstOrCreate(&todo)
  return result.Error
}
func (r *TodoRepository) Delete(todo *model.Todo) error {
  result := r.sqlHandler.Conn.Delete(&todo)

  return result.Error
}
