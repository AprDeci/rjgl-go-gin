package proj

import (
	"time"

	"github.com/aprdec/rjgl/models"
	"gorm.io/gorm"
)

// 审批节点
type ApprovalNode struct {
	gorm.Model
	TemplateID  uint
	NodeOrder   int
	Name        string
	Desc        string `gorm:"column:description"`
	Role        string //审批角色 admin superadmin user
	UserID      []int
	ApproveType string //OR/AND 多人审批类型
	Required    string
}

// 审批模板
type ApprovalTemplate struct {
	gorm.Model
	Type  string //审批类型 "memberJoin"-成员加入
	Name  string
	Desc  string         `gorm:"column:description"`
	Nodes []ApprovalNode `gorm:"foreignKey:TemplateID"`
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

func (a *ApprovalNode) TableName() string {
	return "proj_approval_node"
}

func (a *ApprovalTemplate) TableName() string {
	return "proj_approval_template"
}

func (a *ApprovalInstance) TableName() string {
	return "proj_approval_instance"
}

func InitApproval() {
	models.DB.AutoMigrate(&ApprovalNode{})
	models.DB.AutoMigrate(&ApprovalTemplate{})
	models.DB.AutoMigrate(&ApprovalInstance{})
}
