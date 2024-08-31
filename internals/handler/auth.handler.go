package handler

import (
	"net/http"
	db "rest/api/internals/db/sqlc"
	"rest/api/internals/service"
	"rest/api/internals/utils"

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
func (h *AuthHandler) LoadAuthRoutes(router chi.Router) {
	router.Get("/login", h.Login)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	utils.SuccessMessage(w, utils.Response{
		Message: "Login Message returned.",
	})
}