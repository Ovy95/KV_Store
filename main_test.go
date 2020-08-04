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

// Check tests are passing for the right reason !

// func PostCheckHandler(t *testing.T) {

// 	req, err := http.NewRequest("POST", "/hello", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rr := httptest.NewRecorder()
// 	KvHandlers := newHandlers()
// 	handler := http.HandlerFunc(KvHandlers.SortData)

// 	handler.ServeHTTP(rr, req)

// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// 	get, err := http.NewRequest("GET", "/hello", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	getRR := httptest.NewRecorder()
// 	handler.ServeHTTP(getRR, get)

// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}
// 	expected := ""
// 	if getRR.Body.String() != expected {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			rr.Body.String(), expected)
// 	}
// }

// func PostPostGetCheckHandler(t *testing.T) {

// 	req, err := http.NewRequest("POST", "/hello", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	req2, err := http.NewRequest("POST", "/world", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	postRR := httptest.NewRecorder()

// 	KvHandlers := newHandlers()
// 	handler := http.HandlerFunc(KvHandlers.SortData)

// 	handler.ServeHTTP(postRR, req)
// 	handler.ServeHTTP(postRR, req2)

// 	getreq, err := http.NewRequest("GET", "/", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	getRR := httptest.NewRecorder()
// 	handler.ServeHTTP(getRR, getreq)

// 	if status := getRR.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}
// 	expected := ""
// 	if getRR.Body.String() != expected {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			getRR.Body.String(), expected)
// 	}
// }
