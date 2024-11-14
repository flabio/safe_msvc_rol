package dto

type ModuleDTO struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Order  int    `json:"order"`
	Active bool   `json:"active"`
}
