package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/msvc_rol/globalinterfaces"
	"github.com/msvc_rol/usecases/services"
)

type moduleRolHandler struct {
	rol globalinterfaces.IRolGlobal
}

func NewModuleRolHandler() globalinterfaces.IModuleRolGlobal {

	return &rolHandler{
		rol: services.NewRolService(),
	}
}

func (r *rolHandler) GetModuleRolFindById(c *fiber.Ctx) error {
	return r.rol.GetFindById(c)

}
