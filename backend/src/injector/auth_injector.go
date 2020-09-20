package injector

import (
	usecase_auth "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/application/usecase/auth/signup"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/repository"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/infra/mysql"
	handler "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/interfaces/handler/auth"
)

func NewUserRepository(sqlHandler mysql.SQLHandler) repository.UserRepository {
	return mysql.NewUserRepository(sqlHandler)
}

func NewAuthUsecase(sqlHandler mysql.SQLHandler) usecase_auth.IUserAuthUsecase {
	userRepository := NewUserRepository(sqlHandler)
	return usecase_auth.NewAuthUsecase(userRepository)
}

func InjectAuthHandler(sqlHandler &mysql.SQLHandler) handler.IAuthHandler {
	authUsecase := NewAuthUsecase(sqlHandler)

	return handler.NewAuthHandler(authUsecase)
}
