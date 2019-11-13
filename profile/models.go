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

func (user UserProfile) hasValidProfile() bool {
	return user.hasValidFullname() && user.hasValidLocation() && user.hasValidBio() && user.hasValidWeb()
}

func (user UserProfile) hasValidFullname() bool {
	return len(user.Fullname) >= 3 && len(user.Fullname) <= 128
}

func (user UserProfile) hasValidLocation() bool {
	return len(user.Location) <= 128
}

func (user UserProfile) hasValidBio() bool {
	return len(user.Bio) <= 255
}

func (user UserProfile) hasValidWeb() bool {
	return len(user.Web) <= 128
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
