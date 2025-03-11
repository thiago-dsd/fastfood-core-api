package user_handler

import (
	"errors"

	common_model "github.com/cogniia/core-api-template/src/common/model"
	user_entity "github.com/cogniia/core-api-template/src/user/entity"
	"github.com/gofiber/fiber/v2"
)

// @Summary		Gets current user
// @Description	Returns the currently authenticated user
// @Tags			User
// @Accept			json
// @Produce		json
// @Success		200	{object}	user_entity.User	"Current user"
// @Router			/user/me [get]
// @Security		ApiKeyAuth
func GetCurrentUser(c *fiber.Ctx) error {
	// Retrieve the authenticated user from the context
	user, ok := c.Locals("user").(*user_entity.User)
	if !ok || user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(
			common_model.NewApiError("failed to retrieve user from context locals", errors.New("invalid convertion to type user_entity.User"), "handler").Send(),
		)
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
