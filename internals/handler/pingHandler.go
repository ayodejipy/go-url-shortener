package handler

import (
	"fmt"
	"net/http"
	db "rest/api/internals/db/sqlc"
	"rest/api/internals/service"

	"github.com/go-chi/chi/v5"
)

type PingHandler struct {
	svc *service.PingService
}

func NewPingHandler(store db.Store) *PingHandler {
	svc := &service.PingService{
		Store: store,
	}

	return &PingHandler{
		svc: svc,
	}
}

// load routes
func (h *PingHandler) LoadPingRoute(router *chi.Mux) {
	router.Get("/ping", h.PingServer)
}

// implement route functions
func (h *PingHandler) PingServer(w http.ResponseWriter, r *http.Request) {
	val := h.svc.Ping()
	fmt.Printf("%v\n", val)
	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(http.StatusOK)
	w.Write(val)
}