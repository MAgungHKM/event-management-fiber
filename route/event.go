package route

import (
	"event-management/handler/event"
	"event-management/middleware/jwt"
	"event-management/middleware/request"

	"github.com/gofiber/fiber/v2"
)

func SetupEventRoutes(router fiber.Router) {
	router.Get("/", event.FindAll)

	router.Post("/", jwt.New(), request.New(), event.Create)
	router.Post("/:id/participate", request.New(), event.Participate)

	router.Patch("/:id", request.New(), event.Update)
	router.Patch("/:id/attend/:code", event.Attend)

	router.Delete("/:id", event.Delete)
	router.Delete("/:id/cancel/:code", event.Cancel)
}
