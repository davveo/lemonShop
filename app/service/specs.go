package service

import (
	"errors"

	"github.com/davveo/lemonShop/models/vo"

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
	AddSellerSpec(ctx *gin.Context, categoryId int64, specName string) (*entity.Specs, error)
	GetCatSpecification(ctx *gin.Context, categoryId int64) ([]*vo.SelectVo, error)
	QuerySellerSpec(ctx *gin.Context, categoryId int64) ([]*vo.SpecificationVO, error)
}

type specsService struct {
	specificationMgr *mgr.SpecificationMgr
	categorySpecMgr  *mgr.CategorySpecMgr
	specValuesMgr    *mgr.SpecValuesMgr
}

func NewSpecsService() SpecsService {
	return &specsService{
		specValuesMgr:    mgr.NewSpecValuesMgr(dbLocal.GRpo),
		specificationMgr: mgr.NewSpecificationMgr(dbLocal.GRpo),
		categorySpecMgr:  mgr.NewCategorySpecMgr(dbLocal.GRpo),
	}
}

func (s *specsService) List(ctx *gin.Context, req *entity.SpecsListRequest) (*entity.List, error) {
	ops := []mgr.Option{
		s.specificationMgr.WithDisabled(consts.DefaultEnabled),
		s.specificationMgr.WithSellerID(consts.DefaultSellerId),
		s.specificationMgr.WithSpecName(req.KeyWord),
	}
	page := mgr.NewPage(req.PageSize, req.PageNo, mgr.BuildDesc("spec_id"))
	resultPage, err := s.specificationMgr.SelectPage(page, ops...)
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
		s.specificationMgr.WithDisabled(consts.DefaultEnabled),
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
	specs.Disabled = consts.DefaultEnabled
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
	categorySpecList, err := s.categorySpecMgr.GetBatchFromSpecID(req)
	if err != nil {
		return err
	}
	if len(categorySpecList) > 0 {
		return errors.New("有分类已经绑定要删除的规格，请先解绑分类规格")
	}
	if err = s.specificationMgr.Update(
		map[string]interface{}{
			"disabled": consts.DefaultDisabled,
		}, s.specificationMgr.WithInSpecID(req)); err != nil {
		return err
	}
	return nil
}

func (s *specsService) GetCatSpecification(ctx *gin.Context, categoryId int64) ([]*vo.SelectVo, error) {
	return s.specificationMgr.RawWithSelect(consts.CatSpecificationSqlString, categoryId, categoryId)
}

func (s *specsService) AddSellerSpec(ctx *gin.Context, categoryId int64, specName string) (*entity.Specs, error) {
	if obj, err := s.categorySpecMgr.FetchByPrimaryKey(categoryId); err != nil {
		return nil, err
	} else if obj == nil {
		return nil, errors.New("分类不存在")
	}

	// todo 完成用户session获取
	userContext := entity.UserContext{}
	seller, err := userContext.GetSeller()
	if err != nil {
		return nil, err
	}

	list, err := s.specificationMgr.RawToMap(consts.AddSellerSpecSql, categoryId, seller.SellerId, specName)
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		return nil, errors.New("规格名称重复")
	}

	specs := &entity.Specs{
		SpecName: specName,
		Disabled: consts.DefaultEnabled,
		SpecMemo: "商家自定义",
		SellerID: int(seller.SellerId),
	}

	if err = s.Create(ctx, specs); err != nil {
		return nil, err
	}

	categorySpec := &models.EsCategorySpec{
		CategoryID: int(categoryId),
		SpecID:     specs.SpecID,
	}
	if err := s.categorySpecMgr.Create(categorySpec); err != nil {
		return nil, err
	}

	return specs, nil

}

func (s *specsService) QuerySellerSpec(ctx *gin.Context, categoryId int64) ([]*vo.SpecificationVO, error) {
	var specificationVOList []*vo.SpecificationVO
	// todo 完成用户session获取
	userContext := entity.UserContext{}
	seller, err := userContext.GetSeller()
	if err != nil {
		return nil, err
	}

	specificationList, err := s.specificationMgr.RawToStruct(consts.QuerySpecSql, categoryId, seller.SellerId)
	if err != nil {
		return specificationVOList, err
	}
	if len(specificationList) <= 0 {
		return specificationVOList, nil
	}

	var specIds []int
	for _, specification := range specificationList {
		v := new(vo.SpecificationVO)
		specIds = append(specIds, specification.SpecID)
		if err := copier.Copy(v, &specification); err != nil {
			return specificationVOList, nil
		}
		specificationVOList = append(specificationVOList, v)
	}

	specValuesList, err := s.specValuesMgr.GetBatchFromSpecID(specIds)
	if err != nil {
		return specificationVOList, err
	}

	m := make(map[int][]*models.EsSpecValues)
	for _, specValues := range specValuesList {
		lst, ok := m[specValues.SpecID]
		if !ok {
			lst = []*models.EsSpecValues{}
		}
		lst = append(lst, specValues)
		m[specValues.SpecID] = lst
	}

	for _, specificationVO := range specificationVOList {
		specificationVO.ValueList = m[specificationVO.SpecID]
	}

	return specificationVOList, nil
}

func (s *specsService) i() {}
