package server_service

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/pterm/pterm"
	_ "github.com/thiago-dsd/fastfood-core-api/docs"
	"github.com/thiago-dsd/fastfood-core-api/src/config/env"
)

func makeDocs(app *fiber.App) {
	app.Get("/docs/*", swagger.HandlerDefault)
	pterm.DefaultLogger.Info(
		fmt.Sprintf("Docs available at %s:%s/docs", env.ServerHost, env.ServerPort),
	)
}
