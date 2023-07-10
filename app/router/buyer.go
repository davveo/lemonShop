package router

import (
	"github.com/davveo/lemonShop/app/ctrs"
	"github.com/gin-gonic/gin"
)

func BuyerRouterGroup(router *gin.RouterGroup) {
	// 授权模块
	PassportGroup(router)
	// 会员相关
	MemberGroup(router)
	// 交易模块
	TradeGroup(router)
	// 商品模块
	GoodsGroup(router)
}

func PassportGroup(router *gin.RouterGroup) {
	passportGroup := router.Group("passport/")
	specsController := ctrs.NewSpecsController()
	{
		passportGroup.GET("connect/wechat/auth", specsController.SpecsList)
		passportGroup.GET("connect/wechat/auth/back", specsController.Create)
		passportGroup.GET("connect/wechat/login", specsController.Update)
		passportGroup.GET("connect/pc/:type", specsController.Delete)
		passportGroup.GET("connect/wap/:type", specsController.Specs)
		passportGroup.GET("connect/:port/:type/callback", specsController.Specs)
		passportGroup.GET("account-binder/:type/callback", specsController.Specs)
		passportGroup.PUT("login-binder/pc/:uuid_connect", specsController.Specs)
		passportGroup.POST("mobile-binder/sms-code/:mobile", specsController.Specs)
		passportGroup.POST("mobile-binder/:uuid", specsController.Specs)
		passportGroup.POST("login-binder/wap/:uuid", specsController.Specs)
		passportGroup.GET("connect/app/:type/param", specsController.Specs)
		passportGroup.GET("connect/app/:type/openid", specsController.Specs)
		passportGroup.GET("login-binder/ali/info", specsController.Specs)
		passportGroup.POST("sms-binder/app", specsController.Specs)
		passportGroup.POST("login-binder/app", specsController.Specs)
		passportGroup.POST("register-binder/app", specsController.Specs)
		passportGroup.GET("smscode/:mobile", specsController.Specs)
		passportGroup.GET("username/:username", specsController.Specs)
		passportGroup.GET("mobile/:mobile", specsController.Specs)
		passportGroup.POST("token", specsController.Specs)
		passportGroup.POST("login/smscode/:mobile", specsController.Specs)
		passportGroup.POST("login", specsController.Specs)
		passportGroup.POST("login/:mobile", specsController.Specs)
		passportGroup.GET("auto-login", specsController.Specs)
		passportGroup.GET("decrypt", specsController.Specs)
		passportGroup.POST("register-bind/:uuid", specsController.Specs)
		passportGroup.POST("register/smscode/:mobile", specsController.Specs)
		passportGroup.POST("register/pc", specsController.Specs)
		passportGroup.POST("register/wap", specsController.Specs)
	}
}

func MemberGroup(router *gin.RouterGroup) {
	memberGroup := router.Group("members/")
	memberController := ctrs.NewMemberController()
	{
		// 查询会员商品咨询回复消息列表
		memberGroup.GET("asks/message", memberController.Todo)
		// 将会员商品咨询消息设置为已读
		memberGroup.PUT("asks/message/:ids/read", memberController.Todo)
		// 删除会员商品咨询消息
		memberGroup.DELETE("asks/message/:ids", memberController.Todo)
	}
}

func TradeGroup(router *gin.RouterGroup) {
	tradeGroup := router.Group("trade/")
	tradeController := ctrs.NewTradeController()
	{
		// 购物车相关
		{
			// ⭐️⭐️⭐️⭐️⭐️ 向购物车中添加一个产品
			tradeGroup.POST("carts/", tradeController.Todo)
			// ⭐️⭐️⭐️⭐️⭐️ 清空购物车
			tradeGroup.DELETE("carts/", tradeController.Todo)
			// ⭐️⭐️⭐️⭐️⭐️ 删除购物车中的一个或多个产品
			tradeGroup.DELETE("carts/:sku_ids/sku", tradeController.Todo)
			// ⭐️⭐️⭐️⭐️⭐️ 立即购买
			tradeGroup.POST("carts/buy", tradeController.Todo)
			// ⭐️⭐️⭐️⭐️⭐️ 获取购物车页面购物车详情
			tradeGroup.GET("carts/all", tradeController.Todo)
			// ⭐️⭐️⭐️⭐️⭐️ 获取结算页面购物车详情
			tradeGroup.GET("carts/checked", tradeController.Todo)
			// ⭐️⭐️⭐️⭐️⭐️ 更新购物车中的多个产品
			tradeGroup.POST("carts/sku/:sku_id", tradeController.Todo)
			// ⭐️⭐️⭐️⭐️⭐️ 设置全部商为选中或不选中
			tradeGroup.POST("carts/checked", tradeController.Todo)
			// ⭐️⭐️⭐️⭐️⭐️ 批量设置某商家的商品为选中或不选中
			tradeGroup.POST("carts/seller/:seller_id", tradeController.Todo)
		}

		// 结算参数接口模块
		{
			// ⭐️⭐️⭐️⭐️⭐️ 设置收货地址id
			tradeGroup.POST("checkout-params/address-id/:address_id", tradeController.Todo)
			// ⭐️⭐️⭐️⭐️⭐️ 设置支付类型
			tradeGroup.POST("checkout-params/payment-type", tradeController.Todo)
			// ⭐️⭐️⭐️⭐️⭐️ 设置发票信息
			tradeGroup.POST("checkout-params/receipt", tradeController.Todo)
			// ⭐️⭐️⭐️⭐️⭐️ 取消发票
			tradeGroup.DELETE("checkout-params/receipt", tradeController.Todo)
			// ⭐️⭐️⭐️⭐️⭐️ 设置送货时间
			tradeGroup.POST("checkout-params/receive-time", tradeController.Todo)
			// ⭐️⭐️⭐️⭐️⭐️ 设置订单备注
			tradeGroup.POST("checkout-params/remark", tradeController.Todo)
			// ⭐️⭐️⭐️⭐️⭐️ 获取结算参数
			tradeGroup.GET("checkout-params", tradeController.Todo)
		}

		// 商品交易模块
		{
			// 查询某商品的销售记录
			tradeGroup.GET("goods/:goods_id/sales", tradeController.Todo)
		}

		// 投诉主题相关
		{
			// 查询投诉主题列表
			tradeGroup.GET("order-complains/topics", tradeController.Todo)
		}

		// 物流查询接口
		{
			// 查询物流详细
			tradeGroup.GET("express", tradeController.Todo)
		}

		// 交易快照相关
		{
			// 查询一个交易快照
			tradeGroup.GET("snapshots/:id", tradeController.Todo)
		}

		// 会员订单相关
		{
			// 查询会员订单列表
			tradeGroup.GET("orders", tradeController.Todo)
			// 查询单个订单明细
			tradeGroup.GET("orders/:order_sn", tradeController.Todo)
			// ⭐️⭐️⭐️⭐️⭐️ 确认收货
			tradeGroup.POST("orders/:order_sn/rog", tradeController.Todo)
			// ⭐️⭐️⭐️⭐️⭐️ 取消订单
			tradeGroup.POST("orders/:order_sn/cancel", tradeController.Todo)
			// 查询订单状态的数量
			tradeGroup.GET("orders/status-num", tradeController.Todo)
			// 根据交易编号查询订单列表 {trade_sn}/list --> /list?trade_sn=xxx
			tradeGroup.GET("orders/list", tradeController.Todo)
			// 根据交易编号或者订单编号查询收银台数据
			tradeGroup.GET("orders/cashier", tradeController.Todo)
			// 订单流程图数据
			tradeGroup.GET("orders/:order_sn/flow", tradeController.Todo)
			// 查询订单日志
			tradeGroup.GET("orders/:order_sn/log", tradeController.Todo)
		}

		// 交易投诉相关
		{
			// 查询交易投诉表列表
			tradeGroup.GET("order-complains", tradeController.Todo)
			// 添加交易投诉
			tradeGroup.POST("order-complains", tradeController.Todo)
			// 撤销交易投诉
			tradeGroup.PUT("order-complains/:id", tradeController.Todo)
			// 提交对话
			tradeGroup.PUT("order-complains/:id/communication", tradeController.Todo)
			// 查询一个交易投诉
			tradeGroup.GET("order-complains/:id", tradeController.Todo)
			// 查询一个交易投诉的流程图
			tradeGroup.GET("order-complains/:id/flow", tradeController.Todo)
			// 对话中提交仲裁，随时可以提交仲裁
			tradeGroup.PUT("order-complains/:id/to-arbitration", tradeController.Todo)
		}

		// 交易接口模块
		{
			// ⭐️⭐️⭐️⭐️⭐️ 创建交易
			tradeGroup.POST("create", tradeController.Todo)
		}

		// 购物车价格计算
		{
			// ⭐️⭐️⭐️⭐️⭐️ 选择要参与的促销活动
			tradeGroup.POST("promotion", tradeController.Todo)
			// ⭐️⭐️⭐️⭐️⭐️ 取消参与促销
			tradeGroup.DELETE("promotion", tradeController.Todo)
			// ⭐️⭐️⭐️⭐️⭐️ 设置优惠券
			tradeGroup.POST("promotion/:seller_id/seller/:mc_id/coupon", tradeController.Todo)
		}
	}
}

func GoodsGroup(router *gin.RouterGroup) {
	goodsGroup := router.Group("goods/")
	goodsController := ctrs.NewGoodsController()

	// 商品分类
	{
		// 首页等商品分类数据
		goodsGroup.GET("categories/:parent_id/children", goodsController.Todo)
		// ip
		goodsGroup.GET("ip", goodsController.Todo)
		// 浏览商品的详情,静态部分使用
		goodsGroup.GET(":goods_id", goodsController.Todo)
		// 获取sku信息，商品详情页动态部分
		goodsGroup.GET(":goods_id/skus", goodsController.Todo)
		// 记录浏览器商品次数
		goodsGroup.GET(":goods_id/visit", goodsController.Todo)
		// 查看商品是否在配送区域 1 有货  0 无货
		goodsGroup.GET(":goods_id/area/:area_id", goodsController.Todo)
		// 查询商品列表
		goodsGroup.GET("search", goodsController.Todo)
		// 查询商品选择器
		goodsGroup.GET("search/selector", goodsController.Todo)
		// 查询商品分词对应数量
		goodsGroup.GET("search/words", goodsController.Todo)
		// 查询标签商品列表
		goodsGroup.GET("tags/:mark/goods", goodsController.Todo)
		// 查询商品销量
		goodsGroup.GET("tags/count", goodsController.Todo)
		// 查询某店铺标签商品的数量 tags/{shop_id}/goods-num --> tags/goods-num?shop_id=xxx
		goodsGroup.GET("tags/goods-num", goodsController.Todo)
	}
}
