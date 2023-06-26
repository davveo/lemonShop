package order

import (
	"github.com/davveo/lemonShop/app/entity"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/models/dto"
	"github.com/davveo/lemonShop/models/vo"
)

type OrderQueryService interface {
	// List 查询订单表列表
	List(paramDTO *dto.OrderQueryParam) (*entity.List, error)
	// Export 按条件导出订单数据
	export(paramDTO *dto.OrderQueryParam) ([]*vo.OrderLineVO, error)
	// GetModel 读取一个订单详细
	GetModel(orderSn string, queryParam *dto.OrderDetailQueryParam) (*vo.OrderDetailVO, error)
	// GetModelByOrderId 根据订单ID读取一个订单详细
	GetModelByOrderId(orderId int64) (*models.EsOrder, error)
	// GetOrder 根据订单编号读取一个订单详细
	GetOrder(orderSn string) (*models.EsOrder, error)
	// GetOrderStatusNum 读取订单状态的订单数
	GetOrderStatusNum(memberId, sellerId int64) (*vo.OrderStatusNumVO, error)
	// OrderItems 获取某订单的订单项
	OrderItems(orderSn string) ([]*models.EsOrderItems, error)
	// GetOrderFlow 根据订单sn读取，订单的流程
	GetOrderFlow(orderSn string) ([]*vo.OrderFlowNode, error)
	// GetOrderNumByMemberId 读取会员所有的订单数量
	GetOrderNumByMemberId(memberId int64) (int, error)
	// GetOrderCommentNumByMemberId 读取会员(评论状态)订单数量
	GetOrderCommentNumByMemberId(memberId int64, commentStatus string) (int, error)
	// GetOrderByTradeSn 读取订单列表根据交易编号——系统内部使用
	GetOrderByTradeSn(tradeSn string) ([]*dto.OrderDetailDTO, error)
	// GetOrderByTradeSnAndMember 读取订单列表根据交易编号
	GetOrderByTradeSnAndMember(tradeSn string, memberId int64) ([]*vo.OrderDetailVO, error)
	// GetModelByOrderSn 查询一个订单的详细
	GetModelByOrderSn(orderSn string) (*dto.OrderDetailDTO, error)
	// GetOrderRefundPrice 获取订单可退款总额
	GetOrderRefundPrice(orderSn string) (float64, error)
	// ListOrderByGoods 获取几个月之内购买过相关商品的订单数据
	ListOrderByGoods(goodsId, memberId, month int64) ([]*models.EsOrder, error)
	// GetInvoice 根据订单orderId查询发货单相关信息
	GetInvoice(orderId int64) (*vo.InvoiceVO, error)
}
