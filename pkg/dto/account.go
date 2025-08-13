package dto

type Account struct {
	ID       int    `gorm:"primaryKey" form:"id" json:"id"`
	Name     string `form:"name" json:"name"`
	Password string `form:"password" json:"password"`
	Status   int    `form:"status" json:"status"`
	RoleID   int    `gorm:"column:role_id" form:"role_id" json:"role_id"`
	ShowName string `form:"show_name" json:"show_name"`
}

type Auth struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateAccountReq struct {
	Username string `json:"username" binding:"required,min=2,max=30"`
	Password string `json:"password" binding:"required,min=6"`
}

type GetAccountListReq struct {
	Pagination
	Account
}

type UpdateAccountReq struct {
	Account
}
