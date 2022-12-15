package request

import (
	"event-management/handler/errors"

	"github.com/gofiber/fiber/v2"
)

func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestMethod := string(c.Request().Header.Method())
		contentType := string(c.Request().Header.ContentType())

		if (requestMethod == "POST" || requestMethod == "PATCH") && contentType != "application/json" {
			panic(errors.New("POST or PATCH Request Content-Type must be JSON", fiber.StatusBadRequest))
		}

		return c.Next()
	}
}
