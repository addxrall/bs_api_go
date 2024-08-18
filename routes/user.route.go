package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(api fiber.Router) {
	users := api.Group("/users")

	users.Get("/pr", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})
}
