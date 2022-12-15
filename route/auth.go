package route

import (
	"event-management/handler/auth"
	"event-management/middleware/request"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(router fiber.Router) {
	router.Post("/", request.New(), auth.Auth)
}
