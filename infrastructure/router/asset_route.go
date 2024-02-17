package router

import (
	"goflow/interface/controllers"
	"goflow/interface/middleware"

	"github.com/gofiber/fiber/v2"
)

// AssetRoutes func for describe group of private routes.
func AssetRoutes(a *fiber.App) {

	//Create routes group.
	asset := a.Group("/v1/asset")

	//Asset Routes
	asset.Get("/all", controllers.GetAssets)
	asset.Get("/single/:id", controllers.GetAsset)
	asset.Get("/count", controllers.GetAssetsCount)
	asset.Post("/", middleware.AdminProtected(), controllers.CreateAsset)

	// asset.Post("/upload/:id", middleware.AdminProtected(), controllers.UploadAssetFile)
	// asset.Get("/files/:id", middleware.AdminProtected(), controllers.GetAssetBucketObjectsList)
	// asset.Get("/file/:objectName", middleware.AdminProtected(), controllers.GetAssetBucketObject)

	asset.Put("/update/:id", middleware.AdminProtected(), controllers.UpdateAsset)
	asset.Delete("/remove/:id", middleware.AdminProtected(), controllers.DeleteAsset)

}
