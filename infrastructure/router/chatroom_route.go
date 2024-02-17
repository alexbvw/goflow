package router

import (
	"goflow/interface/controllers"
	"goflow/interface/middleware"

	"github.com/gofiber/fiber/v2"
)

// ChatroomRoutes func for describe group of private routes.
func ChatroomRoutes(a *fiber.App) {

	//Create routes group.
	chatroom := a.Group("/v1/chatroom")

	//Chatroom Statistics Routes
	chatroom.Get("/count", middleware.UserProtected(), controllers.GetChatroomsCount)

	//Chatroom Routes
	chatroom.Get("/all", middleware.UserProtected(), controllers.GetChatrooms)
	chatroom.Get("/single/:id", middleware.UserProtected(), controllers.GetChatroom)

	chatroom.Post("/", middleware.UserProtected(), controllers.CreateChatroom)
	chatroom.Patch("/:id", middleware.UserProtected(), controllers.UpdateChatroom)
	chatroom.Delete("/:id", middleware.UserProtected(), controllers.DeleteChatroom)

	chatroom.Get("/identity/:id", middleware.UserProtected(), controllers.GetChatroomsbyIdentityId)

}
