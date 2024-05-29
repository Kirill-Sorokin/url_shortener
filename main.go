package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

var urlStore = struct {
	sync.RWMutex
	urls map[string]string
}{urls: make(map[string]string)}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func shortenURL(w http.ResponseWriter, r *http.Request) {
	var req map[string]string
	json.NewDecoder(r.Body).Decode(&req)
	originalURL := req["url"]

	shortURL := randString(6)

	urlStore.Lock()
	urlStore.urls[shortURL] = originalURL
	urlStore.Unlock()

	json.NewEncoder(w).Encode(map[string]string{"short_url": shortURL})
}

func resolveURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["short_url"]

	urlStore.RLock()
	originalURL, ok := urlStore.urls[shortURL]
	urlStore.RUnlock()

	if !ok {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	r := mux.NewRouter()
	r.HandleFunc("/shorten", shortenURL).Methods("POST")
	r.HandleFunc("/{short_url}", resolveURL).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.ListenAndServe(":8080", r)
}
