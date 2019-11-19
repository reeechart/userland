package profile

import (
	"net/http"
	"userland/auth"
	ulanderrors "userland/errors"
	"userland/request"
	"userland/response"

	log "github.com/sirupsen/logrus"
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
	log.Info("Get user profile successful")
	response.RespondSuccessWithBody(w, userProfile)
}

func (handler ProfileHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*auth.User)

	var userInfo UserProfile
	err = request.ParseJSON(r.Body, &userInfo)
	if err != nil {
		log.Info(err)
		response.RespondBadRequest(w, ulanderrors.ErrParseBody)
		return
	}

	if !userInfo.hasValidProfile() {
		log.Info("Updated user info is invalid")
		response.RespondBadRequest(w, ulanderrors.ErrUpdateProfileUserInfoInvalid)
		return
	}

	err = handler.ProfileRepo.updateUserProfile(user, userInfo)

	if err != nil {
		log.Info(err)
		response.RespondBadRequest(w, ulanderrors.ErrUpdateProfileQueryExec)
		return
	}

	log.Info("Update user profile successful")
	response.RespondSuccess(w)
}

func (handler ProfileHandler) GetEmail(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*auth.User)
	log.Info("Get user email successful")
	response.RespondSuccessWithBody(w, map[string]string{"email": user.Email})
}

func (handler ProfileHandler) ChangeEmailAddress(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*auth.User)

	var emailReq ChangeEmailRequest
	err = request.ParseJSON(r.Body, &emailReq)

	if err != nil {
		log.Info(err)
		response.RespondBadRequest(w, ulanderrors.ErrParseBody)
		return
	}

	if !emailReq.hasValidEmail() {
		log.Info("User email is invalid")
		response.RespondBadRequest(w, ulanderrors.ErrChangeEmailInvalidEmail)
		return
	}

	err = handler.ProfileRepo.changeUserEmail(user, emailReq.NewEmail)

	if err != nil {
		log.Info(err)
		response.RespondBadRequest(w, ulanderrors.ErrChangeEmailQueryExec)
		return
	}

	log.Info("Change email address successful")
	response.RespondSuccess(w)
}

func (handler ProfileHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*auth.User)

	var passwordReq ChangePasswordRequest
	err = request.ParseJSON(r.Body, &passwordReq)

	if err != nil {
		log.Info(err)
		response.RespondBadRequest(w, ulanderrors.ErrParseBody)
		return
	}

	if !passwordReq.hasValidPassword() {
		log.Info("User change request password has invalid password")
		response.RespondBadRequest(w, ulanderrors.ErrChangePasswordInvalidPassword)
		return
	}

	if !passwordReq.hasMatchingNewPassword() {
		log.Info("User change password request has unmatching passwords")
		response.RespondBadRequest(w, ulanderrors.ErrChangePasswordPasswordUnmatch)
		return
	}

	err = handler.ProfileRepo.changeUserPassword(user, passwordReq.PasswordCurrent, passwordReq.Password)

	if err != nil {
		log.Info(err)
		response.RespondBadRequest(w, ulanderrors.ErrChangePasswordIncorrectCurrentPass)
		return
	}

	log.Info("User change password successful")
	response.RespondSuccess(w)
}

func (handler ProfileHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*auth.User)

	var delReq DeleteAccountRequest
	err = request.ParseJSON(r.Body, &delReq)

	if err != nil {
		log.Info(err)
		response.RespondBadRequest(w, ulanderrors.ErrParseBody)
		return
	}

	err = handler.ProfileRepo.deleteUser(user, delReq.Password)

	if err != nil {
		log.Info(err)
		response.RespondBadRequest(w, ulanderrors.ErrDeleteAccountIncorrectPass)
		return
	}

	log.Info("User delete account successful")
	response.RespondSuccess(w)
}

func (handler ProfileHandler) UpdateProfilePicture(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*auth.User)
	file, fileHeader, err := r.FormFile("file")

	if err != nil {
		log.Info(err)
		response.RespondBadRequest(w, ulanderrors.ErrUpdatePicturePicCantBeFetched)
		return
	}

	defer file.Close()

	picture := make([]byte, fileHeader.Size)
	_, err = file.Read(picture)

	if err != nil {
		log.Info(err)
		response.RespondBadRequest(w, ulanderrors.ErrUpdatePictureCantBeRead)
		return
	}

	err = handler.ProfileRepo.updateUserPicture(user, picture)

	if err != nil {
		log.Info(err)
		response.RespondBadRequest(w, ulanderrors.ErrUpdatePictureQueryExec)
		return
	}

	log.Info("Update profile picture successful")
	response.RespondSuccess(w)
}

func (handler ProfileHandler) DeleteProfilePicture(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*auth.User)

	err = handler.ProfileRepo.deleteUserPicture(user)

	if err != nil {
		log.Info(err)
		response.RespondBadRequest(w, ulanderrors.ErrUpdatePictureQueryExec)
		return
	}

	log.Info("Delete profile picture successful")
	response.RespondSuccess(w)
}
