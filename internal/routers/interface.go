package routers

import (
	"github.com/gofiber/fiber/v2"
)

type RouterService interface {
	RegisterRoutes(routes *fiber.App)
}
