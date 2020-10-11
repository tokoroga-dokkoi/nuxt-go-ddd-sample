package domain_service_users

import (
	domain_model_users "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model/users"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/repository"
)

// EmailVerificationService は仮登録のドメインサービス
type EmailVerificationService struct {
	emailVerificationRepository repository.IUserEmailVerificationRepository
}

// NewEmailVerificationService はコンストラクタ
func NewEmailVerificationService(emailVerificationRepository repository.IUserEmailVerificationRepository) *EmailVerificationService {
	emailVerificationService := EmailVerificationService{
		emailVerificationRepository: emailVerificationRepository,
	}

	return &emailVerificationService
}

// Exist はすでに仮登録されているメールアドレスか検証を行う
func (s *EmailVerificationService) Exist(email domain_model_users.Email) bool {
	result, _ := s.emailVerificationRepository.FindByEmail(email)

	if result == nil {
		return false
	}

	return true
}
