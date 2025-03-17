package order_handler

import (
	"github.com/gofiber/fiber/v2"
	common_model "github.com/thiago-dsd/fastfood-core-api/src/common/model"
	"github.com/thiago-dsd/fastfood-core-api/src/database"
	order_entity "github.com/thiago-dsd/fastfood-core-api/src/order/entity"
	order_model "github.com/thiago-dsd/fastfood-core-api/src/order/model"
	user_entity "github.com/thiago-dsd/fastfood-core-api/src/user/entity"
)

// @Summary		Creates a new order
// @Description	Creates a new order for the user
// @Tags			Order
// @Accept			json
// @Produce		json
// @Param			order	body		order_model.Create	true	"Order data"
// @Success		201		{object}	order_entity.Order		"Created order"
// @Router			/order [post]
// @Security		ApiKeyAuth
func CreateOrder(c *fiber.Ctx) error {
	// Parse the request body
	var newOrder order_model.Create
	if err := c.BodyParser(&newOrder); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common_model.NewParseJsonError(err).Send())
	}

	// Retrieve the userId from the context (set by the middleware)
	user := c.Locals("user").(*user_entity.User)

	// Convert newOrder.Items to order_entity.Items
	items := order_entity.Items(newOrder.Items)

	// Create the new order, using the userId from the context
	newEntity := order_entity.Order{
		UserId:      user.Id, // Use the userId from the context (not from the request body)
		Description: newOrder.Description,
		Items:       &items,
	}

	// Save the new order to the database
	if err := database.Connection().Create(&newEntity).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			common_model.NewApiError("unable to create order", err, "gorm.io/gorm").Send(),
		)
	}

	// Return the created order
	return c.Status(fiber.StatusCreated).JSON(newEntity)
}