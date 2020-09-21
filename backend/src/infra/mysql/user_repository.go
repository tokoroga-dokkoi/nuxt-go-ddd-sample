package mysql

import (
	"errors"
	"time"

	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model/user"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/repository"
)

type UserRepository struct {
	sqlHandler *SQLHandler
}

func NewUserRepository(sqlHandler *SQLHandler) repository.UserRepository {
	userRepository := UserRepository{sqlHandler}

	return &userRepository
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	var user model.User

	sqlResult := r.sqlHandler.Conn.Where("id = ?", id).First(&user)

	if sqlResult.RecordNotFound() {
		return nil, errors.New("Record is not found.")
	}

	return &user, sqlResult.Error
}

func (r *UserRepository) FindByEmail(email user.Email) (*model.User, error) {
	var user model.User

	sqlResult := r.sqlHandler.Conn.Where("email = ?", email.Value()).Find(&user)

	if sqlResult.RecordNotFound() {
		return nil, errors.New("Record is not found.")
	}

	return &user, sqlResult.Error
}

func (r *UserRepository) Save(user *model.User) error {

	result := r.sqlHandler.Conn.Where("id = ?", user.ID).
		Assign(model.User{
			Email:        user.Email,
			Uid:          user.Uid,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			DisplayName:  user.DisplayName,
			LastSigninAt: time.Now(),
		}).FirstOrCreate(&user)

	return result.Error
}
