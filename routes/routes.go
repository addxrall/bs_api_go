package routes

import (
	auth "github.com/addxrall/bs_api_go/services"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/register", auth.Register)
	api.Post("/login", auth.Login)
	api.Post("/logout", auth.Logout)
	api.Get("/session", auth.Session)
}

func handleSomething(c *fiber.Ctx) error {
	return c.SendString("This is /api/test")
}
