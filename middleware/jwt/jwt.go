package jwt

import (
	"event-management/handler/errors"
	"event-management/utils/env"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

var Secret []byte

func Setup() {
	Secret = []byte(env.Get("JWT_SECRET", "PuSR1j4Y4s3LAlU"))
}

func New() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningMethod: "HS512",
		SigningKey:    Secret,
		ErrorHandler:  errors.ErrorHandler,
	})
}
