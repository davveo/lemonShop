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

var _ SpecsValuesService = (*specsValuesService)(nil)

type SpecsValuesService interface {
	i()
	Get(ctx *gin.Context, req *entity.SpecValues) error
	Create(ctx *gin.Context, req *entity.SpecValues) error
	Update(ctx *gin.Context, req *entity.SpecValues) error
	SaveSpecValue(ctx *gin.Context, specId int, valueList []string) ([]*entity.SpecValues, error)
	ListBySpecId(ctx *gin.Context, specId int64, permission consts.Permission) ([]*entity.SpecValues, error)
}

type specsValuesService struct {
	specValuesMgr *mgr.SpecValuesMgr
}

func (s *specsValuesService) ListBySpecId(ctx *gin.Context, specId int64,
	permission consts.Permission) ([]*entity.SpecValues, error) {
	ops := []mgr.Option{
		s.specValuesMgr.WithSpecID(int(specId)),
	}
	if permission == consts.AdminPermission {
		ops = append(ops, s.specValuesMgr.WithSellerID(consts.DefaultSellerId))
	}
	var specValuesList []*entity.SpecValues
	resList, err := s.specValuesMgr.GetByOptions(ops...)
	if err != nil {
		return nil, err
	}

	for _, res := range resList {
		specValues := new(entity.SpecValues)
		if err := copier.Copy(specValues, res); err != nil {
			return nil, err
		}
		specValuesList = append(specValuesList, specValues)
	}
	return specValuesList, nil
}

func (s *specsValuesService) Get(ctx *gin.Context, req *entity.SpecValues) error {
	obj, err := s.specValuesMgr.FetchByPrimaryKey(int(req.SpecValueID))
	if err != nil {
		return err
	}
	if err = copier.Copy(req, obj); err != nil {
		return err
	}
	return nil
}

func (s *specsValuesService) Create(ctx *gin.Context, req *entity.SpecValues) error {
	specValues := new(models.EsSpecValues)
	if err := copier.Copy(specValues, req); err != nil {
		return err
	}
	if err := s.specValuesMgr.Create(specValues); err != nil {
		return err
	}

	req.SpecValueID = specValues.SpecValueID
	return nil
}

func (s *specsValuesService) Update(ctx *gin.Context, req *entity.SpecValues) error {
	specValues := new(models.EsSpecValues)
	if err := copier.Copy(specValues, req); err != nil {
		return err
	}
	if err := s.specValuesMgr.UpdateByModel(specValues,
		s.specValuesMgr.WithSpecValueID(req.SpecValueID)); err != nil {
		return err
	}
	return nil
}

func (s *specsValuesService) SaveSpecValue(ctx *gin.Context,
	specId int, valueList []string) ([]*entity.SpecValues, error) {
	obj, err := s.specValuesMgr.FetchByPrimaryKey(specId)
	if err != nil {
		return nil, err
	}

	ops := []mgr.Option{
		s.specValuesMgr.WithSpecID(int(specId)),
		s.specValuesMgr.WithSellerID(consts.DefaultSellerId),
	}
	if err := s.specValuesMgr.DeleteByOpts(ops...); err != nil {
		return nil, err
	}
	var specValuesList []*entity.SpecValues
	for _, value := range valueList {
		if len(value) > 50 {
			return nil, errors.New("规格值为1到50个字符之间")
		}
		specValue := &models.EsSpecValues{
			SpecID:    specId,
			SpecValue: value,
			SellerID:  consts.DefaultSellerId,
			SpecName:  obj.SpecName,
		}
		if err := s.specValuesMgr.Create(specValue); err != nil {
			return nil, err
		}
		specValueEntity := new(entity.SpecValues)
		if err := copier.Copy(specValueEntity, specValue); err != nil {
			return nil, err
		}
		specValuesList = append(specValuesList, specValueEntity)
	}
	return specValuesList, nil
}

func NewSpecsValuesService() SpecsValuesService {
	return &specsValuesService{
		specValuesMgr: mgr.NewSpecValuesMgr(dbLocal.GRpo),
	}
}

func (s *specsValuesService) i() {}
