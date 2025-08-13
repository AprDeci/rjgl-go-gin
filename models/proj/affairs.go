package proj

import (
	"time"

	"gorm.io/gorm"
)

// 审批节点
type ApprovalNode struct {
	gorm.Model
	TemplateID  uint
	NodeOrder   int
	Name        string
	Desc        string `gorm:"column:description"`
	Role        string
	ApproveType string
	Required    string
}

// 审批模板
type ApprovalTemplate struct {
	gorm.Model
	Code  string
	Name  string
	Desc  string `gorm:"column:description"`
	Nodes []ApprovalNode
}

// 审批实例
type ApprovalInstance struct {
	gorm.Model
	TemplateID uint
	Title      string
	Desc       string `gorm:"column:description"`
	ObjectID   uint
	// 申请人信息
	ApplicantID uint
	Reason      string
	Status      string
	CurrentNode int
	ApprovedAt  *time.Time
	RejectedAt  *time.Time
}
