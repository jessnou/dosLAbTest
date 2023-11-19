package main

import (
	"dosLAbTest/config"
	"dosLAbTest/internal/app"
	"log"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load cofing", err)
	}

	app.Run(cfg)

}
