package profile

import (
	"encoding/json"
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
		response.RespondBadRequest(w, 100, err)
		return
	}

	repo := getProfileRepository()
	err = repo.updateUserProfile(user, userInfo)

	if err != nil {
		response.RespondBadRequest(w, 201, err)
		return
	}

	response.RespondSuccess(w)
}
