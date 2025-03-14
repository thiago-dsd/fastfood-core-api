package order_router

import (
	"github.com/gofiber/fiber/v2"
	auth_middleware "github.com/thiago-dsd/fastfood-core-api/src/auth/middleware"
)

func Route(app *fiber.App) {
    group := app.Group("/order")

    mainRoutes(group)
}

func mainRoutes(group fiber.Router) {
	orderGroup.Post("/", middleware.UserMiddleware, handler.CreateOrder)
	orderGroup.Get("/", middleware.UserMiddleware, handler.GetUserOrders)
	orderGroup.Get("/by-id", auth_middleware.UserMiddleware, handler.GetOrderByID)
	orderGroup.Put("/update", middleware.UserMiddleware, handler.UpdateOrder)
	orderGroup.Delete("/delete", middleware.UserMiddleware, handler.DeleteOrder)
}
