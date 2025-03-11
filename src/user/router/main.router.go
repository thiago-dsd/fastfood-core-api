package user_router

import (
	auth_middleware "github.com/cogniia/core-api-template/src/auth/middleware"
	user_handler "github.com/cogniia/core-api-template/src/user/handler"
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	group := app.Group("/user")

	mainRoutes(group)
	authRoutes(group)
}

func mainRoutes(group fiber.Router) {
	group.Get("/me", auth_middleware.UserMiddleware, user_handler.GetCurrentUser)
	group.Delete("/me", auth_middleware.UserMiddleware, user_handler.DeleteCurrentUser)
	group.Put("/me", auth_middleware.UserMiddleware, user_handler.UpdateCurrentUser)

	group.Post("/", user_handler.CreateUser)
}
