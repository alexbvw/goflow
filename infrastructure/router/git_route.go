package router

import (
	"goflow/interface/controllers"

	"github.com/gofiber/fiber/v2"
)

// GitRoutes func for describe group of private routes.
func GitRoutes(a *fiber.App) {

	//Create routes group.
	repo := a.Group("/v1/repo")

	//Git Repo Routes
	repo.Get("/all", controllers.GetRepos)
}
