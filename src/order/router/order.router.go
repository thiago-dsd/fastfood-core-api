package order_router

import (
	"github.com/gofiber/fiber/v2"
	auth_middleware "github.com/thiago-dsd/fastfood-core-api/src/auth/middleware"
	handler "github.com/thiago-dsd/fastfood-core-api/src/order/handler"
)

func Route(app *fiber.App) {
    group := app.Group("/order")

    mainRoutes(group)
}

func mainRoutes(group fiber.Router) {
	group.Post("/", auth_middleware.UserMiddleware, handler.CreateOrder)
	group.Get("/", auth_middleware.UserMiddleware, handler.GetAllOrders)
	group.Get("/by-id", auth_middleware.UserMiddleware, handler.GetOrderByID)
	group.Put("/update", auth_middleware.UserMiddleware, handler.UpdateOrder)
	group.Delete("/delete", auth_middleware.UserMiddleware, handler.DeleteOrderByID)
}
