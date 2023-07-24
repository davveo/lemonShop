package router

import (
	"github.com/davveo/lemonShop/app/ctrs"
	"github.com/gin-gonic/gin"
)

func Init(route *gin.Engine) {
	route.Use(gin.Logger())
	route.Use(gin.Recovery())

	publicGroup := route.Group("/api/v1")

	checkCTRS := ctrs.NewCheckerController()
	{
		publicGroup.GET("check", checkCTRS.Check)
		publicGroup.GET("test", checkCTRS.Test)
	}

	adminGroup := route.Group("/admin/admin")
	//adminGroup.Use(middleware.JWTAuth())
	{
		AdminRouterGroup(adminGroup)
	}

	buyerGroup := route.Group("/buyer")
	//buyerGroup.Use(middleware.JWTAuth())
	{
		BuyerRouterGroup(buyerGroup)
	}

	sellerGroup := route.Group("/seller/seller")
	//sellerGroup.Use(middleware.JWTAuth())
	{
		SellerRouterGroup(sellerGroup)
	}
}
