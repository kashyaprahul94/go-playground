package web

import (
	"log"
	"net/http"
	"strings"
)

// Server is an instance of http request handler
type Server struct{}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "get called"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "post called"}`))
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "put called"}`))
	case "DELETE":
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(`{"message": "delete called"}`))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"message": "method not allowed"}`))
	}
}

// StartServer will start listening
func StartServer() {
	s := &Server{}

	port := "4444"
	addr := []string{"", port}

	a := strings.Join(addr, ":")

	http.Handle("/", s)

	log.Fatal(http.ListenAndServe(a, nil))
}
