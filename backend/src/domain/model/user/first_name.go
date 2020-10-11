package user

import (
	"errors"
	"unicode/utf8"
)

type FirstName string

func NewFirstName(v string) (FirstName, error) {
	// Name is Required

	if v == "" {
		return "", errors.New("姓が入力されていません")
	}

	// Name is less than 101
	if utf8.RuneCountInString(v) > 100 {
		return "", errors.New("姓は100文字以内で入力してください")
	}

	firstNameEntity := FirstName(v)

	return firstNameEntity, nil
}
