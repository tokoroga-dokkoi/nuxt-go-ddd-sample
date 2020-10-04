package usecase_auth

import application_common "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/application/common"

// IUserAuthSignUpRequestUsecase は仮登録のユースケースを実現するインターフェース
// 実装は実装元の処理に任せる
type IUserAuthSignUpRequestUsecase interface {
	SignUpRequest(command UserSignUpRequestInputCommand) (*UserAuthSignUpRequestDto, *application_common.IErrorResponse)
}
