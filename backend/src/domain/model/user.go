package model

import (
	"time"

	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model/user"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Uid          string
	Email        user.Email
	FirstName    user.FirstName
	LastName     user.LastName
	DisplayName  user.DisplayName
	LastSigninAt time.Time `type:date`
}

func NewUser(email user.Email, uid string, firstName user.FirstName, lastName user.LastName, displayName user.DisplayName) *User {
	return &User{
		Email:       email,
		Uid:         uid,
		FirstName:   firstName,
		LastName:    lastName,
		DisplayName: displayName,
	}
}
