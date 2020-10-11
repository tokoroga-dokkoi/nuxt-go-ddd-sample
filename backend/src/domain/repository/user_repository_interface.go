package repository

import (
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model/user"
)

type UserRepository interface {
	Find(id int) (*model.User, error)
	FindByEmail(email user.Email) (*model.User, error)
	Save(user *model.User) error
}
