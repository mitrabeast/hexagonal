package config

import (
	"encoding/json"
	"log"
	"os"
)

type App struct {
	Web      *Web      `json:"web"`
	Database *Database `json:"database"`
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
