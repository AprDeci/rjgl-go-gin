package api

import (
	"net/http"

	"github.com/aprdec/rjgl/models/ucenter"
	"github.com/aprdec/rjgl/pkg/dto"
	"github.com/gin-gonic/gin"
)

// @Summary 登录
// @Schemes
// @Description login
// @Tags auth
// @Param req body dto.Auth true "req"
// @Accept json
// @Produce json
// @Success 200
// @Router /auth [post]
func GetAuth(ctx *gin.Context) {
	var req dto.Auth

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}

	// 校验
	isExist := ucenter.CheckAuth(req)

	var (
		code    int
		message string
	)
	if isExist {
		code = 0
		message = "登录成功"
	} else {
		code = 1
		message = "登录失败"
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
	})

}
