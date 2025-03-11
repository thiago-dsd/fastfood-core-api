package user_router

import (
	user_handler "github.com/cogniia/core-api-template/src/user/handler"
	"github.com/gofiber/fiber/v2"
)

func authRoutes(group fiber.Router) {
	authGroup := group.Group("/auth")
	authGroup.Post("/token", user_handler.LoginHandler)
	authGroup.Post("/refresh-token", user_handler.RefreshTokenHandler)
}
