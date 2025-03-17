package order_handler

import (
	"github.com/gofiber/fiber/v2"
	common_model "github.com/thiago-dsd/fastfood-core-api/src/common/model"
	order_entity "github.com/thiago-dsd/fastfood-core-api/src/order/entity"
	order_model "github.com/thiago-dsd/fastfood-core-api/src/order/model"
	"github.com/thiago-dsd/fastfood-core-api/src/repository"
	user_entity "github.com/thiago-dsd/fastfood-core-api/src/user/entity"
)

// @Summary		Update an order
// @Description	Updates the details of an order based on the data sent in the request body
// @Tags			Order
// @Accept			json
// @Produce			json
// @Param			body	body		order_model.UpdateOrder	true	"Order data to update"
// @Success			200		{object}	order_entity.Order		"Order updated successfully"
// @Router			/order [put]
// @Security		ApiKeyAuth
func UpdateOrder(c *fiber.Ctx) error {
	// Parse the body to extract the order data
	var orderData order_model.UpdateOrder
	if err := c.BodyParser(&orderData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common_model.NewParseJsonError(err).Send())
	}

	// Retrieve the user from the context (set by the middleware)
	user, ok := c.Locals("user").(*user_entity.User)
	if !ok || user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(common_model.NewApiError("failed to retrieve user from context", nil, "handler").Send())
	}

	// Fetch the order using the provided ID in the request body
	order, err := repository.First(
		order_entity.Order{
			Audit: common_model.Audit{
				Id: user.Id,
			},
		},
		0, nil, nil, "", nil,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			common_model.NewApiError("unable to find order", err, "repository").Send(),
		)
	}

	// Ensure the order belongs to the current user
	if order.UserId != user.Id {
		return c.Status(fiber.StatusUnauthorized).JSON(
			common_model.NewApiError("You are not authorized to update this order", nil, "handler").Send(),
		)
	}

	// Prepare updated order data
	updatedOrder := order_entity.Order{
		UserId:     user.Id,  // The userID is already set by middleware, no need to send in the body
		Description: orderData.Description,
		Items:      &orderData.Items,
	}

	// Update the order using the repository
	updatedOrderEntity, err := repository.Updates(updatedOrder, &order_entity.Order{
		Audit: common_model.Audit{
			Id: order.Id, // Use the correct order ID here
		},
	}, nil)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			common_model.NewApiError("unable to update order", err, "repository").Send(),
		)
	}

	// Return the updated order
	return c.Status(fiber.StatusOK).JSON(updatedOrderEntity)
}
