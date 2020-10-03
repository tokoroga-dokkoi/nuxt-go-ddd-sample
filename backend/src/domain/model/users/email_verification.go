package domain_model_users

import (
	"time"

	"github.com/jinzhu/gorm"
)

// EmailVerification はユーザの仮登録データを表すドメインモデル
type EmailVerification struct {
	gorm.Model
	Email                   Email
	RegistrationURLToken    RegistrationUrlToken
	Status                  EmailVerificationStatus
	RegistrationEmailSentAt time.Time `type:datetime`
	URLTokenExpiredAt       time.Time `type:datetime`
}

// NewEmailVerification はユーザの仮登録オブジェクトを生成する
func NewEmailVerification(email Email) *EmailVerification {
	token := NewRegistrationUrlToken()
	status, _ := NewEmailVerificationStatus("waiting_registration")
	// デフォルトは１週間
	tokenExpiredAt := time.Now().AddDate(0, 0, 7)
	return &EmailVerification{
		Email:                   email,
		RegistrationURLToken:    token,
		Status:                  status,
		RegistrationEmailSentAt: time.Now(),
		URLTokenExpiredAt:       tokenExpiredAt,
	}
}
