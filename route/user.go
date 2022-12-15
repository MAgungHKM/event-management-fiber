package route

import (
	"event-management/handler/user"
	"event-management/middleware/request"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	router.Get("/", user.FindAll)
	router.Post("/", request.New(), user.Create)
	router.Patch("/:id", request.New(), user.Update)
	router.Delete("/:id", user.Delete)
}
