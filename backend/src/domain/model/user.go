package model

import (
  "time"
  "github.com/jinzhu/gorm"
  "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model/user"
)

type User struct (
  gorm.model
  Email user.Email
  FirstName user.FirstName
  LastName user.LastName
  DisplayName user.DisplayName
  LastSigninAt time.Time `type:date`
)


func NewUser(email todo.Email, firstName user.FirstName, lastName user.LastName, displayName DisplayName) (*User) {
  return &User{
    Email: email,
    FirstName: firstName,
    LastName: lastName,
    DisplayName: displayName
  }
}
