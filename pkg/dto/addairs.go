package dto

import "github.com/aprdec/rjgl/models/proj"

type GetApprovalTemplateListReq struct {
	Pagination
	proj.ApprovalTemplate
}
