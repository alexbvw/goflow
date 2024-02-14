package main

import (
	"goflow/config"
	"goflow/controller"
	"goflow/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("loading: %v", err)
	}

	config := config.FiberConfig()
	app := fiber.New(config)
	middleware.Logger(app)
	middleware.CorsMiddleware(app)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hellos, ")
	})

	api := app.Group("/api")
	api.Post("/token", controller.RequestTokenHandler)
	api.Get("/sites", controller.FetchSitesHandler)
	api.Get("/collections", controller.FetchCollectionsHandler)
	api.Get("/items", controller.FetchCollectionItemsHandler)

	app.Listen(":3000")
}
