package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"rest/api/internals/cache"
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
	router := chi.NewRouter()
	ctx := context.Background()

	// init cache
	cache.Init(config)
	cacheErr := cache.Set("Name", "superb important value!")
	fmt.Println("SetCache error: ", cacheErr)

	conn, err := connectToDB(ctx, config.Dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	store := db.NewStore(conn)

	return &Server{
		config: config,
		router: router,
		store: store,
	}
}

func (s *Server) Start(port string) {
	// fmt.Printf("Server struct: %v \n", s.router);

	SetupRoutes(s)

	// start and listen to server on port
	svr := &http.Server{
		Handler: s.router,
		Addr: ":" + port,
	}
	err := svr.ListenAndServe()
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}

	fmt.Printf("Server started and running on http://localhost%s \n", port);
}

func connectToDB(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	fmt.Println("Database connected successfully!")

	return pool, nil
}