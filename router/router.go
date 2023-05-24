package router

import (
	"github.com/gin-gonic/gin"
	"mygin/service"
	"net/http"
)

func Load(g *gin.Engine) {
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "找不到该页面")
	})

	//分组管理路由
	u := g.Group("/v1/user")
	{
		u.GET("", service.FindUser)
		u.POST("", service.Create)
		u.PUT("/:id", service.Modify)
		u.DELETE("/:id", service.Delete)

	}
}
