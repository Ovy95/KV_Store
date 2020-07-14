package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	urlChannel := make(chan string)
	// go readingstrings()
	go readingstrings(urlChannel, "T Testing")
	go readingstrings(urlChannel, "2 Testing")
	urlListing(urlChannel)

}

func readingstrings(out chan<- string, w string) {

	time.Sleep(2 * time.Second)
	out <- w + " --> " + strings.ToUpper(w)
	time.Sleep(2 * time.Second)
	close(out)
}

func urlListing(in <-chan string) { // Prints urls from Using channel strings to terminal
	newdata := in
	storeddata := []byte{}
	storeddata = append(storeddata, newdata)

	for urls := range in {
		fmt.Println(urls)
	}
}
