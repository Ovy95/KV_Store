package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetErrorMethod(t *testing.T) {

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

func TestSortDataDefaultCase(t *testing.T) {
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

func TestPost(t *testing.T) {
	bodyReader := strings.NewReader("Body Testing")

	req, err := http.NewRequest("POST", "/hello", bodyReader)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	KvHandlers := newHandlers()
	handler := http.HandlerFunc(KvHandlers.SortData)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	get, err := http.NewRequest("GET", "/hello", http.NoBody)
	if err != nil {
		t.Fatal(err)
	}

	getRR := httptest.NewRecorder()
	handler.ServeHTTP(getRR, get)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Body Testing"
	if getRR.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got%v want%v", getRR.Body.String(), expected)
	}
}

func TestDeletePost(t *testing.T) {
	bodyReader := strings.NewReader("Body Testing")

	req, err := http.NewRequest("POST", "/hello", bodyReader)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	KvHandlers := newHandlers()
	handler := http.HandlerFunc(KvHandlers.SortData)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	get, err := http.NewRequest("GET", "/hello", http.NoBody)
	if err != nil {
		t.Fatal(err)
	}

	getRR := httptest.NewRecorder()
	handler.ServeHTTP(getRR, get)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedBT := "Body Testing"
	if getRR.Body.String() != expectedBT {
		t.Errorf("handler returned unexpected body: got%v want%v", getRR.Body.String(), expectedBT)
	}

	delete, err := http.NewRequest("DELETE", "/hello", http.NoBody)
	if err != nil {
		t.Fatal(err)
	}

	deleteRR := httptest.NewRecorder()
	handler.ServeHTTP(deleteRR, delete)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedDeleted := ""
	if deleteRR.Body.String() != expectedDeleted {
		t.Errorf("handler returned unexpected body: got%v want%v", deleteRR.Body.String(), expectedDeleted)
	}
	/////
	req, err = http.NewRequest("GET", "/hello", http.NoBody)
	if err != nil {
		t.Fatal(err)
	}

	rrG := httptest.NewRecorder()
	handler.ServeHTTP(rrG, req)

	if status := rrG.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}

	expected := "No Key could be found, try http://localhost:8080/ to see the list of stored keys"
	if rrG.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rrG.Body.String(), expected)
	}

}
