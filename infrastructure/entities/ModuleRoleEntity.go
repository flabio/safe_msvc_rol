package entities

type ModuleRole struct {
	Id       uint   `gorm:"primary_key:auto_increment" json:"id"`
	RoleId   uint   `gorm:"type:integer" json:"role_id"`
	Rol      Rol    `gorm:"foreignKey:RoleId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"rol"`
	ModuleId uint   `json:"module_id"`
	Module   Module `gorm:"foreignKey:ModuleId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"module"`
	Active   bool   `gorm:"type:boolean" json:"active"`
}
