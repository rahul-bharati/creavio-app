package main

import (
	"log"
	"net/http"

	"github.com/rahul-bharati/creavio-app/services/user/internal/handler"
	repository "github.com/rahul-bharati/creavio-app/services/user/internal/repository/memory"
	"github.com/rahul-bharati/creavio-app/services/user/internal/router"
	"github.com/rahul-bharati/creavio-app/services/user/internal/service"
)

func main() {
	repo := repository.NewMemoryHelloRepository()
	helloSvc := service.NewUserService(repo)
	h := handler.NewHandler(helloSvc)

	r := router.New(h)

	addr := ":8001"
	log.Printf("Listening on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
