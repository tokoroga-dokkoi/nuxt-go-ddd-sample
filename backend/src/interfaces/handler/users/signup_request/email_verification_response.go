package users_signup_handler

import "time"

// emailVerificationJsonResponse は201レスポンス
type emailVerificationJsonResponse struct {
	// 登録ID
	Id uint `json:"id"`
	// 仮登録時のメールアドレス
	Email string `json:"email"`
	// 仮登録日時
	CreatedAt time.Time `json:"createdAt"`
}
