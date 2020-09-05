package user

import (
	"errors"
	"unicode/utf8"
)

type LastName string

func NewLastName(v string) (LastName, error) {

	if v == "" {
		return "", errors.New("名前が入力されていません")
	}

	// Name is less than 101
	if utf8.RuneCountInString(v) > 100 {
		return "", errors.New("姓は100文字以内で入力してください")
	}

	lastNameEntity := LastName(v)

	return lastNameEntity, nil
}
