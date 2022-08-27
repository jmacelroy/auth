package main

import (
	"log"
	"net/http"
)

func auth(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("x-api-key")
	if apiKey == "" {
		http.Error(w, "x-api-key header is not set", http.StatusUnauthorized)
		return
	}

	w.Header().Set("x-auth-user", "neo")
	log.Println("Hard coding user to neo.")
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", auth)
	log.Println("starting auth listener on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
