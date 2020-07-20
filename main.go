package main

import (
	"fmt"
	"log"
	"net/http"
)

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
	case "GET":
		fmt.Fprintf(w, "THIS IS A %s method ON THE LOCALHOST AT THE INDEX%s\n", r.Method, r.URL)
		fmt.Fprintf(w, "----\n")

		verb := r.Method
		method := httpVerbs{verb}
		httpVerbs.storeInterface(method)

		urlstr := r.URL.String()
		url := httpVerbs{urlstr}
		fmt.Println(url)
		fmt.Println("This works")
		//httpVerbs.storeInterface(url)
		break
	case "POST":
		fmt.Fprintf(w, "THIS IS A %s method ON THE LOCALHOST AT THE INDEX%s\n", r.Method, r.URL)
		fmt.Println(r.Method)
		fmt.Fprintf(w, "Post")

		verb := r.Method
		method := httpVerbs{verb}
		httpVerbs.storeInterface(method)
		fmt.Println("This works")
		break
	case "PUT":
		fmt.Fprintf(w, "THIS IS A %s method ON THE LOCALHOST AT THE INDEX%s\n", r.Method, r.URL)
		fmt.Println(r.Method)
		fmt.Fprintf(w, "PUT")

		verb := r.Method
		method := httpVerbs{verb}
		httpVerbs.storeInterface(method)

		urlstr := r.URL.String()
		url := httpVerbs{urlstr}
		fmt.Println(url)

		httpVerbs.storeInterface(url)
		break
	case "DELETE":
		fmt.Fprintf(w, "THIS IS A %s method ON THE LOCALHOST AT THE INDEX%s\n", r.Method, r.URL)
		fmt.Println(r.Method)
		fmt.Fprintf(w, "Post")

		verb := r.Method
		method := httpVerbs{verb}
		httpVerbs.storeInterface(method)
		fmt.Println(method)
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

	} else if v.verb == "DELETE" {
		fmt.Println("DELETE- Deletes(key string): this will delete the key from the store")
	} else {
		fmt.Println(v.verb)
		fmt.Println("This is the value that got print out")
	}

}
