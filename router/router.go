package router

import (
	"net/http"
	"userland/auth"
	"userland/ping"
	"userland/profile"

	"github.com/gorilla/mux"
)

var (
	authHandler auth.AuthHandler
)

func GetRouter() *mux.Router {
	router := mux.NewRouter()

	initHandler()
	setupRouteHandler(router)

	return router
}

func initHandler() {
	authHandler = auth.AuthHandler{UserRepo: auth.GetUserRepository()}
}

func setupRouteHandler(router *mux.Router) {
	router.HandleFunc("/api/ping", ping.Ping).Methods(http.MethodGet)
	router.HandleFunc("/api/auth/register", authHandler.Register).Methods(http.MethodPost)
	router.HandleFunc("/api/auth/verification", authHandler.Verify).Methods(http.MethodPost)
	router.HandleFunc("/api/auth/login", authHandler.Login).Methods(http.MethodPost)
	router.HandleFunc("/api/auth/password/forgot", authHandler.ForgetPassword).Methods(http.MethodPost)
	router.HandleFunc("/api/auth/password/reset", authHandler.ResetPassword).Methods(http.MethodPost)

	router.HandleFunc("/api/me", auth.WithVerifyJWT(profile.GetProfile)).Methods(http.MethodGet)
	router.HandleFunc("/api/me", auth.WithVerifyJWT(profile.UpdateProfile)).Methods(http.MethodPut)
	router.HandleFunc("/api/me/email", auth.WithVerifyJWT(profile.GetEmail)).Methods(http.MethodGet)
	router.HandleFunc("/api/me/email", auth.WithVerifyJWT(profile.ChangeEmailAddress)).Methods(http.MethodPut)
	router.HandleFunc("/api/me/password", auth.WithVerifyJWT(profile.ChangePassword)).Methods(http.MethodPost)
	router.HandleFunc("/api/me/delete", auth.WithVerifyJWT(profile.DeleteAccount)).Methods(http.MethodPost)
	router.HandleFunc("/api/me/picture", auth.WithVerifyJWT(profile.UpdateProfilePicture)).Methods(http.MethodPut)
	router.HandleFunc("/api/me/picture", auth.WithVerifyJWT(profile.DeleteProfilePicture)).Methods(http.MethodDelete)
}
