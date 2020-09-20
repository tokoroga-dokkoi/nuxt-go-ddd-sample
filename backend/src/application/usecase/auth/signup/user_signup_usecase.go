package usecase_auth

import (
	"fmt"

	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model/user"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/repository"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/service/user"
	domain_service_users "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/service/users"
)

type IUserSignUpUsecase interface {
	SignUp() (*model.User, error)
}

type UserSignUpUsecase struct {
	userRepository repository.UserRepository
	userService    domain_service_users.UserService
}

func NewUserSignUpUsecase(useRepository repository.UserRepository, userService domain_service_users.UserService) *UserSignUpUsecase {
	userSignUpUsecase := UserSignUpUsecase{
		userRepository: useRepository,
		userService:    userService,
	}

	return &userSignUpUsecase
}

func (usecase *UserSignUpUsecase) SignUp(command UserSignUpInputCommand) (*model.User, error) {
	// create object
	email, err := user.NewEmail(command.Email)
	firstName, err := user.NewFirstName("new user")
	lastName, err := user.NewLastName("new user")
	displayName, err := user.NewDisplayName("new user displayname")

	if err != nil {
		return nil, err
	}

	user := model.NewUser(email, firstName, lastName, displayName)
	// Find user
	isAlreadyRegister := usecase.userService.EmailExists(user.Email)

	if isAlreadyRegister {
		return nil, fmt.Errorf("%s is already exists", command.Email)
	}

	usecase.userRepository.Save(user)

	return user, nil
}
