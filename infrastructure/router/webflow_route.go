package router

import (
	"goflow/interface/controllers"

	"github.com/gofiber/fiber/v2"
)

// WebflowRoutes func for describe group of private routes.
func WebflowRoutes(a *fiber.App) {

	//Create routes group.
	webflow := a.Group("/v1/webflow")

	webflow.Post("/token", controllers.RequestTokenHandler)
	webflow.Get("/sites", controllers.FetchSitesHandler)
	webflow.Get("/collections", controllers.FetchCollectionsHandler)
	webflow.Get("/items", controllers.FetchCollectionItemsHandler)

}
