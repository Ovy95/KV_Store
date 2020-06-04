package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test text to appear on page")
}
func main() {
	http.HandleFunc("/", indexHandler)
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
