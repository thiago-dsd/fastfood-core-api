package user_handler

import (
	"github.com/gofiber/fiber/v2"
	common_model "github.com/thiago-dsd/fastfood-core-api/src/common/model"
	"github.com/thiago-dsd/fastfood-core-api/src/repository"
	user_entity "github.com/thiago-dsd/fastfood-core-api/src/user/entity"
	user_model "github.com/thiago-dsd/fastfood-core-api/src/user/model"
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
