package main

import (
	"fmt"
	"log"
	"net/http"

	v1 "circuitbreaker-golang/api/v1"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/{userID}", v1.GetUserInfo).Methods(http.MethodGet)

	log.Printf("[INFO]Service listening on Port : %s", v1.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", v1.Port), r))

}
