package aftersale

import (
	"github.com/davveo/lemonShop/app/entity"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/models/dto"
	"github.com/davveo/lemonShop/models/vo"
)

type AfterSaleQueryService interface {
	// List 获取申请售后服务记录列表
	List(param *dto.AfterSaleQueryParam) (*entity.List, error)
	// ApplyOrderInfo 获取申请售后页面订单收货信息和商品信息
	ApplyOrderInfo(orderSn string, skuId int) (*dto.AfterSaleOrderDTO, error)
	// Detail 获取售后服务详细
	Detail(serviceSn string) (*vo.ApplyAfterSaleVO, error)
	// ExportAfterSale 根据条件导出售后服务信息
	ExportAfterSale(param *dto.AfterSaleQueryParam) ([]*vo.AfterSaleExportVO, error)
	// GetAfterSaleCount 获取还未完成的售后服务
	GetAfterSaleCount(memberId, sellerId int64) (int, error)
	// GetService 根据售后服务单编号获取售后服务单信息
	GetService(serviceSn string) (*models.EsAsOrder, error)
	// GetCancelService 根据订单编号获取取消订单售后服务信息
	GetCancelService(orderSn string) (*models.EsAsOrder, error)
	// GetExpress 根据售后服务单号获取物流相关信息
	GetExpress(serviceSn string) (*models.EsAsExpress, error)
}
