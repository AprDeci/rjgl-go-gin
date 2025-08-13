package service

import (
	"github.com/aprdec/rjgl/models"
	"github.com/aprdec/rjgl/models/proj"
	"github.com/aprdec/rjgl/pkg/dto"
)

func CreateApprovalTemplate(template *proj.ApprovalTemplate) (bool, error) {
	if err := models.DB.Create(template).Error; err != nil {
		return false, err
	}
	return true, nil
}

func GetApprovalTemplateList(req *dto.GetApprovalTemplateListReq) ([]*proj.ApprovalTemplate, error) {
	var list []*proj.ApprovalTemplate
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	query := models.DB
	if req.Name != "" {
		query = query.Where("name like ?", "%"+req.Name+"%")
	}
	if req.Desc != "" {
		query = query.Where("description like ?", "%"+req.Desc+"%")
	}
	if req.Type != "" {
		query = query.Where("type = ?", req.Type)
	}
	if err := query.Order("id desc").Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil

}

func DeleteApprovalTemplate(id uint) (bool, error) {

	if err := models.DB.Delete(&proj.ApprovalTemplate{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func UpdateApprovalTemplate(template *proj.ApprovalTemplate) (bool, error) {
	if err := models.DB.Save(template).Error; err != nil {
		return false, err
	}
	return true, nil
}
