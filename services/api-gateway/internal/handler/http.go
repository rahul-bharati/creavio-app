package handler

import (
	"net/http"

	"github.com/rahul-bharati/creavio-app/services/api-gateway/internal/client"
)

type Handler struct {
	User         *client.UserClient
	Notification *client.NotificationClient
}

func (h *Handler) Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		return
	}
}

func (h *Handler) UserGreeting(w http.ResponseWriter, r *http.Request) {
	greeting, err := h.User.GetGreeting()
	if err != nil {
		http.Error(w, "Failed to get user greeting", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(greeting))
	if err != nil {
		return
	}
}

func (h *Handler) NotificationGreeting(w http.ResponseWriter, r *http.Request) {
	greeting, err := h.Notification.GetGreeting()
	if err != nil {
		http.Error(w, "Failed to get notification greeting", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(greeting))
	if err != nil {
		return
	}
}
