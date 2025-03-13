package server_service

import (
	"fmt"

	"github.com/cogniia/core-api-template/src/config/env"
	user_router "github.com/cogniia/core-api-template/src/user/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pterm/pterm"
)

func Serve() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type,Authorization",
		AllowCredentials: true,
	}))
	
	makeDocs(app)
	user_router.Route(app)

	err := app.Listen(fmt.Sprintf(":%s", env.ServerPort))
	pterm.DefaultLogger.Fatal(
		fmt.Sprintf("%v", err),
	)
}
