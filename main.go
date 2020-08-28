package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

type Store interface {
	get(key string) ([]byte, bool)
	post(key string, value []byte)
	delete(key string)
}

type Data struct {
	sync.Mutex
	DataStore map[string][]byte
}

func HttpHandlerFunc(w http.ResponseWriter, r *http.Request) {

	url := r.URL.String()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()
	key := strings.Trim(url, "/")
	value := body

	switch r.Method {
	case "GET":
		Store.get(key)
		w.WriteHeader(http.StatusOK)

		// This has got to be in here as its being used from a WR function
		//w.WriteHeader(404)
		//w.Write([]byte("No Key could be found, try http://localhost:8080/ to see the list of stored keys"))
		return
	case "POST":
		Store.post(key, []byte(value))
		w.WriteHeader(http.StatusOK)
		return
	case "DELETE":
		Store.delete(key)
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}

func (s *Data) get(key string) {
	fmt.Println("get called")
	s.Lock()

	s.Unlock()

}

func (s *Data) post(key string, value []byte) {
	fmt.Println("post called")

	s.Lock()

	s.Unlock()
}
func (s *Data) delete(st, key string) {
	fmt.Println("delete called")
	s.Lock()
	// From feedback might look something like this
	//delete(s.Data, key)
	s.Unlock()
}

func initialise() {
	// this Needs fixing
	// return &Data{
	// 	DataStore: map[string][]byte{},
	// }
}

func main() {

	http.HandleFunc("/", HttpHandlerFunc)
	fmt.Println("Server 8080 is up")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
