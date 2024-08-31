package handler

import (
	db "rest/api/internals/db/sqlc"
	"rest/api/internals/service"

	"github.com/go-chi/chi/v5"
)

type AuthHandler struct {
	svc *service.AuthService
}

func NewAuthHandler(store db.Store) *AuthHandler {
	svc := &service.AuthService{
		Store: store,
	}

	return &AuthHandler{
		svc: svc,
	}
}

// Load routes
func LoadAuthRoutes(router *chi.Mux) {
	// router.
}