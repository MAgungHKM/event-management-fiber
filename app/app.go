package app

import (
	"event-management/config"
	"event-management/handler/errors"
	"event-management/middleware"
	"event-management/route"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Setup(cfg config.Main) {
	app := fiber.New(fiber.Config{
		AppName:      *cfg.AppName,
		Prefork:      *cfg.Prod,
		ErrorHandler: errors.ErrorHandler,
	})

	middleware.Setup(app)
	route.Setup(app)

	err := app.Listen(fmt.Sprintf(":%s", *cfg.Port))

	if err != nil {
		log.Fatalf(fmt.Sprintf("Error Starting Webserver: %s", err.Error()))
	}
}
