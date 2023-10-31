package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/status", statusHandler)

	router.HandleFunc("/{code}", redirectHandler)

	router.HandleFunc("/", indexPageHandler)
	router.HandleFunc("/web/url", postUrlHandler)

	http.Handle("/", router)
	http.ListenAndServe("127.0.0.1:3000", router)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	w.WriteHeader(http.StatusOK)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	code := mux.Vars(r)["code"]

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Code: %s", code)
}

func indexPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	t, _ := template.ParseFiles("../internal/web/public/index.html")
	t.Execute(w, nil)
}

func postUrlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	// Handle form submission

	w.WriteHeader(http.StatusOK)
}
