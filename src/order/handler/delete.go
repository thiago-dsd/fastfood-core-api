package order_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thiago-dsd/fastfood-core-api/src/common/model"
	order_entity "github.com/thiago-dsd/fastfood-core-api/src/order/entity"
	"github.com/thiago-dsd/fastfood-core-api/src/repository"
)

// @Summary		Delete an order
// @Description	Deletes a specific order based on the ID sent in the request body
// @Tags			Order
// @Accept			json
// @Produce			json
// @Param			body	body		model.RequiredId	true	"Order ID to delete"
// @Success			204		"Order deleted successfully"
// @Router			/api/orders/delete [delete]
// @Security		ApiKeyAuth
func DeleteOrder(c *fiber.Ctx) error {
	// Retrieve the userId and role from the context (set by the middleware)
	userId := c.Locals("userId").(string)
	role := c.Locals("role").(string)

	// Define the model that will be passed in the body to delete the order
	var reqBody model.RequiredId
	if err := c.BodyParser(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.NewParseJsonError(err).Send())
	}

	// Retrieve the order from the database by ID
	order, err := repository.First(
		order_entity.Order{
			Audit: model.Audit{
				Id: reqBody.Id,
			},
		},
		0, nil, nil, "", nil,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			model.NewApiError("unable to find order", err, "repository").Send(),
		)
	}

	// If the user is not an admin, ensure that they can only delete their own orders
	if role != "admin" && order.UserId != userId {
		return c.Status(fiber.StatusUnauthorized).JSON(
			model.NewApiError("You are not authorized to delete this order", nil, "handler").Send(),
		)
	}

	// Delete the order
	err = repository.DeleteById[order_entity.Order](reqBody.Id, nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			model.NewApiError("unable to delete order", err, "repository").Send(),
		)
	}

	// Return a success status with no content
	return c.SendStatus(fiber.StatusNoContent)
}
