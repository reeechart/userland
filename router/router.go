package router

import (
	"net/http"
	"userland/auth"
	"userland/ping"
	"userland/profile"

	"github.com/gorilla/mux"
)

var (
	authHandler    auth.AuthHandler
	authMiddleware auth.AuthMiddleware
	profileHandler profile.ProfileHandler
)

func GetRouter() *mux.Router {
	router := mux.NewRouter()

	initHandlersAndMiddlewares()
	setupRouteHandler(router)

	return router
}

func initHandlersAndMiddlewares() {
	authHandler = auth.AuthHandler{UserRepo: auth.GetUserRepository()}
	profileHandler = profile.ProfileHandler{ProfileRepo: profile.GetProfileRepository()}
	authMiddleware = auth.AuthMiddleware{UserRepo: auth.GetUserRepository()}
}

func setupRouteHandler(router *mux.Router) {
	router.HandleFunc("/api/ping", ping.Ping).Methods(http.MethodGet)
	router.HandleFunc("/api/auth/register", authHandler.Register).Methods(http.MethodPost)
	router.HandleFunc("/api/auth/verification", authHandler.Verify).Methods(http.MethodPost)
	router.HandleFunc("/api/auth/login", authHandler.Login).Methods(http.MethodPost)
	router.HandleFunc("/api/auth/password/forgot", authHandler.ForgetPassword).Methods(http.MethodPost)
	router.HandleFunc("/api/auth/password/reset", authHandler.ResetPassword).Methods(http.MethodPost)

	router.HandleFunc("/api/me", authMiddleware.WithVerifyJWT(profileHandler.GetProfile)).Methods(http.MethodGet)
	router.HandleFunc("/api/me", authMiddleware.WithVerifyJWT(profileHandler.UpdateProfile)).Methods(http.MethodPut)
	router.HandleFunc("/api/me/email", authMiddleware.WithVerifyJWT(profileHandler.GetEmail)).Methods(http.MethodGet)
	router.HandleFunc("/api/me/email", authMiddleware.WithVerifyJWT(profileHandler.ChangeEmailAddress)).Methods(http.MethodPut)
	router.HandleFunc("/api/me/password", authMiddleware.WithVerifyJWT(profileHandler.ChangePassword)).Methods(http.MethodPost)
	router.HandleFunc("/api/me/delete", authMiddleware.WithVerifyJWT(profileHandler.DeleteAccount)).Methods(http.MethodPost)
	router.HandleFunc("/api/me/picture", authMiddleware.WithVerifyJWT(profileHandler.UpdateProfilePicture)).Methods(http.MethodPut)
	router.HandleFunc("/api/me/picture", authMiddleware.WithVerifyJWT(profileHandler.DeleteProfilePicture)).Methods(http.MethodDelete)
}
