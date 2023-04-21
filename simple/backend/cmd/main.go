package main

import (
	"os"

	"idiomatic/config"
	"idiomatic/postgres"
	"idiomatic/server"
)

func main() {
	appConfig := config.Load(getConfigFilename())
	db := postgres.NewPostgres(appConfig.Postgres)
	defer db.Close()
	web := server.NewWebServer(appConfig.Server.Address(), db)
	web.Run()
}

func getConfigFilename() string {
	if len(os.Args) >= 2 {
		return os.Args[1]
	}
	return "config.json"
}
