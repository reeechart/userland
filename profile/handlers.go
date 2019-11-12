package profile

import (
	"encoding/json"
	"errors"
	"net/http"
	"userland/auth"
	"userland/response"
)

var err error

func GetProfile(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*auth.User)
	userProfile := UserProfile{
		Id:             user.Id,
		Fullname:       user.Fullname,
		Location:       user.Location.String,
		Bio:            user.Bio.String,
		Web:            user.Web.String,
		ProfilePicture: user.ProfilePicture,
		CreatedAt:      user.CreatedAt,
	}
	response.RespondSuccessWithBody(w, userProfile)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*auth.User)

	var userInfo UserProfile
	err = json.NewDecoder(r.Body).Decode(&userInfo)

	if err != nil {
		response.RespondBadRequest(w, REQUEST_BODY_UNDECODABLE, err)
		return
	}

	repo := getProfileRepository()
	err = repo.updateUserProfile(user, userInfo)

	if err != nil {
		response.RespondBadRequest(w, UNABLE_TO_UPDATE_PROFILE, err)
		return
	}

	response.RespondSuccess(w)
}

func GetEmail(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*auth.User)
	response.RespondSuccessWithBody(w, map[string]string{"email": user.Email})
}

func ChangeEmailAddress(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*auth.User)

	var emailReq ChangeEmailRequest
	err = json.NewDecoder(r.Body).Decode(&emailReq)

	if err != nil {
		response.RespondBadRequest(w, REQUEST_BODY_UNDECODABLE, err)
		return
	}

	repo := getProfileRepository()
	err = repo.changeUserEmail(user, emailReq.NewEmail)

	if err != nil {
		response.RespondBadRequest(w, UNABLE_TO_EXEC_UPDATE_EMAIL_QUERY, err)
		return
	}

	response.RespondSuccess(w)
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*auth.User)

	var passwordReq ChangePasswordRequest
	err = json.NewDecoder(r.Body).Decode(&passwordReq)

	if err != nil {
		response.RespondBadRequest(w, REQUEST_BODY_UNDECODABLE, err)
		return
	}

	if !passwordReq.hasMatchingNewPassword() {
		err = errors.New("Passwords don't match")
		response.RespondBadRequest(w, CHANGE_PASSWORD_PASSWORD_NOT_MATCH, err)
		return
	}

	repo := getProfileRepository()
	err = repo.changeUserPassword(user, passwordReq.PasswordCurrent, passwordReq.Password)

	if err != nil {
		response.RespondBadRequest(w, CHANGE_PASSWORD_INCORRECT_CURRENT_PASSWORD, err)
		return
	}

	response.RespondSuccess(w)
}
