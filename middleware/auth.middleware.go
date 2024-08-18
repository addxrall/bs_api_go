package middleware

import (
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = os.Getenv("JWT")

func AuthCheckHandler(c *fiber.Ctx) error {
	tokenString := c.Cookies("token")
	if tokenString == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if err != nil || !token.Valid {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	return c.Next()
}
