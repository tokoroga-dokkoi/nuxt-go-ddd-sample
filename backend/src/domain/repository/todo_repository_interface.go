package repository

import (
  "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model"
)
type TodoRepository interface {
  Find(id int) (*model.Todo, error)
  FindAll(limit int) ([]*model.Todo, error)
  Create(todo *model.Todo) error
  Save(todo *model.Todo) error
  Delete(todo *model.Todo) error
}