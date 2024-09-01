package api

import (
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

	// register handlers
	pHandler := handler.NewPingHandler(s.store)
	authHandler := handler.NewAuthHandler(s.store, s.config)

	// Group routes
	router.Route("/", pHandler.LoadPingRoute)
	router.Route("/auth", authHandler.LoadAuthRoutes)
}
