package profile

import (
	"fmt"
	"net/http"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	fmt.Println(user)
	w.Write([]byte("success"))
}
