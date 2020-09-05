package user_test

import (
	"strings"
	"testing"

	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model/user"
)

func TestDisplayNameSucceeds(t *testing.T) {
	var displayName = strings.Repeat("a", 50)

	_, err := user.NewDisplayName(displayName)

	if err != nil {
		t.Errorf("expected %s, value %s", "nil", err)
	}
}

func TestDisplayNameFailedIfEmpty(t *testing.T) {
	var displayName = ""
	var expected = "表示名が入力されていません"

	_, err := user.NewDisplayName(displayName)

	if err == nil {
		t.Errorf("expected %s, value %s", expected, err)
	}
}

func TestDisplayNameFailedIfOver51(t *testing.T) {
	var displayName = strings.Repeat("a", 51)
	var expected = "表示名は50文字以内で入力してください"

	_, err := user.NewDisplayName(displayName)

	if err == nil {
		t.Errorf("expected %s, value %s", expected, err)
	}
}
