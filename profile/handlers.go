package profile

import (
	"encoding/json"
	"errors"
	"net/http"
	"userland/auth"
	"userland/response"
)

var err error

type ProfileHandler struct {
	ProfileRepo profileRepositoryInterface
}

func (handler ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
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

func (handler ProfileHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*auth.User)

	var userInfo UserProfile
	err = json.NewDecoder(r.Body).Decode(&userInfo)

	if err != nil {
		response.RespondBadRequest(w, REQUEST_BODY_UNDECODABLE, err)
		return
	}

	if !userInfo.hasValidProfile() {
		err = errors.New("Invalid user profile")
		response.RespondBadRequest(w, USER_INFO_INVALID, err)
		return
	}

	err = handler.ProfileRepo.updateUserProfile(user, userInfo)

	if err != nil {
		response.RespondBadRequest(w, UNABLE_TO_UPDATE_PROFILE, err)
		return
	}

	response.RespondSuccess(w)
}

func (handler ProfileHandler) GetEmail(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*auth.User)
	response.RespondSuccessWithBody(w, map[string]string{"email": user.Email})
}

func (handler ProfileHandler) ChangeEmailAddress(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*auth.User)

	var emailReq ChangeEmailRequest
	err = json.NewDecoder(r.Body).Decode(&emailReq)

	if err != nil {
		response.RespondBadRequest(w, REQUEST_BODY_UNDECODABLE, err)
		return
	}

	if !emailReq.hasValidEmail() {
		err = errors.New("Invalid email address")
		response.RespondBadRequest(w, EMAIL_INVALID, err)
		return
	}

	err = handler.ProfileRepo.changeUserEmail(user, emailReq.NewEmail)

	if err != nil {
		response.RespondBadRequest(w, UNABLE_TO_EXEC_UPDATE_EMAIL_QUERY, err)
		return
	}

	response.RespondSuccess(w)
}

func (handler ProfileHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*auth.User)

	var passwordReq ChangePasswordRequest
	err = json.NewDecoder(r.Body).Decode(&passwordReq)

	if err != nil {
		response.RespondBadRequest(w, REQUEST_BODY_UNDECODABLE, err)
		return
	}

	if !passwordReq.hasValidPassword() {
		err = errors.New("Password is invalid")
		response.RespondBadRequest(w, CHANGE_PASSWORD_PASSWORD_INVALID, err)
		return
	}

	if !passwordReq.hasMatchingNewPassword() {
		err = errors.New("Passwords don't match")
		response.RespondBadRequest(w, CHANGE_PASSWORD_PASSWORD_NOT_MATCH, err)
		return
	}

	err = handler.ProfileRepo.changeUserPassword(user, passwordReq.PasswordCurrent, passwordReq.Password)

	if err != nil {
		response.RespondBadRequest(w, CHANGE_PASSWORD_INCORRECT_CURRENT_PASSWORD, err)
		return
	}

	response.RespondSuccess(w)
}

func (handler ProfileHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*auth.User)

	var delReq DeleteAccountRequest
	err = json.NewDecoder(r.Body).Decode(&delReq)

	if err != nil {
		response.RespondBadRequest(w, REQUEST_BODY_UNDECODABLE, err)
		return
	}

	err = handler.ProfileRepo.deleteUser(user, delReq.Password)

	if err != nil {
		response.RespondBadRequest(w, DELETE_ACCOUNT_INCORRECT_PASSWORD, err)
		return
	}

	response.RespondSuccess(w)
}

func (handler ProfileHandler) UpdateProfilePicture(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*auth.User)
	file, fileHeader, err := r.FormFile("file")

	if err != nil {
		response.RespondBadRequest(w, PICTURE_CANNOT_BE_FETCHED_FROM_FORM, err)
		return
	}

	defer file.Close()

	picture := make([]byte, fileHeader.Size)
	_, err = file.Read(picture)

	if err != nil {
		response.RespondBadRequest(w, PICTURE_CANNOT_BE_READ, err)
		return
	}

	err = handler.ProfileRepo.updateUserPicture(user, picture)

	if err != nil {
		response.RespondBadRequest(w, PICTURE_FAILED_TO_EXEC_QUERY, err)
		return
	}

	response.RespondSuccess(w)
}

func (handler ProfileHandler) DeleteProfilePicture(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*auth.User)

	err = handler.ProfileRepo.deleteUserPicture(user)

	if err != nil {
		response.RespondBadRequest(w, PICTURE_FAILED_TO_EXEC_QUERY, err)
		return
	}

	response.RespondSuccess(w)
}
