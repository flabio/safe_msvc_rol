package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/msvc_rol/globalinterfaces"
	"github.com/msvc_rol/usecases/services"
)

type rolHandler struct {
	rol globalinterfaces.IRolGlobal
}

func NewRolHandler() globalinterfaces.IRolGlobal {

	return &rolHandler{
		rol: services.NewRolService(),
	}
}

func (r *rolHandler) GetFindAll(c *fiber.Ctx) error {
	return r.rol.GetFindAll(c)
}

func (r *rolHandler) GetFindById(c *fiber.Ctx) error {
	return r.rol.GetFindById(c)

}

func (r *rolHandler) Create(c *fiber.Ctx) error {
	return r.rol.Create(c)
}
func (r *rolHandler) Update(c *fiber.Ctx) error {
	return r.rol.Update(c)
}

func (r *rolHandler) Delete(c *fiber.Ctx) error {
	return r.rol.Delete(c)
}
