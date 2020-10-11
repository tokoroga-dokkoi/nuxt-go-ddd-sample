package users_signup_handler

import "time"

// emailVerificationJsonResponse は201レスポンス
type emailVerificationJsonResponse struct {
	// 登録ID
	id uint `json:"id"`
	// 仮登録時のメールアドレス
	email string `json:"email"`
	// 仮登録日時
	createdAt time.Time `json:"createdAt"`
}
