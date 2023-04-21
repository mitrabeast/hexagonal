package server

import (
	"database/sql"
	"log"
	"net/http"

	"idiomatic/api"
	"idiomatic/postgres"
	"idiomatic/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	address string
	db      *sql.DB
}

func NewWebServer(address string, db *sql.DB) *WebServer {
	return &WebServer{
		address: address,
		db:      db,
	}
}

func (w *WebServer) Run() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(setContentType("application/json"))
	mountRoutes(r, w.db)

	log.Printf("app :: starting on %s", w.address)
	if err := http.ListenAndServe(w.address, r); err != nil {
		log.Fatalf("app :: can't listen and serve :: %v", err)
	}
}

func mountRoutes(r chi.Router, db *sql.DB) {
	api.NewUserController(
		user.NewUserService(
			postgres.NewUserRepo(db),
		),
	).MountRoutes(r)
}
