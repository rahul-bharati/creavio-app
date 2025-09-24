package main

import (
	"log"
	"net/http"

	"github.com/rahul-bharati/creavio-app/services/notification/internal/handler"
	"github.com/rahul-bharati/creavio-app/services/notification/internal/repository/memory"
	"github.com/rahul-bharati/creavio-app/services/notification/internal/router"
	"github.com/rahul-bharati/creavio-app/services/notification/internal/service"
)

func main() {
	repo := memory.NewMemoryNotificationRepository()
	svc := service.NewNotificationService(repo)
	h := handler.NewNotificationHandler(repo)

	r := router.New(h)

	addr := ":8002"

	log.Printf("Listening on %s", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("could not start server: %v", err)
	}

}
