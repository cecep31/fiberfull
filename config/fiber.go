package config

import (
	"fiberfull/exception"

	"github.com/gofiber/fiber/v2"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
		Prefork:      true,
		AppName:      "fiberfull",
	}
}
