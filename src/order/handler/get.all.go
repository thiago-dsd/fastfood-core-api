package order_handler

import (
	"github.com/gofiber/fiber/v2"
	common_model "github.com/thiago-dsd/fastfood-core-api/src/common/model"
	order_entity "github.com/thiago-dsd/fastfood-core-api/src/order/entity"
	order_model "github.com/thiago-dsd/fastfood-core-api/src/order/model"
	"github.com/thiago-dsd/fastfood-core-api/src/repository"
	user_entity "github.com/thiago-dsd/fastfood-core-api/src/user/entity"
)

// @Summary		Get all orders for the current user
// @Description	Retrieves a paginated list of orders for the user making the request
// @Tags			Order
// @Accept			json
// @Produce			json
// @Param			body	body		model.QueryPaginated	true	"Query parameters for pagination and filtering orders"
// @Success			200		{array}		order_entity.Order	"List of orders"
// @Router			/api/orders [get]
// @Security		ApiKeyAuth
func GetAllOrders(c *fiber.Ctx) error {
	// Retrieve the userId from the context using the middleware
	user := c.Locals("user").(*user_entity.User)

	// Parse the query parameters from the body
	query := new(order_model.QueryPaginated)
	if err := c.BodyParser(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common_model.NewParseJsonError(err).Send())
	}

	// Query for paginated orders based on userId and query parameters
	orders, err := repository.GetPaginated(
		order_entity.Order{
			UserID: user.Id, // Use the userId from the context
		},
		&query.Paginate,
		&query.DateOrder,
		&query.DateWhere,
		"",
		nil,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common_model.NewApiError("unable to get paginated orders", err, "repository").Send())
	}

	// Return the list of orders
	return c.Status(fiber.StatusOK).JSON(orders)
}
