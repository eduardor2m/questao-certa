package main

import (
	"errors"

	"github.com/eduardor2m/questao-certa/internal/adapters/delivery/http"
	"github.com/eduardor2m/questao-certa/tools/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("", errors.New("Error loading .env file"))
	}

	api := http.NewAPI(&http.Options{})
	api.Serve()
}
