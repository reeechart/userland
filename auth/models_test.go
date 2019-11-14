package auth

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	user User
)

func testAuthModelReset() {
	user = User{
		Id:                 1,
		Fullname:           "user",
		Email:              "user@example.com",
		Password:           "password",
		Location:           sql.NullString{String: "Jakarta, Indonesia", Valid: true},
		Bio:                sql.NullString{String: "Hi there", Valid: true},
		Web:                sql.NullString{String: "https://example.com", Valid: true},
		Verified:           true,
		ProfilePicture:     nil,
		VerificationToken:  sql.NullString{String: "", Valid: false},
		ResetPasswordToken: sql.NullString{String: "", Valid: false},
		CreatedAt:          time.Now(),
	}
}

func TestUserAbleToLogin(t *testing.T) {
	testAuthModelReset()
	assert.True(t, user.ableToLogin(), "User should be able to login when email and password are provided")

	user.Email = ""
	assert.False(t, user.ableToLogin(), "User should be unable to login when email is empty")

	testAuthModelReset()
	user.Password = ""
	assert.False(t, user.ableToLogin(), "User should be unable to login when password is empty")

	user.Email = ""
	assert.False(t, user.ableToLogin(), "User should be unable to login when email and password are empty")

	testAuthModelReset()
}
