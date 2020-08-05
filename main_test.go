package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/string", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	KvHandlers := newHandlers()
	handler := http.HandlerFunc(KvHandlers.SortData)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}

	expected := "No Key could be found, try http://localhost:8080/ to see the list of stored keys"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestSortdataMethodHandler(t *testing.T) {
	// This is the check the error if the method isnt equal to the three cases
	put, err := http.NewRequest("PUT", "/", http.NoBody)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	KvHandlers := newHandlers()
	handler := http.HandlerFunc(KvHandlers.SortData)

	handler.ServeHTTP(rr, put)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusMethodNotAllowed)
	}

	expected := "method not allowed"
	if rr.Body.String() != expected {
		t.Errorf("handler returned wrong status code: got %s want %s",
			rr.Body, expected)
	}
}
