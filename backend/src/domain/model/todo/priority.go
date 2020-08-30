package todo

import (
  "errors"
)

type Priority string

func NewPriority(v string) (Priority, error) {
  if v == "" {
    return "", errors.New("重要度を入力してください")
  }
  // Priority is enum value
  switch v {
  case "high", "medium", "low", "unset":
    n := Priority(v)
    return n, nil
  default:
    return "", errors.New("重要度のフォーマットが正しくありません")
  }
}

func (priority Priority) Value() string {
  return string(priority)
}
