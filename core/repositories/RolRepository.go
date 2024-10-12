package repositories

import (
	"sync"

	"github.com/msvc_rol/core/interfaces"
	"github.com/msvc_rol/infrastructure/database"
	"github.com/msvc_rol/infrastructure/entities"
	"github.com/msvc_rol/infrastructure/utils"
	"gorm.io/gorm"
)

type OpenConnection struct {
	connection *gorm.DB
	mux        sync.Mutex
}

func GetRolInstance() interfaces.IRol {
	var (
		_OPEN *OpenConnection
		_ONCE sync.Once
	)

	_ONCE.Do(func() {
		_OPEN = &OpenConnection{
			connection: database.DatabaseConnection(),
		}
	})
	return _OPEN
}

func (db *OpenConnection) GetFindAll() ([]entities.Rol, error) {
	var roles []entities.Rol
	db.mux.Lock()
	result := db.connection.Order(utils.DB_ORDER_DESC).Find(&roles)
	defer database.Closedb()
	defer db.mux.Unlock()
	return roles, result.Error
}

func (db *OpenConnection) GetFindById(id int) (entities.Rol, error) {
	db.mux.Lock()
	var rol entities.Rol
	result := db.connection.Find(&rol, id)
	defer database.Closedb()
	defer db.mux.Unlock()
	return rol, result.Error
}

/*
param:rol is a struct
*/
func (db *OpenConnection) Create(rol entities.Rol) (entities.Rol, error) {
	db.mux.Lock()
	err := db.connection.Create(&rol).Error
	defer database.Closedb()
	defer db.mux.Unlock()
	return rol, err
}

/*
@Params: rol Rol is a struct, id is an integer
*/
func (db *OpenConnection) Update(id uint, rol entities.Rol) (entities.Rol, error) {
	db.mux.Lock()
	result := db.connection.Where(utils.DB_EQUAL_ID, id).Updates(&rol)
	defer database.Closedb()
	defer db.mux.Unlock()
	return rol, result.Error

}

/*
@param: id is an int
*/
func (db *OpenConnection) Delete(id uint) (bool, error) {
	db.mux.Lock()
	var rol entities.Rol
	result := db.connection.Where(utils.DB_EQUAL_ID, id).Delete(&rol)

	defer database.Closedb()
	defer db.mux.Unlock()
	if result.RowsAffected == 0 {
		return true, result.Error
	}
	return false, result.Error
}

/*
@params: id is an uint number and name is a string
*/
func (db *OpenConnection) GetFindByName(id uint, name string) (bool, error) {
	db.mux.Lock()
	var rol entities.Rol

	query := db.connection.Where(utils.DB_NAME, name)
	if id > 0 {
		query = query.Where(utils.DB_DIFF_ID, id)
	}
	query = query.Find(&rol)

	defer database.Closedb()
	defer db.mux.Unlock()
	if query.RowsAffected == 1 {
		return true, query.Error
	}
	return false, query.Error

}
