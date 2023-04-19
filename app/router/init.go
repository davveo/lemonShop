package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init(route *gin.Engine) {
	route.Use(gin.Logger())
	route.Use(gin.Recovery())

	publicGroup := route.Group("/api/v1")
	{
		publicGroup.GET("health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"ok": true})
		})
		AdminRouterGroup(publicGroup)
	}

	//privateGroup := r.Group("/api/v1")
	//{
	//
	//}
}
