package routers

import (
	"github.com/aprdec/rjgl/pkg/setting"
	"github.com/aprdec/rjgl/routers/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	// 中间件
	gin.SetMode(setting.RunMode)

	r.POST("/auth", api.GetAuth)
	accountGroup := r.Group("/account")
	{
		accountGroup.POST("/create", api.CreateAccount)
		accountGroup.GET("/list", api.GetAccountList)
		accountGroup.GET("/delete", api.DeleteAccount)
		accountGroup.POST("/update", api.UpdateAccount)

	}
	affairGroup := r.Group("/affair")
	{
		affairGroup.POST("/create", api.CreateApprovalTemplate)
	}

	return r
}
