package impl

import (
	"errors"

	"github.com/davveo/lemonShop/app/entity"
	"github.com/davveo/lemonShop/app/service/aftersale"
	"github.com/davveo/lemonShop/app/service/system"
	"github.com/davveo/lemonShop/app/service/trade/order"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/models/dto"
	"github.com/davveo/lemonShop/models/vo"
)

type afterSaleDataCheckServiceImpl struct {
	orderQueryService       *order.OrderQueryService
	logisticsCompanyService *system.LogisticsCompanyService
	afterSaleQueryService   *aftersale.AfterSaleQueryService
}

func (a afterSaleDataCheckServiceImpl) CheckApplyService(applyAfterSaleVO *vo.ApplyAfterSaleVO) (*dto.AfterSaleApplyDTO, error) {
	userContext := entity.UserContext{}
	buyer, _ := userContext.GetBuyer()
	if buyer == nil {
		return nil, errors.New("当前会员已经退出登录")
	}

	return nil, nil
}

func (a afterSaleDataCheckServiceImpl) CheckCancelOrder(refundApplyVO *vo.RefundApplyVO) (*dto.AfterSaleApplyDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (a afterSaleDataCheckServiceImpl) CheckRefundInfo(refundDO *models.EsAsRefund) error {
	//TODO implement me
	panic("implement me")
}

func (a afterSaleDataCheckServiceImpl) CheckAfterSaleExpress(afterSaleExpressDO *models.EsAsExpress) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (a afterSaleDataCheckServiceImpl) CheckAudit(serviceSn, auditStatus string, refundPrice float64, returnAddr, auditRemark string) (*vo.ApplyAfterSaleVO, error) {
	//TODO implement me
	panic("implement me")
}

func (a afterSaleDataCheckServiceImpl) CheckPutInWarehouse(remark, serviceSn string, storageList []*dto.PutInWarehouseDTO) error {
	//TODO implement me
	panic("implement me")
}

func (a afterSaleDataCheckServiceImpl) CheckCloseAfterSale(serviceSn, remark string) error {
	//TODO implement me
	panic("implement me")
}

func (a afterSaleDataCheckServiceImpl) CheckAdminRefund(remark, serviceSn string, refundPrice float64) error {
	//TODO implement me
	panic("implement me")
}

func NewAfterSaleDataCheckService() aftersale.AfterSaleDataCheckService {
	return &afterSaleDataCheckServiceImpl{}
}
