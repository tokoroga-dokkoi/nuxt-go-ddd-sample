package todo

import (
  "errors"
  "unicode/utf8"
)

type Name string

func NewName(v string) (Name, error) {
  // Name is Required
  if v == "" {
    return "", errors.New("名前の入力が必要です")
  }
  // Name is less than 40
  if utf8.RuneCountInString(v) > 40 {
    return "", errors.New("名前は40文字以下で入力してください。")
  }

  n := Name(v)

  return n, nil
}

func (name Name) Value() string {
  return string(name)
}
