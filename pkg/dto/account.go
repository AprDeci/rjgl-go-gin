package dto

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Name     string
	Password string
	Status   int
	RoleID   int `gorm:"column:role_id"`
	ShowName string
}

type CreateAccountReq struct {
	Username string `json:"username" binding:"required,min=2,max=30"`
	Password string `json:"password" binding:"required,min=6"`
}

type GetAccountListReq struct {
	pagination
	Account
}
