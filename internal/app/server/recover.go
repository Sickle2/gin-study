package server

import (
	"fengcheqiu/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"runtime"
	"runtime/debug"
)

/**
 * @program: fengcheqiu
 * @description:
 * @author: sickle
 * @create: 2021-08-06 11:27
 **/
func RecoverWrapper(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
			var buf [4096]byte
			n := runtime.Stack(buf[:], false)
			utils.Logger.Error("发现全局异常 " + string(buf[:n]))
			utils.Logger.Error("全局异常信息 ", zap.Reflect("error", r))
			//TODO  发送邮件通知
		}
	}()
	ctx.Next()
}
