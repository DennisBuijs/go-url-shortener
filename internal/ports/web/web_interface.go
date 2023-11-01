package web

import (
	"encoding/json"
	"html/template"
	"net/http"
	"url-shortener/internal/use_cases"
)

type WebInterface struct {
	CreateUrlUseCase *use_cases.CreateUrlUseCase
}

func (webInterface *WebInterface) CreateUrlViaWebUiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	r.ParseForm()
	urlString := r.FormValue("url")

	url := webInterface.CreateUrlUseCase.CreateUrl(urlString)
	webInterface.CreateUrlUseCase.UrlRepository.Add(url)

	w.WriteHeader(http.StatusOK)
	t, _ := template.ParseFiles("../internal/web/public/url_detail.html")
	t.Execute(w, url)
}

func (webInterface *WebInterface) CreateUrlViaApiHandler(w http.ResponseWriter, r *http.Request) {
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
	url := webInterface.CreateUrlUseCase.CreateUrl(urlString)
	webInterface.CreateUrlUseCase.UrlRepository.Add(url)

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
