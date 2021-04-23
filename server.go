package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/thomasbreydo/trieapi/tries"
)

var trie = tries.New()
var mutex sync.Mutex

func main() {
	router := mux.NewRouter().StrictSlash(true)

	// operations when word == ""
	router.HandleFunc("/api/v1/add/", addHandler)
	router.HandleFunc("/api/v1/delete/", deleteHandler)
	router.HandleFunc("/api/v1/complete/", completeHandler)
	router.HandleFunc("/api/v1/search/", searchHandler)

	// operations when word != ""
	router.HandleFunc("/api/v1/add/{word}", addHandler)
	router.HandleFunc("/api/v1/delete/{word}", deleteHandler)
	router.HandleFunc("/api/v1/complete/{word}", completeHandler)
	router.HandleFunc("/api/v1/search/{word}", searchHandler)

	// operations without word
	router.HandleFunc("/api/v1/display", displayHandler)
	router.HandleFunc("/api/v1/clear", clearHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	word, ok := mux.Vars(r)["word"]
	if !ok {
		word = ""
	}
	mutex.Lock()
	mod := trie.Add(word)
	mutex.Unlock()
	w.Header().Set("Content-Type", "application/json")
	out, err := json.Marshal(map[string]bool{"modified": mod})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(out)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	word, ok := mux.Vars(r)["word"]
	if !ok {
		word = ""
	}
	mutex.Lock()
	mod := trie.Delete(word)
	mutex.Unlock()
	w.Header().Set("Content-Type", "application/json")
	out, err := json.Marshal(map[string]bool{"modified": mod})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(out)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	word, ok := mux.Vars(r)["word"]
	if !ok {
		word = ""
	}
	mutex.Lock()
	found := trie.Search(word)
	mutex.Unlock()
	w.Header().Set("Content-Type", "application/json")
	out, err := json.Marshal(map[string]bool{"found": found})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(out)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func completeHandler(w http.ResponseWriter, r *http.Request) {
	word, ok := mux.Vars(r)["word"]
	if !ok {
		word = ""
	}
	mutex.Lock()
	words := trie.Complete(word)
	mutex.Unlock()
	w.Header().Set("Content-Type", "application/json")
	out, err := json.Marshal(words)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(out)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func displayHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	words := trie.AllWords()
	mutex.Unlock()
	w.Header().Set("Content-Type", "application/json")
	out, err := json.Marshal(words)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(out)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func clearHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	mod := trie.Clear()
	mutex.Unlock()
	w.Header().Set("Content-Type", "application/json")
	out, err := json.Marshal(map[string]bool{"modified": mod})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(out)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
