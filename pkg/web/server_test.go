package web

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kashyaprahul94/go-playground/pkg/web/common"
	"github.com/kashyaprahul94/go-playground/pkg/web/mux"
	"github.com/kashyaprahul94/go-playground/pkg/web/native"
)

func assertStatusCode(t *testing.T, expected, received int) {
	t.Helper()

	if expected != received {
		t.Errorf("handler returned wrong status code: got %v want %v",
			received, expected)
	}

}

func assertBody(t *testing.T, expected, received string) {
	t.Helper()

	if expected != received {
		t.Errorf("handler returned wrong body: got %v want %v",
			received, expected)
	}

}

func TestWebServer(t *testing.T) {

	getRequest, err := http.NewRequest("GET", "/", nil)
	postRequest, err := http.NewRequest("POST", "/", nil)
	putRequest, err := http.NewRequest("PUT", "/", nil)
	patchRequest, err := http.NewRequest("PATCH", "/", nil)
	deleteRequest, err := http.NewRequest("DELETE", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	webServerTests := []struct {
		name           string
		request        *http.Request
		expectedBody   string
		expectedStatus int
	}{
		{name: "GET", request: getRequest, expectedBody: `{"message": "get called"}`, expectedStatus: http.StatusOK},
		{name: "POST", request: postRequest, expectedBody: `{"message": "post called"}`, expectedStatus: http.StatusCreated},
		{name: "PUT", request: putRequest, expectedBody: `{"message": "put called"}`, expectedStatus: http.StatusAccepted},
		{name: "DELETE", request: deleteRequest, expectedBody: `{"message": "delete called"}`, expectedStatus: http.StatusNoContent},
		{name: "Not Implemented", request: patchRequest, expectedBody: `{"message": "method not allowed"}`, expectedStatus: http.StatusMethodNotAllowed},
	}

	strategies := []struct {
		name   string
		server common.Server
	}{
		{name: "Native", server: native.GetHandler()},
		{name: "Mux", server: mux.GetHandler()},
	}

	for _, s := range strategies {

		t.Run(s.name, func(ttt *testing.T) {

			handler := http.HandlerFunc(s.server.ServeHTTP)

			for _, tt := range webServerTests {

				testName := strings.Join([]string{s.name, tt.name}, "/")

				t.Run(testName, func(t *testing.T) {

					rr := httptest.NewRecorder()
					handler.ServeHTTP(rr, tt.request)

					assertStatusCode(t, tt.expectedStatus, rr.Code)
					assertBody(t, tt.expectedBody, rr.Body.String())
				})
			}
		})
	}

}
