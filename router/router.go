package router

import (
	"userland/auth"
	"userland/ping"
	"userland/profile"

	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	router := mux.NewRouter()

	setupRouteHandler(router)

	return router
}

func setupRouteHandler(router *mux.Router) {
	router.HandleFunc("/api/ping", ping.Ping).Methods("GET")
	router.HandleFunc("/api/auth/register", auth.Register).Methods("POST")
	router.HandleFunc("/api/auth/verification", auth.Verify).Methods("POST")
	router.HandleFunc("/api/auth/login", auth.Login).Methods("POST")
	router.HandleFunc("/api/auth/password/forgot", auth.ForgetPassword).Methods("POST")
	router.HandleFunc("/api/auth/password/reset", auth.ResetPassword).Methods("POST")

	router.HandleFunc("/api/me", auth.WithVerifyJWT(profile.GetProfile)).Methods("GET")
	router.HandleFunc("/api/me", auth.WithVerifyJWT(profile.UpdateProfile)).Methods("PUT")
	router.HandleFunc("/api/me/email", auth.WithVerifyJWT(profile.GetEmail)).Methods("GET")
	router.HandleFunc("/api/me/email", auth.WithVerifyJWT(profile.ChangeEmailAddress)).Methods("PUT")
}
