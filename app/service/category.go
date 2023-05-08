package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/davveo/lemonShop/models/vo"

	"github.com/davveo/lemonShop/app/event_handler"

	"github.com/davveo/lemonShop/pkg/event"

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
	Update(ctx *gin.Context, req *entity.Category) error
	List(parentId int64, format string) (list interface{}, err error)
	ListAllChildren(parentId int) ([]*vo.CategoryVo, error)
	Get(ctx *gin.Context, req *entity.Category) error
}

type categoryService struct {
	categoryMgr      *mgr.CategoryMgr
	categorySpecMgr  *mgr.CategorySpecMgr
	categoryBrandMgr *mgr.CategoryBrandMgr
	eventMgr         event.EventManager
}

func (c *categoryService) Create(ctx *gin.Context, req *entity.Category) error {
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

	// 发送mq消息 todo 补充消费者逻辑
	c.eventMgr.Bind(consts.GoodsCategoryChange, &event_handler.CategoryEventHandler{}).
		Trigger(consts.GoodsCategoryChange, map[string]interface{}{
			"categoryId":    category.CategoryID,
			"operationType": consts.CategoryAddOperation,
		})
	return nil
}

func (c *categoryService) Get(ctx *gin.Context, req *entity.Category) error {
	obj, err := c.categoryMgr.FetchByPrimaryKey(int(req.CategoryID))
	if err != nil {
		return err
	}
	if err = copier.Copy(req, obj); err != nil {
		return err
	}
	return nil
}

func (c *categoryService) Update(ctx *gin.Context, req *entity.Category) error {
	category, err := c.categoryMgr.FetchByPrimaryKey(int(req.CategoryID))
	if err != nil {
		return err
	}

	//不能添加重复的分类名称
	res, err := c.categoryMgr.RawToStruct(consts.SqlFive, req.Name, req.CategoryID)
	if err != nil {
		return err
	}
	if len(res) > 0 {
		return errors.New("该分类名称已存在")
	}

	// 如果有子分类则不能更换上级分类
	// 更换上级分类
	if req.ParentID != category.ParentID {
		// 查看是否有子分类
		lst, err := c.List(int64(req.CategoryID), "")
		if err != nil {
			return err
		}
		voList, ok := lst.([]vo.CategoryVo)
		if !ok {
			return errors.New("分类异常")
		}

		if len(voList) > 0 {
			return errors.New("当前分类有子分类，不能更换上级分类")
		}

		parent := &entity.Category{
			CategoryID: req.CategoryID,
		}
		err = c.Get(ctx, parent)
		if err != nil {
			return err
		}

		strCategory := strings.Split(parent.CategoryPath, "|")
		// 如果当前的catPath length 大于4 证明当前分类级别大于五级 提示
		if len(strCategory) >= 4 {
			return errors.New("最多为三级分类,添加失败")
		}
		req.CategoryPath = parent.CategoryPath + string(req.CategoryID) + "|"
	}
	updates := new(models.EsCategory)
	if err := copier.Copy(updates, req); err != nil {
		return err
	}
	if err := c.categoryMgr.UpdateByModel(category,
		c.categoryMgr.WithCategoryID(req.CategoryID)); err != nil {
		return err
	}
	// 清楚缓存
	cache.Cache.Del(c.getCacheKey())

	// 发送mq消息 todo 补充消费者逻辑
	c.eventMgr.Bind(consts.GoodsCategoryChange, &event_handler.CategoryEventHandler{}).
		Trigger(consts.GoodsCategoryChange, map[string]interface{}{
			"categoryId":    category.CategoryID,
			"operationType": consts.CategoryUpdateOperation,
		})
	return nil
}

func (c *categoryService) Delete(ctx *gin.Context, req *entity.Category) error {
	var count int64
	list, err := c.ListAllChildren(req.CategoryID)
	if err != nil {
		return err
	}

	if len(list) > 0 {
		return errors.New("此类别下存在子类别不能删除")
	}

	sql := "select count(0) from es_goods where category_id = ? and disabled != -1"
	if err := c.categoryMgr.CountByCondition(sql, &count, req.CategoryID); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("此类别下存在商品不能删除")
	}

	cache.Cache.Del(fmt.Sprintf("%s%d", consts.GoodsCat, req.CategoryID))
	cache.Cache.Del(c.getCacheKey())

	c.categoryMgr.Delete(c.categoryMgr.WithCategoryID(req.CategoryID))
	c.categoryMgr.Delete(c.categorySpecMgr.WithCategoryID(req.CategoryID))
	c.categoryMgr.Delete(c.categoryBrandMgr.WithCategoryID(req.CategoryID))

	return nil
}

func (c *categoryService) List(parentId int64, format string) (list interface{}, err error) {
	if parentId == 0 {
		return
	}
	sql := "select c.* from es_category c  where c.parent_id = ? order by c.category_order asc"
	if format != "" {
		var categoryPluginVoList []*vo.CategoryPluginVo
		categoryList, err := c.categoryMgr.RawToStruct(sql, parentId)
		if err != nil {
			return categoryPluginVoList, err
		}

		for _, category := range categoryList {
			var cpv *vo.CategoryPluginVo
			err := copier.Copy(cpv, category)
			if err != nil {
				return categoryPluginVoList, err
			}
			cpv.Id = category.CategoryID
			cpv.Text = category.Name
			categoryPluginVoList = append(categoryPluginVoList, cpv)
		}
		return categoryPluginVoList, nil
	} else {
		var catList []vo.CategoryVo
		res, toStructErr := c.categoryMgr.RawToStruct(sql, parentId)
		if toStructErr != nil {
			return catList, toStructErr
		}
		for _, category := range res {
			categoryVoTmp := new(vo.CategoryVo)
			if copyErr := copier.Copy(categoryVoTmp, category); copyErr != nil {
				return catList, copyErr
			}
			children, listErr := c.listChildren(category.CategoryID)
			if listErr != nil {
				return catList, listErr
			}
			if len(children) > 0 {
				categoryVoTmp.Children = children
			}
			catList = append(catList, *categoryVoTmp)
		}

		return catList, nil
	}
}

func (c *categoryService) ListAllChildren(parentId int) ([]*vo.CategoryVo, error) {
	var topCatList []*vo.CategoryVo
	var categoryList []*models.EsCategory
	// 从缓存取所有的分类
	l, err := cache.Cache.Get(c.getCacheKey())
	if err != nil {
		return topCatList, err
	}

	err = json.Unmarshal([]byte(l), &categoryList)
	if err != nil {
		return topCatList, err
	}
	if len(categoryList) <= 0 {
		categoryList, err = c.initCategory()
		if err != nil {
			return topCatList, err
		}
	}

	for _, cat := range categoryList {
		var categoryVoTmp *vo.CategoryVo
		err := copier.Copy(cat, categoryVoTmp)
		if err != nil {
			return topCatList, err
		}
		if cat.ParentID == parentId {
			categoryVoTmp.Children = c.getChildren(categoryList, cat.CategoryID)
			topCatList = append(topCatList, categoryVoTmp)
		}
	}

	return topCatList, nil
}

func (c *categoryService) getSellerCategory(ctx *gin.Context, req *entity.Category) (categoryList []*models.EsCategory, err error) {
	// todo 完成用户session
	//context := entity.UserContext{}
	//seller, err := context.GetSeller()
	//if err != nil {
	//	return nil, err
	//}
	//
	return nil, nil
}

func (c *categoryService) listChildren(parentId int) ([]*vo.CategoryVo, error) {
	var categoryList []*models.EsCategory
	var topCatList []*vo.CategoryVo
	// 从缓存取所有的分类
	l, err := cache.Cache.Get(c.getCacheKey())
	if err != nil {
		return topCatList, err
	}

	err = json.Unmarshal([]byte(l), &categoryList)
	if err != nil {
		return topCatList, err
	}

	if len(categoryList) <= 0 {
		categoryList, err = c.initCategory()
		if err != nil {
			return topCatList, err
		}
	}

	for _, cat := range categoryList {
		var categoryVoTmp *vo.CategoryVo
		copier.Copy(cat, categoryVoTmp)
		if cat.ParentID == parentId {
			topCatList = append(topCatList, categoryVoTmp)
		}
	}

	return topCatList, nil
}

func (c *categoryService) getChildren(catList []*models.EsCategory, parentId int) []*vo.CategoryVo {
	var children []*vo.CategoryVo
	for _, cat := range catList {
		categoryVo := new(vo.CategoryVo)
		copier.Copy(categoryVo, cat)
		if cat.ParentID == parentId {
			categoryVo.Children = c.getChildren(catList, cat.ParentID)
			children = append(children, categoryVo)
		}
	}
	return children
}

func (c *categoryService) initCategory() (categoryList []*models.EsCategory, err error) {
	categoryList, err = c.getCategoryList()
	if err != nil {
		return
	}
	if len(categoryList) > 0 {
		for _, category := range categoryList {
			k := fmt.Sprintf("%s%d", consts.GoodsCat, category.CategoryID)
			if str, marshalErr := json.Marshal(category); marshalErr != nil {
				return
			} else {
				cache.Cache.Set(k, string(str), 0)
			}
		}
		if m, marshalErr := json.Marshal(categoryList); marshalErr != nil {
			return
		} else {
			cache.Cache.Set(c.getCacheKey(), string(m), 0)
		}

	}
	return
}

func (c *categoryService) getCategoryList() (res []*models.EsCategory, err error) {
	sql := "select * from es_category order by category_order asc"
	res, err = c.categoryMgr.RawToStruct(sql)
	if err != nil {
		return
	}
	return
}

func (c *categoryService) getCacheKey() string {
	return consts.GoodsCat + "ALL"
}

func NewCategoryService() CategoryService {
	return &categoryService{
		categorySpecMgr:  mgr.NewCategorySpecMgr(dbLocal.GRpo),
		categoryBrandMgr: mgr.NewCategoryBrandMgr(dbLocal.GRpo),
		categoryMgr:      mgr.NewCategoryMgr(dbLocal.GRpo),
	}
}

func (c *categoryService) i() {}
