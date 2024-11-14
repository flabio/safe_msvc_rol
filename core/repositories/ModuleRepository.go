package repositories

import (
	"sync"

	constants "github.com/flabio/safe_constants"
	"github.com/msvc_rol/infrastructure/database"
	"github.com/msvc_rol/infrastructure/entities"
	"github.com/msvc_rol/infrastructure/ui/uicore"
)

func GetModuleInstance() uicore.UIModuleCore {
	var (
		_OPEN *OpenConnection
		_ONCE sync.Once
	)
	_ONCE.Do(func() {
		_OPEN = &OpenConnection{
			connection: database.GetDatabaseInstance(),
		}
	})
	return _OPEN
}
func (c *OpenConnection) GetModuleFindAll() ([]entities.Module, error) {
	var modules []entities.Module
	c.mux.Lock()
	defer c.mux.Unlock()
	query := c.connection.Preload("ModuleRole.Module").Order(constants.DB_ORDER_DESC).Find(&modules)
	return modules, query.Error
}

func (c *OpenConnection) GetModuleFindById(id uint) (entities.Module, error) {
	var module entities.Module
	c.mux.Lock()
	defer c.mux.Unlock()
	query := c.connection.Where(constants.DB_EQUAL_ID, id).Find(&module)
	return module, query.Error
}
func (c *OpenConnection) CreateModule(module entities.Module) (entities.Module, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	query := c.connection.Create(&module)
	return module, query.Error
}
func (c *OpenConnection) UpdateModule(id uint, module entities.Module) (entities.Module, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	query := c.connection.Where(constants.DB_EQUAL_ID, id).Updates(&module)
	return module, query.Error
}
func (c *OpenConnection) DeleteModule(id uint) (bool, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	query := c.connection.Delete(&entities.Module{}, id)
	return query.RowsAffected > 0, query.Error
}
func (c *OpenConnection) GetModuleFindByName(id uint, name string) (bool, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	var module entities.Module
	query := c.connection.Where(constants.DB_EQUAL_NAME, name)
	if id > 0 {
		query = query.Where(constants.DB_DIFF_ID, id).First(&module)
	} else {
		query = query.First(&module)
	}
	return module.Id > 0, query.Error
}

// Module whith Role

func (c *OpenConnection) GetModuleRoleFindAll() ([]entities.ModuleRole, error) {
	var roles []entities.ModuleRole
	c.mux.Lock()
	defer c.mux.Unlock()
	query := c.connection.Find(&roles)
	return roles, query.Error
}

func (c *OpenConnection) CreateModuleRole(role entities.ModuleRole) (entities.ModuleRole, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	query := c.connection.Create(&role)
	return role, query.Error
}

func (c *OpenConnection) DeleteModuleRole(id uint) (bool, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	query := c.connection.Delete(&entities.ModuleRole{}, id)
	return query.RowsAffected > 0, query.Error
}
