package usecase_auth

import (
	common "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/application/common"
	domain_model_users "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model/users"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/repository"
	domain_service_users "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/service/users"
)

// UserSignUpRequestUsecase はユーザの仮登録を行うユースケース
type UserSignUpRequestUsecase struct {
	emailVerificationRepository repository.IUserEmailVerificationRepository
	emailVerificationService    domain_service_users.EmailVerificationService
}

// NewUserSignUpRequestUsecase はコンストトラクタ
func NewUserSignUpRequestUsecase(repo repository.IUserEmailVerificationRepository, domain_service domain_service_users.EmailVerificationService) IUserAuthSignUpRequestUsecase {
	usecase := UserSignUpRequestUsecase{
		emailVerificationRepository: repo,
		emailVerificationService:    domain_service,
	}

	return &usecase
}

// SignUpRequest は仮登録を実装する
func (uc *UserSignUpRequestUsecase) SignUpRequest(command UserSignUpRequestInputCommand) (*UserAuthSignUpRequestDto, *common.IErrorResponse) {
	email, err := domain_model_users.NewEmail(command.Email)
	if err != nil {
		return nil, common.NewUnProcessableErrorResponse(err)
	}
	status, err := domain_model_users.NewUserEmailVerificationStatus("waiting_registration")

	if err != nil {
		return nil, common.NewUnProcessableErrorResponse(err)
	}
	token := domain_model_users.NewRegistrationUrlToken()

	// すでに存在しているか検証
	if uc.emailVerificationService.Exist(email) {
		return nil, common.NewBadRequestErrorResponse(err)
	}

	// Create Entity
	emailVerification := domain_model_users.NewUserEmailVerification(email, status, token)

	// Save
	err = uc.emailVerificationRepository.Save(emailVerification)
	// Failed Save
	if err != nil {
		return nil, common.NewInternalServerErrorResponse(err)
	}

	// Create DTO
	dto := NewUserAuthSignUpRequestDto(emailVerification)

	return &dto, nil
}
