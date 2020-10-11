package mysql

import (
	"errors"
	"fmt"

	domain_model_users "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model/users"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/repository"
)

// UserEmailVerificationRepository はsqlHandlerをフィールドとしてもつ
type UserEmailVerificationRepository struct {
	sqlHandler *SQLHandler
}

// NewUserEmailVerificationRepository はコンストラクタ
func NewUserEmailVerificationRepository(sqlHandler *SQLHandler) repository.IUserEmailVerificationRepository {
	repo := UserEmailVerificationRepository{sqlHandler}

	return &repo
}

// FindByEmail はメールアドレスの検索メソッド
func (r *UserEmailVerificationRepository) FindByEmail(email domain_model_users.Email) (*domain_model_users.UserEmailVerification, error) {
	var emailVerification domain_model_users.UserEmailVerification

	result := r.sqlHandler.Conn.Where("email = ?", email.Value()).First(&emailVerification)

	if result.RecordNotFound() {
		return nil, errors.New("Record Is Not Found")
	}

	return &emailVerification, nil
}

// Save は1件レコードを生成するメソッド
func (r *UserEmailVerificationRepository) Save(emailVerification *domain_model_users.UserEmailVerification) error {

	result := r.sqlHandler.Conn.Where("id = ?", emailVerification.ID).
		Assign(emailVerification).FirstOrCreate(&emailVerification)

	if result.Error != nil {
		errMes := fmt.Sprintf("[UserEmailVerificationRepository.Save] failed to save email verification. email: %s", emailVerification.Email.Value())
		errObj := NewMySQLError(errMes, result.Error)
		return errObj
	}

	return nil
}
