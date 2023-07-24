package aftersale

import (
	"context"
	"fmt"

	"github.com/davveo/lemonShop/models/vo"

	"github.com/davveo/lemonShop/app/consts"
	"github.com/davveo/lemonShop/models"
)

var _ SellerCreateTradeService = (*sellerCreateTradeService)(nil)

type (
	SellerCreateTradeService interface {
		i()
		// SystemCreateTrade 系统自动为售后服务创建新订单
		SystemCreateTrade(serviceSn string) (*models.EsOrder, error)
		// SellerCreateTrade 商家手动为售后服务创建新订单
		SellerCreateTrade(ctx context.Context, serviceSn string) (*models.EsOrder, error)
	}
	sellerCreateTradeService struct {
		afterSaleLogService   AfterSaleLogService
		afterSaleQueryService AfterSaleQueryService
		afterSaleService      AfterSaleService
	}
)

func (s sellerCreateTradeService) i() {}

func (s sellerCreateTradeService) SystemCreateTrade(serviceSn string) (*models.EsOrder, error) {
	applyAfterSaleVO, err := s.afterSaleQueryService.Detail(serviceSn)
	if err != nil {
		return nil, err
	}
	orderDo, err := s.createTrade(applyAfterSaleVO, "系统")
	if err != nil {
		s.afterSaleService.UpdateServiceStatus(serviceSn, string(consts.ERROR_EXCEPTION_SERVICE_STATUS))
		s.afterSaleLogService.Add(serviceSn,
			"系统自动生成新订单失败，原因："+err.Error()+"，等待商家手动创建新订单。", "系统")
		return nil, err
	}
	return orderDo, nil
}

func (s sellerCreateTradeService) SellerCreateTrade(ctx context.Context, serviceSn string) (*models.EsOrder, error) {
	applyAfterSaleVO, err := s.afterSaleQueryService.Detail(serviceSn)
	if err != nil {
		return nil, err
	}
	// todo
	fmt.Println(applyAfterSaleVO)
	return nil, nil
}

func (s sellerCreateTradeService) createTrade(applyAfterSaleVO *vo.ApplyAfterSaleVO, operator string) (*models.EsOrder, error) {
	return nil, nil
}

func NewSellerCreateTradeService() SellerCreateTradeService {
	return &sellerCreateTradeService{
		afterSaleLogService:   NewAfterSaleLogService(),
		afterSaleService:      NewAfterSaleService(),
		afterSaleQueryService: NewAfterSaleQueryService(),
	}
}
