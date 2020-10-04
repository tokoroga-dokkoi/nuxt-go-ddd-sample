package usecase_auth_signup_request

import (
	"time"

	domain_model_users "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model/users"
)

// UserAuthSignUpRequestDto は仮登録のレスポンス用DTO
type UserAuthSignUpRequestDto struct {
	id        uint      `json:"id"`
	email     string    `json:"email"`
	createdAt time.Time `json:"createdAt"`
}

// NewUserAuthSignUpRequestDto はコンストラクタ
func NewUserAuthSignUpRequestDto(emailVerification *domain_model_users.UserEmailVerification) UserAuthSignUpRequestDto {
	return UserAuthSignUpRequestDto{
		id:        emailVerification.ID,
		email:     emailVerification.Email.Value(),
		createdAt: emailVerification.CreatedAt,
	}
}
