package user

import (
	"errors"
	"unicode/utf8"
)

type DisplayName string

func NewDisplayName(v string) (DisplayName, error) {
	if v == "" {
		return "", errors.New("表示名が入力されていません")
	}

	if utf8.RuneCountInString(v) > 50 {
		return "", errors.New("表示名は50文字以内で入力してください")
	}

	displayNameEntity := DisplayName(v)

	return displayNameEntity, nil
}
