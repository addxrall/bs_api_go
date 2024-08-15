package main

import (
	"github.com/addxrall/bs_api_go/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	routes.SetupRoutes(app)

	app.Listen(":2137")
}
