package ucenter

import "github.com/aprdec/rjgl/models"

func init() {
	models.DB.AutoMigrate(&Account{})
}
