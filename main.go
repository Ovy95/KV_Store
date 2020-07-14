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

func storeInterface(r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("This is running the get if statment")
	} else if r.Method == "PUT" {
		fmt.Println("- Set(key string, value []byte): this will set or update the key with the specified value")
	} else if r.Method == delete {
		fmt.Println("- Deletes(key string): this will delete the key from the store")
	}
}

// This is this to be a receiver function of type string, then do the rest .
// func storeInterface {
// 	if r.Method == "GET" {
// 		fmt.Println("This is running the get if statment")
// 	} else if r.Method == put {
// 		fmt.Println("- Set(key string, value []byte): this will set or update the key with the specified value")
// 	} else if r.Method == delete {
// 		fmt.Println("- Deletes(key string): this will delete the key from the store")
// 	}
// }
