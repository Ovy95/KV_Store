package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

const put = "PUT"
const get = "GET"

var waitG sync.WaitGroup

func main() {

	http.HandleFunc("/", handler)
	fmt.Println("Server is starting on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	urlChannel := make(chan []string)
	//bank := make(chan string)

	switch r.Method {
	case put:
		fmt.Fprintf(w, "THIS IS A %s method ON THE LOCALHOST AT THE INDEX%s\n", r.Method, r.URL)
		fmt.Println(r.Method)

		url := r.URL.String()

		waitG.Add(1)
		go readingstrings(urlChannel, url)

		waitG.Done()

		waitG.Wait()
		printchannel(urlChannel)

	default:
		http.Error(w, "BadRequest status code", 400)
	}
}

func readingstrings(c chan []string, url string) {
	url = strings.Trim(url, "/")
	var website []string
	website = strings.SplitAfter(url, "?")

	website = append(website)

	for i := 0; i < 1; i++ {

		fmt.Println("INCOUNTER loop")

		waitG.Add(1)
		go func() {
			c <- append(website)
		}()
		fmt.Println("Loop done")
		waitG.Done()

		waitG.Wait()
		fmt.Println("break out of function")
	}
}

func printchannel(c chan []string) {
	time.Sleep(2 * time.Second)
	for strings := range c {
		fmt.Println("Printed channelfunction stings")
		fmt.Println(strings)
		fmt.Println("Printed function complete")
		break
	}
}
