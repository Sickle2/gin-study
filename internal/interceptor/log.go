package interceptor

import (
	"fengcheqiu/pkg/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		method := c.Request.Method
		uri := c.Request.RequestURI
		utils.Logger.Info("---request begin---")
		utils.Logger.Info("[" + method + "] " + uri)
		auth := c.GetHeader("Authorization")
		utils.Logger.Info("auth: " + auth)
		c.Next()

		end := time.Now()
		duration := end.Sub(start)
		code := c.Writer.Status()
		utils.Logger.Info("[" + method + "] " + uri + " code:" + strconv.Itoa(code) + " duration:" + duration.String())
		utils.Logger.Info("---request end---")
	}
}
