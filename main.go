package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

type KvData struct {
	Key   string
	Value []byte
}

type KvHandlers struct {
	sync.Mutex
	store map[string]KvData
}

func (h *KvHandlers) HttpHandlerFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
		w.WriteHeader(http.StatusOK)
		return
	case "POST":
		h.post(w, r)
		w.WriteHeader(http.StatusOK)
		return
	case "DELETE":
		h.delete(w, r)
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}

func (h *KvHandlers) get(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	url = strings.Trim(url, "/")
	KValue := make([]KvData, len(h.store))
	h.Lock()

	i := 0
	for _, stored := range h.store {
		KValue[i] = stored
		i++
	}
	h.Unlock()

	if url == "" {
		fmt.Fprintln(w, KValue)
		return
	}
	for _, item := range KValue {
		if url == item.Key {
			w.Write([]byte(item.Value))
			return
		}
	}
	w.WriteHeader(404)
	w.Write([]byte("No Key could be found, try http://localhost:8080/ to see the list of stored keys"))
}

func (h *KvHandlers) post(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	url = strings.Trim(url, "/")

	var KeyValue KvData
	KeyValue.Key = url

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	KeyValue.Value = append(KeyValue.Value, body...)
	h.Lock()
	h.store[KeyValue.Key] = KeyValue
	defer h.Unlock()
}
func (h *KvHandlers) delete(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	url = strings.Trim(url, "/")
	for _, slices := range h.store {
		if url == slices.Key {
			delete(h.store, slices.Key)
		}
	}
}

func newHandlers() *KvHandlers {
	return &KvHandlers{
		store: map[string]KvData{},
	}
}

func main() {
	KvHandlers := newHandlers()
	http.HandleFunc("/", KvHandlers.HttpHandlerFunc)
	fmt.Println("Server 8080 is up")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
