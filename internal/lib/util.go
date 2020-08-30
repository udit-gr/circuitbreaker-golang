package lib

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/eapache/go-resiliency/breaker"
)

// StartNewCircuitBreaker starts new circuit breaker from a closed state with the input paramters
func StartNewCircuitBreaker(errorThreshold, successThreshold int, timeout time.Duration) *breaker.Breaker {

	breaker := breaker.New(errorThreshold, successThreshold, timeout)

	return breaker
}

// RequestGet makes GET request to the parameter endpoint using simple http-client pool
func RequestGet(url string, timeout time.Duration) (buffer []byte, err error) {

	var response *http.Response

	response, err = (&http.Client{Timeout: timeout}).Get(url)
	if err != nil {
		log.Printf("[ERR][RequestGet] Error while response from GET request on endpoint : %s, Err : %v", url, err)
		return
	}

	defer response.Body.Close()

	buffer, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("[ERR][RequestGet] Error while reading body response from GET request on endpoint : %s, Err : %v", url, err)
		return
	}

	return
}
