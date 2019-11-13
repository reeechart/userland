package profile

import (
	"regexp"
	"time"
)

type UserProfile struct {
	Id             int       `json:"id"`
	Fullname       string    `json:"fullname"`
	Location       string    `json:"location"`
	Bio            string    `json:"bio"`
	Web            string    `json:"web"`
	ProfilePicture []byte    `json:"picture" db:"picture"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

type ChangeEmailRequest struct {
	NewEmail string `json:"email"`
}

func (req ChangeEmailRequest) hasValidEmail() bool {
	emailFormatValid := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString(req.NewEmail)
	return len(req.NewEmail) <= 128 && emailFormatValid
}

type ChangePasswordRequest struct {
	PasswordCurrent string `json:"password_current"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

func (req ChangePasswordRequest) hasMatchingNewPassword() bool {
	return req.Password == req.PasswordConfirm
}

func (req ChangePasswordRequest) hasValidPassword() bool {
	return len(req.Password) >= 6 && len(req.Password) <= 128
}

type DeleteAccountRequest struct {
	Password string `json:"password"`
}
