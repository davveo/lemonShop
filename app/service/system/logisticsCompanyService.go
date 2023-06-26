package system

import (
	"github.com/davveo/lemonShop/app/entity"
	"github.com/davveo/lemonShop/models"
)

type LogisticsCompanyService interface {
	// Page 查询物流公司列表
	Page(page, pageSize int, name string) (*entity.List, error)
	// Add 添加物流公司
	Add(mo *models.EsLogisticsCompany) (*models.EsLogisticsCompany, error)
	// Edit 修改物流公司
	Edit(mo *models.EsLogisticsCompany, id int64) (*models.EsLogisticsCompany, error)
	// Delete 删除物流公司
	Delete(id int64) error
	// GetModel 获取物流公司
	GetModel(id int64) (*models.EsLogisticsCompany, error)
	// GetLogiByCode 通过code获取物流公司
	GetLogiByCode(code string) (*models.EsLogisticsCompany, error)
	// GetLogiBykdCode 通过快递鸟物流code获取物流公司
	GetLogiBykdCode(kdcode string) (*models.EsLogisticsCompany, error)
	// GetLogiByName 根据物流名称查询物流信息
	GetLogiByName(name string) (*models.EsLogisticsCompany, error)
	// List 查询物流公司列表(不分页)
	List() ([]*models.EsLogisticsCompany, error)
	// OpenCloseLogi 开启或禁用物流公司
	openCloseLogi(id int64, disabled string) error
	// ListAllNormal 查询平台添加的全部物流公司（正常使用未删除的）
	ListAllNormal() ([]*models.EsLogisticsCompany, error)
}
