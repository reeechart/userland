package main

import (
	"net/http"

	"userland/appcontext"
	"userland/router"

	log "github.com/sirupsen/logrus"
)

func main() {
	appcontext.InitContext()
	router := router.GetRouter()

	log.Info("Server is listening at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
