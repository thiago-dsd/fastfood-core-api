package order_handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	common_model "github.com/thiago-dsd/fastfood-core-api/src/common/model" // Ajuste conforme o seu repositório
	order_entity "github.com/thiago-dsd/fastfood-core-api/src/order/entity" // Ajuste conforme o seu repositório
	"github.com/thiago-dsd/fastfood-core-api/src/repository"
	user_entity "github.com/thiago-dsd/fastfood-core-api/src/user/entity"
)

// @Summary		Get order by ID
// @Description	Returns a specific order by ID, based on the user's role
// @Tags			Order
// @Accept			json
// @Produce			json
// @Param			body	body		common_model.RequiredId	true	"Order ID to get"
// @Success			200		{object}	order_entity.Order	"Order details"
// @Failure			403		"Forbidden: User is not authorized to access this order"
// @Failure			404		"Order not found"
// @Router			/order/by-id [get]
// @Security		ApiKeyAuth
func GetOrderByID(c *fiber.Ctx) error {
	// Define a model to receive the order ID from the body
	var reqBody common_model.RequiredId
	if err := c.BodyParser(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common_model.NewParseJsonError(err).Send())
	}

	// Retrieve the authenticated user from the context
	user, ok := c.Locals("user").(*user_entity.User) // Ajuste conforme o tipo do usuário no contexto
	if !ok || user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(
			common_model.NewApiError("failed to retrieve user from context locals", errors.New("invalid conversion to user entity"), "handler").Send(),
		)
	}

	// Fetch the order by its ID
	order, err := repository.First(
		order_entity.Order{
			Audit: common_model.Audit{
				Id: reqBody.Id,
			},
		},
		0, nil, nil, "", nil,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common_model.NewApiError("unable to find order", err, "repository").Send())
	}

	// Check if the user is allowed to view this order
	if user.Id != order.UserId {
		return c.Status(fiber.StatusForbidden).JSON(
			common_model.NewApiError("user not authorized to access this order", errors.New("permission denied"), "handler").Send(),
		)
	}

	// Return the order details
	return c.Status(fiber.StatusOK).JSON(order)
}
