package web

import (
	"net/http"
	"net/http/httptest"
	"testing"
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

func TestHealthCheckHandler(t *testing.T) {

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

	server := &Server{}
	handler := http.HandlerFunc(server.ServeHTTP)

	for _, tt := range webServerTests {

		t.Run(tt.name, func(t *testing.T) {

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, tt.request)

			assertStatusCode(t, tt.expectedStatus, rr.Code)
			assertBody(t, tt.expectedBody, rr.Body.String())
		})
	}

}
