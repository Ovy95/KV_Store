package main

import (
	"fmt"
	"log"
	"net/http"
)

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
	case post:
		fmt.Fprintf(w, "THIS IS A %s method ON THE LOCALHOST AT THE INDEX%s\n", r.Method, r.URL)
		fmt.Println(r.Method)
		break
	case delete:
		fmt.Fprintf(w, "THIS IS A %s method ON THE LOCALHOST AT THE INDEX%s\n", r.Method, r.URL)
		fmt.Println(r.Method)
		break
	case get:
		fmt.Fprintf(w, "THIS IS A %s method ON THE LOCALHOST AT THE INDEX%s\n", r.Method, r.URL)
		fmt.Println(r.Method)
		break
	default:
		http.Error(w, "BadRequest status code", 400)
		//log the errors default case on the switch statement to log and error and return a BadRequest status code
	}
	// fmt.Println(r.Method)
}
