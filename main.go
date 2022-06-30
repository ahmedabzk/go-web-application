package main

import (
	"log"
	"net/http"

	router "github.com/ahmed/go-web/routers"
)

func main() {
	r := router.Router()

	log.Fatal(http.ListenAndServe(":8080",r))
}