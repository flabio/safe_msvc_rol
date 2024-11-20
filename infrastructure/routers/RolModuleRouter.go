package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/msvc_rol/globalinterfaces"
	"github.com/msvc_rol/handler"
)

var (
	rolModule globalinterfaces.IModuleRolGlobal = handler.NewModuleRolHandler()
)

func NewRolModuleRouter(app *fiber.App) {
	apiModule := app.Group("/api/rol_module")
	apiModule.Get("/rol/:id", func(c *fiber.Ctx) error {
		return rolModule.GetModuleRolFindById(c)
	})
}
