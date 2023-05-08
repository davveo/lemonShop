package mgr

import (
	"context"
	"fmt"

	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type CategorySpecMgr struct {
	*_BaseMgr
}

// NewCategorySpecMgr open func
func NewCategorySpecMgr(db db.Repo) *CategorySpecMgr {
	if db == nil {
		panic(fmt.Errorf("CategorySpecMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &CategorySpecMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_category_spec"),
		wdb:       db.GetDbW().Table("es_category_spec"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *CategorySpecMgr) GetTableName() string {
	return "es_category_spec"
}

func (obj *CategorySpecMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsCategorySpec{}).Count(count)
}

// WithCategoryID category_id获取 分类id
func (obj *CategorySpecMgr) WithCategoryID(categoryID int) Option {
	return optionFunc(func(o *options) { o.query["category_id"] = categoryID })
}

// WithSpecID spec_id获取 规格id
func (obj *CategorySpecMgr) WithSpecID(specID int) Option {
	return optionFunc(func(o *options) { o.query["spec_id"] = specID })
}

// GetByOption 功能选项模式获取
func (obj *CategorySpecMgr) GetByOption(opts ...Option) (result models.EsCategorySpec, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCategorySpec{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *CategorySpecMgr) GetByOptions(opts ...Option) (results []*models.EsCategorySpec, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCategorySpec{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *CategorySpecMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsCategorySpec, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsCategorySpec{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

// GetFromCategoryID 通过category_id获取内容 分类id
func (obj *CategorySpecMgr) GetFromCategoryID(categoryID int) (results []*models.EsCategorySpec, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCategorySpec{}).Where("`category_id` = ?", categoryID).Find(&results).Error

	return
}

// GetBatchFromCategoryID 批量查找 分类id
func (obj *CategorySpecMgr) GetBatchFromCategoryID(categoryIDs []int) (results []*models.EsCategorySpec, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCategorySpec{}).Where("`category_id` IN (?)", categoryIDs).Find(&results).Error

	return
}

// GetFromSpecID 通过spec_id获取内容 规格id
func (obj *CategorySpecMgr) GetFromSpecID(specID int) (results []*models.EsCategorySpec, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCategorySpec{}).Where("`spec_id` = ?", specID).Find(&results).Error

	return
}

// GetBatchFromSpecID 批量查找 规格id
func (obj *CategorySpecMgr) GetBatchFromSpecID(specIDs []int64) (results []*models.EsCategorySpec, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCategorySpec{}).Where("`spec_id` IN (?)", specIDs).Find(&results).Error

	return
}

// FetchByPrimaryKey 通过id查询
func (obj *CategorySpecMgr) FetchByPrimaryKey(categoryID int64) (result *models.EsCategorySpec, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCategorySpec{}).Where("`category_id` = ?", categoryID).First(&result).Error

	return
}

func (obj *CategorySpecMgr) Create(categorySpec *models.EsCategorySpec) (err error) {
	err = obj.wdb.WithContext(obj.ctx).Create(categorySpec).Error

	return
}

func (obj *CategorySpecMgr) DeleteByOptions(opts ...Option) (err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Where(options.query).Delete(&models.EsCategorySpec{}).Error

	return
}
