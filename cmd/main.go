package main

import (
	"song-library-api/internal/api"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	
	v1 := app.Group("/api/v1")

	apiService := api.NewAPIService()
	apiService.RegisterGateway(v1)

	app.Listen(":3000")
}
