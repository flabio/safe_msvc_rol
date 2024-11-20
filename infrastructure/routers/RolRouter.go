package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/msvc_rol/globalinterfaces"
	"github.com/msvc_rol/handler"
	"github.com/msvc_rol/middleware"
)

var (
	rol globalinterfaces.IRolGlobal = handler.NewRolHandler()
)

func NewRouter(app *fiber.App) {

	api := app.Group("/api/rol")
	api.Use(middleware.ValidateToken)
	api.Get("/", func(c *fiber.Ctx) error {
		return rol.GetFindAll(c)
	})
	api.Get("/:id", func(c *fiber.Ctx) error {
		return rol.GetFindById(c)
	})
	api.Post("/", func(c *fiber.Ctx) error {
		return rol.Create(c)
	})
	api.Put("/:id", func(c *fiber.Ctx) error {
		return rol.Update(c)
	})
	api.Delete("/:id", func(c *fiber.Ctx) error {
		return rol.Delete(c)
	})

}
