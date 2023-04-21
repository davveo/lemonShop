package router

import (
	"github.com/davveo/lemonShop/app/ctrs"
	"github.com/gin-gonic/gin"
)

func AdminRouterGroup(router *gin.RouterGroup) {
	adminGroup := router.Group("admin/")
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

	goodsGroup := router.Group("goods")
	specsController := ctrs.NewSpecsController()
	{
		// done 查询规格项列表
		goodsGroup.GET("/specs", specsController.SpecsList)
		// done 添加规格项
		goodsGroup.POST("admin/goods/specs", specsController.Create)
		// done 修改规格项
		goodsGroup.PUT("admin/goods/specs/:spec_id", specsController.Update)
		// done 删除规格项
		goodsGroup.DELETE("admin/goods/specs/:spec_id", specsController.Delete)
		// done 查询一个规格项
		goodsGroup.GET("/specs/:spec_id", specsController.Specs)
		//// done 查询规格值列表
		//goodsGroup.GET("admin/goods/specs/:spec_id/values", admin.SpecsValues)
		//// done 添加某规格的规格值
		//goodsGroup.POST("admin/goods/specs/:spec_id/values", admin.UpdateSpecsValues)
	}
}
