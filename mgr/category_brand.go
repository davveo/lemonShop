package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsCategoryBrandMgr struct {
	*_BaseMgr
}

// EsCategoryBrandMgr open func
func EsCategoryBrandMgr(db *gorm.DB) *_EsCategoryBrandMgr {
	if db == nil {
		panic(fmt.Errorf("EsCategoryBrandMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsCategoryBrandMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_category_brand"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsCategoryBrandMgr) GetTableName() string {
	return "es_category_brand"
}

// Reset 重置gorm会话
func (obj *_EsCategoryBrandMgr) Reset() *_EsCategoryBrandMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsCategoryBrandMgr) Get() (result models.EsCategoryBrand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategoryBrand{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsCategoryBrandMgr) Gets() (results []*models.EsCategoryBrand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategoryBrand{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsCategoryBrandMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsCategoryBrand{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithCategoryID category_id获取 分类id
func (obj *_EsCategoryBrandMgr) WithCategoryID(categoryID int) Option {
	return optionFunc(func(o *options) { o.query["category_id"] = categoryID })
}

// WithBrandID brand_id获取 品牌id
func (obj *_EsCategoryBrandMgr) WithBrandID(brandID int) Option {
	return optionFunc(func(o *options) { o.query["brand_id"] = brandID })
}

// GetByOption 功能选项模式获取
func (obj *_EsCategoryBrandMgr) GetByOption(opts ...Option) (result models.EsCategoryBrand, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategoryBrand{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsCategoryBrandMgr) GetByOptions(opts ...Option) (results []*models.EsCategoryBrand, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategoryBrand{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsCategoryBrandMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsCategoryBrand, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsCategoryBrand{}).Where(options.query)
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

// GetFromCategoryID 通过category_id获取内容 分类id
func (obj *_EsCategoryBrandMgr) GetFromCategoryID(categoryID int) (results []*models.EsCategoryBrand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategoryBrand{}).Where("`category_id` = ?", categoryID).Find(&results).Error

	return
}

// GetBatchFromCategoryID 批量查找 分类id
func (obj *_EsCategoryBrandMgr) GetBatchFromCategoryID(categoryIDs []int) (results []*models.EsCategoryBrand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategoryBrand{}).Where("`category_id` IN (?)", categoryIDs).Find(&results).Error

	return
}

// GetFromBrandID 通过brand_id获取内容 品牌id
func (obj *_EsCategoryBrandMgr) GetFromBrandID(brandID int) (results []*models.EsCategoryBrand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategoryBrand{}).Where("`brand_id` = ?", brandID).Find(&results).Error

	return
}

// GetBatchFromBrandID 批量查找 品牌id
func (obj *_EsCategoryBrandMgr) GetBatchFromBrandID(brandIDs []int) (results []*models.EsCategoryBrand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategoryBrand{}).Where("`brand_id` IN (?)", brandIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
