package lib

import (
	"encoding/json"
	"log"
	"net"
	"net/http"

	v1 "circuitbreaker-golang/api/v1"

	"github.com/eapache/go-resiliency/breaker"
)

// ServiceContext defines all the variables/parameters which can be service specific e.g different circuit breaker
// for different API's, database pool, session store etc
type ServiceContext struct {
	CircuitBreaker *breaker.Breaker
}

// ServiceHandler used as a struct
// 1. Similar to httpHandler and implements ServeHTTP function
// 2. embedded field of ServiceContext to access the service specific variables
type ServiceHandler struct {
	*ServiceContext
	Handle func(w http.ResponseWriter, r *http.Request) v1.HTTPResponse
}

// Our custom ServeHTTP performs the similar operation as http's ServeHTTP the only difference being it can access our
// service specific parameters
func (sh ServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var data v1.HTTPResponse
	errStatus := http.StatusOK
	var bytesBuffer []byte
	var err error

	// checking if the circut breaker is available for the endpoint/api
	if sh.CircuitBreaker != nil {
		breakerResult := sh.CircuitBreaker.Run(func() error {
			data = sh.Handle(w, r)

			if err, ok := data.Error.(net.Error); ok && err.Timeout() {
				// checking for network error
				log.Printf("Network Error : %v", err)
				return err
			} else if data.Error != nil {
				return data.Error
			}
			return nil
		})

		if breakerResult != nil {
			// if error is due to circuit breaker being in open state the log and return error message for same
			if breakerResult == breaker.ErrBreakerOpen {
				log.Printf("Circuit Open : %v", breakerResult)
				errStatus = http.StatusServiceUnavailable
				bytesBuffer = []byte("Unable to process request as circuit is open")
			}
		}
	} else {
		data = sh.Handle(w, r)
	}

	// appedning proper headers in the response
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Accept-Charset", "utf-8")
	w.Header().Set("Access-Control-Allow-Headers", "X-Device, X-User-ID")

	if string(bytesBuffer) == "" {
		bytesBuffer, err = json.Marshal(data.Data)
		if err != nil {
			log.Printf("Error while unmarshalling API-Reponse interface : %+v, Err : %v", data.Data, err)
			errStatus = http.StatusInternalServerError
			bytesBuffer = []byte("Unable to process request, Error while unmarshalling the api-response")
		}
	}

	// appending byte response from the API/endpoint with the appropriate http status code
	w.WriteHeader(errStatus)
	w.Write(bytesBuffer)
}
