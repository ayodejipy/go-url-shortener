package handler

import (
	"net/http"
	"rest/api/internals/config"
	db "rest/api/internals/db/sqlc"
	"rest/api/internals/dto"
	"rest/api/internals/logger"
	"rest/api/internals/service"

	// "rest/api/internals/service"
	"rest/api/internals/utils"

	"github.com/go-chi/chi/v5"
)

type AuthHandler struct {
	svc *service.AuthService
}

func NewAuthHandler(store db.Store, config *config.AppConfig, logger *logger.Logger) *AuthHandler {
	svc := &service.AuthService{
		Store: store,
		Config: config,
		Logger: logger,
	}

	return &AuthHandler{
		svc: svc,
	}
}

// Load routes
func (h *AuthHandler) LoadAuthRoutes(router chi.Router) {
	router.Post("/login", h.login)
	router.Post("/register", h.register)
	router.Post("/forgot-password", h.forgotPassword)
}


func (h *AuthHandler) register(w http.ResponseWriter, r *http.Request) {
	req := dto.CreateUser{}

	// parse the request body into json
	if err := utils.ParseJSON(r, &req); err != nil {
		utils.BadRequestError(w, err)
		return
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
		return
	}

	utils.SuccessMessage(w, utils.Response{
		Message: "User created successfully",
		Data: map[string]string{
			"accessToken": token,
		},
	})
}


func (h *AuthHandler) login(w http.ResponseWriter, r *http.Request) {
	req := dto.LoginPayload{}

	// parse the request body
	if err := utils.ParseJSON(r, &req); err != nil {
		utils.BadRequestError(w, err)
		return
	}

	// send the body to the service
	token, err := h.svc.Login(r.Context(), req)
	if err != nil {
		utils.ErrorMessage(w, err)
		return
	}

	utils.SuccessMessage(w, utils.Response{
		Message: "User logged in successfully",
		Data: map[string]string{
			"accessToken": token,
		},
	})
}

func (h *AuthHandler) forgotPassword(w http.ResponseWriter, r *http.Request) {
	req := dto.ForgotPasswordPayload{}

	// parse the request body
	if err := utils.ParseJSON(r, &req); err != nil {
		utils.BadRequestError(w, err)
		return
	}

	// send the body to the service
	err := h.svc.ForgotPassword(r.Context(), req)
	if err != nil {
		utils.ErrorMessage(w, err)
		return
	}

	utils.SuccessMessage(w, utils.Response{
		Message: "Password reset email sent",
		Data: map[string]string{},
	})
}

func (h *AuthHandler) resetPassword(w http.ResponseWriter, r *http.Request) {
	
}