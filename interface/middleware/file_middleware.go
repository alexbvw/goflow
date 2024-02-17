package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"

	"github.com/gobuffalo/packr/v2"
)

func FileMiddleware(a *fiber.App) {
	a.Use(
		"/images", filesystem.New(filesystem.Config{
			Root:         packr.New("Images ", "/images"),
			PathPrefix:   "",
			Browse:       false,
			Index:        "",
			MaxAge:       0,
			NotFoundFile: "",
		}),
	)
}
