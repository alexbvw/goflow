package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CorsMiddleware(a *fiber.App) {
	a.Use(
		// Add CORS to each route.
		cors.New(),
	)
}
