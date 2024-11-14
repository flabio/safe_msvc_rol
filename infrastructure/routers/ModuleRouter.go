package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/msvc_rol/handler"
	"github.com/msvc_rol/middleware"
)

var (
	HandlerModule = handler.NewInstanceHandler()
)

func NewInstanceModuleRouter(app *fiber.App) {

	api := app.Group("/api/module")
	api.Use(middleware.ValidateToken)
	api.Get("/", func(c *fiber.Ctx) error {
		return HandlerModule.GetModuleFindAll(c)
	})
	api.Post("/", func(c *fiber.Ctx) error {
		return HandlerModule.CreateModule(c)
	})
	api.Put("/:id", func(c *fiber.Ctx) error {
		return HandlerModule.UpdateModule(c)
	})
	api.Delete("/:id", func(c *fiber.Ctx) error {
		return HandlerModule.DeleteModule(c)
	})
	api.Get("/:id", func(c *fiber.Ctx) error {
		return HandlerModule.GetModuleFindById(c)
	})
	api.Post("/role", func(c *fiber.Ctx) error {
		return HandlerModule.CreateModuleRole(c)
	})
	api.Delete("/role/:id", func(c *fiber.Ctx) error {
		return HandlerModule.DeleteModuleRole(c)
	})
	api.Get("/role/", func(c *fiber.Ctx) error {
		return HandlerModule.GetModuleRoleFindAll(c)
	})
}
