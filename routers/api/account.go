package api

import (
	"log"
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
// @Success 200
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

// @Summary 获取账号列表
// @Schemes
// @Description get account list
// @Tags account
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Param name query string false "name"
// @Param status query int false "status"
// @Param roleID query int false "roleID"
// @Accept json
// @Produce json
// @Success 200
// @Router /account/list [get]
func GetAccountList(c *gin.Context) {
	var req dto.GetAccountListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("req: %v", req)

	accounts, err := ucenter.GetAccountList(req.Page, req.PageSize, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "获取成功",
		"data":    accounts,
	})
}

// @Summary 删除账号
// @Schemes
// @Description delete account
// @Tags account
// @Param id query string true "id"
// @Accept json
// @Produce json
// @Success 200
// @Router /account/delete [delete]
func DeleteAccount(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "id is required"})
		return
	}
	ok, err := ucenter.DeleteAccount(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": err.Error()})
		return
	}
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "删除失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "删除成功",
	})

}

// @Summary 更新账号
// @Schemes
// @Description update account
// @Tags account
// @Param account body dto.UpdateAccountReq true "account"
// @Accept json
// @Produce json
// @Success 200
// @Router /account/update [put]
func UpdateAccount(c *gin.Context) {
	var req dto.UpdateAccountReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": err.Error()})
		return

	}
	if req.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "id is required"})
		return
	}
	ok, err := ucenter.UpdateAccount(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": err.Error()})
		return
	}
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "更新失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "更新成功",
	})

}
