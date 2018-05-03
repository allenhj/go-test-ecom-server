package main

import (
	"log"
	"net/http"
	"os"
	"rest-and-go/store"

	"github.com/gorilla/handlers"
)

func main() {
	// Get the "PORT" env variable
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := store.NewRouter() // create routes

	//allowing front-end to utilize the API methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	// Launch server with CORS validations
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
