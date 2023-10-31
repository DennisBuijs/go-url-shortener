package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/status", statusHandler)
	http.ListenAndServe("127.0.0.1:3000", nil)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	w.WriteHeader(http.StatusOK)
}
