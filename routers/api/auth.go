package api

import (
	"net/http"

	"github.com/aprdec/rjgl/models/ucenter"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password"`
}

func GetAuth(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	// 校验
	isExist := ucenter.CheckAuth(username, password)

	var (
		code    int
		message string
	)
	if isExist {

	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
	})

}
