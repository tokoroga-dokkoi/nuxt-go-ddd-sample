package domain_model_users

import "github.com/dgrijalva/jwt-go"

type RegistrationUrlToken *jwt.Token

// 認証用URLのトークンを発行する
func NewRegistrationUrlToken() RegistrationUrlToken {

	token := jwt.New(jwt.SigningMethodHS256)

	return RegistrationUrlToken(token)
}
