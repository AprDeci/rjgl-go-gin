package api

import (
	"net/http"

	"github.com/aprdec/rjgl/models/ucenter"
	"github.com/aprdec/rjgl/pkg/dto"
	"github.com/gin-gonic/gin"
)

// @Summary 创建账号
// @Schemes
// @Description create account
// @Tags account
// @Param account body dto.CreateAccountReq true "account"
// @Accept json
// @Produce json
// @Success 200 {object} dto.CreateAccountResp
// @Router /account/create [post]
func CreateAccount(c *gin.Context) {
	var req dto.CreateAccountReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ok, err := ucenter.CreateAccount(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "创建成功",
	})
}
