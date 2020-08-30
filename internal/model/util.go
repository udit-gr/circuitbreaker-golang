package model

import (
	"time"

	"github.com/eapache/go-resiliency/breaker"
)

// StartNewCircuitBreaker starts new circuit breaker from a closed state with the input paramters
func StartNewCircuitBreaker(errorThreshold, successThreshold int, timeout time.Duration) *breaker.Breaker {

	breaker := breaker.New(errorThreshold, successThreshold, timeout)

	return breaker
}
