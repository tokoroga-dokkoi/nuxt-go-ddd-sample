package usecase_auth

import (
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model/user"
)

type TestUserSignUpUsecase struct{}

func NewTestAuthUsecase() IUserAuthUsecase {
	userSignUpUsecase := TestUserSignUpUsecase{}

	return &userSignUpUsecase
}

func (usecase *TestUserSignUpUsecase) SignUp(command *UserSignUpInputCommand) (*model.User, error) {

	// test create object
	email, err := user.NewEmail(command.Email)
	firstName, err := user.NewFirstName("hogehoge")
	lastName, err := user.NewLastName("hogehoge")
	displayName, err := user.NewDisplayName("hogehoge")

	if err != nil {
		return nil, err
	}

	user := model.NewUser(email, firstName, lastName, displayName)

	return user, nil
}
