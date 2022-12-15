package handler

import (
	"github.com/gofiber/fiber/v2"
)

func ApiResponse(c *fiber.Ctx, data interface{}, statusCode int) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	// Return statuscode with error message
	return c.Status(statusCode).JSON(data)
}

func ApiResponse200(c *fiber.Ctx, data interface{}) error {
	return ApiResponse(c, data, 200)
}

func ApiResponseHandler(data interface{}, statusCode int) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return ApiResponse(c, data, 200)
	}
}
