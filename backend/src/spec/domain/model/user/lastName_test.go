pacakge user_test

import (
	"strings"
	"testing"

	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model/user"
)


// Success Test Case
func TestLastNameEntitySucceeds(t *testing.T) {
  var lastName = strings.Repeat("a", 100)

  _, err := user.NewLastName(lastName)

  if err != nil {
    t.Errorf("expected %s, value %s", "nil", err)
  }
}

// Failed Test Case
func TestLastNameEntityFailedIfEmpty(t *testing.T) {
  var lastName = ""
  var expected = "姓が入力されていません"
  _, err := user.NewLastName(lastName)

  if err == nil {
    t.Errorf("expected %s, value %s", expected, err)
  }
}

func TestLastNameEntityFailedIfOver101(t *testing.T) {
  var lastName = strings.Repeat("a", 101)
  var expected = "姓は100文字以内で入力してください"

  _, err := user.NewLastName(lastName)

  if err == nil {
    t.Errorf("expected %s, value %s", expected, err)
  }
}
