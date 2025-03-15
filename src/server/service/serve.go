package server_service

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pterm/pterm"
	"github.com/thiago-dsd/fastfood-core-api/src/config/env"
	order_router "github.com/thiago-dsd/fastfood-core-api/src/order/router"
	user_router "github.com/thiago-dsd/fastfood-core-api/src/user/router"
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
	order_router.Route(app)

	err := app.Listen(fmt.Sprintf(":%s", env.ServerPort))
	pterm.DefaultLogger.Fatal(
		fmt.Sprintf("%v", err),
	)
}
