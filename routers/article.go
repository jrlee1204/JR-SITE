package routers

import (
	"github.com/gin-gonic/gin"
	"jr_site/result"
)

func InitArticleRouter(router *gin.RouterGroup) {
	routerGroup := router.Group("/article")
	routerGroup.GET("/state", func(c *gin.Context) {
		result.Success(c)
	})
}
