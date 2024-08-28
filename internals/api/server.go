package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
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

func (s *Server) Start(port string) {
	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/text")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello world!"))
	})

	// run server
	listenPort := fmt.Sprintf(":%v", port)
	log.Fatal(http.ListenAndServe(listenPort, s.router))

	fmt.Printf("Server started and running on http://localhost%s \n", listenPort);
}

func connectToDB(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	fmt.Println("Database connected successfully!")

	return pool, nil
}