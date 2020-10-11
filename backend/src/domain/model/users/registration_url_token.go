package domain_model_users

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type RegistrationUrlToken string

// 認証用URLのトークンを発行する
func NewRegistrationUrlToken() (RegistrationUrlToken, error) {
	// tokenを変えるため時刻文字列をkeyにする
	secret := time.Now()
	token := jwt.New(jwt.SigningMethodHS256)

	strToken, err := token.SignedString([]byte(secret.String()))

	if err != nil {
		return "", errors.New("トークンの発行に失敗しました")
	}
	return RegistrationUrlToken(strToken), nil
}
