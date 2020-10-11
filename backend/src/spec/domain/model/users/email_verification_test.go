package test_domain_model_users

import (
	"testing"

	domain_model_users "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model/users"
	"github.com/bxcodec/faker/v3"
)

// UserEmailVerificationが生成されること
func TestUserEmailVerificationEntitySucceeds(t *testing.T) {
	email, _ := domain_model_users.NewEmail(
		faker.Email(),
	)
	status, _ := domain_model_users.NewUserEmailVerificationStatus("waiting_registration")
	token := domain_model_users.NewRegistrationUrlToken()

	result := domain_model_users.NewUserEmailVerification(email, status, token)

	// トークンの執行期限が登録の7日後になっていること
	expectedExpiredAt := result.RegistrationEmailSentAt.AddDate(0, 0, 7)
	if result.URLTokenExpiredAt != expectedExpiredAt {
		t.Errorf("got %v\n expected %v", result.URLTokenExpiredAt, expectedExpiredAt)
	}
}
