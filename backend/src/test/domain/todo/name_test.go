package test

import (
  "testing"
  "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model/todo"
  "github.com/stretchr/testify/assert"
)

func TestTodoName(t *testing.T) {
  t.Run("success NewName()", func(t *testing.T){
    var name = 'テストName'
    todoName, err := todo.NewName(name)
    assert.Equal(todoName.Value(), name)
  })

  t.Run("failed NewName() when name is empty", func(t *testing.T){
    var name = ''
    assert.Error(t, todo.NewName(name))
  })

  t.Run("failed NewName() when name over then 40", func(t *testing.T){
    var name = ''
  })
}

func TestDomainTodoNameSuccess(t *testing.T) {
  result, err := todo.NewName('テストName')
  if err == nil {
    t.Fatalf("\n ValueObject TodoNameが生成されませんでした")
  }
}

func TestDomainTodoNameFailedWhenIfEmpty(t *testing.T) {
  result, err := todo.NewName('')
  if err != nil {
    t.Fatalf("\n ValueObject TodoNameが空文字で生成されました")
  }
}

func TestDomainTodoNameFailedWhenIfOverFortyCharLength(t *testing.T) {
  result, err := todo.NewName()
}
