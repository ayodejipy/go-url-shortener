package handler

import (
	"net/http"
	"rest/api/internals/config"
	db "rest/api/internals/db/sqlc"
	"rest/api/internals/logger"
	"rest/api/internals/middleware"
	"rest/api/internals/service"
	"rest/api/internals/utils"

	"github.com/go-chi/chi/v5"
)

type UrlHandler struct {
	svc *service.UrlService
	mw  *middleware.Middleware
}

func NewUrlHandler(store db.Store, config *config.AppConfig, logger *logger.Logger, mw *middleware.Middleware) *UrlHandler {
	svc := &service.UrlService{
		Store:  store,
		Logger: logger,
	}

	return &UrlHandler{
		svc: svc,
		mw:  mw,
	}
}

// load routes
func (h *UrlHandler) LoadUrlRoutes(router chi.Router) {
	// protected routes
	router.Group(func(r chi.Router) {
		r.Use(h.mw.AuthorizeUser())

		r.Post("/short", h.ShortenUrl)
		r.Get("/{shortCode}", h.GetUrlByCodeAndRedirect)
	})
}

func (h *UrlHandler) ShortenUrl(w http.ResponseWriter, r *http.Request) {

}

func (h *UrlHandler) GetUrlByCodeAndRedirect(w http.ResponseWriter, r *http.Request) {
	// get shortcode from url param
	shortCode := chi.URLParam(r, "shortCode")

	url, err := h.svc.GetUrlByShortCode(r.Context(), shortCode)
	if err != nil {
		utils.ErrorMessage(w, http.StatusNotFound, err)
		return
	}

	http.Redirect(w, r, url.OriginalUrl, http.StatusFound)
}

func (h *UrlHandler) GetUsersUrl(w http.ResponseWriter, r *http.Request) {

}
