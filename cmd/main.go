package main

import (
	"dogma_test_task/internal"
	"dogma_test_task/internal/handler"
	"dogma_test_task/internal/repository"
	"dogma_test_task/internal/service"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// @title dogma_test_task API
// @version 1.0
// @BasePath /api/
func main() {

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	getEnv := func(key string) string {
		s, status := os.LookupEnv(key)
		if !status {
			log.Fatalf("Failed get env variable %s", key)
		}
		return s
	}

	db, err := repository.NewPostgresql(repository.Config{
		Host:     getEnv("POSTGRES_HOST"),
		Port:     getEnv("POSTGRES_PORT"),
		User:     getEnv("POSTGRES_USER"),
		Password: getEnv("POSTGRES_PASSWORD"),
		Database: getEnv("POSTGRES_DB"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := internal.Server{Server: *handlers.InitRoutes()}
	err = srv.Run()

	if err != nil {
		log.Fatalf("failed to run server: %s", err.Error())
	}

	log.Print("domainStorage Started")

}
