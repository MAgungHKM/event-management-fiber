package route

import (
	"event-management/handler/errors"
	"event-management/middleware/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/docs/")
	})

	app.Get("/docs/*", swagger.HandlerDefault)

	auth := app.Group("/auth")
	user := app.Group("/user", jwt.New())
	tag := app.Group("/tag", jwt.New())
	event := app.Group("/event")
	SetupAuthRoutes(auth)
	SetupUserRoutes(user)
	SetupTagRoutes(tag)
	SetupEventRoutes(event)

	// 404 handler
	app.Use(errors.Error404)
}
