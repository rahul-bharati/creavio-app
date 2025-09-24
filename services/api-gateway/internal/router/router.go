package router

import (
	"net/http"

	"github.com/rahul-bharati/creavio-app/services/api-gateway/internal/handler"
)

func NewRouter(h *handler.Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", h.Healthz)
	mux.HandleFunc("/user/hello", h.UserGreeting)
	mux.HandleFunc("/notification/hello", h.NotificationGreeting)

	return mux
}
