package router

import (
	"net/http"

	"github.com/rahul-bharati/creavio-app/services/user/internal/handler"
)

func New(h *handler.Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", h.Healthz)
	mux.HandleFunc("/hello", h.Greet)
	return mux
}
