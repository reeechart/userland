package auth

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	user             User
	registrationData userRegistration
)

func resetUserModel() {
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

func resetRegistrationDataModel() {
	registrationData = userRegistration{
		Fullname:        "user",
		Email:           "user@example.com",
		Password:        "password",
		PasswordConfirm: "password",
	}
}

func TestUserAbleToLogin(t *testing.T) {
	resetUserModel()
	assert.True(t, user.ableToLogin(), "User should be able to login when email and password are provided")

	user.Email = ""
	assert.False(t, user.ableToLogin(), "User should be unable to login when email is empty")

	resetUserModel()
	user.Password = ""
	assert.False(t, user.ableToLogin(), "User should be unable to login when password is empty")

	user.Email = ""
	assert.False(t, user.ableToLogin(), "User should be unable to login when email and password are empty")

	resetUserModel()
}

func TestUserRegistrationHasCompleteData(t *testing.T) {
	resetRegistrationDataModel()
	assert.True(t, registrationData.hasCompleteData(), "Registration data is complete when all attributes has value")

	registrationData.Fullname = ""
	assert.False(t, registrationData.hasCompleteData(), "Registration data should not be complete when fullname is empty")

	resetRegistrationDataModel()
	registrationData.Email = ""
	assert.False(t, registrationData.hasCompleteData(), "Registration data should not be complete when email is empty")

	resetRegistrationDataModel()
	registrationData.Password = ""
	assert.False(t, registrationData.hasCompleteData(), "Registration data should not be complete when password is empty")

	resetRegistrationDataModel()
	registrationData.PasswordConfirm = ""
	assert.False(t, registrationData.hasCompleteData(), "Registration data should not be complete when password confirm is empty")

	resetRegistrationDataModel()
}

func TestUserRegistrationHasMatchingPassword(t *testing.T) {
	resetRegistrationDataModel()
	assert.True(t, registrationData.hasMatchingPassword(), "Registration data should have matching password when Password==PasswordConfirm")

	registrationData.Password = "passwordchanged"
	assert.False(t, registrationData.hasMatchingPassword(), "Registration data should not have matching password when Password!=PasswordConfirm")

	resetRegistrationDataModel()
}
