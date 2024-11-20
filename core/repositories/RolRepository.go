package repositories

import (
	"sync"

	constants "github.com/flabio/safe_constants"
	"github.com/msvc_rol/core/interfaces"
	"github.com/msvc_rol/infrastructure/database"
	"github.com/msvc_rol/infrastructure/entities"
)

func GetRolInstance() interfaces.IRol {
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

func (db *OpenConnection) GetFindAll() ([]entities.Rol, error) {
	var roles []entities.Rol
	db.mux.Lock()
	defer db.mux.Unlock()
	result := db.connection.Preload("RoleModule.Module").Order(constants.DB_ORDER_DESC).Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}
	return roles, nil
}

func (db *OpenConnection) GetFindById(id int) (entities.Rol, error) {
	var rol entities.Rol
	db.mux.Lock()
	defer db.mux.Unlock()
	result := db.connection.Preload("RoleModule.Module").Find(&rol, id)
	return rol, result.Error
}

/*
param:rol is a struct
*/
func (db *OpenConnection) Create(rol entities.Rol) (entities.Rol, error) {
	db.mux.Lock()
	defer db.mux.Unlock()
	err := db.connection.Create(&rol).Error
	return rol, err
}

/*
@Params: rol Rol is a struct, id is an integer
*/
func (db *OpenConnection) Update(id uint, rol entities.Rol) (entities.Rol, error) {
	db.mux.Lock()
	defer db.mux.Unlock()
	result := db.connection.Where(constants.DB_EQUAL_ID, id).Updates(&rol)
	return rol, result.Error

}

/*
@param: id is an int
*/
func (db *OpenConnection) Delete(id uint) (bool, error) {
	var rol entities.Rol
	db.mux.Lock()
	defer db.mux.Unlock()

	result := db.connection.Where(constants.DB_EQUAL_ID, id).Delete(&rol)
	if result.RowsAffected == 0 {
		return true, result.Error
	}
	return false, result.Error
}

/*
@params: id is an uint number and name is a string
*/
func (db *OpenConnection) GetFindByName(id uint, name string) (bool, error) {

	var rol entities.Rol
	db.mux.Lock()
	defer db.mux.Unlock()
	query := db.connection.Where(constants.DB_NAME, name)
	if id > 0 {
		query = query.Where(constants.DB_DIFF_ID, id)
	}
	query = query.Find(&rol)

	if query.RowsAffected == 1 {
		return true, query.Error
	}
	return false, query.Error

}
