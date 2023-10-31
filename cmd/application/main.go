package main

import (
	"errors"
	"os"

	"github.com/eduardor2m/questao-certa/internal/adapters/delivery/http"
	"github.com/eduardor2m/questao-certa/tools/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if os.Getenv("DEVELOPMENT") == "true" || os.Getenv("DEVELOPMENT") == "" {
		if err != nil {
			logger.Fatal("Error loading .env file", errors.New("error loading .env file"))
		}
		logger.Info("Running in development mode")
	} else {
		logger.Info("Running in production mode or docker mode")
	}

	api := http.NewAPI(&http.Options{})
	api.Serve()
}
