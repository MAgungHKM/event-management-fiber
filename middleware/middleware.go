package middleware

import (
	"event-management/handler/errors"
	"event-management/middleware/jwt"
	"event-management/utils/env"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Setup(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		TimeFormat: env.Get("LOG_FORMAT", "02-Jan-2006 15:04:05.000"),
		TimeZone:   env.Get("LOG_TIMEZONE", "Asia/Jakarta"),
	}))

	app.Use(favicon.New())

	defer app.Use(limiter.New(limiter.Config{
		Max:        30,
		Expiration: time.Second * 30,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			panic(errors.New("Timed Out. Try again after 30 seconds.", fiber.StatusTooManyRequests))
		},
	}))

	defer app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	defer app.Use(recover.New())

	jwt.Setup()

	// app.Use(response.New())
}
