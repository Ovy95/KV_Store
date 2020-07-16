package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const put = "PUT"
const get = "GET"
const post = "POST"
const delete = "DELETE"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website! Create your Key and value pair you would like to save")
		bank := make(chan []string, 5)
	})
	// http.HandleFunc("/", handler))
	// fmt.Println("Server is starting on port 8080")

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

		url := r.URL.String()
		url = strings.Trim(url, "/")
		storeInterface(r.Method, url)
		fmt.Fprintf(bank)
		fmt.Fprintf(w, "put") // Writing on the actually page
		break
	case post:
		fmt.Fprintf(w, "THIS IS A %s method ON THE LOCALHOST AT THE INDEX%s\n", r.Method, r.URL)
		fmt.Println(r.Method)
		fmt.Fprintf(w, "Post")
		url := r.URL.String()
		url = strings.Trim(url, "/")
		storeInterface(r.Method, url)
		break
	case put:
		fmt.Fprintf(w, "THIS IS A %s method ON THE LOCALHOST AT THE INDEX%s\n", r.Method, r.URL)
		fmt.Println(r.Method)
		fmt.Fprintf(w, "put")
		url := r.URL.String()
		url = strings.Trim(url, "/")
		storeInterface(r.Method, url)
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

func storeInterface(method string, url string) {
	//bank := make(chan []string, 5)
	keystore := []string{}
	if method == "GET" {
		fmt.Println("This is running the get if statment")
	} else if method == "PUT" {
		fmt.Println("- Set(key string, value []byte): this will set or update the key with the specified value")
		fmt.Println(url)

		keystore = append(keystore, "coco")
		fmt.Println(keystore)
		keystore = append(keystore, url)
		fmt.Println(keystore)
	} else if method == delete {
		fmt.Println("- Deletes(key string): this will delete the key from the store")
	}

}
