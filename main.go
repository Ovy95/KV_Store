package main

import (
	"KV_STORE/store"
	"strings"

	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var st *store.Data

func HttpHandlerFunc(w http.ResponseWriter, r *http.Request) {

	url := getKeyFromURL(r.URL)
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()
	key := url
	value := body

	switch r.Method {

	case "GET":
		data, ok := store.Get(st, key)
		if ok {
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No Key could be found, try http://localhost:8080/ to see the list of stored keys"))
		}

	case "POST":
		store.Post(st, key, []byte(value))
		w.WriteHeader(http.StatusCreated)

	case "DELETE":
		store.Delete(st, key)
		w.WriteHeader(http.StatusOK)

	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func init() {
	st = &store.Data{
		Data: map[string][]byte{},
	}
}

func getKeyFromURL(url *url.URL) string {
	path := url.EscapedPath()
	return strings.TrimSpace(string(path[1:]))
}

func main() {

	http.HandleFunc("/", HttpHandlerFunc)

	fmt.Println("Listening at localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalln(err)
	}
}
