package aftersale

import (
	"github.com/davveo/lemonShop/app/consts"
	"github.com/davveo/lemonShop/app/entity"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/models/dto"
	"github.com/davveo/lemonShop/models/vo"
)

var _ AfterSaleRefundService = (*afterSaleRefundService)(nil)

type (
	AfterSaleRefundService interface {
		i()
		// List 查询退款单列表
		List(dto *dto.RefundQueryParam) (*entity.List, error)
		// AdminRefund 平台退款操作 主要用于平台人工线下退款
		AdminRefund(serviceSn string, refundPrice float64, remark string, typeEnum consts.ServiceOperateTypeEnum) error
		// SellerRefund 在线支付订单商家退款操作
		SellerRefund(serviceSn string) error
		// FillRefund 填充并创建退款单
		FillRefund(refundPrice float64, applyAfterSaleVO vo.ApplyAfterSaleVO) (*models.EsRefund, error)
		// AddRefund 新增退款单
		AddRefund(refundDO *models.EsRefund) (int, error)
		// GetModel 获取退款单
		GetModel(serviceSn string) (*models.EsRefund, error)
		// FillAfterSaleRefund 填充并入库申请售后服务时的退款相关信息
		FillAfterSaleRefund(serviceSn string, returnNum int, orderDO *models.EsOrder, itemsDO *models.EsOrderItems, refundDO *models.EsAsRefund) error
		// FillCancelOrderRefund 填充并入库取消订单时的退款相关信息
		FillCancelOrderRefund(serviceSn string, orderDO *models.EsOrder, refundApplyVO *vo.RefundApplyVO) error
		// AddAfterSaleRefund 添加申请售后服务时的退款相关信息
		AddAfterSaleRefund(refundDO *models.EsAsRefund)
		// GetAfterSaleRefund 根据售后服务单号获取售后退款账户信息
		GetAfterSaleRefund(serviceSn string) (*models.EsAsRefund, error)
	}
	afterSaleRefundService struct {
	}
)

func (a afterSaleRefundService) i() {}

func (a afterSaleRefundService) List(dto *dto.RefundQueryParam) (*entity.List, error) {
	//TODO implement me
	panic("implement me")
}

func (a afterSaleRefundService) AdminRefund(serviceSn string, refundPrice float64, remark string, typeEnum consts.ServiceOperateTypeEnum) error {
	//TODO implement me
	panic("implement me")
}

func (a afterSaleRefundService) SellerRefund(serviceSn string) error {
	//TODO implement me
	panic("implement me")
}

func (a afterSaleRefundService) FillRefund(refundPrice float64, applyAfterSaleVO vo.ApplyAfterSaleVO) (*models.EsRefund, error) {
	//TODO implement me
	panic("implement me")
}

func (a afterSaleRefundService) AddRefund(refundDO *models.EsRefund) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (a afterSaleRefundService) GetModel(serviceSn string) (*models.EsRefund, error) {
	//TODO implement me
	panic("implement me")
}

func (a afterSaleRefundService) FillAfterSaleRefund(serviceSn string, returnNum int, orderDO *models.EsOrder, itemsDO *models.EsOrderItems, refundDO *models.EsAsRefund) error {
	//TODO implement me
	panic("implement me")
}

func (a afterSaleRefundService) FillCancelOrderRefund(serviceSn string, orderDO *models.EsOrder, refundApplyVO *vo.RefundApplyVO) error {
	//TODO implement me
	panic("implement me")
}

func (a afterSaleRefundService) AddAfterSaleRefund(refundDO *models.EsAsRefund) {
	//TODO implement me
	panic("implement me")
}

func (a afterSaleRefundService) GetAfterSaleRefund(serviceSn string) (*models.EsAsRefund, error) {
	//TODO implement me
	panic("implement me")
}

func NewAfterSaleRefundService() AfterSaleRefundService {
	return &afterSaleRefundService{}
}
