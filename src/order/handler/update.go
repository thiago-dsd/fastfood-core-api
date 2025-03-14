package order_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thiago-dsd/fastfood-core-api/src/common/model"
	order_entity "github.com/thiago-dsd/fastfood-core-api/src/order/entity"
	order_model "github.com/thiago-dsd/fastfood-core-api/src/order/model"
	"github.com/thiago-dsd/fastfood-core-api/src/repository"
)

// @Summary		Update an order
// @Description	Updates the details of an order based on the data sent in the request body
// @Tags			Order
// @Accept			json
// @Produce			json
// @Param			body	body		order_model.UpdateOrder	true	"Order data to update"
// @Success			200		{object}	order_entity.Order		"Order updated successfully"
// @Router			/api/orders/update [put]
// @Security		ApiKeyAuth
func UpdateOrder(c *fiber.Ctx) error {
	// Parse the body to extract the order data
	var orderData order_model.UpdateOrder
	if err := c.BodyParser(&orderData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.NewParseJsonError(err).Send())
	}

	// Retrieve the userId from the context (set by the middleware)
	userId := c.Locals("userId").(string)

	// Fetch the order using the provided ID in the request body
	order, err := repository.First(
		order_entity.Order{
			Audit: model.Audit{
				Id: orderData.Id,
			},
		},
		0, nil, nil, "", nil,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			model.NewApiError("unable to find order", err, "repository").Send(),
		)
	}

	// Ensure the order belongs to the current user
	if order.UserId != userId {
		return c.Status(fiber.StatusUnauthorized).JSON(
			model.NewApiError("You are not authorized to update this order", nil, "handler").Send(),
		)
	}

	// Prepare updated order data
	updatedOrder := order_entity.Order{
		UserId:     userId,  // No need to include userId in the request body, it's already set by middleware
		Description: orderData.Description,
		Items:      orderData.Items,
	}

	// Update the order using repository
	updatedOrderEntity, err := repository.Updates(
		updatedOrder,
		&order_entity.Order{
			Audit: model.Audit{
				Id: orderData.Id,
			},
		},
		nil,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			model.NewApiError("unable to update order", err, "repository").Send(),
		)
	}

	// Return the updated order
	return c.Status(fiber.StatusOK).JSON(updatedOrderEntity)
}
