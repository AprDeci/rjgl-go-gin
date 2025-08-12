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

func CreateAccount(data map[string]interface{}) bool {
	models.DB.Create(&Account{
		Name:     data["name"].(string),
		Password: data["password"].(string),
		Status:   1,
		RoleID:   1,
		ShowName: data["show_name"].(string),
	})
	return true
}
