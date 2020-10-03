package domain_model_users

import (
	"time"

	"github.com/jinzhu/gorm"
)

// EmailVerification はユーザの仮登録データを表すドメインモデル
type EmailVerification {
  gorm.Model
  Email Email
  RegistrationUrlToken RegistrationUrlToken
}

// NewEmailVerification はユーザの仮登録オブジェクトを生成する
func NewEmailVerification(email Email) *EmailVerification {
  token := NewRegistrationUrlToken()
  return &EmailVerification {
    Email: email,
    RegistrationUrlToken: token
  }
}
