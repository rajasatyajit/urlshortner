package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/lithammer/shortuuid/v4"
)

type Mapper struct {
	Mapping map[string]string
	Lock    sync.Mutex
}

var urlMapper Mapper

func init() {
	urlMapper.Mapping = make(map[string]string)
}

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	r.Post("/short-it", createShortURLHandler)
	r.Get("/short/{key}", redirectHandler)
	http.ListenAndServe(":3000", r)
}

func createShortURLHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	u := r.Form.Get("url")
	if u == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("url is required"))
		return
	}
	if _, ok := urlMapper.Mapping[u]; ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("url already exists"))
		return
	}
	// generate key
	key := shortuuid.NewWithNamespace(u)
	insertMapping(key, u)
	log.Println("key: " + key + " url: " + u)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("http://localhost:3000/short/" + key))
}

func insertMapping(key string, url string) {
	urlMapper.Lock.Lock() // lock the mutex to prevent race conditions and deadlocks when accessing shared resources across multiple goroutines simultaneously (i.e., in parallel)
	defer urlMapper.Lock.Unlock()

	urlMapper.Mapping[key] = url
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("key is required"))
		return
	}

	u := fetchMapping(key)
	if u == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("url has not been shortened yet"))
		return
	}
	log.Println("Redirecting to: " + u)
	http.Redirect(w, r, u, http.StatusFound)
}

func fetchMapping(key string) string {
	urlMapper.Lock.Lock() // lock the mutex to prevent race conditions and deadlocks when accessing shared resources across multiple goroutines simultaneously (
	defer urlMapper.Lock.Unlock()

	log.Printf("Full mapping: %v\n", urlMapper.Mapping)
	return urlMapper.Mapping[key]
}
