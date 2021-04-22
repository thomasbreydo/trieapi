package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/gorilla/mux"

	"github.com/thomasbreydo/trieapi/tries"
)

var trie = tries.New()
var mutex sync.Mutex

func main() {
	router := mux.NewRouter().StrictSlash(true)

	// operations when keyword == ""
	router.HandleFunc("/api/v1/add/", addHandler)
	router.HandleFunc("/api/v1/delete/", deleteHandler)
	router.HandleFunc("/api/v1/complete/", completeHandler)
	router.HandleFunc("/api/v1/search/", searchHandler)

	// operations when keyword != ""
	router.HandleFunc("/api/v1/add/{keyword}", addHandler)
	router.HandleFunc("/api/v1/delete/{keyword}", deleteHandler)
	router.HandleFunc("/api/v1/complete/{keyword}", completeHandler)
	router.HandleFunc("/api/v1/search/{keyword}", searchHandler)

	// operations without keyword
	router.HandleFunc("/api/v1/display", displayHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else {
		port = ":" + port // todo might not work
	}
	log.Printf("See http://localhost%s/api/v1/", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	keyword, ok := mux.Vars(r)["keyword"]
	if !ok {
		keyword = ""
	}
	mutex.Lock()
	modified := trie.Add(keyword)
	mutex.Unlock()
	if modified {
		_, err := fmt.Fprintf(w, "Keyword (%s) added", keyword)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		_, err := fmt.Fprintf(w, "Keyword (%s) present", keyword)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	keyword, ok := mux.Vars(r)["keyword"]
	if !ok {
		keyword = ""
	}
	mutex.Lock()
	deleted := trie.Delete(keyword)
	mutex.Unlock()
	if deleted {
		_, err := fmt.Fprintf(w, "Keyword (%s) deleted", keyword)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		_, err := fmt.Fprintf(w, "Keyword (%s) missing", keyword)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	keyword, ok := mux.Vars(r)["keyword"]
	if !ok {
		keyword = ""
	}
	mutex.Lock()
	found := trie.Search(keyword)
	mutex.Unlock()
	if found {
		_, err := fmt.Fprintf(w, "Keyword (%s) found", keyword)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		_, err := fmt.Fprintf(w, "Keyword (%s) not found", keyword)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func completeHandler(w http.ResponseWriter, r *http.Request) {
	keyword, ok := mux.Vars(r)["keyword"]
	if !ok {
		keyword = ""
	}
	mutex.Lock()
	words := trie.Complete(keyword)
	mutex.Unlock()
	_, err := fmt.Fprintf(w, "Keyword (%s) found", strings.Join(words, "\n"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func displayHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	_, err := fmt.Fprint(w, trie)
	mutex.Unlock()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
