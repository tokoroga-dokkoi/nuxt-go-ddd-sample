package domain_service_users

import (
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model/user"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	userService := UserService{
		userRepository: userRepository,
	}

	return &userService
}

func (userSerice *UserService) EmailExists(email user.Email) bool {
	userRes, _ := userSerice.userRepository.FindByEmail(email)

	if userRes == nil {
		return false
	}

	return true
}
