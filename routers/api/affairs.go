package api

import (
	"net/http"

	"github.com/aprdec/rjgl/models/proj"
	"github.com/aprdec/rjgl/pkg/dto"
	"github.com/aprdec/rjgl/service"
	"github.com/gin-gonic/gin"
)

// @Summary 创建审批模板
// @Description 创建审批模板
// @Tags affairs
// @Accept json
// @Produce json
// @Param template body proj.ApprovalTemplate true "审批模板"
// @Success 200 {object} gin.H{code=int,msg=string}
// @Router /affair/create [post]
func CreateApprovalTemplate(c *gin.Context) {
	var template proj.ApprovalTemplate
	if err := c.ShouldBind(&template); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	ok, err := service.CreateApprovalTemplate(&template)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "创建失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "创建成功",
	})
}

// @Summary 获取审批模板列表
// @Description 获取审批模板列表
// @Tags affairs
// @Accept json
// @Produce json
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Param name query string false "name"
// @Param desc query string false "desc"
// @Param type query string false "type"
// @Success 200 {object} gin.H{code=int,msg=string,data=[]proj.ApprovalTemplate}
// @Router /affair/list [get]
func GetApprovalTemplateList(c *gin.Context) {
	var req dto.GetApprovalTemplateListReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	list, err := service.GetApprovalTemplateList(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "获取成功",
		"data": list,
	})
}

// @Summary 删除审批模板
// @Description 删除审批模板
// @Tags affairs
// @Accept json
// @Produce json
// @Param id query int true "id"
// @Success 200 {object} gin.H{code=int,msg=string}
// @Router /affair/delete [delete]
func DeleteApprovalTemplate(c *gin.Context) {
	var req dto.CommonReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	ok, err := service.DeleteApprovalTemplate(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "删除成功",
	})

}
