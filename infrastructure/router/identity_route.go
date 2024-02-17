package router

import (
	"goflow/interface/controllers"
	"goflow/interface/middleware"

	"github.com/gofiber/fiber/v2"
)

// IdentityRoutes func for describe group of private routes.
func IdentityRoutes(a *fiber.App) {

	//Create routes group.
	identity := a.Group("/v1/identity")

	//Identity Routes
	identity.Get("/all", controllers.GetIdentities)
	identity.Get("/single/:id", controllers.GetIdentity)
	identity.Get("/count", controllers.GetIdentitiesCount)

	identity.Put("/update/:id", middleware.AdminProtected(), controllers.UpdateIdentity)
	identity.Delete("/remove/:id", middleware.AdminProtected(), controllers.DeleteIdentity)

	//Identity Authentication Routes
	identity.Post("/login", controllers.Login)
	identity.Post("/register", controllers.Register)

}
