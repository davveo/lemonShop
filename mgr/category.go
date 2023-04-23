package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsCategoryMgr struct {
	*_BaseMgr
}

// EsCategoryMgr open func
func EsCategoryMgr(db *gorm.DB) *_EsCategoryMgr {
	if db == nil {
		panic(fmt.Errorf("EsCategoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsCategoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_category"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsCategoryMgr) GetTableName() string {
	return "es_category"
}

// Reset 重置gorm会话
func (obj *_EsCategoryMgr) Reset() *_EsCategoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsCategoryMgr) Get() (result models.EsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsCategoryMgr) Gets() (results []*models.EsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategory{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsCategoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsCategory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithCategoryID category_id获取 主键
func (obj *_EsCategoryMgr) WithCategoryID(categoryID int) Option {
	return optionFunc(func(o *options) { o.query["category_id"] = categoryID })
}

// WithName name获取 分类名称
func (obj *_EsCategoryMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithParentID parent_id获取 分类父id
func (obj *_EsCategoryMgr) WithParentID(parentID int) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithCategoryPath category_path获取 分类父子路径
func (obj *_EsCategoryMgr) WithCategoryPath(categoryPath string) Option {
	return optionFunc(func(o *options) { o.query["category_path"] = categoryPath })
}

// WithGoodsCount goods_count获取 该分类下商品数量
func (obj *_EsCategoryMgr) WithGoodsCount(goodsCount int) Option {
	return optionFunc(func(o *options) { o.query["goods_count"] = goodsCount })
}

// WithCategoryOrder category_order获取 分类排序
func (obj *_EsCategoryMgr) WithCategoryOrder(categoryOrder int) Option {
	return optionFunc(func(o *options) { o.query["category_order"] = categoryOrder })
}

// WithImage image获取 分类图标
func (obj *_EsCategoryMgr) WithImage(image string) Option {
	return optionFunc(func(o *options) { o.query["image"] = image })
}

// GetByOption 功能选项模式获取
func (obj *_EsCategoryMgr) GetByOption(opts ...Option) (result models.EsCategory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsCategoryMgr) GetByOptions(opts ...Option) (results []*models.EsCategory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsCategoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsCategory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsCategory{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromCategoryID 通过category_id获取内容 主键
func (obj *_EsCategoryMgr) GetFromCategoryID(categoryID int) (result models.EsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`category_id` = ?", categoryID).First(&result).Error

	return
}

// GetBatchFromCategoryID 批量查找 主键
func (obj *_EsCategoryMgr) GetBatchFromCategoryID(categoryIDs []int) (results []*models.EsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`category_id` IN (?)", categoryIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 分类名称
func (obj *_EsCategoryMgr) GetFromName(name string) (results []*models.EsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 分类名称
func (obj *_EsCategoryMgr) GetBatchFromName(names []string) (results []*models.EsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容 分类父id
func (obj *_EsCategoryMgr) GetFromParentID(parentID int) (results []*models.EsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量查找 分类父id
func (obj *_EsCategoryMgr) GetBatchFromParentID(parentIDs []int) (results []*models.EsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromCategoryPath 通过category_path获取内容 分类父子路径
func (obj *_EsCategoryMgr) GetFromCategoryPath(categoryPath string) (results []*models.EsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`category_path` = ?", categoryPath).Find(&results).Error

	return
}

// GetBatchFromCategoryPath 批量查找 分类父子路径
func (obj *_EsCategoryMgr) GetBatchFromCategoryPath(categoryPaths []string) (results []*models.EsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`category_path` IN (?)", categoryPaths).Find(&results).Error

	return
}

// GetFromGoodsCount 通过goods_count获取内容 该分类下商品数量
func (obj *_EsCategoryMgr) GetFromGoodsCount(goodsCount int) (results []*models.EsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`goods_count` = ?", goodsCount).Find(&results).Error

	return
}

// GetBatchFromGoodsCount 批量查找 该分类下商品数量
func (obj *_EsCategoryMgr) GetBatchFromGoodsCount(goodsCounts []int) (results []*models.EsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`goods_count` IN (?)", goodsCounts).Find(&results).Error

	return
}

// GetFromCategoryOrder 通过category_order获取内容 分类排序
func (obj *_EsCategoryMgr) GetFromCategoryOrder(categoryOrder int) (results []*models.EsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`category_order` = ?", categoryOrder).Find(&results).Error

	return
}

// GetBatchFromCategoryOrder 批量查找 分类排序
func (obj *_EsCategoryMgr) GetBatchFromCategoryOrder(categoryOrders []int) (results []*models.EsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`category_order` IN (?)", categoryOrders).Find(&results).Error

	return
}

// GetFromImage 通过image获取内容 分类图标
func (obj *_EsCategoryMgr) GetFromImage(image string) (results []*models.EsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`image` = ?", image).Find(&results).Error

	return
}

// GetBatchFromImage 批量查找 分类图标
func (obj *_EsCategoryMgr) GetBatchFromImage(images []string) (results []*models.EsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`image` IN (?)", images).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsCategoryMgr) FetchByPrimaryKey(categoryID int) (result models.EsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`category_id` = ?", categoryID).First(&result).Error

	return
}

// FetchIndexByIndGoodsCatParentid  获取多个内容
func (obj *_EsCategoryMgr) FetchIndexByIndGoodsCatParentid(parentID int) (results []*models.EsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}
