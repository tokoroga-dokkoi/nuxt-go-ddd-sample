package usershandler

import (
	USignUpHandler "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/interfaces/handler/users/signup_request"
)

type UserHandler struct {
	SignUpRequestHandler USignUpHandler.ISignUpRequestHandler
}
