package main

import "net/http"

func main() {
	http.HandleFunc("/status", statusHandler)
	http.ListenAndServe("127.0.0.1:3000", nil)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
