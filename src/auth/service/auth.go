package auth_service

import (
	"errors"
	"fmt"
	"time"

	auth_model "github.com/cogniia/core-api-template/src/auth/model"
	"github.com/cogniia/core-api-template/src/config/env"
	crypto_service "github.com/cogniia/core-api-template/src/crypto/service"
	"github.com/cogniia/core-api-template/src/database"
	user_entity "github.com/cogniia/core-api-template/src/user/entity"
	"github.com/golang-jwt/jwt/v4"
)

func Login(email, password string) (*auth_model.LoginResponse, error) {
	var user user_entity.User
	err := database.Connection().First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}

	if !crypto_service.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("incorrect password")
	}

	// Generate JWT tokens
	accessToken, err := generateToken(user.Id.String(), 1*time.Hour)
	if err != nil {
		return nil, err
	}

	refreshToken, err := generateToken(user.Id.String(), 7*24*time.Hour)
	if err != nil {
		return nil, err
	}

	return &auth_model.LoginResponse{
		Token:        accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func RefreshToken(refreshToken string) (*auth_model.LoginResponse, error) {
	// Parse the refresh token
	token, err := ParseToken(refreshToken)
	if err != nil {
		return nil, err
	}

	// Get user ID from token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to parse claims from token")
	}
	userID := claims["sub"].(string)

	// Generate new access token
	accessToken, err := generateToken(userID, 1*time.Hour)
	if err != nil {
		return nil, err
	}

	return &auth_model.LoginResponse{
		Token:        accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func generateToken(userID string, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(duration).Unix(),
		"iss": "core-api-template", // Issuer
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(env.JwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(env.JwtSecret), nil
		},
	)
}
