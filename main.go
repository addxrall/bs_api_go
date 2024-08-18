package main

import (
	"context"
	"log"
	"os"

	"github.com/addxrall/bs_api_go/middleware"
	"github.com/addxrall/bs_api_go/routes"
	"github.com/addxrall/bs_api_go/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v5"
)

func main() {
	app := fiber.New()
	dbString := os.Getenv("GOOSE_DBSTRING")
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, dbString)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	app.Use(logger.New())

	services.InitServices(conn)

	api := app.Group("/api")

	routes.AuthRoutes(api)
	api.Use(middleware.AuthCheckHandler)
	routes.UserRoutes(api)

	log.Fatal(app.Listen(":2137"))
}
