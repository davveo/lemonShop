package aftersale

import (
	"github.com/davveo/lemonShop/mgr"
	"github.com/davveo/lemonShop/models"
	dbLocal "github.com/davveo/lemonShop/pkg/db"
)

var _ AfterSaleChangeService = (*afterSaleChangeService)(nil)

// AfterSaleChangeService 售后服务退货地址业务接口
type (
	AfterSaleChangeService interface {
		i()
		// Add 新增售后服务退货地址信息
		Add(asChange *models.AsChange) error
		// FillChange 填充售后服务退货地址信息
		FillChange(serviceSn string, asChange *models.AsChange) (*models.AsChange, error)
		// GetModel 根据售后服务单号获取收货地址相关信息
		GetModel(serviceSn string) (*models.AsChange, error)
	}
	afterSaleChangeService struct {
		asChangeMgr *mgr.AsChangeMgr
	}
)

func (a afterSaleChangeService) Add(asChange *models.AsChange) error {
	return a.asChangeMgr.Create(asChange)
}

func (a afterSaleChangeService) FillChange(serviceSn string, asChange *models.AsChange) (*models.AsChange, error) {
	asChange.ServiceSn = serviceSn
	if err := a.Add(asChange); err != nil {
		return nil, err
	}
	return asChange, nil
}

func (a afterSaleChangeService) GetModel(serviceSn string) (*models.AsChange, error) {
	return a.asChangeMgr.GetByOption(a.asChangeMgr.WithServiceSn(serviceSn))
}

func (a afterSaleChangeService) i() {}

func NewAfterSaleChangeService() AfterSaleChangeService {
	return &afterSaleChangeService{
		asChangeMgr: mgr.NewAsChangeMgr(dbLocal.GRpo),
	}
}
