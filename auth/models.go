package auth

import (
	"database/sql"
	"time"
)

type User struct {
	Id                 int            `json:"id"`
	Fullname           string         `json:"fullname"`
	Email              string         `json:"email"`
	Password           string         `json:"password"`
	Location           sql.NullString `json:"location"`
	Bio                sql.NullString `json:"bio"`
	Web                sql.NullString `json:"web"`
	Verified           bool           `json:"verified"`
	ProfilePicture     []byte         `json:"picture" db:"picture"`
	VerificationToken  sql.NullString `json:"verification_token" db:"verification_token"`
	ResetPasswordToken sql.NullString `json:"reset_password_token" db:"reset_password_token"`
	CreatedAt          time.Time      `json:"created_at" db:"created_at"`
}

func (u *User) ableToLogin() bool {
	return u.Email != "" && u.Password != ""
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

type verificationRequest struct {
	Type              string `json:"type"`
	Recipient         string `json:"recipient"`
	VerificationToken string `json:"verification_token"`
}

func (req verificationRequest) isValid() bool {
	return req.Type != "" && req.Recipient != "" && req.VerificationToken != ""
}

type resetPasswordRequest struct {
	Token           string `json:"token"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

func (req resetPasswordRequest) hasMatchingPassword() bool {
	return req.Password == req.PasswordConfirm
}
