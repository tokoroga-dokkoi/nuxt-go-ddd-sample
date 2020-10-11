package domain_model_users

import (
	"time"

	"github.com/jinzhu/gorm"
)

// UserEmailVerification はメールアドレスの認証を表すドメインモデル
type UserEmailVerification struct {
	gorm.Model
	Email                   Email
	RegistrationURLToken    RegistrationUrlToken
	Status                  UserEmailVerificationStatus
	RegistrationEmailSentAt time.Time `type:datetime`
	URLTokenExpiredAt       time.Time `type:datetime`
}

// NewUserEmailVerification はユーザの仮登録オブジェクトを生成する
func NewUserEmailVerification(email Email, status UserEmailVerificationStatus, registrationURLToken RegistrationUrlToken) *UserEmailVerification {
	emailSentAt := time.Now()
	urlTokenExpiredAt := emailSentAt.AddDate(0, 0, 7)

	return &UserEmailVerification{
		Email:                   email,
		RegistrationURLToken:    registrationURLToken,
		Status:                  status,
		RegistrationEmailSentAt: emailSentAt,
		URLTokenExpiredAt:       urlTokenExpiredAt,
	}
}
