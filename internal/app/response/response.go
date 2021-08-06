package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @program: fengcheqiu
 * @description:
 * @author: sickle
 * @create: 2021-08-06 11:30
 **/

type Meta struct {
	Code    ResultCode  `json:"code"`
	Message interface{} `json:"date,omitempty"`
}

type Response struct {
	Meta *Meta       `json:"meta,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func echo(ctx *gin.Context, response Response) {
	ctx.Set("contentType", "application/json")
	ctx.JSON(http.StatusOK, response)
}

func ResponseSuccess(ctx *gin.Context) {
	meta := &Meta{
		Code:    Success,
		Message: GetMessage(Success),
	}
	response := Response{
		Meta: meta,
	}
	echo(ctx, response)
}

func ResponseSuccessWithDate(ctx *gin.Context, data interface{}) {
	meta := &Meta{
		Code:    Success,
		Message: GetMessage(Success),
	}
	response := Response{
		Meta: meta,
		Data: data,
	}
	echo(ctx, response)
}

func ResponseError(ctx *gin.Context) {
	meta := &Meta{
		Code:    Failed,
		Message: GetMessage(Failed),
	}
	response := Response{
		Meta: meta,
	}
	echo(ctx, response)
}

func ResponseErrorWithData(ctx *gin.Context, data interface{}) {
	meta := &Meta{
		Code:    Failed,
		Message: GetMessage(Failed),
	}
	response := Response{
		Meta: meta,
		Data: data,
	}
	echo(ctx, response)
}
