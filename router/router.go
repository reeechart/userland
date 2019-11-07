package router

import (
	"userland/util"

	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	router := mux.NewRouter()

	setupRouteHandler(router)

	return router
}

func setupRouteHandler(router *mux.Router) {
	router.HandleFunc("/ping", util.Ping).Methods("GET")
}
