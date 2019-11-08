package router

import (
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
}
