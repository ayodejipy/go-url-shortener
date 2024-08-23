package api

import (
	"log"
	"database/sql"
	"rest/api/internals/config"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	config *config.AppConfig
	router *chi.Mux
	db *sql.DB
}

func NewServer(config *config.AppConfig) *Server {
	r := chi.NewRouter()
	// ctx := context.Background()

	db, err := connectToDB(config.Dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	return &Server{
		config: config,
		router: r,
		db: db,
	}
}

func connectToDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
		// panic(err)
	}
	
	defer db.Close()

	return db, nil
}