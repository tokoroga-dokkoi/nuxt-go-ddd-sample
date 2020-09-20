package usecase_auth

type UserSignUpInputCommand struct {
	Email   string `json:"email" validate: "required"`
	IdToken string `json:"idToken" validate: "required"`
}
