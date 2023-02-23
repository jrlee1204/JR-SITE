package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"jr_site/config"
	"jr_site/db/mysql"
	"jr_site/routers"
)

func main() {
	// 初始化MySQL
	mysql.InitMysql()
	err := config.InitLogger(config.LogConfig{
		Filename:   "./test.log",
		Level:      "DEBUG",
		MaxSize:    10,
		MaxAge:     30,
		MaxBackups: 10,
	})
	if err != nil {
		fmt.Println("初始化zap异常：", err)
	}
	r := gin.New()
	r.Use(config.GinLogger(), config.GinRecovery(true))
	// 初始化路由
	routers.InitRouter(r)
	err = r.Run()
	if err != nil {
		zap.L().Error("启动服务异常：", zap.String("error", err.Error()))
	}
}
