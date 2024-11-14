package dto

type ModuleRoleDTO struct {
	Id       uint `json:"id"`
	RoleId   uint `json:"role_id"`
	ModuleId uint `json:"module_id"`
	Active   bool `json:"active"`
}
