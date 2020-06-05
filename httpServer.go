package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is starting on port 8080")
	http.ListenAndServe(":8080", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test text to appear on page    ")
	fmt.Fprintf(w, "")
	fmt.Fprintf(w, "%s is the method. This is the URL %s\n", r.Method, r.URL)
}
