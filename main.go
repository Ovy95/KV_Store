package main

import (
	"fmt"
	"log"
	"net/http"
)

const put = "PUT"
const get = "GET"
const post = "POST"
const delete = "DELETE"

type storeMethods interface {
	storeInterface()
}

type httpVerbs struct {
	verb string
}

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
		fmt.Fprintf(w, "----\n")
		fmt.Printf(r.Method)
		fmt.Println(r.URL)

		Url := r.URL.String()
		// method := r.Method
		keystore := []string{}
		keystore = append(keystore, "coco")
		// keystore = append(keystore, method)
		keystore = append(keystore, Url)
		fmt.Println(keystore)

		fmt.Fprintf(w, "put") // Writing on the actually page
		break
	case post:
		fmt.Fprintf(w, "THIS IS A %s method ON THE LOCALHOST AT THE INDEX%s\n", r.Method, r.URL)
		fmt.Println(r.Method)
		fmt.Fprintf(w, "Post")

		break
	case put:
		fmt.Fprintf(w, "THIS IS A %s method ON THE LOCALHOST AT THE INDEX%s\n", r.Method, r.URL)
		fmt.Println(r.Method)
		fmt.Fprintf(w, "put")
		urlstring := r.Method
		url := httpVerbs{urlstring}
		fmt.Println(url)
		break
	case delete:
		fmt.Fprintf(w, "THIS IS A %s method ON THE LOCALHOST AT THE INDEX%s\n", r.Method, r.URL)
		fmt.Println(r.Method)
		fmt.Fprintf(w, "Post")
		break
	default:
		http.Error(w, "BadRequest status code", 400)
	}
}

func (v httpVerbs) storeInterface() {
	if v.verb == "GET" {

		fmt.Println("This is running the GET if statment")

	} else if v.verb == "PUT" {
		fmt.Println("PUT- Set(key string, value []byte): this will set or update the key with the specified value")
	} else if v.verb == delete {
		fmt.Println("DELETE- Deletes(key string): this will delete the key from the store")
	} else {
		fmt.Println(v.verb)
		fmt.Println("This is the value that got print out")
	}

}
