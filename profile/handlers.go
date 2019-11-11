package profile

import (
	"fmt"
	"net/http"
	"userland/auth"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*auth.User)
	userProfile := UserProfile{
		Id:             user.Id,
		Fullname:       user.Fullname,
		Location:       user.Location,
		Bio:            user.Bio,
		Web:            user.Web,
		ProfilePicture: user.ProfilePicture,
		CreatedAt:      user.CreatedAt,
	}
	fmt.Println(userProfile)
	w.Write([]byte("success"))
}
