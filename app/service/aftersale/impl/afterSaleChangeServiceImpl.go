package impl

import (
	"github.com/davveo/lemonShop/app/service/aftersale"
	"github.com/davveo/lemonShop/mgr"
	"github.com/davveo/lemonShop/models"
	dbLocal "github.com/davveo/lemonShop/pkg/db"
)

type afterSaleChangeServiceImpl struct {
	asChangeMgr *mgr.AsChangeMgr
}

func (a afterSaleChangeServiceImpl) Add(asChange *models.AsChange) error {
	return a.asChangeMgr.Create(asChange)
}

func (a afterSaleChangeServiceImpl) FillChange(serviceSn string, asChange *models.AsChange) (*models.AsChange, error) {
	asChange.ServiceSn = serviceSn
	if err := a.Add(asChange); err != nil {
		return nil, err
	}
	return asChange, nil
}

func (a afterSaleChangeServiceImpl) GetModel(serviceSn string) (*models.AsChange, error) {
	return a.asChangeMgr.GetByOption(a.asChangeMgr.WithServiceSn(serviceSn))
}

func NewAfterSaleChangeService() aftersale.AfterSaleChangeService {
	return &afterSaleChangeServiceImpl{
		asChangeMgr: mgr.NewAsChangeMgr(dbLocal.GRpo),
	}
}
