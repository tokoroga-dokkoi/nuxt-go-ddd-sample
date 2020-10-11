package signup_request

import (
	"time"

	domain_model_users "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model/users"
)

// UserAuthSignUpRequestDto は仮登録のレスポンス用DTO
type UserAuthSignUpRequestDto struct {
	Id        uint
	Email     string
	CreatedAt time.Time
}

// NewUserAuthSignUpRequestDto はコンストラクタ
func NewUserAuthSignUpRequestDto(emailVerification *domain_model_users.UserEmailVerification) UserAuthSignUpRequestDto {
	return UserAuthSignUpRequestDto{
		Id:        emailVerification.ID,
		Email:     emailVerification.Email.Value(),
		CreatedAt: emailVerification.CreatedAt,
	}
}
