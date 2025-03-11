package auth_service

import (
	"errors"

	"github.com/cogniia/core-api-template/src/database"
	user_entity "github.com/cogniia/core-api-template/src/user/entity"
	"github.com/golang-jwt/jwt/v4"
)

// Retrieves the user associated with the provided JWT token.
func GetFromToken(tokenString string) (*user_entity.User, error) {
	token, err := ParseToken(tokenString)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to parse claims from token")
	}

	userID := claims["sub"].(string)
	var user user_entity.User

	// Find user in the database based on userID
	err = database.Connection().First(&user, "id = ?", userID).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
