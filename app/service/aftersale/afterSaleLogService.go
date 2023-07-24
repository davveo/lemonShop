package aftersale

import "github.com/davveo/lemonShop/models"

var _ AfterSaleLogService = (*afterSaleLogService)(nil)

type (
	AfterSaleLogService interface {
		i()
		// Add 新增售后日志
		Add(serviceSn, logDetail, operator string) error
		// List 根据售后服务单号获取售后日志信息集合
		List(serviceSn string) ([]*models.EsAsLog, error)
	}
	afterSaleLogService struct {
	}
)

func (a afterSaleLogService) i() {}

func (a afterSaleLogService) Add(serviceSn, logDetail, operator string) error {
	//TODO implement me
	panic("implement me")
}

func (a afterSaleLogService) List(serviceSn string) ([]*models.EsAsLog, error) {
	//TODO implement me
	panic("implement me")
}

func NewAfterSaleLogService() AfterSaleLogService {
	return &afterSaleLogService{}
}
