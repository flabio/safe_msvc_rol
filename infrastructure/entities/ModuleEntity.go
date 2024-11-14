package entities

type Module struct {
	Id          uint          `gorm:"primary_key:auto_increment"`
	Name        string        `gorm:"type:varchar(100);not null" json:"name"`
	Description string        `gorm:"type:text" json:"description"`
	Order       int           `gorm:"type:integer" json:"order"`
	Active      bool          `gorm:"type:boolean" json:"active"`
	ModuleRole  *[]ModuleRole `json:"module_role"`
	CreatedAt   string        `gorm:"<-:created_at" json:"created"`
	UpdatedAt   *string       `gorm:"type:TIMESTAMP(6)" json:"updated"`
}
