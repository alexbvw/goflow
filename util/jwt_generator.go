package util

import (
	"goflow/model"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

// GenerateNewAccessToken func for generate a new Access token.
func GenerateNewAdminAccessToken(checkexists model.CheckIdentityExist) (string, error) {
	godotenv.Load()
	// Set secret key from .env file.
	secret := os.Getenv("JWT_SECRET_KEY")
	// Set expires minutes count for secret key from .env file.
	minutesCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	// Create a new claims.
	claims := jwt.MapClaims{}

	// Set public claims:
	claims["id"] = checkexists.ID
	claims["role"] = checkexists.Role
	claims["full_name"] = checkexists.FullName
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()
	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return t, nil
}

func GenerateNewUserAccessToken(checkexists model.CheckIdentityExist) (string, error) {
	godotenv.Load()
	// Set secret key from .env file.
	secret := os.Getenv("JWT_SECRET_KEY")
	// Set expires minutes count for secret key from .env file.
	minutesCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	// Create a new claims.
	claims := jwt.MapClaims{}

	// Set public claims:
	claims["id"] = checkexists.ID
	claims["role"] = checkexists.Role
	claims["full_name"] = checkexists.FullName
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()
	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return t, nil
}
