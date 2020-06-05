package main

import (
	"fmt"
	"net/http"
)

const get string = "GET"
const post string = "POST"
const delete string = "DELETE"

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is starting on port 8080")
	http.ListenAndServe(":8080", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case get:
		fmt.Fprintf(w, "CASE METHOD   ")
		fmt.Fprintf(w, "")
		fmt.Fprintf(w, "THIS IS A %s method ON THE LOCALHOST AT THE INDEX%s\n", r.Method, r.URL)
		fmt.Println(r.Method)
	case post:
		fmt.Fprintf(w, "THIS IS A %s method ON THE LOCALHOST AT THE INDEX%s\n", r.Method, r.URL)
		fmt.Println(r.Method)
	case delete:
		fmt.Println(r.Method)
	}
	// fmt.Println(r.Method)
}
