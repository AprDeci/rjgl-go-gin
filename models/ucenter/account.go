package ucenter

import (
	"github.com/aprdec/rjgl/models"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name     string
	Password string
	Status   int
	RoleID   int `gorm:"column:role_id"`
	ShowName string
}

func (a *Account) TableName() string {
	return "ucenter_account"
}

func CheckAuth(username, password string) bool {
	var account Account
	models.DB.Select("id,role_id,name").Where("name = ? and password = ?", username, password).First(&account)
	return account.ID > 0
}

func CreateAccount(account *Account) error {
	return models.DB.Create(account).Error
}
