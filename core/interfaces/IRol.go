package interfaces

import "github.com/msvc_rol/infrastructure/entities"

type IRol interface {
	GetFindAll() ([]entities.Rol, error)
	GetFindById(id int) (entities.Rol, error)
	GetFindByName(id uint, name string) (bool, error)
	Create(rol entities.Rol) (entities.Rol, error)
	Update(id uint, rol entities.Rol) (entities.Rol, error)
	Delete(id uint) (bool, error)
}
