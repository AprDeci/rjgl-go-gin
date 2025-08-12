package routers

import (
	"github.com/aprdec/rjgl/pkg/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	// 中间件
	gin.SetMode(setting.RunMode)

}
