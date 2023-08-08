package main

import (
	"log"

	"github.com/alexlzrv/note-service/config"
	"github.com/alexlzrv/note-service/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.RunServer(cfg)
}
