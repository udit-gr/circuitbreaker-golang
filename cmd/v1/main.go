package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	v1 "circuitbreaker-golang/api/v1"
	lib "circuitbreaker-golang/internal/lib"

	"github.com/gorilla/mux"
)

// getUserInfo return info for the user-id in request payload
func getUserInfo(w http.ResponseWriter, r *http.Request) v1.HTTPResponse {

	var (
		userID       int
		err          error
		httpResponse v1.HTTPResponse
	)

	userID, err = strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		log.Printf("[ERR][getUserInfo] Error while converting user-id : %s to int, Err : %v", r.FormValue("user_id"), err)
		httpResponse.Data = v1.ErrorMessage{Message: v1.MesssageInvalidUserID}
		httpResponse.Error = err
		return httpResponse
	}

	log.Printf("request came to the function")

	httpResponse.Data = v1.UserInfo{Message: userID}
	httpResponse.Error = nil

	return httpResponse
}

func main() {
	r := mux.NewRouter()

	cb := lib.StartNewCircuitBreaker(10, 1, time.Second*2)
	apiV1ServiceContext := &lib.ServiceContext{CircuitBreaker: cb}

	r.Handle("/api/v1", lib.ServiceHandler{apiV1ServiceContext, getUserInfo}).Methods("GET")

	log.Printf("[INFO]Service listening on Port : %s", v1.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", v1.Port), r))

}
