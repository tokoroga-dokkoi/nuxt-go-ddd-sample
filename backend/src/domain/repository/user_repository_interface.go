package repository

import (
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model"
)

type UserRepository interface {
	Find(id int) (*model.User, error)
	Save(todo *model.User) error
}
