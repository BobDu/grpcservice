package middleware

type CircuitBreakerCallGet struct {
	Next callGetter
}