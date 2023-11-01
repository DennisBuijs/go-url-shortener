package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"url-shortener/internal/adapters/secondary/sqlite"
	"url-shortener/internal/ports"
	"url-shortener/internal/ports/web"
	"url-shortener/internal/use_cases"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "../data/database.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	urlRepository, _ := sqlite.NewUrlRepository(db)
	createUrlUseCase := use_cases.NewCreateUrlUseCase(urlRepository)
	webInterface := &web.WebInterface{CreateUrlUseCase: createUrlUseCase}

	router := mux.NewRouter()

	router.HandleFunc("/status", statusHandler)

	router.HandleFunc("/", indexPageHandler)
	router.HandleFunc("/web/url", webInterface.CreateUrlViaWebUiHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/url", webInterface.CreateUrlViaApiHandler).Methods(http.MethodPost)

	router.HandleFunc("/{code}", redirectHandler(urlRepository))

	fs := http.FileServer(http.Dir("../internal/web/public/"))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	http.Handle("/", router)
	http.ListenAndServe("127.0.0.1:3000", router)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	w.WriteHeader(http.StatusOK)
}

func redirectHandler(urlRepository ports.UrlRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}

		code := mux.Vars(r)["code"]
		url, err := urlRepository.FindByCode(code)

		if err != nil {
			http.Error(w, "URL with code '"+code+"' not found", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, url.Url, http.StatusMovedPermanently)
		return
	}
}

func indexPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	t, _ := template.ParseFiles("../internal/web/public/index.html")
	t.Execute(w, nil)
}
