package globalinterfaces

import (
	"github.com/gofiber/fiber/v2"
)

type IRolGlobal interface {
	GetFindAll(c *fiber.Ctx) error
	GetFindById(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
type IModuleRolGlobal interface {
	GetModuleRolFindById(c *fiber.Ctx) error
}
