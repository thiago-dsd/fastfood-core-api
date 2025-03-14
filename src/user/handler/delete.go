package user_handler

import (
	"github.com/gofiber/fiber/v2"
	common_model "github.com/thiago-dsd/fastfood-core-api/src/common/model"
	"github.com/thiago-dsd/fastfood-core-api/src/repository"
	user_entity "github.com/thiago-dsd/fastfood-core-api/src/user/entity"
)

// @Summary		Delete current user
// @Description	Deletes the user who made the request
// @Tags			User
// @Success		204 "User deleted successfully"
// @Router			/user/me [delete]
// @Security		ApiKeyAuth
func DeleteCurrentUser(c *fiber.Ctx) error {
	// Retrieve the user from the context
	user := c.Locals("user").(*user_entity.User)

	if err := repository.DeleteById[user_entity.User](user.Id, nil); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			common_model.NewApiError("unable to delete user", err, "repository").Send(),
		)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// @Summary		Delete user by ID
// @Description	Deletes a user by their ID (only accessible by admins). You cannot delete su@user
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			body	body		common_model.RequiredId	true	"User ID to delete"
// @Success		204 "User deleted successfully"
// @Router			/user [delete]
// @Security		ApiKeyAuth
func DeleteUserByID(c *fiber.Ctx) error {
	var reqBody common_model.RequiredId
	if err := c.BodyParser(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			common_model.NewParseJsonError(err).Send(),
		)
	}

	user, err := repository.First(
		user_entity.User{
			Audit: common_model.Audit{
				Id: reqBody.Id,
			},
		},
		0, nil, nil, "", nil,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			common_model.NewApiError("unable to find user", err, "repository").Send(),
		)
	}
	if user.Email == "su@sudo" {
		return c.Status(fiber.StatusUnauthorized).JSON(
			common_model.NewApiError("one cannot delete su@sudo user", err, "handler").Send(),
		)
	}

	err = repository.DeleteById[user_entity.User](reqBody.Id, nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			common_model.NewApiError("unable to delete user", err, "repository").Send(),
		)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
