package order_handler

import (
	"github.com/gofiber/fiber/v2"
	common_model "github.com/thiago-dsd/fastfood-core-api/src/common/model"
	order_entity "github.com/thiago-dsd/fastfood-core-api/src/order/entity"
	order_model "github.com/thiago-dsd/fastfood-core-api/src/order/model"
	"github.com/thiago-dsd/fastfood-core-api/src/repository"
	user_entity "github.com/thiago-dsd/fastfood-core-api/src/user/entity"
)

// @Summary		Get all orders for the current user (paginated)
// @Description	Retrieves a paginated list of orders for the user making the request
// @Tags			Order
// @Accept			json
// @Produce			json
// @Param			user	query		order_model.QueryPaginated	true	"Query parameters for pagination and filtering orders"
// @Success			200		{array}		order_entity.Order	"List of orders"
// @Router			/order [get]
// @Security		ApiKeyAuth
func GetAllOrders(c *fiber.Ctx) error {

	// Parse the query parameters from the URL
	query := new(order_model.QueryPaginated)
	if err := c.QueryParser(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common_model.NewParseJsonError(err).Send())
	}

	// Gets the authenticated user from the context
	user, ok := c.Locals("user").(*user_entity.User)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(common_model.NewApiError("Usuário não autenticado", nil, "handler").Send())
	}

	// Fetch user's paginated orders
	orders, err := repository.GetPaginated(
		order_entity.Order{
			UserId: user.Id,
		},
		&query.Paginate,  // Pagination parameters
		&query.DateOrder, // Sort by date
		&query.DateWhere, // Filter by date
		"",
		nil,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common_model.NewApiError("Erro ao buscar pedidos", err, "repository").Send())
	}

	// Returns found orders
	return c.Status(fiber.StatusOK).JSON(orders)
}
