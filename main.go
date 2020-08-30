package main

import (
	"log"
	"net/http"

	v1 "circuitbreaker-golang/api/v1"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/{userID}", v1.GetUserInfo).Methods(http.MethodGet)

	log.Printf("[INFO]Started Service on Port : %s", v1.Port)
	log.Fatal(http.ListenAndServe(v1.Port, r))

}
