package native

import (
	"log"
	"net/http"
	"strings"

	"github.com/kashyaprahul94/go-playground/pkg/web/common"
)

// handler is basic handler
type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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

// GetHandler returns the router instance for all the paths handled
func GetHandler() common.Server {

	h := &handler{}

	return h
}

// StartServer will start listening
func StartServer(port string) {

	addr := []string{"", port}

	http.Handle("/", GetHandler())

	log.Printf("> Server stated ---> http://localhost:%v", port)

	log.Fatal(http.ListenAndServe(strings.Join(addr, ":"), nil))
}
