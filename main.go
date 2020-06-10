package main

import (
	"fmt"
	"log"
	"net/http"
)

const set = "SET"
const get = "GET"
const post = "POST"
const delete = "DELETE"

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is starting on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case get:
		fmt.Fprintf(w, "THIS IS A %s method ON THE LOCALHOST AT THE INDEX%s\n", r.Method, r.URL)
		break
	case set:
		fmt.Fprintf(w, "THIS IS A %s method ON THE LOCALHOST AT THE INDEX%s\n", r.Method, r.URL)
		fmt.Println(r.Method)
		fmt.Fprintf(w, "Delete")
		break
	case post:
		fmt.Fprintf(w, "THIS IS A %s method ON THE LOCALHOST AT THE INDEX%s\n", r.Method, r.URL)
		fmt.Println(r.Method)
		fmt.Fprintf(w, "Post")
		break
	case delete:
		fmt.Fprintf(w, "THIS IS A %s method ON THE LOCALHOST AT THE INDEX%s\n", r.Method, r.URL)
		fmt.Println(r.Method)
		fmt.Fprintf(w, "Delete")
		break
	default:
		http.Error(w, "BadRequest status code", 400)
	}
}

func storeInterface() {
	The store interface should support the following operations:
	- Get(key string) []byte : This will return the value for the given key
	  or will return nil if it doesn’t exist. Your http handler should return
	  a NOT FOUND status code if that is the case
	- Set(key string, value []byte): this will set or update the key with the specified value
	- Delete(key string): this will delete the key from the store
}
