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
