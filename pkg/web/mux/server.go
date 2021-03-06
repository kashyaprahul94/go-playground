package mux

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/kashyaprahul94/go-playground/pkg/web/common"
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(`{"message": "delete called"}`))
}

func notAllowed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte(`{"message": "method not allowed"}`))
}

// GetHandler returns the router instance for all the paths handled
func GetHandler() common.Server {

	r := mux.NewRouter()

	r.HandleFunc("/", get).Methods(http.MethodGet)
	r.HandleFunc("/", post).Methods(http.MethodPost)
	r.HandleFunc("/", put).Methods(http.MethodPut)
	r.HandleFunc("/", delete).Methods(http.MethodDelete)
	r.HandleFunc("/", notAllowed)

	return r
}

// StartServer will start listening
func StartServer(port string) {

	addr := []string{"", port}

	http.Handle("/", GetHandler())

	log.Printf("> Server stated ---> http://localhost:%v", port)

	log.Fatal(http.ListenAndServe(strings.Join(addr, ":"), nil))
}
