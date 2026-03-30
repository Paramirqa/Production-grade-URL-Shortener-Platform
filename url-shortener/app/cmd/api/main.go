package main

import (
	"log"
	"net/http"

	"github.com/Paramirqa/url-shortener/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	h := handler.NewHandler()
	mux.HandleFunc("/health", h.Health)
	mux.HandleFunc("/shorten", h.Shorten)
	mux.HandleFunc("/", h.Redirect)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
