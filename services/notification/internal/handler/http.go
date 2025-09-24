package handler

import (
	"fmt"
	"net/http"

	"github.com/rahul-bharati/creavio-app/services/notification/internal/repository"
)

type NotificationHandler struct {
	notificationService repository.NotificationRepository
}

func NewNotificationHandler(notificationService repository.NotificationRepository) *NotificationHandler {
	return &NotificationHandler{notificationService}
}

func (h *NotificationHandler) Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprint(w, "OK")
	if err != nil {
		return
	}
}

func (h *NotificationHandler) Greet(w http.ResponseWriter, r *http.Request) {
	msg, err := h.notificationService.GetGreeting()
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
