package order_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thiago-dsd/fastfood-core-api/src/common/model"
	"github.com/thiago-dsd/fastfood-core-api/src/database"
	order_entity "github.com/thiago-dsd/fastfood-core-api/src/order/entity"
	order_model "github.com/thiago-dsd/fastfood-core-api/src/order/model"
)

// @Summary		Creates a new order
// @Description	Creates a new order for the user
// @Tags			Order
// @Accept			json
// @Produce		json
// @Param			order	body		order_model.Create	true	"Order data"
// @Success		201		{object}	order_entity.Order		"Created order"
// @Router			/order [post]
func CreateOrder(c *fiber.Ctx) error {
	// Parse the request body
	var newOrder order_model.Create
	if err := c.BodyParser(&newOrder); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.NewParseJsonError(err).Send())
	}

	// Retrieve the userId from the context (set by the middleware)
	userId := c.Locals("userId").(string)

	// Create the new order, using the userId from the context
	newEntity := order_entity.Order{
		UserID:      userId, // Use the userId from the context (not from the request body)
		Description: newOrder.Description,
		Items:       newOrder.Items,
	}

	// Save the new order to the database
	if err := database.Connection().Create(&newEntity).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			model.NewApiError("unable to create order", err, "gorm.io/gorm").Send(),
		)
	}

	// Return the created order
	return c.Status(fiber.StatusCreated).JSON(newEntity)
}
