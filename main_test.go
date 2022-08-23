package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloServer(t *testing.T) {
	type args struct {
		n        string
		expected string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"world",
			args{"world", "Hello, world!"},
		},
		{
			"Fred",
			args{"Fred%20Flintstone", "Hello, Fred Flintstone!"},
		},
		{
			"Blank",
			args{"", "Hello, !"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/"+tt.args.n, nil)
			if err != nil {
				t.Fatal(err)
			}
			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(HelloServer)
			// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
			// directly and pass in our Request and ResponseRecorder.
			handler.ServeHTTP(rr, req)
			// Check the status code is what we expect.
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
			// Check the response body is what we expect.
			if rr.Body.String() != tt.args.expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), tt.args.expected)
			}
		})
	}
}
