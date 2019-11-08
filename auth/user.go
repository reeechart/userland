package auth

type User struct {
	Id                 int    `json:"id"`
	Fullname           string `json:"fullname"`
	Email              string `json:"email"`
	Password           string `json:"password"`
	Location           string `json:"location"`
	Bio                string `json:"bio"`
	Web                string `json:"web"`
	ProfilePicture     []byte `json:"profile_picture"`
	VerificationToken  string `json:"verification_token"`
	ResetPasswordToken string `json:"reset_password_token"`
}

type userRegistration struct {
	Fullname        string `json:"fullname"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

func (u *userRegistration) hasValidData() bool {
	return u.Fullname != "" && u.Email != "" && u.Password != "" && u.PasswordConfirm != ""
}

func (u *userRegistration) hasMatchingPassword() bool {
	return u.Password == u.PasswordConfirm
}
