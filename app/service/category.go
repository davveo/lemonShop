package service

import (
	"github.com/davveo/lemonShop/mgr"
	dbLocal "github.com/davveo/lemonShop/pkg/db"
)

var _ CategoryService = (*categoryService)(nil)

type CategoryService interface {
	i()
}

type categoryService struct {
	categoryMgr     *mgr.CategoryMgr
	categorySpecMgr *mgr.CategorySpecMgr
}

func NewCategoryService() CategoryService {
	return &categoryService{
		categorySpecMgr: mgr.NewCategorySpecMgr(dbLocal.GRpo),
		categoryMgr:     mgr.NewCategoryMgr(dbLocal.GRpo),
	}
}

func (c categoryService) i() {}
