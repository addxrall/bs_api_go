package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/test", handleSomething)
}

func handleSomething(c *fiber.Ctx) error {
	return c.SendString("This is /api/test")
}
