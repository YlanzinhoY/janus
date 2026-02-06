package web

import (
	"github.com/gofiber/fiber/v2"
)

type Port interface {
	PostOperator(c fiber.Ctx) error
}
