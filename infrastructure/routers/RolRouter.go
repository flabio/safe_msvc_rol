package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/msvc_rol/globalinterfaces"
	"github.com/msvc_rol/handler"
)

var (
	rol globalinterfaces.IRolGlobal = handler.NewRolHandler()
)

// CreateOrder Creating Order
//
//	@Summary		Creating Order
//	@Description	Creating Order with given request
//	@Tags			Rol
//	@Accept			json
//	@Produce		json
//	@Success		200				{string}	string
//	@Failure		400				{string}	string	"Bad Request"
//	@Router			/rols [post]

func NewRouter(app *fiber.App) {

	api := app.Group("/api/rol")
	//api.Use(middleware.ValidateToken)
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
