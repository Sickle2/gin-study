package interceptor

import (
	conf "fengcheqiu/configs"
	"fengcheqiu/internal/app/response"
	"fengcheqiu/internal/constant"
	"fengcheqiu/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	prefix = "Bearer "
)

func Auth(ctx *gin.Context) {
	for _, k := range conf.Conf.App.UnAuthUrl {
		if strings.Contains(ctx.Request.URL.Path, k) {
			return
		}
	}
	auth := ctx.Request.Header.Get(constant.Authorization)

	if auth == "" {
		utils.Logger.Error("没有token")
		ctx.Set("contentType", "application/json")
		meta := &response.Meta{
			Code:    response.Unauthorized,
			Message: "认证失败",
		}
		resp := &response.Response{
			Meta: meta,
		}
		ctx.JSON(http.StatusOK, resp)
		//e.EchoError(ctx, e.Unauthorized)
		ctx.Abort()
		return
	}

	if !strings.HasPrefix(auth, prefix) {
		utils.Logger.Error("格式不正确")
		ctx.Set("contentType", "application/json")
		meta := &response.Meta{
			Code:    response.Unauthorized,
			Message: "认证失败",
		}
		resp := &response.Response{
			Meta: meta,
		}
		ctx.JSON(http.StatusOK, resp)
		ctx.Abort()
		return
	}
	//_, b := utils.ValidJwtToken(auth[len(prefix):])
	b := true
	if !b {
		utils.Logger.Error("验证失败")
		ctx.Set("contentType", "application/json")
		meta := &response.Meta{
			Code:    response.Unauthorized,
			Message: "认证失败",
		}
		resp := &response.Response{
			Meta: meta,
		}
		ctx.JSON(http.StatusOK, resp)
		ctx.Abort()
		return
	}

	return
}
