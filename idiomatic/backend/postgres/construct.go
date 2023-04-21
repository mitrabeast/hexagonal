package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgres(config *Config) *sql.DB {
	db, err := sql.Open("postgres", config.DSN())
	if err != nil {
		log.Fatalf("app : can't connect to database :: %v", err)
	}
	db.SetMaxOpenConns(config.MaxOpenConns)
	return db
}
