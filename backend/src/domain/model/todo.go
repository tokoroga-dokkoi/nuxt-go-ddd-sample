package model

import (
  "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model/todo"
  "github.com/jinzhu/gorm"
)
type Todo struct {
  gorm.Model
  Name todo.Name
  Priority todo.Priority
  Description todo.Description
}


func NewTodo(name todo.Name, priority todo.Priority, description todo.Description) (*Todo) {
  return &Todo{
    Name: name,
    Priority: priority,
    Description: description}
}

// Getter
func (todo Todo) GetId() uint {
  return todo.ID
}
func (todo Todo) GetName() string {
  return todo.Name.Value()
}

func (todo Todo) GetPriority() string {
  return todo.Priority.Value()
}

func (todo Todo) GetDescription() string {
  return todo.Description.Value()
}

// Setter

func (todo *Todo) ChangeName(name todo.Name) {
  todo.Name = name
}

func (todo *Todo) ChangePriority(priority todo.Priority) {
  todo.Priority = priority
}

func (todo *Todo) ChangeDescription(description todo.Description) {
  todo.Description = description
}
