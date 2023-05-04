package service

import (
	"errors"
	"strconv"
	"strings"

	"github.com/davveo/lemonShop/pkg/cache"

	"github.com/davveo/lemonShop/app/consts"

	"github.com/davveo/lemonShop/models"

	"github.com/davveo/lemonShop/app/entity"
	"github.com/davveo/lemonShop/mgr"
	dbLocal "github.com/davveo/lemonShop/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

var _ CategoryService = (*categoryService)(nil)

type CategoryService interface {
	i()
	Create(ctx *gin.Context, req *entity.Category) error
	Get(ctx *gin.Context, req *entity.Category) error
}

type categoryService struct {
	categoryMgr     *mgr.CategoryMgr
	categorySpecMgr *mgr.CategorySpecMgr
}

func (c categoryService) Create(ctx *gin.Context, req *entity.Category) error {
	list, err := c.categorySpecMgr.GetByOptions(c.categoryMgr.WithName(req.Name))
	if err != nil {
		return err
	}
	if len(list) > 0 {
		return errors.New("该分类名称已存在")
	}
	var parent *models.EsCategory
	// 非顶级分类
	if req.ParentID != 0 {
		parent, err = c.categoryMgr.FetchByPrimaryKey(int(req.ParentID))
		if err != nil {
			return err
		}
		if len(strings.Split(parent.CategoryPath, "|")) > 4 {
			return errors.New("最多为三级分类,添加失败")
		}
	}
	category := new(models.EsCategory)
	if err := copier.Copy(category, req); err != nil {
		return err
	}
	if err := c.categoryMgr.Create(category); err != nil {
		return err
	}
	req.CategoryID = category.CategoryID

	if parent != nil {
		req.CategoryPath = parent.CategoryPath + strconv.Itoa(category.CategoryID) + "|"
	} else {
		req.CategoryPath = "0|" + strconv.Itoa(category.CategoryID) + "|"
	}

	if err = c.categoryMgr.UpdateByModel(
		&models.EsCategory{CategoryPath: req.CategoryPath},
		c.categoryMgr.WithCategoryID(category.CategoryID)); err != nil {
		return err
	}
	if !cache.Cache.Del(c.getCacheKey()) {
		return errors.New("清除缓存失败")
	}
	// todo 发送mq消息
	return nil
}

func (c categoryService) Get(ctx *gin.Context, req *entity.Category) error {
	obj, err := c.categoryMgr.FetchByPrimaryKey(int(req.CategoryID))
	if err != nil {
		return err
	}
	if err = copier.Copy(req, obj); err != nil {
		return err
	}
	return nil
}

func (c categoryService) getCacheKey() string {
	return consts.GoodsCat + "ALL"
}

func NewCategoryService() CategoryService {
	return &categoryService{
		categorySpecMgr: mgr.NewCategorySpecMgr(dbLocal.GRpo),
		categoryMgr:     mgr.NewCategoryMgr(dbLocal.GRpo),
	}
}

func (c categoryService) i() {}
