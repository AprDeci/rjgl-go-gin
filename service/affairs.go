package service

import (
	"github.com/aprdec/rjgl/models"
	"github.com/aprdec/rjgl/models/proj"
)

func CreateApprovalTemplate(template *proj.ApprovalTemplate) (bool, error) {
	if err := models.DB.Create(template).Error; err != nil {
		return false, err
	}
	return true, nil
}
