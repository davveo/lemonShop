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
	{
		AdminRouterGroup(adminGroup)
	}

	buyerGroup := route.Group("/buyer")
	{
		BuyerRouterGroup(buyerGroup)
	}

	sellerGroup := route.Group("/seller/seller")
	{
		SellerRouterGroup(sellerGroup)
	}
}
