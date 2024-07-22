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
	webflow.Get("/site/:id", controllers.FetchSiteHandler)
	webflow.Get("/collections", controllers.FetchCollectionsHandler)
	webflow.Get("/collections/:collectionId/items/:itemId", controllers.FetchCollectionItemHandler)
	webflow.Get("/collection/:id", controllers.FetchCollectionHandler)
	webflow.Get("/items", controllers.FetchCollectionItemsHandler)

}
