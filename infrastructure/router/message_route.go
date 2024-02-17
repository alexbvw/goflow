package router

import (
	"goflow/interface/controllers"
	"goflow/interface/middleware"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

// MessageRoutes func for describe group of private routes.
func MessageRoutes(a *fiber.App) {

	// Create routes group.
	message := a.Group("/v1/message")

	//Message Statistics Routes
	message.Get("/count", middleware.UserProtected(), controllers.GetMessagesCount)

	//Message Routes
	message.Patch("/:id", controllers.UpdateMessage)
	message.Delete("/:id", controllers.DeleteMessage)
	message.Post("/send", middleware.UserProtected(), controllers.CreateMessage)

	message.Get("/all", middleware.UserProtected(), controllers.GetMessages)
	message.Get("/single/:id", middleware.UserProtected(), controllers.GetMessage)

	message.Get("/chatroom", adaptor.HTTPHandler(controllers.Handler(controllers.ChatroomMessagesHandler)))

}
