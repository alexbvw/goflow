package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtMiddleware "github.com/gofiber/jwt/v2"
	"github.com/joho/godotenv"
)

// JWTProtected func for specify routes group with JWT authentication.
func UserProtected() func(*fiber.Ctx) error {
	godotenv.Load()
	// Create config for JWT authentication middleware.
	config := jwtMiddleware.Config{
		SigningMethod: "HS256",
		SigningKey:    []byte(os.Getenv("JWT_SECRET_KEY")),
		ContextKey:    "jwt", // used in private routes
		ErrorHandler:  jwtError,
	}
	return jwtMiddleware.New(config)
}

func AdminProtected() func(*fiber.Ctx) error {
	godotenv.Load()
	// Create config for JWT authentication middleware.
	config := jwtMiddleware.Config{
		SigningMethod: "HS512",
		SigningKey:    []byte(os.Getenv("JWT_SECRET_KEY")),
		ContextKey:    "jwt", // used in private routes
		ErrorHandler:  jwtError,
	}
	return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	// Return status 401 and failed authentication error.
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error":   false,
		"message": err.Error(),
	})
}
