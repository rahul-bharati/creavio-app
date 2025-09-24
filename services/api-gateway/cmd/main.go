package main

import (
	"log"
	"net/http"
	"time"

	"github.com/rahul-bharati/creavio-app/services/api-gateway/internal/client"
	"github.com/rahul-bharati/creavio-app/services/api-gateway/internal/handler"
	"github.com/rahul-bharati/creavio-app/services/api-gateway/internal/router"
)

func main() {
	notificationURL := "http://localhost:8002/"
	userURL := "http://localhost:8001/"

	httpClient := &http.Client{Timeout: time.Second * 3}

	h := &handler.Handler{
		User:         &client.UserClient{BaseURL: userURL, HTTP: httpClient},
		Notification: &client.NotificationClient{BaseURL: notificationURL, HTTP: httpClient},
	}

	mux := router.NewRouter(h)

	addr := ":8000"
	log.Printf("API Gateway listening on %s (user= %s, notification=%s)", addr, notificationURL, notificationURL)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
