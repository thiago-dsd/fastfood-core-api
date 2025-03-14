package user_router

import (
	"github.com/gofiber/fiber/v2"
	auth_middleware "github.com/thiago-dsd/fastfood-core-api/src/auth/middleware"
	user_handler "github.com/thiago-dsd/fastfood-core-api/src/user/handler"
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
