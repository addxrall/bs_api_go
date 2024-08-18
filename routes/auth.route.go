package routes

import (
	"github.com/addxrall/bs_api_go/services"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(api fiber.Router) {
	api.Post("/register", services.Register)
	api.Post("/login", services.Login)
	api.Post("/logout", services.Logout)
	api.Get("/session", services.Session)
}
