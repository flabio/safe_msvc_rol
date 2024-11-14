package globalinterfaces

import "github.com/gofiber/fiber/v2"

type UIModuleGlobal interface {
	GetModuleFindAll(c *fiber.Ctx) error
	GetModuleFindById(c *fiber.Ctx) error
	CreateModule(c *fiber.Ctx) error
	UpdateModule(c *fiber.Ctx) error
	DeleteModule(c *fiber.Ctx) error
	GetModuleRoleFindAll(c *fiber.Ctx) error
	CreateModuleRole(c *fiber.Ctx) error
	DeleteModuleRole(c *fiber.Ctx) error
}
