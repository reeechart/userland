package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"userland/response"
)

var err error

func Register(w http.ResponseWriter, r *http.Request) {
	var userRegistrationData userRegistration
	err = json.NewDecoder(r.Body).Decode(&userRegistrationData)
	defer r.Body.Close()

	if err != nil {
		response.RespondWithError(w, REGISTRATION_BODY_UNDECODABLE, err)
		return
	}

	if !userRegistrationData.hasValidData() {
		err = errors.New("Registration data incomplete")
		response.RespondWithError(w, REGISTRATION_BODY_INCOMPLETE, err)
		return
	}

	if !userRegistrationData.hasMatchingPassword() {
		err = errors.New("Passwords doesn't match")
		response.RespondWithError(w, REGISTRATION_PASSWORD_NOT_MATCH, err)
		return
	}

	userRepo := getUserRepository()
	err = userRepo.createNewUser(userRegistrationData)

	if err != nil {
		response.RespondWithError(w, REGISTRATION_UNABLE_TO_EXEC_QUERY, err)
		return
	}

	response.RespondSuccess(w)
}

func Verify(w http.ResponseWriter, r *http.Request) {
	var verifReq verificationRequest
	err = json.NewDecoder(r.Body).Decode(&verifReq)
	defer r.Body.Close()

	if !verifReq.isValid() {
		err = errors.New("Verification request incomplete")
		response.RespondWithError(w, VERIFICATION_BODY_INCOMPLETE, err)
		return
	}

	userRepo := getUserRepository()
	err = userRepo.verifyUser(verifReq.Recipient)

	if err != nil {
		response.RespondWithError(w, VERIFICATION_UNABLE_TO_EXEC_QUERY, err)
		return
	}

	response.RespondSuccess(w)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginUser User
	err = json.NewDecoder(r.Body).Decode(&loginUser)
	defer r.Body.Close()

	if !loginUser.ableToLogin() {
		err = errors.New("Email must not be empty")
		response.RespondWithError(w, LOGIN_EMAIL_NOT_PROVIDED, err)
		return
	}

	userRepo := getUserRepository()
	err = userRepo.loginUser(loginUser.Email, loginUser.Password)

	if err != nil {
		response.RespondWithError(w, LOGIN_PASSWORD_DOES_NOT_MATCH, err)
		return
	}

	response.RespondSuccess(w)
}
