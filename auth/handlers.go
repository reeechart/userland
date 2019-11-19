package auth

import (
	"encoding/json"
	"net/http"
	"time"
	ulanderrors "userland/errors"
	"userland/request"
	"userland/response"

	log "github.com/sirupsen/logrus"
)

var err error

type AuthHandler struct {
	UserRepo userRepositoryInterface
}

func (handler AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var userRegistrationData userRegistration
	err = request.ParseJSON(r.Body, &userRegistrationData)

	if err != nil {
		log.Info(err)
		response.RespondBadRequest(w, ulanderrors.ErrParseBody)
		return
	}

	if !userRegistrationData.hasCompleteData() {
		log.Info("User doesn't have complete data")
		response.RespondBadRequest(w, ulanderrors.ErrRegistrationIncomplete)
		return
	}

	if !userRegistrationData.hasValidData() {
		log.Info("User registration data is invalid")
		response.RespondBadRequest(w, ulanderrors.ErrRegistrationInvalid)
		return
	}

	if !userRegistrationData.hasMatchingPassword() {
		log.Info("User registration data has unmatching passwords")
		response.RespondBadRequest(w, ulanderrors.ErrRegistrationUnmatchingPassword)
		return
	}

	err = handler.UserRepo.createNewUser(userRegistrationData)

	if err != nil {
		log.Info(err)
		response.RespondBadRequest(w, ulanderrors.ErrRegistrationQueryExec)
		return
	}

	log.Info("User registration successful")
	response.RespondSuccess(w)
}

func (handler AuthHandler) Verify(w http.ResponseWriter, r *http.Request) {
	var verifReq verificationRequest
	err = json.NewDecoder(r.Body).Decode(&verifReq)

	if err != nil {
		log.Info(err)
		response.RespondBadRequest(w, ulanderrors.ErrParseBody)
		return
	}

	if !verifReq.isValid() {
		log.Info("Verification request data is invalid")
		response.RespondBadRequest(w, ulanderrors.ErrVerificationIncomplete)
		return
	}

	err = handler.UserRepo.verifyUser(verifReq.Recipient, verifReq.VerificationToken)

	if err != nil {
		log.Info(err)
		response.RespondBadRequest(w, ulanderrors.ErrVerificationQueryExec)
		return
	}

	log.Info("Verification successful")
	response.RespondSuccess(w)
}

func (handler AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginUser User
	err = request.ParseJSON(r.Body, &loginUser)

	if err != nil {
		log.Info(err)
		response.RespondBadRequest(w, ulanderrors.ErrParseBody)
		return
	}

	if !loginUser.ableToLogin() {
		log.Info("Login data is incomplete")
		response.RespondBadRequest(w, ulanderrors.ErrLoginIncomplete)
		return
	}

	err = handler.UserRepo.loginUser(loginUser.Email, loginUser.Password)

	if err != nil {
		log.Info(err)
		response.RespondUnauthorized(w, ulanderrors.ErrLoginUnmatch)
		return
	}

	user, _ := handler.UserRepo.getUserByEmail(loginUser.Email)
	if !user.Verified {
		log.Info("User hasn't been verified by the system")
		response.RespondUnauthorized(w, ulanderrors.ErrLoginUnverified)
		return
	}

	expirationTime := time.Now().Add(HOURS_IN_DAY * time.Hour)
	token, err := generateJWT(*user, expirationTime)
	if err != nil {
		log.Info(err)
		response.RespondInternalError(w, ulanderrors.ErrLoginJWT)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expirationTime,
	})

	log.Info("Login successful")
	response.RespondSuccessWithBody(w, map[string]bool{"require_tfa": false})
}

func (handler AuthHandler) ForgetPassword(w http.ResponseWriter, r *http.Request) {
	var user User
	err = request.ParseJSON(r.Body, &user)

	if err != nil {
		log.Info(err)
		response.RespondBadRequest(w, ulanderrors.ErrParseBody)
		return
	}

	if user.Email == "" {
		log.Info("User email is empty")
		response.RespondBadRequest(w, ulanderrors.ErrForgetPassIncomplete)
		return
	}

	err = handler.UserRepo.forgetPassword(user.Email)

	if err != nil {
		log.Info(err)
		response.RespondBadRequest(w, ulanderrors.ErrForgetPassQueryExec)
		return
	}

	log.Info("Forget password execution successful")
	response.RespondSuccess(w)
}

func (handler AuthHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var req resetPasswordRequest
	err = request.ParseJSON(r.Body, &req)

	if err != nil {
		log.Info(err)
		response.RespondBadRequest(w, ulanderrors.ErrParseBody)
		return
	}

	if !req.isValid() {
		log.Info("Reset password request is invalid")
		response.RespondBadRequest(w, ulanderrors.ErrResetPassInvalid)
		return
	}

	if !req.hasValidPassword() {
		log.Info("Reset password request's password is invalid")
		response.RespondBadRequest(w, ulanderrors.ErrResetPassInvalid)
		return
	}

	if !req.hasMatchingPassword() {
		log.Info("Reset password request has unmatching passwords")
		response.RespondBadRequest(w, ulanderrors.ErrResetPassUnmatchPass)
		return
	}

	err = handler.UserRepo.resetPassword(req.Token, req.Password)

	if err != nil {
		log.Info(err)
		response.RespondBadRequest(w, ulanderrors.ErrResetPassQueryExec)
		return
	}

	log.Info("Reset password successful")
	response.RespondSuccess(w)
}
