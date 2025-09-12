package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("OK"))
		if err != nil {
			log.Println("Failed to write response:", err)
			return
		}
	})

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
