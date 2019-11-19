package auth

import (
	"encoding/json"
	"net/http"
	"time"
	ulanderrors "userland/errors"
	"userland/response"
)

var err error

type AuthHandler struct {
	UserRepo userRepositoryInterface
}

func (handler AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var userRegistrationData userRegistration
	err = json.NewDecoder(r.Body).Decode(&userRegistrationData)

	if err != nil {
		response.RespondBadRequest(w, ulanderrors.ErrParseBody)
		return
	}

	if !userRegistrationData.hasCompleteData() {
		response.RespondBadRequest(w, ulanderrors.ErrRegistrationIncomplete)
		return
	}

	if !userRegistrationData.hasValidData() {
		response.RespondBadRequest(w, ulanderrors.ErrRegistrationInvalid)
		return
	}

	if !userRegistrationData.hasMatchingPassword() {
		response.RespondBadRequest(w, ulanderrors.ErrRegistrationUnmatchingPassword)
		return
	}

	err = handler.UserRepo.createNewUser(userRegistrationData)

	if err != nil {
		response.RespondBadRequest(w, ulanderrors.ErrRegistrationQueryExec)
		return
	}

	response.RespondSuccess(w)
}

func (handler AuthHandler) Verify(w http.ResponseWriter, r *http.Request) {
	var verifReq verificationRequest
	err = json.NewDecoder(r.Body).Decode(&verifReq)

	if err != nil {
		response.RespondBadRequest(w, ulanderrors.ErrParseBody)
		return
	}

	if !verifReq.isValid() {
		response.RespondBadRequest(w, ulanderrors.ErrVerificationIncomplete)
		return
	}

	err = handler.UserRepo.verifyUser(verifReq.Recipient, verifReq.VerificationToken)

	if err != nil {
		response.RespondBadRequest(w, ulanderrors.ErrVerificationQueryExec)
		return
	}

	response.RespondSuccess(w)
}

func (handler AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginUser User
	err = json.NewDecoder(r.Body).Decode(&loginUser)

	if err != nil {
		response.RespondBadRequest(w, ulanderrors.ErrParseBody)
		return
	}

	if !loginUser.ableToLogin() {
		response.RespondBadRequest(w, ulanderrors.ErrLoginIncomplete)
		return
	}

	err = handler.UserRepo.loginUser(loginUser.Email, loginUser.Password)

	if err != nil {
		response.RespondUnauthorized(w, ulanderrors.ErrLoginUnmatchUnverified)
		return
	}

	user, _ := handler.UserRepo.getUserByEmail(loginUser.Email)
	expirationTime := time.Now().Add(HOURS_IN_DAY * time.Hour)
	token, err := generateJWT(*user, expirationTime)
	if err != nil {
		response.RespondInternalError(w, ulanderrors.ErrLoginJWT)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expirationTime,
	})

	response.RespondSuccessWithBody(w, map[string]bool{"require_tfa": false})
}

func (handler AuthHandler) ForgetPassword(w http.ResponseWriter, r *http.Request) {
	var user User
	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		response.RespondBadRequest(w, ulanderrors.ErrParseBody)
		return
	}

	if user.Email == "" {
		response.RespondBadRequest(w, ulanderrors.ErrForgetPassIncomplete)
		return
	}

	err = handler.UserRepo.forgetPassword(user.Email)

	if err != nil {
		response.RespondBadRequest(w, ulanderrors.ErrForgetPassQueryExec)
		return
	}

	response.RespondSuccess(w)
}

func (handler AuthHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var req resetPasswordRequest
	err = json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		response.RespondBadRequest(w, ulanderrors.ErrParseBody)
		return
	}

	if !req.isValid() {
		response.RespondBadRequest(w, ulanderrors.ErrResetPassInvalid)
		return
	}

	if !req.hasValidPassword() {
		response.RespondBadRequest(w, ulanderrors.ErrResetPassInvalid)
		return
	}

	if !req.hasMatchingPassword() {
		response.RespondBadRequest(w, ulanderrors.ErrResetPassUnmatchPass)
		return
	}

	err = handler.UserRepo.resetPassword(req.Token, req.Password)

	if err != nil {
		response.RespondBadRequest(w, ulanderrors.ErrResetPassQueryExec)
		return
	}

	response.RespondSuccess(w)
}
