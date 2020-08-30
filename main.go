package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/{userID}", getUserInfo).Methods(http.MethodGet)

	log.Printf("[INFO]Started Service on Port : %s", v1.port)
	log.Fatal(http.ListenAndServe(v1.port, r))

}
