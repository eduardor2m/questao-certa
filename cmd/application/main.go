package main

import (
	"log"

	"github.com/eduardor2m/questao-certa/internal/adapters/delivery/http"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("./cmd/application/.env"); err != nil {
		log.Println("Error loading .env file")
	}

	api := http.NewAPI(&http.Options{})
	api.Serve()
}
