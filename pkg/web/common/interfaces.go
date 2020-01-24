package common

import "net/http"

// Server interface ensures that any implementation exposes needed methods
type Server interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}
