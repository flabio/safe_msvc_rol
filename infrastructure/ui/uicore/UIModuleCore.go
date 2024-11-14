package uicore

import "github.com/msvc_rol/infrastructure/entities"

type UIModuleCore interface {
	GetModuleFindAll() ([]entities.Module, error)
	GetModuleFindById(id uint) (entities.Module, error)
	CreateModule(module entities.Module) (entities.Module, error)
	UpdateModule(id uint, module entities.Module) (entities.Module, error)
	DeleteModule(id uint) (bool, error)
	GetModuleFindByName(id uint, name string) (bool, error)
	GetModuleRoleFindAll() ([]entities.ModuleRole, error)
	CreateModuleRole(module entities.ModuleRole) (entities.ModuleRole, error)
	DeleteModuleRole(id uint) (bool, error)
}
