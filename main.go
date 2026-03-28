package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang-microservice/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
