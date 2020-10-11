package signup_request

import (
	"errors"
	"fmt"

	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/application/usecase"
	domain_model_users "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model/users"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/repository"
	domain_service_users "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/service/users"
)

// UserSignUpRequestUsecase はユーザの仮登録を行うユースケース
type UserSignUpRequestUsecase struct {
	emailVerificationRepository repository.IUserEmailVerificationRepository
	emailVerificationService    *domain_service_users.EmailVerificationService
}

// NewUserSignUpRequestUsecase はコンストトラクタ
func NewUserSignUpRequestUsecase(repo repository.IUserEmailVerificationRepository, domain_service *domain_service_users.EmailVerificationService) IUserAuthSignUpRequestUsecase {
	usecase := UserSignUpRequestUsecase{
		emailVerificationRepository: repo,
		emailVerificationService:    domain_service,
	}

	return &usecase
}

// SignUpRequest は仮登録を実装する
func (uc *UserSignUpRequestUsecase) SignUpRequest(command *UserSignUpRequestInputCommand) (*UserAuthSignUpRequestDto, error) {
	entityErrors := []string{}

	email, err := domain_model_users.NewEmail(command.Email)
	if err != nil {
		entityErrors = append(entityErrors, err.Error())
	}
	status, err := domain_model_users.NewUserEmailVerificationStatus("waiting_registration")
	if err != nil {
		entityErrors = append(entityErrors, err.Error())
	}
	token, err := domain_model_users.NewRegistrationUrlToken()
	if err != nil {
		entityErrors = append(entityErrors, err.Error())
	}
	// Value Object Create Error
	if len(entityErrors) > 0 {
		errorObj := usecase.NewUsecaseError("400", entityErrors, err)
		return nil, errorObj
	}

	// すでに存在しているか検証
	if uc.emailVerificationService.Exist(email) {
		errMes := []string{"すでに仮登録済です。本登録メールをご確認ください。"}
		orgErr := errors.New(fmt.Sprintf("[UserSignUpRequestUsecase.SignUpRequest] User is already registered.\n user email: %s", email))
		errorObj := usecase.NewUsecaseError("422", errMes, orgErr)
		return nil, errorObj
	}

	// Create Entity
	emailVerification := domain_model_users.NewUserEmailVerification(email, status, token)

	// Save
	err = uc.emailVerificationRepository.Save(emailVerification)
	// Failed Save
	if err != nil {
		// DB層のエラーは内部エラーとして処理する
		return nil, err
	}

	// Create DTO
	dto := NewUserAuthSignUpRequestDto(emailVerification)

	return &dto, nil
}
