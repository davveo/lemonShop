package aftersale

import "github.com/davveo/lemonShop/models"

// AfterSaleChangeService 售后服务退货地址业务接口
type AfterSaleChangeService interface {
	// Add 新增售后服务退货地址信息
	Add(asChange *models.AsChange) error
	// FillChange 填充售后服务退货地址信息
	FillChange(serviceSn string, asChange *models.AsChange) (*models.AsChange, error)
	// GetModel 根据售后服务单号获取收货地址相关信息
	GetModel(serviceSn string) (*models.AsChange, error)
}
