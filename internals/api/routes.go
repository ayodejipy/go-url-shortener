package api

import (
	"rest/api/internals/email"
	"rest/api/internals/handler"
	"rest/api/internals/middleware"

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

	// TODO: Pass email handler to routes to avoid reinstantiation
	// email handler to be passed down routes
	emailHandler := email.NewSendEmailHandler(s.config, s.logger)

	// middleware
	params := middleware.Middleware{
		Config: s.config,
		Store:  s.store,
		Logger: s.logger,
	}
	middlewareHandler := middleware.NewMiddleware(params)

	// register handlers
	pHandler := handler.NewPingHandler(s.store)
	authHandler := handler.NewAuthHandler(s.store, s.config, s.logger, middlewareHandler, emailHandler)
	urlHandler := handler.NewUrlHandler(s.store, s.config, s.logger, middlewareHandler)

	// Group routes
	router.Route("/ping", pHandler.LoadPingRoute)
	router.Route("/auth", authHandler.LoadAuthRoutes)
	router.Route("/", urlHandler.LoadUrlRoutes)

	// Private Routes
	// Require Authentication
	// router.Group(func(r chi))

}
