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
	webflow.Post("/sites/:siteId/assets", controllers.UploadAssetHandler)

	webflow.Get("/user", controllers.GetAuthorizedUserInfo)

	webflow.Get("/collections", controllers.FetchCollectionsHandler)
	webflow.Get("/collections/:collectionId/items/:itemId", controllers.FetchCollectionItemHandler)
	webflow.Get("/collection/:id", controllers.FetchCollectionHandler)

	webflow.Post("/items/:collectionId", controllers.CreateCollectionItemHandler)
	webflow.Get("/items", controllers.FetchCollectionItemsHandler)
	webflow.Put("/collection/:collectionId/items", controllers.UpdateCollectionItemsHandler)
	webflow.Post("/collections/:collectionId/items", controllers.PublishCollectionItemHandler)
}
