package router

import (
	"github.com/davveo/lemonShop/app/ctrs"
	"github.com/gin-gonic/gin"
)

func AdminRouterGroup(router *gin.RouterGroup) {
	adminGroup := router.Group("/api/v1")
	articleController := ctrs.NewArticleController()
	{
		//获取文章列表
		adminGroup.GET("/articles", articleController.GetArticles)
		//获取指定文章
		adminGroup.GET("/articles/:id", articleController.GetArticle)
		//新建文章
		adminGroup.POST("/articles", articleController.AddArticle)
		//更新指定文章
		adminGroup.PUT("/articles/:id", articleController.EditArticle)
		//删除指定文章
		adminGroup.DELETE("/articles/:id", articleController.DeleteArticle)
	}
}
