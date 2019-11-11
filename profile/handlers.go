package profile

import (
	"net/http"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("middleware works :)"))
}
