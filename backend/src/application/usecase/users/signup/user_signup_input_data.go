package usecase_users

type UserSignUpInputCommand struct {
	Email   string `json:"email"`
	IdToken string `json:"idToken"`
}
