package dto

type ModuleDTO struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Icon   string `json:"icon"`
	Order  int    `json:"order"`
	Active bool   `json:"active"`
}
