package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"url-shortener/internal/adapters/secondary/memory"
	"url-shortener/internal/ports"
)

func main() {
	createUrlUseCase := &ports.CreateUrlUseCase{}
	webInterface := &WebInterface{createUrlUseCase}

	router := mux.NewRouter()

	router.HandleFunc("/status", statusHandler)

	router.HandleFunc("/{code}", redirectHandler)

	router.HandleFunc("/", indexPageHandler)
	router.HandleFunc("/web/url", webInterface.createUrlViaWebUiHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/url", webInterface.createUrlViaApiHandler).Methods(http.MethodPost)

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

type WebInterface struct {
	createUrlUseCase *ports.CreateUrlUseCase
}

func (webInterface *WebInterface) createUrlViaWebUiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	r.ParseForm()
	urlString := r.FormValue("url")

	url := webInterface.createUrlUseCase.CreateUrl(urlString)

	repository, _ := memory.NewUrlRepository()
	repository.Add(url)

	w.WriteHeader(http.StatusOK)
	t, _ := template.ParseFiles("../internal/web/public/url_detail.html")
	t.Execute(w, url)
}

func (webInterface *WebInterface) createUrlViaApiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	decoder := json.NewDecoder(r.Body)
	var requestData map[string]string

	if err := decoder.Decode(&requestData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	urlString := requestData["url"]
	url := webInterface.createUrlUseCase.CreateUrl(urlString)

	repository, _ := memory.NewUrlRepository()
	repository.Add(url)

	responseStruct := struct {
		Url string `json:"url"`
	}{Url: url.GetShortUrl()}

	responseJson, err := json.Marshal(responseStruct)
	if err != nil {
		http.Error(w, "Failed to create response body", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}
