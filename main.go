package main

import (
	"goflow/config"
	"goflow/infrastructure/router"
	"goflow/interface/controllers"
	"goflow/interface/middleware"
	"goflow/util"
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

	controllers.AssetMigration()
	controllers.MessageMigration()
	controllers.IdentityMigration()
	controllers.ChatroomMigration()

	router.WebflowRoutes(app)
	router.GitRoutes(app)
	router.AcountRoutes(app)
	router.IdentityRoutes(app)
	router.AssetRoutes(app)
	router.ChatroomRoutes(app)
	router.MessageRoutes(app)
	router.NotFoundRoute(app)

	util.StartServer(app)
}
