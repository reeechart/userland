package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
	"userland/response"
)

var err error

func Register(w http.ResponseWriter, r *http.Request) {
	var userRegistrationData userRegistration
	err = json.NewDecoder(r.Body).Decode(&userRegistrationData)

	if err != nil {
		response.RespondBadRequest(w, REQUEST_BODY_UNDECODABLE, err)
		return
	}

	if !userRegistrationData.hasCompleteData() {
		err = errors.New("Registration data incomplete")
		response.RespondBadRequest(w, REGISTRATION_BODY_INCOMPLETE, err)
		return
	}

	if !userRegistrationData.hasMatchingPassword() {
		err = errors.New("Passwords don't match")
		response.RespondBadRequest(w, REGISTRATION_PASSWORD_NOT_MATCH, err)
		return
	}

	userRepo := getUserRepository()
	err = userRepo.createNewUser(userRegistrationData)

	if err != nil {
		response.RespondBadRequest(w, REGISTRATION_UNABLE_TO_EXEC_QUERY, err)
		return
	}

	response.RespondSuccess(w)
}

func Verify(w http.ResponseWriter, r *http.Request) {
	var verifReq verificationRequest
	err = json.NewDecoder(r.Body).Decode(&verifReq)

	if err != nil {
		response.RespondBadRequest(w, REQUEST_BODY_UNDECODABLE, err)
		return
	}

	if !verifReq.isValid() {
		err = errors.New("Verification request incomplete")
		response.RespondBadRequest(w, VERIFICATION_BODY_INCOMPLETE, err)
		return
	}

	userRepo := getUserRepository()
	err = userRepo.verifyUser(verifReq.Recipient, verifReq.VerificationToken)

	if err != nil {
		response.RespondBadRequest(w, VERIFICATION_UNABLE_TO_EXEC_QUERY, err)
		return
	}

	response.RespondSuccess(w)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginUser User
	err = json.NewDecoder(r.Body).Decode(&loginUser)

	if err != nil {
		response.RespondBadRequest(w, REQUEST_BODY_UNDECODABLE, err)
		return
	}

	if !loginUser.ableToLogin() {
		err = errors.New("Incomplete provided credentials")
		response.RespondBadRequest(w, LOGIN_INCOMPLETE_CREDENTIALS, err)
		return
	}

	userRepo := getUserRepository()
	err = userRepo.loginUser(loginUser.Email, loginUser.Password)

	if err != nil {
		response.RespondUnauthorized(w, LOGIN_PASSWORD_NOT_MATCH_OR_UNVERIFIED, err)
		return
	}

	user, _ := userRepo.getUserByEmail(loginUser.Email)
	expirationTime := time.Now().Add(HOURS_IN_DAY * time.Hour)
	token, err := generateJWT(*user, expirationTime)
	if err != nil {
		response.RespondInternalError(w, LOGIN_JWT_ERROR, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expirationTime,
	})

	response.RespondSuccessWithBody(w, map[string]bool{"require_tfa": false})
}

func ForgetPassword(w http.ResponseWriter, r *http.Request) {
	var user User
	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		response.RespondBadRequest(w, REQUEST_BODY_UNDECODABLE, err)
		return
	}

	if user.Email == "" {
		err = errors.New("Incomplete credentials to forget password")
		response.RespondBadRequest(w, FORGET_PASSWORD_INCOMPLETE_CREDENTIALS, err)
		return
	}

	userRepo := getUserRepository()
	err = userRepo.forgetPassword(user.Email)

	if err != nil {
		err = errors.New("Error on executing query")
		response.RespondBadRequest(w, FORGET_PASSWORD_UNABLE_TO_EXEC_QUERY, err)
		return
	}

	response.RespondSuccess(w)
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	var req resetPasswordRequest
	err = json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		response.RespondBadRequest(w, REQUEST_BODY_UNDECODABLE, err)
		return
	}

	if !req.hasMatchingPassword() {
		err = errors.New("Passwords don't match")
		response.RespondBadRequest(w, RESET_PASSWORD_PASSWORD_NOT_MATCH, err)
		return
	}

	userRepo := getUserRepository()
	err = userRepo.resetPassword(req.Token, req.Password)

	if err != nil {
		response.RespondBadRequest(w, RESET_PASSWORD_UNABLE_TO_EXEC_QUERY, err)
		return
	}

	response.RespondSuccess(w)
}
