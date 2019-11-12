package profile

import (
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

type ChangePasswordRequest struct {
	PasswordCurrent string `json:"password_current"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

func (req ChangePasswordRequest) hasMatchingNewPassword() bool {
	return req.Password == req.PasswordConfirm
}

type DeleteAccountRequest struct {
	Password string `json:"password"`
}
