package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	// all endpoints
	router.HandleFunc("/", welcomeToAnimalPocketApi)
	router.HandleFunc("/fish/{id}", getOneFishById).Methods("GET")
	router.HandleFunc("/bug/{id}", getOneBugById).Methods("GET")

	// start the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", router))
}

func welcomeToAnimalPocketApi(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome!")
	fmt.Fprintf(w, "Welcome to the Animal Pocket Api!")
}
