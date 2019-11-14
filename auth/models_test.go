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
	verifRequest     verificationRequest
	resetPassRequest resetPasswordRequest
)

const (
	STR_LEN_LESS_THAN_6   = "aaaaa"
	STR_LEN_MORE_THAN_128 = "NAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMEA"
	INVALID_EMAIL_SAMPLE  = "example@example/com"
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

func resetVerificationRequestModel() {
	verifRequest = verificationRequest{
		Type:              "email.verify",
		Recipient:         "user@example.com",
		VerificationToken: "newtokennewtokennewtokennewtoken",
	}
}

func resetResetPasswordRequestModel() {
	resetPassRequest = resetPasswordRequest{
		Token:           "newtokennewtokennewtokennewtoken",
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

func TestUserRegistrationValidity(t *testing.T) {
	resetRegistrationDataModel()
	assert.True(t, registrationData.hasValidData(), "Registration data should have valid fullname, email, and password")

	registrationData.Fullname = STR_LEN_MORE_THAN_128
	assert.False(t, registrationData.hasValidData(), "Registration data should not be valid when fullname is longer than 128")

	resetRegistrationDataModel()
	registrationData.Email = STR_LEN_MORE_THAN_128
	assert.False(t, registrationData.hasValidData(), "Registration data should not be valid when email length is longer than 128")
	registrationData.Email = INVALID_EMAIL_SAMPLE
	assert.False(t, registrationData.hasValidData(), "Registration data should not be valid when email does not match regex")

	resetRegistrationDataModel()
	registrationData.Password = STR_LEN_LESS_THAN_6
	assert.False(t, registrationData.hasValidData(), "Registration data should not be valid when password is shorter than 6")
	registrationData.PasswordConfirm = STR_LEN_MORE_THAN_128
	assert.False(t, registrationData.hasValidData(), "Registration data should not be valid when password is longer than 128")

	resetRegistrationDataModel()
}

func TestUserVerificationRequestValidity(t *testing.T) {
	resetVerificationRequestModel()
	assert.True(t, verifRequest.isValid(), "Verification request is valid when data is complete")

	verifRequest.Type = ""
	assert.False(t, verifRequest.isValid(), "Verification request should not be valid when verification type is empty")

	resetVerificationRequestModel()
	verifRequest.Recipient = ""
	assert.False(t, verifRequest.isValid(), "Verification request should not be valid when recipient is empty")

	resetVerificationRequestModel()
	verifRequest.VerificationToken = ""
	assert.False(t, verifRequest.isValid(), "Verification request should not be valid when verification token is empty")

	resetVerificationRequestModel()
}

func TestResetPasswordRequestHasValidPassword(t *testing.T) {
	resetResetPasswordRequestModel()
	assert.True(t, resetPassRequest.hasValidPassword(), "Reset password request should be valid when password length is between 6-128")

	resetPassRequest.Password = STR_LEN_LESS_THAN_6
	assert.False(t, resetPassRequest.hasValidPassword(), "Reset password request should not be valid when password is shorter than 6")

	resetResetPasswordRequestModel()
	resetPassRequest.Password = STR_LEN_MORE_THAN_128
	assert.False(t, resetPassRequest.hasValidPassword(), "Reset password request should not be valid when password is longer than 128")

	resetResetPasswordRequestModel()
}

func TestResetPasswordRequestHasMatchingPassword(t *testing.T) {
	resetResetPasswordRequestModel()
	assert.True(t, resetPassRequest.hasMatchingPassword(), "Reset password request should have matching password when Password==PasswordConfirm")

	resetPassRequest.Password = "passwordchanged"
	assert.False(t, resetPassRequest.hasMatchingPassword(), "Reset password request should not have matching password when Password!=PasswordConfirm")

	resetResetPasswordRequestModel()
}
