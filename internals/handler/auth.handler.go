package handler

import (
	"net/http"
	"rest/api/internals/config"
	db "rest/api/internals/db/sqlc"
	"rest/api/internals/dto"
	"rest/api/internals/logger"
	"rest/api/internals/middleware"
	"rest/api/internals/service"

	// "rest/api/internals/service"
	"rest/api/internals/utils"

	"github.com/go-chi/chi/v5"
)

type AuthHandler struct {
	mw  *middleware.Middleware
	svc *service.AuthService
}

func NewAuthHandler(store db.Store, config *config.AppConfig, logger *logger.Logger, mw *middleware.Middleware) *AuthHandler {
	svc := &service.AuthService{
		Store:  store,
		Config: config,
		Logger: logger,
	}

	return &AuthHandler{
		svc: svc,
		mw:  mw,
	}
}

// Load routes
func (h *AuthHandler) LoadAuthRoutes(router chi.Router) {
	router.Post("/login", h.login)
	router.Post("/register", h.register)
	router.Post("/forgot-password", h.forgotPassword)
	router.Post("/reset-password", h.resetPassword)

	// protected routes
	router.Group(func(r chi.Router) {
		r.Use(h.mw.AuthorizeUser())
		r.Post("/verify", h.verifyUser)
	})
}

func (h *AuthHandler) register(w http.ResponseWriter, r *http.Request) {
	req := dto.CreateUser{}

	// parse the request body into json
	if err := utils.ParseJSON(r, &req); err != nil {
		utils.BadRequestError(w, err)
		return
	}

	user := db.CreateUserParams{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  req.Password,
	}

	err := h.svc.Register(r.Context(), user)
	if err != nil {
		utils.InternalError(w, err)
		return
	}

	utils.SuccessMessage(w, utils.Response{
		Message: "User created successfully. Please verify your email next",
	})
}

func (h *AuthHandler) verifyUser(w http.ResponseWriter, r *http.Request) {
	req := dto.VerifyEmailPayload{}

	// parse the request body into json
	if err := utils.ParseJSON(r, &req); err != nil {
		utils.BadRequestError(w, err)
		return
	}

	err := h.svc.VerifyUser(r.Context(), req.Code)
	if err != nil {
		utils.InternalError(w, err)
		return
	}

	utils.SuccessMessage(w, utils.Response{
		Message: "Email verified successfully.",
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
		utils.InternalError(w, err)
		return
	}

	// set cookie
	cookie := http.Cookie{
		Name:     "Authorization",
		Value:    token,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   3600 * 24 * 30,
		Secure:   false, // set to TRUE in production
		Path:     "/",
		Domain:   "",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

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
		utils.BadRequestError(w, err)
		return
	}

	utils.SuccessMessage(w, utils.Response{
		Message: "Password reset email sent",
		Data:    map[string]string{},
	})
}

func (h *AuthHandler) resetPassword(w http.ResponseWriter, r *http.Request) {
	payload := dto.ResetPasswordPayload{}

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.BadRequestError(w, err)
		return
	}

	err := h.svc.ResetPassword(r.Context(), payload)
	if err != nil {
		utils.BadRequestError(w, err)
		return
	}

	utils.SuccessMessage(w, utils.Response{
		Message: "Password reset succesfully.",
		Data:    map[string]string{},
	})
}
