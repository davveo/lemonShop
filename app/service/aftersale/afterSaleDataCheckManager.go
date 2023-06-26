package aftersale

import (
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/models/dto"
	"github.com/davveo/lemonShop/models/vo"
)

// AfterSaleDataCheckService 售后服务数据校验接口
type AfterSaleDataCheckService interface {
	// CheckApplyService 校验售后服务申请参数并获取构成售后服务单的相关数据
	CheckApplyService(applyAfterSaleVO *vo.ApplyAfterSaleVO) (*dto.AfterSaleApplyDTO, error)
	// CheckCancelOrder 校验取消订单申请参数并获取构成售后服务单的相关数据
	CheckCancelOrder(refundApplyVO *vo.RefundApplyVO) (*dto.AfterSaleApplyDTO, error)
	// CheckRefundInfo 校验退货或取消订单申请参数中的退款相关信息
	CheckRefundInfo(refundDO *models.EsAsRefund) error
	// CheckAfterSaleExpress 校验用户退还售后商品填写的物流信息并返回物流公司名称
	CheckAfterSaleExpress(afterSaleExpressDO *models.EsAsExpress) (string, error)
	// CheckAudit 商家审核售后服务申请校验并返回售后服务单详细信息
	CheckAudit(serviceSn, auditStatus string, refundPrice float64, returnAddr, auditRemark string) (*vo.ApplyAfterSaleVO, error)
	// CheckPutInWarehouse 商家入库操作校验
	CheckPutInWarehouse(remark, serviceSn string, storageList []*dto.PutInWarehouseDTO) error
	// CheckCloseAfterSale 校验关闭售后服务单参数
	CheckCloseAfterSale(serviceSn, remark string) error
	// CheckAdminRefund 校验平台退款参数
	CheckAdminRefund(remark, serviceSn string, refundPrice float64) error
}
