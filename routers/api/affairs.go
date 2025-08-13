package api

import (
	"net/http"

	"github.com/aprdec/rjgl/models/proj"
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
