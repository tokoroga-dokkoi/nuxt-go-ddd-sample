package user_test

import (
  "testing"
  "github.com/bxcodec/faker/v3"
  "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model/user"
)

// Success Case
func TestEmailEntitySucceeds (t *testing.T) {

  var email = faker.Email()
  var expErr = nil

  emailEntity, err = user.NewEmail(email)

  if err != nil {
    t.Errorf("expected %s, got %s", expErr, err)
  }
}

// Failed Case
func TestEmailEntityFailedIfEmptyEmail (t *testing.T) {
  var email = ""
  var expErr = "メールアドレスの入力は必須です"

  emailEntity, err = user.NewEmail(email)

  if err == nil {
    t.Errorf("expected %s, got %s", expected, err)
  }
}

func TestEmailEntityFailedIfValidPattern (t *testing.T) {
  var email = "hogehogeho"
  var expErr = "メールアドレスが正しくありません"

  emailEntity, err = user.NewEmail(email)

  if err == nil {
    t.Errorf("expected %s, got %s", expected, err)
  }
}
