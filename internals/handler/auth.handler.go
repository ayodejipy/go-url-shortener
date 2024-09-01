package handler

import (
	"net/http"
	"rest/api/internals/config"
	db "rest/api/internals/db/sqlc"
	"rest/api/internals/dto"
	"rest/api/internals/service"
	"rest/api/internals/utils"

	"github.com/go-chi/chi/v5"
)

type AuthHandler struct {
	svc *service.AuthService
}

func NewAuthHandler(store db.Store, config *config.AppConfig) *AuthHandler {
	svc := &service.AuthService{
		Store: store,
		Config: config,
	}

	return &AuthHandler{
		svc: svc,
	}
}

// Load routes
func (h *AuthHandler) LoadAuthRoutes(router chi.Router) {
	router.Get("/login", h.login)
	router.Get("/register", h.register)
}


func (h *AuthHandler) register(w http.ResponseWriter, r *http.Request) {
	req := dto.CreateUser{}

	// parse the request body into json
	if err := utils.ParseJSON(r, &req); err != nil {
		utils.BadRequestError(w, err)
	}

	user := db.CreateUserParams{
		Email: req.Email,
		FirstName: req.FirstName,
		LastName: req.LastName,
		Password: req.Password,
	}

	token, err := h.svc.Register(r.Context(), user)
	if err != nil {
		utils.ErrorMessage(w, err)
	}

	utils.SuccessMessage(w, utils.Response{
		Message: "User created successfully",
		Data: map[string]string{
			"accessToken": token,
		},
	})
}


func (h *AuthHandler) login(w http.ResponseWriter, r *http.Request) {
	utils.SuccessMessage(w, utils.Response{
		Message: "Login Message returned.",
	})
}