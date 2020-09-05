package user_test

import (
	"strings"
	"testing"

	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model/user"
)

// Success Test Case
func TestFirstNameEntitySucceeds(t *testing.T) {
	var firstName = strings.Repeat("a", 100)

	_, err := user.NewFirstName(firstName)

	if err != nil {
		t.Errorf("expected %s, got %s", "nil", err)
	}
}

// failed Test Case
func TestFirstNameEntityFailedIfEmtpty(t *testing.T) {
	var firstName = ""
	var expErr = "姓が入力されていません"

	_, err := user.NewFirstName(firstName)

	if err == nil {
		t.Errorf("expected %s, got %s", expErr, err)
	}
}

func TestFirstNameEntityFailedIfOver101(t *testing.T) {
	var firstName = strings.Repeat("a", 101)
	var expErr = "姓は100文字以内で入力してください"

	_, err := user.NewFirstName(firstName)

	if err == nil {
		t.Errorf("expected %s, got %s", expErr, err)
	}
}
