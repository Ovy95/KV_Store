package store

import (
	"fmt"
	"sync"
)

type Store interface {
	get(key string) ([]byte, bool)
	post(key string, value []byte)
	del(key string)
}

type Data struct {
	sync.Mutex
	Data map[string][]byte
}

func (s *Data) post(key string, value []byte) {
	fmt.Println("post called")
	s.Lock()

	s.Data[key] = append(s.Data[key], value...)

	s.Unlock()

}

func (s *Data) get(key string) ([]byte, bool) {

	fmt.Println("get called")
	var data []byte

	if key == "" {
		var tmp string
		for key := range s.Data {
			tmp += key + ","
		}
		return []byte(tmp), true
	}

	if data, ok := s.Data[key]; ok {
		return data, ok
	}
	return data, false
}

func (s *Data) del(key string) {

	s.Lock()

	delete(s.Data, key)

	s.Unlock()
}

func Get(s Store, key string) ([]byte, bool) {
	return s.get(key)
}

func Post(s Store, key string, value []byte) {
	s.post(key, value)
	return
}

func Delete(s Store, key string) {
	s.del(key)
	return
}
