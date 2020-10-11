package repository

import domain_model_users "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model/users"

// IUserEmailVerificationRepository は仮登録リポジトリのインタフェース
type IUserEmailVerificationRepository interface {
	FindByEmail(email domain_model_users.Email) (*domain_model_users.UserEmailVerification, error)
	Save(emailVerification *domain_model_users.UserEmailVerification) error
}
