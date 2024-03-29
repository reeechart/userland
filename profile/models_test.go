package profile

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	userProfile   UserProfile
	changePassReq ChangePasswordRequest
)

const (
	STR_LEN_LESS_THAN_3   = "na"
	STR_LEN_LESS_THAN_6   = "aaaaa"
	STR_LEN_MORE_THAN_128 = "NAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMENAMEA"
	INVALID_WEB           = "htt://example.co"
)

func resetUserProfileModel() {
	userProfile = UserProfile{
		Id:             1,
		Fullname:       "user",
		Location:       "Jakarta, Indonesia",
		Bio:            "hello",
		Web:            "https://example.com",
		ProfilePicture: nil,
		CreatedAt:      time.Now(),
	}
}

func resetChangePasswordRequestModel() {
	changePassReq = ChangePasswordRequest{
		PasswordCurrent: "password",
		Password:        "password",
		PasswordConfirm: "password",
	}
}

func TestUserProfileValidity(t *testing.T) {
	resetUserProfileModel()
	assert.True(t, userProfile.hasValidProfile(), "User profile model should be valid when its attributes are valid")

	userProfile.Fullname = STR_LEN_LESS_THAN_3
	assert.False(t, userProfile.hasValidProfile(), "User profile should not be valid when fullname is shorter than 3")
	userProfile.Fullname = STR_LEN_MORE_THAN_128
	assert.False(t, userProfile.hasValidProfile(), "User profile should not be valid when fullname is longer than 128")

	resetUserProfileModel()
	userProfile.Location = STR_LEN_MORE_THAN_128
	assert.False(t, userProfile.hasValidProfile(), "User profile should not be valid when location is longer than 128")

	resetUserProfileModel()
	userProfile.Bio = STR_LEN_MORE_THAN_128 + STR_LEN_MORE_THAN_128
	assert.False(t, userProfile.hasValidProfile(), "User profile should not be valid when bio is longer than 255")

	resetUserProfileModel()
	userProfile.Web = STR_LEN_MORE_THAN_128
	assert.False(t, userProfile.hasValidProfile(), "User profile should not be valid when web is longer than 128")
	userProfile.Web = INVALID_WEB
	assert.False(t, userProfile.hasValidProfile(), "User profile should not be valid when web does not match its regex")

	resetUserProfileModel()
}

func TestChangeEmailRequestValidity(t *testing.T) {
	emailReq := ChangeEmailRequest{
		NewEmail: "user@example.com",
	}
	assert.True(t, emailReq.hasValidEmail(), "Change email request is valid when email is valid")

	emailReq.NewEmail = "userexample.com"
	assert.False(t, emailReq.hasValidEmail(), "Change email request is invalid when email local is invalid")

	emailReq.NewEmail = "user@examplecom"
	assert.False(t, emailReq.hasValidEmail(), "Change email request is invalid when email domain is invalid")
}

func TestChangePasswordRequestHasMatchingPassword(t *testing.T) {
	resetChangePasswordRequestModel()
	assert.True(t, changePassReq.hasMatchingNewPassword(), "Change password request should be valid when Password==PasswordConfirm")

	changePassReq.Password = "passwordchanged"
	assert.False(t, changePassReq.hasMatchingNewPassword(), "Change password request should not be valid when Password!=PasswordConfirm")
}

func TestChangePasswordRequestHasValidPassword(t *testing.T) {
	resetChangePasswordRequestModel()
	assert.True(t, changePassReq.hasValidPassword(), "Change password request should be valid when password has valid length")

	changePassReq.Password = STR_LEN_LESS_THAN_6
	assert.False(t, changePassReq.hasValidPassword(), "Change password request should not be valid when password is shorter than 6")

	resetChangePasswordRequestModel()
	changePassReq.Password = STR_LEN_MORE_THAN_128
	assert.False(t, changePassReq.hasValidPassword(), "Change password request should not be valid when password is longer than 128")

	resetChangePasswordRequestModel()
}
