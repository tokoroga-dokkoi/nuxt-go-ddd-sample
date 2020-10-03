package domain_model_users

import "errors"

// EmailVerificationStatus は仮登録の状態を示す(waiting_registration, registered, expired)
type EmailVerificationStatus string

// NewEmailVerificationStatus はコンストラクタ
func NewEmailVerificationStatus(status string) (EmailVerificationStatus, error) {

	if status == "" {
		return "", errors.New("ステータスが入力されていません")
	}

	if !isStatusValid(status) {
		return "", errors.New("仮登録のステータスが不正です")
	}

	return EmailVerificationStatus(status), nil
}

// isStatusValid は引数で受け取ったstatusがABLE_STATUSに含まれているか判定する
func isStatusValid(status string) bool {
	for _, val := range ableStatusList() {
		if status == val {
			return true
		}
	}

	return false
}

func ableStatusList() []string {
	// statusに入れることができる文字列
	var status = []string{
		"waiting_registration",
		"registered",
		"expired",
	}
	return status
}
