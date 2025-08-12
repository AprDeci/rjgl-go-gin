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

func CreateAccount(c *gin.Context) {
	var auth auth
	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
