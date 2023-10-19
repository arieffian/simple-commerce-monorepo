package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type HealthcheckService interface {
	HealthCheckHandler(c *fiber.Ctx) error
}
