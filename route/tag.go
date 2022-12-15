package route

import (
	"event-management/handler/tag"
	"event-management/middleware/request"

	"github.com/gofiber/fiber/v2"
)

func SetupTagRoutes(router fiber.Router) {
	router.Get("/", tag.FindAll)
	router.Post("/", request.New(), tag.Create)
	router.Patch("/:id", request.New(), tag.Update)
	router.Delete("/:id", tag.Delete)
}
