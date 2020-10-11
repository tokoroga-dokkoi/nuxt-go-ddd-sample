package injector

import (
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/application/usecase/signup_request"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/repository"
	domain_service_users "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/service/users"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/infra/mysql"
	handler "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/interfaces/handler/users"
)

func NewEmailVerificationRepository(sqlHandler *mysql.SQLHandler) repository.IUserEmailVerificationRepository {
	return mysql.NewUserEmailVerificationRepository(sqlHandler)
}

func NewSignUpRequestUsecase(sqlHandler *mysql.SQLHandler) signup_request.IUserAuthSignUpRequestUsecase {
	repo := NewEmailVerificationRepository(sqlHandler)
	ds := domain_service_users.NewEmailVerificationService(repo)
	usecase := signup_request.NewUserSignUpRequestUsecase(repo, ds)

	return usecase
}

func SignUpRequestInjector(sqlHandler *mysql.SQLHandler) handler.ISignUpRequestHandler {
  usecase := NewSignUpRequestUsecase(sqlHandler)

  return handler.SignUpRequestHandler {
    signupRequestUsecase: usecase
  }
}
