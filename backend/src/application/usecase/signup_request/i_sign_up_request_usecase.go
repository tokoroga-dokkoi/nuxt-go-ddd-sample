package signup_request

// IUserAuthSignUpRequestUsecase は仮登録のユースケースを実現するインターフェース
// 実装は実装元の処理に任せる
type IUserAuthSignUpRequestUsecase interface {
	SignUpRequest(command UserSignUpRequestInputCommand) (*UserAuthSignUpRequestDto, error)
}
