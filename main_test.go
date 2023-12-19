// main_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/hello?name=John", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	expected := `{"text":"Hello, John!","FetchTime":"`
	if body := rr.Body.String(); body[:len(expected)] != expected {
		t.Errorf("Handler returned unexpected body: got %v, want %v", body[:len(expected)], expected)
	}
}
