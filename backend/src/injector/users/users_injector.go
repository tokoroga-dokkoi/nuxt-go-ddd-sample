package usersinjector

import (
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/infra/mysql"
	handler "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/interfaces/handler/users"
)

func InjectUserHandler(sqlHandler *mysql.SQLHandler) handler.UserHandler {
	// 仮登録
	signupRequestHandler := NewSignUpRequestInjector(sqlHandler)

	// UserHandler
	return handler.UserHandler{
		SignUpRequestHandler: signupRequestHandler,
	}
}
