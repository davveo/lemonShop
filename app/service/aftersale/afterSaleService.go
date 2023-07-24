package aftersale

import (
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/models/dto"
	"github.com/davveo/lemonShop/models/vo"
)

var _ AfterSaleService = (*afterSaleService)(nil)

type (
	AfterSaleService interface {
		i()
		// Apply 申请售后服务
		Apply(applyAfterSaleVO *vo.ApplyAfterSaleVO) error
		// ApplyCancelOrder 申请取消订单
		ApplyCancelOrder(refundApplyVO *vo.RefundApplyVO) error
		// ApplyShip 保存用户填写的退货物流信息
		ApplyShip(afterSaleExpressDO *models.EsAsExpress) error
		// Audit 商家审核售后服务
		Audit(serviceSn, auditStatus, returnAddr, auditRemark string, refundPrice float64) error
		// PutInWarehouse 商家将申请售后服务退还的商品入库
		PutInWarehouse(serviceSn, remark string, storageList []dto.PutInWarehouseDTO) error
		// CloseAfterSale 关闭售后服务单
		CloseAfterSale(serviceSn, remark string) error
		// EditAfterSaleShopName 修改售后服务单中的店铺名称 当商家修改店铺名称是调用此方法
		EditAfterSaleShopName(shopId int, shopName string) error
		// CancelPintuanOrder 系统对拼团订单执行取消操作 此方法针对的是拼团订单在拼团活动结束时没有成团的情况下使用
		CancelPintuanOrder(orderSn, cancelReason string) error
		// AddAfterSaleService 添加售后服务单基础信息
		AddAfterSaleService(serviceDO *models.EsAsOrder) error
		// UpdateServiceStatus 修改售后服务单状态
		UpdateServiceStatus(serviceSn, serviceStatus string) error
		// UpdateServiceStatusMore 修改售后服务单状态和备注
		UpdateServiceStatusMore(serviceSn, serviceStatus, auditRemark, storageRemark, refundRemark, closeReason string) error
		// SetServiceNewOrderSn 修改售后服务单新订单号 换货或补发商品重新生成订单后调用
		SetServiceNewOrderSn(serviceSn, newOrderSn string)
	}
	afterSaleService struct {
	}
)

func (a afterSaleService) i() {}

func (a afterSaleService) Apply(applyAfterSaleVO *vo.ApplyAfterSaleVO) error {
	//TODO implement me
	panic("implement me")
}

func (a afterSaleService) ApplyCancelOrder(refundApplyVO *vo.RefundApplyVO) error {
	//TODO implement me
	panic("implement me")
}

func (a afterSaleService) ApplyShip(afterSaleExpressDO *models.EsAsExpress) error {
	//TODO implement me
	panic("implement me")
}

func (a afterSaleService) Audit(serviceSn, auditStatus, returnAddr, auditRemark string, refundPrice float64) error {
	//TODO implement me
	panic("implement me")
}

func (a afterSaleService) PutInWarehouse(serviceSn, remark string, storageList []dto.PutInWarehouseDTO) error {
	//TODO implement me
	panic("implement me")
}

func (a afterSaleService) CloseAfterSale(serviceSn, remark string) error {
	//TODO implement me
	panic("implement me")
}

func (a afterSaleService) EditAfterSaleShopName(shopId int, shopName string) error {
	//TODO implement me
	panic("implement me")
}

func (a afterSaleService) CancelPintuanOrder(orderSn, cancelReason string) error {
	//TODO implement me
	panic("implement me")
}

func (a afterSaleService) AddAfterSaleService(serviceDO *models.EsAsOrder) error {
	//TODO implement me
	panic("implement me")
}

func (a afterSaleService) UpdateServiceStatus(serviceSn, serviceStatus string) error {
	//TODO implement me
	panic("implement me")
}

func (a afterSaleService) UpdateServiceStatusMore(serviceSn, serviceStatus, auditRemark, storageRemark, refundRemark, closeReason string) error {
	//TODO implement me
	panic("implement me")
}

func (a afterSaleService) SetServiceNewOrderSn(serviceSn, newOrderSn string) {
	//TODO implement me
	panic("implement me")
}

func NewAfterSaleService() AfterSaleService {
	return &afterSaleService{}
}
