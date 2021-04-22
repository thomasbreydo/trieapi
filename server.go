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

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else {
		port = ":" + port
	}
	log.Printf("See http://localhost%s/api/v1/", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	word, ok := mux.Vars(r)["word"]
	if !ok {
		word = ""
	}
	mutex.Lock()
	add := trie.Add(word)
	mutex.Unlock()
	if add {
		_, err := fmt.Fprintf(w, "Keyword (%s) added", word)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		_, err := fmt.Fprintf(w, "Keyword (%s) present", word)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	word, ok := mux.Vars(r)["word"]
	if !ok {
		word = ""
	}
	mutex.Lock()
	del := trie.Delete(word)
	mutex.Unlock()
	if del {
		_, err := fmt.Fprintf(w, "Keyword (%s) deleted", word)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		_, err := fmt.Fprintf(w, "Keyword (%s) missing", word)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
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
	if found {
		_, err := fmt.Fprintf(w, "Keyword (%s) found", word)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		_, err := fmt.Fprintf(w, "Keyword (%s) not found", word)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
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
	_, err := fmt.Fprint(w, strings.Join(words, "\n"))
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

func clearHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	trie.Clear()
	mutex.Unlock()
}
