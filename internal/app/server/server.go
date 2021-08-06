package server

import (
	conf "fengcheqiu/configs"
	"fengcheqiu/internal/interceptor"
	"github.com/gin-gonic/gin"
	"strconv"
)

/**
 * @program: fengcheqiu
 * @description:
 * @author: sickle
 * @create: 2021-08-06 10:25
 **/

func Run() {
	engine := gin.Default()
	engine.Use(interceptor.Log())
	engine.Use(interceptor.Cors())
	initRouter(engine)
	var port = ":" + strconv.FormatInt(conf.Conf.App.Port, 10)
	if err := engine.Run(port); err != nil {
		panic(err)
	}
}
