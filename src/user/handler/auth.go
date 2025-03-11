package user_handler

import (
	auth_model "github.com/cogniia/core-api-template/src/auth/model"
	auth_service "github.com/cogniia/core-api-template/src/auth/service"
	common_model "github.com/cogniia/core-api-template/src/common/model"
	"github.com/gofiber/fiber/v2"
)

// @Summary		Login
// @Description	Authenticates a user and returns access and refresh tokens.
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			login	body		auth_model.LoginRequest		true	"Login data"
// @Success		200		{object}	auth_model.LoginResponse	"Login successful"
// @Router			/user/auth/token [post]
func LoginHandler(c *fiber.Ctx) error {
	var loginRequest auth_model.LoginRequest
	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			common_model.NewParseJsonError(err).Send(),
		)
	}

	loginResponse, err := auth_service.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(
			common_model.NewApiError(
				"unable to login",
				err,
				"auth_service",
			).Send(),
		)
	}

	return c.Status(fiber.StatusOK).JSON(loginResponse)
}

// @Summary		Refresh Access Token
// @Description	Refreshes the access token using a refresh token.
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			refresh	body		auth_model.RefreshRequest	true	"Refresh token data"
// @Success		200		{object}	auth_model.LoginResponse	"Refresh successful"
// @Router			/user/auth/refresh-token [post]
func RefreshTokenHandler(c *fiber.Ctx) error {
	var refreshRequest auth_model.RefreshRequest
	if err := c.BodyParser(&refreshRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common_model.NewParseJsonError(err).Send())
	}

	accessToken, err := auth_service.RefreshToken(refreshRequest.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(
			common_model.NewApiError(
				"unable to refresh token",
				err,
				"auth_service",
			).Send(),
		)
	}

	return c.Status(fiber.StatusOK).JSON(accessToken)
}
