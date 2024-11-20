package entities

import "time"

type Rol struct {
	Id         uint         `gorm:"primary_key:auto_increment"  json:"id" `
	Name       string       `gorm:"type:varchar(100);not null" json:"name" `
	Active     bool         `gorm:"type:boolean"  json:"active"`
	RoleModule []ModuleRole `gorm:"foreignKey:RoleId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"role_module"`

	CreatedAt time.Time  `gorm:"<-:created_at"  json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:TIMESTAMP(6)" json:"updated_at" `
}
