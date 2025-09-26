package router

import (
	"net/http"

	"github.com/rahul-bharati/creavio-app/services/notification/internal/handler"
)

func New(h *handler.NotificationHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", h.Healthz)
	mux.HandleFunc("/hello", h.Greet)

	return mux
}
