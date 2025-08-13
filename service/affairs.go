package service

import (
	"github.com/aprdec/rjgl/models/proj"
	"gorm.io/gorm"
)

type AffairsService struct {
	DB *gorm.DB
}

func (s *AffairsService) CreateApprovalTemplate(template *proj.ApprovalTemplate) (bool, error) {
	if err := s.DB.Create(template).Error; err != nil {
		return false, err
	}
	return true, nil
}
