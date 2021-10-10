package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/office-aska/lwgc/pkg/environ"
	"github.com/office-aska/lwgc/server/app"
)

func main() {
	if environ.IsLocal() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	app.Bootstrap()
}
