package usecase_auth

type UserSignUpInputCommand struct {
	Email   string `json:"email"`
	IdToken string `json:"idToken"`
}
