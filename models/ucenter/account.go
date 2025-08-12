package ucenter

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name     string
	Password string
	Status   int
	RoleID   int "gorm:column:role_id"
	ShowName string
}
