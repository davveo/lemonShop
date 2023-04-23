package service

import (
	"errors"
	"github.com/davveo/lemonShop/app/consts"
	"github.com/davveo/lemonShop/app/entity"
	"github.com/davveo/lemonShop/mgr"
	"github.com/davveo/lemonShop/models"
	dbLocal "github.com/davveo/lemonShop/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

var _ SpecsService = (*specsService)(nil)

type SpecsService interface {
	i()
	List(ctx *gin.Context, req *entity.SpecsListRequest) (*entity.List, error)
	Get(ctx *gin.Context, req *entity.Specs) error
	Create(ctx *gin.Context, req *entity.Specs) error
	Update(ctx *gin.Context, req *entity.Specs) error
	Delete(ctx *gin.Context, req []int64) error
}

type specsService struct {
	specificationMgr *mgr.SpecificationMgr
}

func NewSpecsService() SpecsService {
	return &specsService{
		specificationMgr: mgr.NewSpecificationMgr(dbLocal.GRpo),
	}
}

func (s *specsService) List(ctx *gin.Context, req *entity.SpecsListRequest) (*entity.List, error) {
	ops := []mgr.Option{
		s.specificationMgr.WithDisabled(consts.DefaultDisabled),
		s.specificationMgr.WithSellerID(consts.DefaultSellerId),
	}
	if req.KeyWord != "" {
		ops = append(ops, s.specificationMgr.WithSpecName(req.KeyWord))
	}

	resultPage, err := s.specificationMgr.SelectPage(mgr.NewPage(
		req.PageSize, req.PageNo, mgr.BuildDesc("spec_id")), ops...)
	if err != nil {
		return nil, err
	}

	return &entity.List{
		PageNo:    req.PageNo,
		PageSize:  req.PageSize,
		DataTotal: resultPage.GetTotal(),
		Data:      resultPage.GetRecords(),
	}, nil
}

func (s *specsService) Get(ctx *gin.Context, specs *entity.Specs) error {
	obj, err := s.specificationMgr.FetchByPrimaryKey(int(specs.SpecID))
	if err != nil {
		return err
	}
	if err = copier.Copy(specs, obj); err != nil {
		return err
	}
	return nil
}

func (s *specsService) Create(ctx *gin.Context, specs *entity.Specs) error {
	ops := []mgr.Option{
		s.specificationMgr.WithDisabled(consts.DefaultDisabled),
		s.specificationMgr.WithSellerID(consts.DefaultSellerId),
		s.specificationMgr.WithSpecName(specs.SpecName),
	}
	list, err := s.specificationMgr.GetByOptions(ops...)
	if err != nil {
		return err
	}
	if len(list) > 0 {
		return errors.New("规格名称重复")
	}
	specs.Disabled = consts.DefaultDisabled
	specification := new(models.EsSpecification)
	if err := copier.Copy(specification, specs); err != nil {
		return err
	}
	if err := s.specificationMgr.Create(specification); err != nil {
		return err
	}

	specs.SpecID = specification.SpecID
	return nil
}

func (s *specsService) Update(ctx *gin.Context, specs *entity.Specs) error {
	obj, err := s.specificationMgr.FetchByPrimaryKey(int(specs.SpecID))
	if err != nil {
		return err
	}
	if obj == nil {
		return errors.New("规格不存在")
	}

	return nil
}

func (s *specsService) Delete(ctx *gin.Context, req []int64) error {
	s.specificationMgr.GetByOptions()
	return nil
}

func (s *specsService) i() {}
