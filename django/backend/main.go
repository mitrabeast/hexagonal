package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"

	controller "django-hexagonal/api/http"
	"django-hexagonal/config"
	repo "django-hexagonal/hexagon/user/repo/postgres"
	userservice "django-hexagonal/hexagon/user/service"
)

func main() {
	appConfig := config.Load(getConfigFilename())

	db, err := sql.Open("postgres", appConfig.Database.DSN())
	if err != nil {
		log.Fatalf("app : can't connect to database :: %v", err)
	}
	defer db.Close()
	db.SetMaxOpenConns(appConfig.Database.MaxOpenConns)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(setContentType("application/json"))
	mountRoutes(r, db)

	log.Printf("app :: starting on %s", appConfig.Web.Address())
	if err := http.ListenAndServe(appConfig.Web.Address(), r); err != nil {
		log.Fatalf("app :: can't listen and serve :: %v", err)
	}
}

func getConfigFilename() string {
	if len(os.Args) >= 2 {
		return os.Args[1]
	}
	return "config.json"
}

func setContentType(contentType string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", contentType)
			next.ServeHTTP(w, r)
		})
	}
}

func mountRoutes(r chi.Router, db *sql.DB) {
	controller.NewUserController(
		userservice.NewUserService(
			repo.NewUserRepo(db),
		),
	).MountRoutes(r)
}
