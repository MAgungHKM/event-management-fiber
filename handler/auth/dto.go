package auth

import (
	mJwt "event-management/middleware/jwt"
	"event-management/utils"
	"event-management/utils/env"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}

type AuthRequest struct {
	Username string `json:"username" validate:"required" extensions:"x-order=1"`
	Password string `json:"password" validate:"required" extensions:"x-order=2"`
}

func (request AuthRequest) GenerateAccessToken() (string, error) {
	duration := time.Minute * 15
	if env.Get("APP_ENV", "production") == "development" {
		duration = time.Hour * 24
	}

	now := utils.NowInJKT()
	// Create the Claims
	claims := jwt.MapClaims{
		"iss":      env.Get("APP_URL", "http://localhost"),
		"iat":      now.Unix(),
		"exp":      now.Add(duration).Unix(),
		"nbf":      now.Unix(),
		"username": request.Username,
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// Generate encoded token and send it as response.
	return token.SignedString(mJwt.Secret)
}
