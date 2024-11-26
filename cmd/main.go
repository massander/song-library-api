package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"song-library-api/internal/api"
	"song-library-api/internal/storage/postgres"
)

func main() {
	godotenv.Load()

	storage, err := postgres.New(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	if err := storage.Migrate(os.Getenv("MIGRATIONS_FOLDER")); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Use(logger.New())

	v1 := app.Group("/api/v1")

	apiService := api.NewAPIService(storage)
	apiService.RegisterGateway(v1)

	app.Listen(":3000")
}
