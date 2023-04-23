package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsCategorySpecMgr struct {
	*_BaseMgr
}

// EsCategorySpecMgr open func
func EsCategorySpecMgr(db *gorm.DB) *_EsCategorySpecMgr {
	if db == nil {
		panic(fmt.Errorf("EsCategorySpecMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsCategorySpecMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_category_spec"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsCategorySpecMgr) GetTableName() string {
	return "es_category_spec"
}

// Reset 重置gorm会话
func (obj *_EsCategorySpecMgr) Reset() *_EsCategorySpecMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsCategorySpecMgr) Get() (result models.EsCategorySpec, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategorySpec{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsCategorySpecMgr) Gets() (results []*models.EsCategorySpec, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategorySpec{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsCategorySpecMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsCategorySpec{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithCategoryID category_id获取 分类id
func (obj *_EsCategorySpecMgr) WithCategoryID(categoryID int) Option {
	return optionFunc(func(o *options) { o.query["category_id"] = categoryID })
}

// WithSpecID spec_id获取 规格id
func (obj *_EsCategorySpecMgr) WithSpecID(specID int) Option {
	return optionFunc(func(o *options) { o.query["spec_id"] = specID })
}

// GetByOption 功能选项模式获取
func (obj *_EsCategorySpecMgr) GetByOption(opts ...Option) (result models.EsCategorySpec, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategorySpec{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsCategorySpecMgr) GetByOptions(opts ...Option) (results []*models.EsCategorySpec, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategorySpec{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsCategorySpecMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsCategorySpec, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsCategorySpec{}).Where(options.query)
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
func (obj *_EsCategorySpecMgr) GetFromCategoryID(categoryID int) (results []*models.EsCategorySpec, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategorySpec{}).Where("`category_id` = ?", categoryID).Find(&results).Error

	return
}

// GetBatchFromCategoryID 批量查找 分类id
func (obj *_EsCategorySpecMgr) GetBatchFromCategoryID(categoryIDs []int) (results []*models.EsCategorySpec, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategorySpec{}).Where("`category_id` IN (?)", categoryIDs).Find(&results).Error

	return
}

// GetFromSpecID 通过spec_id获取内容 规格id
func (obj *_EsCategorySpecMgr) GetFromSpecID(specID int) (results []*models.EsCategorySpec, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategorySpec{}).Where("`spec_id` = ?", specID).Find(&results).Error

	return
}

// GetBatchFromSpecID 批量查找 规格id
func (obj *_EsCategorySpecMgr) GetBatchFromSpecID(specIDs []int) (results []*models.EsCategorySpec, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCategorySpec{}).Where("`spec_id` IN (?)", specIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
