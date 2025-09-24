package handler

import (
	"fmt"
	"net/http"

	"github.com/rahul-bharati/creavio-app/services/user/internal/repository"
	"github.com/rahul-bharati/creavio-app/services/user/internal/service"
)

type Handler struct {
	userService *service.UserService
}

func NewHandler(handlerService *service.UserService) *Handler {
	return &Handler{userService: handlerService}
}

func (h *Handler) Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprint(w, "ok")
	if err != nil {
		return
	}
}

func (h *Handler) Greet(w http.ResponseWriter, r *http.Request) {
	msg, err := h.userService.GetGreeting()
	if err != nil {
		http.Error(w, "Failed to get greeting", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprint(w, msg)
	if err != nil {
		return
	}
}
