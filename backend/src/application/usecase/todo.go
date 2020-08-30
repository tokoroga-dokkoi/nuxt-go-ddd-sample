package usecase

import (
  "fmt"
  "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/repository"
  "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/application/command"
  "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model"
  "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model/todo"
)

// Create Usecase
type TodoUsecase interface {
  FindAll() ([]*model.Todo, error)
  Add(*command.TodoCreateCommand) (*model.Todo, error)
  Update(*command.TodoUpdateCommand) (*model.Todo, error)
  Delete(int) error
}

type TypeTodoUsecase struct {
  todoRepo repository.TodoRepository
}

func NewTodoUsecase(todoRepo repository.TodoRepository) TodoUsecase {
  todoUsecase := TypeTodoUsecase{todoRepo: todoRepo}
  return &todoUsecase
}

// Operation

func (usecase *TypeTodoUsecase) FindAll() ([]*model.Todo, error) {
  todoList, err := usecase.todoRepo.FindAll(10)
  if err != nil {
    return nil, err
  }

  return todoList, nil
}

func (usecase *TypeTodoUsecase) Add(command *command.TodoCreateCommand) (*model.Todo, error) {
  // Create Entity
  name, errName := todo.NewName(command.Name)
  if errName != nil {
    return nil, errName
  }
  priority, errPriority := todo.NewPriority(command.Priority)
  if errPriority != nil {
    return nil, errPriority
  }
  descrition, errDesc := todo.NewDescription(command.Descrition)
  if errDesc != nil {
    return nil, errDesc
  }
  todo := model.NewTodo(name, priority, descrition)
  err  := usecase.todoRepo.Create(todo)
  if err != nil {
    return nil, err
  }

  return todo, nil
}

func (usecase *TypeTodoUsecase) Update(command *command.TodoUpdateCommand) (*model.Todo, error) {
  // Find Entity
  resultTodo, err := usecase.todoRepo.Find(command.Id)
  // If DB Error happened
  if resultTodo == nil || err != nil {
    fmt.Println(err)
    return nil, err
  }
  // Update Name
  name, _ := todo.NewName(command.Name)
  resultTodo.ChangeName(name)
  // Update Priority
  priority, _ := todo.NewPriority(command.Priority)
  resultTodo.ChangePriority(priority)
  // Update description
  description, _ := todo.NewDescription(command.Descrition)
  resultTodo.ChangeDescription(description)

  // Store Data
  err = usecase.todoRepo.Save(resultTodo)
  if err != nil {
    return nil, err
  }

  return resultTodo, nil
}

func (usecase *TypeTodoUsecase) Delete(todoId int) error {
  // Find Todo
  targetTodo, err := usecase.todoRepo.Find(todoId)

  if targetTodo == nil || err != nil {
    fmt.Println(err)
    return err
  }

  // Delete Component

  deleteErr := usecase.todoRepo.Delete(targetTodo)

  if deleteErr != nil {
    return deleteErr
  }

  return nil
}
