package ucenter

import (
	"errors"
	"log"
	"strconv"

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

func CheckAuth(req dto.Auth) bool {
	var account Account
	log.Printf("username: %s, password: %s", req.Username, req.Password)
	models.DB.Select("id,role_id,name").Where("name = ? and password = ?", req.Username, req.Password).First(&account)
	if account.ID > 0 {
		return true
	}
	return false
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

func GetAccount(id string) (*Account, error) {
	var account Account
	if err := models.DB.Where("id=?", id).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func GetAccountByUsername(username string) (*Account, error) {
	var account Account
	if err := models.DB.Where("name=?", username).First(&account).Error; err != nil {
		return nil, err
	}
	if account.ID <= 0 {
		return nil, errors.New("account not found")
	}
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

func DeleteAccount(id string) (bool, error) {
	account, err := GetAccount(id)
	if err != nil {
		return false, err
	}
	if account.ID <= 0 {
		return false, errors.New("account not found")
	}
	if err := models.DB.Delete(&account).Error; err != nil {
		return false, err
	}
	return true, nil
}

func UpdateAccount(req *dto.UpdateAccountReq) (bool, error) {
	account, err := GetAccount(strconv.Itoa(req.ID))
	if err != nil {
		return false, err
	}
	if req.Name != "" {
		account.Name = req.Name
	}
	if req.Password != "" {
		account.Password = req.Password
	}
	if req.Status > 0 {
		account.Status = req.Status
	}
	if req.RoleID > 0 {
		account.RoleID = req.RoleID
	}
	if req.ShowName != "" {
		account.ShowName = req.ShowName
	}
	if err := models.DB.Save(&account).Error; err != nil {
		return false, err
	}
	return true, nil
}
