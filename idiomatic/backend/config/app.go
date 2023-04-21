package config

import (
	"encoding/json"
	"log"
	"os"

	"idiomatic/postgres"
	"idiomatic/server"
)

type App struct {
	Server   *server.Config   `json:"web"`
	Postgres *postgres.Config `json:"postgres"`
}

func Load(filename string) *App {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("config :: can't read file :: %v", err)
	}
	var app App
	if err := json.Unmarshal(data, &app); err != nil {
		log.Fatalf("config :: can't parse file :: %v", err)
	}
	return &app
}
