package user_handler

import (
	common_model "github.com/cogniia/core-api-template/src/common/model"
	"github.com/cogniia/core-api-template/src/repository"
	user_entity "github.com/cogniia/core-api-template/src/user/entity"
	user_model "github.com/cogniia/core-api-template/src/user/model"
	"github.com/gofiber/fiber/v2"
)

// @Summary		Get users paginated
// @Description	Returns a paginated list of users
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			user	query		user_model.QueryPaginated	true	"Pagination and query parameters"
// @Success		200			{array}		user_entity.User			"List of users"
// @Router			/user [get]
// @Security		ApiKeyAuth
func Get(c *fiber.Ctx) error {
	query := new(user_model.QueryPaginated)
	if err := c.QueryParser(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common_model.NewParseJsonError(err).Send())
	}

	users, err := repository.GetPaginated(
		user_entity.User{
			Name:  query.Name,
			Email: query.Email,
			Audit: common_model.Audit{Id: query.Id},
			Role:  query.Role,
		},
		&query.Paginate,
		&query.DateOrder,
		&query.DateWhere,
		"",
		nil,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common_model.NewApiError("unable to get paginated", err, "repository").Send())
	}

	return c.Status(fiber.StatusOK).JSON(users)
}
