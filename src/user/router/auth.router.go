package user_router

import (
	"github.com/gofiber/fiber/v2"
	user_handler "github.com/thiago-dsd/fastfood-core-api/src/user/handler"
)

func authRoutes(group fiber.Router) {
	authGroup := group.Group("/auth")
	authGroup.Post("/token", user_handler.LoginHandler)
	authGroup.Post("/refresh-token", user_handler.RefreshTokenHandler)
}
