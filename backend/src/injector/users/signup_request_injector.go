package usersinjector

import (
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/application/usecase/signup_request"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/repository"
	domain_service_users "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/service/users"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/infra/mysql"
	users_signup_handler "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/interfaces/handler/users/signup_request"
)

// NewEmailVerificationRepository は本番用のSQLHandlerを作成する
func NewEmailVerificationRepository(sqlHandler *mysql.SQLHandler) repository.IUserEmailVerificationRepository {
	return mysql.NewUserEmailVerificationRepository(sqlHandler)
}

// NewSignUpRequestUsecase はユースケースをDIする
func NewSignUpRequestUsecase(sqlHandler *mysql.SQLHandler) signup_request.IUserAuthSignUpRequestUsecase {
	repo := NewEmailVerificationRepository(sqlHandler)
	ds := domain_service_users.NewEmailVerificationService(repo)
	usecase := signup_request.NewUserSignUpRequestUsecase(repo, ds)

	return usecase
}

// NewSignUpRequestInjector は仮登録のDIコンテナ
func NewSignUpRequestInjector(sqlHandler *mysql.SQLHandler) users_signup_handler.ISignUpRequestHandler {
	usecase := NewSignUpRequestUsecase(sqlHandler)

	return users_signup_handler.NewSignUpRequestHandler(usecase)
}
