package main

import (
	"fmt"
	"log"
	"net/http"

	"userland/appcontext"
	"userland/router"
)

func main() {
	appcontext.InitContext()
	router := router.GetRouter()

	fmt.Println("Server is listening at 8080")
	log.Println(http.ListenAndServe(":8080", router))
}
