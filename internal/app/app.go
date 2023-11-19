package app

import (
	"dosLAbTest/config"
	"dosLAbTest/internal/usecase"
	"dosLAbTest/pkg/httpserver"
	"dosLAbTest/pkg/postgres"
	"log"
)

func Run(cfg config.Config) {
	pg, err := postgres.New(cfg)
	if err != nil {
		log.Fatalf("cannot create db %v", err)
	}

	server, err := httpserver.NewServer(cfg, *pg)
	if err != nil {
		log.Fatalf("cannot create server %v", err)
	}

	go usecase.UpdateStatistics(pg)
	if err = server.Start(cfg.ServerAddress); err != nil {
		log.Fatalf("cannot start server %v", err)
	}

}
