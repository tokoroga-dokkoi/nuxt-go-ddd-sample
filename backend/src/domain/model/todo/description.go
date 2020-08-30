package todo

import (
  "errors"
  "unicode/utf8"
)

type Description string

func NewDescription(v string) (Description, error) {
  // Description is less then 200
  if utf8.RuneCountInString(v) > 200 {
    return "", errors.New("概要は200文字以下で入力してください")
  }

  n:= Description(v)

  return n, nil
}


func (description Description) Value() string {
  return string(description)
}
