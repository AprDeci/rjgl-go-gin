package ucenter

import (
	"github.com/aprdec/rjgl/models"
	"github.com/aprdec/rjgl/pkg/dto"
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

func CreateAccount(req *dto.CreateAccountReq) (bool, error) {
	models.DB.Create(&Account{
		Name:     req.Username,
		Password: req.Password,
		Status:   1,
		RoleID:   1,
	})
	return true, nil
}

func GetAccount(id uint) (*Account, error) {
	var account Account
	models.DB.Where("id=?", id).First(&account)
	return &account, nil
}

func GetAccountByUsername(username string) (*Account, error) {
	var account Account
	models.DB.Where("name=?", username).First(&account)
	return &account, nil
}

func GetAccountList(page, pageSize int, req *dto.GetAccountListReq) ([]*Account, error) {
	var accounts []*Account
	query := models.DB.Model(&accounts)
	if req.ID > 0 {
		query = query.Where("id=?", req.ID)
	}
	if req.Name != "" {
		query = query.Where("name like ?", "%"+req.Name+"%")
	}
	if req.Status > 0 {
		query = query.Where("status=?", req.Status)
	}
	if req.RoleID > 0 {
		query = query.Where("role_id=?", req.RoleID)
	}
	query = query.Order("id desc")
	query = query.Limit(pageSize)
	query = query.Offset((page - 1) * pageSize)
	if err := query.Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}
