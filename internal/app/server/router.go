package server

import (
	"fengcheqiu/api"
	_ "fengcheqiu/internal/interceptor"
	"github.com/gin-gonic/gin"
)

func initRouter(router *gin.Engine) {
	router.Use(RecoverWrapper)
	api1 := router.Group("/api")
	//api.Use(interceptor.Auth)
	//上传文件在内存中保存
	//router.MaxMultipartMemory = 1 << 23
	api1.GET("/h", api.HelloWorld)
}
