package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsShopLogiRelMgr struct {
	*_BaseMgr
}

// EsShopLogiRelMgr open func
func EsShopLogiRelMgr(db *gorm.DB) *_EsShopLogiRelMgr {
	if db == nil {
		panic(fmt.Errorf("EsShopLogiRelMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsShopLogiRelMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_shop_logi_rel"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsShopLogiRelMgr) GetTableName() string {
	return "es_shop_logi_rel"
}

// Reset 重置gorm会话
func (obj *_EsShopLogiRelMgr) Reset() *_EsShopLogiRelMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsShopLogiRelMgr) Get() (result models.EsShopLogiRel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopLogiRel{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsShopLogiRelMgr) Gets() (results []*models.EsShopLogiRel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopLogiRel{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsShopLogiRelMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsShopLogiRel{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithLogiID logi_id获取 物流公司ID
func (obj *_EsShopLogiRelMgr) WithLogiID(logiID int) Option {
	return optionFunc(func(o *options) { o.query["logi_id"] = logiID })
}

// WithShopID shop_id获取 店铺ID
func (obj *_EsShopLogiRelMgr) WithShopID(shopID int) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// GetByOption 功能选项模式获取
func (obj *_EsShopLogiRelMgr) GetByOption(opts ...Option) (result models.EsShopLogiRel, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopLogiRel{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsShopLogiRelMgr) GetByOptions(opts ...Option) (results []*models.EsShopLogiRel, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopLogiRel{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsShopLogiRelMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsShopLogiRel, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsShopLogiRel{}).Where(options.query)
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

// GetFromLogiID 通过logi_id获取内容 物流公司ID
func (obj *_EsShopLogiRelMgr) GetFromLogiID(logiID int) (results []*models.EsShopLogiRel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopLogiRel{}).Where("`logi_id` = ?", logiID).Find(&results).Error

	return
}

// GetBatchFromLogiID 批量查找 物流公司ID
func (obj *_EsShopLogiRelMgr) GetBatchFromLogiID(logiIDs []int) (results []*models.EsShopLogiRel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopLogiRel{}).Where("`logi_id` IN (?)", logiIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺ID
func (obj *_EsShopLogiRelMgr) GetFromShopID(shopID int) (results []*models.EsShopLogiRel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopLogiRel{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺ID
func (obj *_EsShopLogiRelMgr) GetBatchFromShopID(shopIDs []int) (results []*models.EsShopLogiRel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopLogiRel{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
