package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/msvc_rol/globalinterfaces"
	"github.com/msvc_rol/usecases/services"
)

type HandlerModule struct {
	handler globalinterfaces.UIModuleGlobal
}

func NewInstanceHandler() globalinterfaces.UIModuleGlobal {
	return &HandlerModule{handler: services.ModuleInstance()}
}
func (h *HandlerModule) GetModuleFindAll(c *fiber.Ctx) error {
	return h.handler.GetModuleFindAll(c)
}
func (h *HandlerModule) GetModuleFindById(c *fiber.Ctx) error {
	return h.handler.GetModuleFindById(c)
}
func (h *HandlerModule) CreateModule(c *fiber.Ctx) error {
	return h.handler.CreateModule(c)
}
func (h *HandlerModule) DeleteModule(c *fiber.Ctx) error {
	return h.handler.DeleteModule(c)
}
func (h *HandlerModule) UpdateModule(c *fiber.Ctx) error {
	return h.handler.UpdateModule(c)
}
func (h *HandlerModule) CreateModuleRole(c *fiber.Ctx) error {
	return h.handler.CreateModuleRole(c)
}
func (h *HandlerModule) GetModuleRoleFindAll(c *fiber.Ctx) error {
	return h.handler.GetModuleRoleFindAll(c)
}
func (h *HandlerModule) DeleteModuleRole(c *fiber.Ctx) error {
	return h.handler.DeleteModuleRole(c)
}
