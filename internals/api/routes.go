package api

import (
	"rest/api/internals/email"
	"rest/api/internals/handler"

	"github.com/go-chi/cors"
)

func SetupRoutes(s *Server) {
	router := s.router

	// setup CORS
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	email.NewSendEmailHandler(s.config, s.logger)

	// register handlers
	pHandler := handler.NewPingHandler(s.store)
	authHandler := handler.NewAuthHandler(s.store, s.config, s.logger)

	// Group routes
	router.Route("/", pHandler.LoadPingRoute)
	router.Route("/auth", authHandler.LoadAuthRoutes)
}
