package router

import (
	"userland/auth"
	"userland/ping"

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
}
