package order_handler

import (
	"github.com/gofiber/fiber/v2"
	common_model "github.com/thiago-dsd/fastfood-core-api/src/common/model"
	order_entity "github.com/thiago-dsd/fastfood-core-api/src/order/entity"
	"github.com/thiago-dsd/fastfood-core-api/src/repository"
	user_entity "github.com/thiago-dsd/fastfood-core-api/src/user/entity"
)

// @Summary		Delete order by ID
// @Description	Deletes an order by its ID. Users can only delete their own orders.
// @Tags			Order
// @Accept			json
// @Produce		json
// @Param			body	body		common_model.RequiredId	true	"Order ID to delete"
// @Success		204	"Order deleted successfully"
// @Router			/order [delete]
// @Security		ApiKeyAuth
func DeleteOrderByID(c *fiber.Ctx) error {
	var reqBody common_model.RequiredId

	if err := c.BodyParser(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			common_model.NewParseJsonError(err).Send(),
		)
	}

	// Retrieve the authenticated user from the middleware
	user := c.Locals("user").(*user_entity.User)

	// Fetch the order from the database using the provided ID
	order, err := repository.First(
		order_entity.Order{
			Audit: common_model.Audit{
				Id: reqBody.Id,
			},
		},
		0, nil, nil, "", nil,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			common_model.NewApiError("unable to find order", err, "repository").Send(),
		)
	}

	// Check if the user is authorized to delete the order (users can only delete their own orders)
	if order.UserId != user.Id {
		return c.Status(fiber.StatusUnauthorized).JSON(
			common_model.NewApiError("You are not authorized to delete this order", nil, "handler").Send(),
		)
	}

	// Delete the order from the database
	err = repository.DeleteById[order_entity.Order](reqBody.Id, nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			common_model.NewApiError("unable to delete order", err, "repository").Send(),
		)
	}

	// Return HTTP 204 No Content on successful deletion
	return c.SendStatus(fiber.StatusNoContent)
}
