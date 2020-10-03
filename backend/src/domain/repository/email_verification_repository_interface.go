package repository

import domain_model_users "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model/users"

type IUserEmailVerificationRepository interface {
	FindByEmail(email domain_model_users.Email) (*domain_model_users.UserEmailVerification, error)
	Save(emailVerification domain_model_users.UserEmailVerification) error
}
