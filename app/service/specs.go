package service

import (
	"github.com/davveo/lemonShop/app/entity"
	"github.com/davveo/lemonShop/mgr"
	dbLocal "github.com/davveo/lemonShop/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

var _ SpecsService = (*specsService)(nil)

type SpecsService interface {
	i()
	List(ctx *gin.Context, req *entity.SpecsListRequest) (*entity.SpecsListResp, error)
	Get(ctx *gin.Context, req *entity.SpecsRequest) (*entity.SpecsBase, error)
	Create(ctx *gin.Context, req *entity.SpecsCreateRequest) (*entity.SpecsBase, error)
	Update(ctx *gin.Context, req *entity.SpecsUpdateRequest) (*entity.SpecsBase, error)
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

func (s *specsService) List(ctx *gin.Context, req *entity.SpecsListRequest) (*entity.SpecsListResp, error) {
	ops := []mgr.Option{
		s.specificationMgr.WithDisabled(1),
		s.specificationMgr.WithSellerID(0),
	}
	if req.KeyWord != "" {
		ops = append(ops, s.specificationMgr.WithSpecName(req.KeyWord))
	}

	resultPage, err := s.specificationMgr.SelectPage(mgr.NewPage(
		req.PageSize, req.PageNo, mgr.BuildDesc("spec_id")), ops...)
	if err != nil {
		return nil, err
	}

	return &entity.SpecsListResp{
		PageNo:    req.PageNo,
		PageSize:  req.PageSize,
		DataTotal: resultPage.GetTotal(),
		Data:      resultPage.GetRecords(),
	}, nil
}

func (s *specsService) Get(ctx *gin.Context, req *entity.SpecsRequest) (*entity.SpecsBase, error) {
	rt := new(entity.SpecsBase)
	obj, err := s.specificationMgr.FetchByPrimaryKey(int(req.SpecId))
	if err != nil {
		return nil, err
	}
	if err = copier.Copy(rt, &obj); err != nil {
		return nil, err
	}
	return rt, nil
}

func (s *specsService) Create(ctx *gin.Context, req *entity.SpecsCreateRequest) (*entity.SpecsBase, error) {
	// extra params
	//mapData["seller_id"] = "0"
	//mapData["disabled"] = 1
	return nil, nil
}

func (s *specsService) Update(ctx *gin.Context, req *entity.SpecsUpdateRequest) (*entity.SpecsBase, error) {
	return nil, nil
}

func (s *specsService) Delete(ctx *gin.Context, req []int64) error {
	return nil
}

func (s *specsService) i() {}
