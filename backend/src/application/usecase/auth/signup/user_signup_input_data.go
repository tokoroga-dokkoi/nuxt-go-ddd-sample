package usecase_auth

type UserSignUpInputCommand struct {
	Email   string `json:"email"`
	IDToken string `json:"idToken"`
}
