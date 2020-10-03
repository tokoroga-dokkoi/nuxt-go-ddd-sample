package test_domain_model_users

import (
	"testing"

	domain_model_users "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model/users"
)

func TestUserEmailVerificationStatusSucceeds(t *testing.T) {
	var validStatuses = []string{
		"waiting_registration",
		"registered",
		"expired",
	}

	for _, status := range validStatuses {
		_, err := domain_model_users.NewUserEmailVerificationStatus(status)
		if err != nil {
			t.Errorf("cannot create object. status: %s", status)
		}
	}
}

func TestUserEmailVerificationStatusFailedInValidStatus(t *testing.T) {
	status := "invalid"

	_, err := domain_model_users.NewUserEmailVerificationStatus(status)

	if err == nil {
		t.Errorf("can create object. status: %s", status)
	}
}
