package routers

import (
	"github.com/gin-gonic/gin"
	"jr_site/result"
)

// InitRouter 初始化项目总路由
func InitRouter(c *gin.Engine) {
	routerGroup := c.Group("/api/v1")
	// 健康检查接口
	routerGroup.GET("health_check", func(c *gin.Context) {
		result.SuccessWithData(c, "health check ok!")
	})
	InitArticleRouter(routerGroup)
}
