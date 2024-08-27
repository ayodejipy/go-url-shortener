package api

import (
	"context"
	"fmt"
	"log"
	"rest/api/internals/config"
	db "rest/api/internals/db/sqlc"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	config *config.AppConfig
	router *chi.Mux
	store db.Store
}

func NewServer(config *config.AppConfig) *Server {
	r := chi.NewRouter()
	ctx := context.Background()

	conn, err := connectToDB(ctx, config.Dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	store := db.NewStore(conn)

	return &Server{
		config: config,
		router: r,
		store: store,
	}
}

func connectToDB(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	fmt.Println("Database connected successfully!")

	return pool, nil
}