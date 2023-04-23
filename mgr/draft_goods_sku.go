package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsDraftGoodsSkuMgr struct {
	*_BaseMgr
}

// EsDraftGoodsSkuMgr open func
func EsDraftGoodsSkuMgr(db *gorm.DB) *_EsDraftGoodsSkuMgr {
	if db == nil {
		panic(fmt.Errorf("EsDraftGoodsSkuMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsDraftGoodsSkuMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_draft_goods_sku"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsDraftGoodsSkuMgr) GetTableName() string {
	return "es_draft_goods_sku"
}

// Reset 重置gorm会话
func (obj *_EsDraftGoodsSkuMgr) Reset() *_EsDraftGoodsSkuMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsDraftGoodsSkuMgr) Get() (result models.EsDraftGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsSku{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsDraftGoodsSkuMgr) Gets() (results []*models.EsDraftGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsSku{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsDraftGoodsSkuMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsSku{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithDraftSkuID draft_sku_id获取 主键ID
func (obj *_EsDraftGoodsSkuMgr) WithDraftSkuID(draftSkuID int) Option {
	return optionFunc(func(o *options) { o.query["draft_sku_id"] = draftSkuID })
}

// WithDraftGoodsID draft_goods_id获取 草稿id
func (obj *_EsDraftGoodsSkuMgr) WithDraftGoodsID(draftGoodsID int) Option {
	return optionFunc(func(o *options) { o.query["draft_goods_id"] = draftGoodsID })
}

// WithSn sn获取 货号
func (obj *_EsDraftGoodsSkuMgr) WithSn(sn string) Option {
	return optionFunc(func(o *options) { o.query["sn"] = sn })
}

// WithQuantity quantity获取 总库存
func (obj *_EsDraftGoodsSkuMgr) WithQuantity(quantity int) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

// WithPrice price获取 价格
func (obj *_EsDraftGoodsSkuMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithSpecs specs获取 规格
func (obj *_EsDraftGoodsSkuMgr) WithSpecs(specs string) Option {
	return optionFunc(func(o *options) { o.query["specs"] = specs })
}

// WithCost cost获取 成本
func (obj *_EsDraftGoodsSkuMgr) WithCost(cost float64) Option {
	return optionFunc(func(o *options) { o.query["cost"] = cost })
}

// WithWeight weight获取 重量
func (obj *_EsDraftGoodsSkuMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// GetByOption 功能选项模式获取
func (obj *_EsDraftGoodsSkuMgr) GetByOption(opts ...Option) (result models.EsDraftGoodsSku, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsSku{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsDraftGoodsSkuMgr) GetByOptions(opts ...Option) (results []*models.EsDraftGoodsSku, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsSku{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsDraftGoodsSkuMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsDraftGoodsSku, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsSku{}).Where(options.query)
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

// GetFromDraftSkuID 通过draft_sku_id获取内容 主键ID
func (obj *_EsDraftGoodsSkuMgr) GetFromDraftSkuID(draftSkuID int) (result models.EsDraftGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsSku{}).Where("`draft_sku_id` = ?", draftSkuID).First(&result).Error

	return
}

// GetBatchFromDraftSkuID 批量查找 主键ID
func (obj *_EsDraftGoodsSkuMgr) GetBatchFromDraftSkuID(draftSkuIDs []int) (results []*models.EsDraftGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsSku{}).Where("`draft_sku_id` IN (?)", draftSkuIDs).Find(&results).Error

	return
}

// GetFromDraftGoodsID 通过draft_goods_id获取内容 草稿id
func (obj *_EsDraftGoodsSkuMgr) GetFromDraftGoodsID(draftGoodsID int) (results []*models.EsDraftGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsSku{}).Where("`draft_goods_id` = ?", draftGoodsID).Find(&results).Error

	return
}

// GetBatchFromDraftGoodsID 批量查找 草稿id
func (obj *_EsDraftGoodsSkuMgr) GetBatchFromDraftGoodsID(draftGoodsIDs []int) (results []*models.EsDraftGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsSku{}).Where("`draft_goods_id` IN (?)", draftGoodsIDs).Find(&results).Error

	return
}

// GetFromSn 通过sn获取内容 货号
func (obj *_EsDraftGoodsSkuMgr) GetFromSn(sn string) (results []*models.EsDraftGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsSku{}).Where("`sn` = ?", sn).Find(&results).Error

	return
}

// GetBatchFromSn 批量查找 货号
func (obj *_EsDraftGoodsSkuMgr) GetBatchFromSn(sns []string) (results []*models.EsDraftGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsSku{}).Where("`sn` IN (?)", sns).Find(&results).Error

	return
}

// GetFromQuantity 通过quantity获取内容 总库存
func (obj *_EsDraftGoodsSkuMgr) GetFromQuantity(quantity int) (results []*models.EsDraftGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsSku{}).Where("`quantity` = ?", quantity).Find(&results).Error

	return
}

// GetBatchFromQuantity 批量查找 总库存
func (obj *_EsDraftGoodsSkuMgr) GetBatchFromQuantity(quantitys []int) (results []*models.EsDraftGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsSku{}).Where("`quantity` IN (?)", quantitys).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 价格
func (obj *_EsDraftGoodsSkuMgr) GetFromPrice(price float64) (results []*models.EsDraftGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsSku{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 价格
func (obj *_EsDraftGoodsSkuMgr) GetBatchFromPrice(prices []float64) (results []*models.EsDraftGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsSku{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromSpecs 通过specs获取内容 规格
func (obj *_EsDraftGoodsSkuMgr) GetFromSpecs(specs string) (results []*models.EsDraftGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsSku{}).Where("`specs` = ?", specs).Find(&results).Error

	return
}

// GetBatchFromSpecs 批量查找 规格
func (obj *_EsDraftGoodsSkuMgr) GetBatchFromSpecs(specss []string) (results []*models.EsDraftGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsSku{}).Where("`specs` IN (?)", specss).Find(&results).Error

	return
}

// GetFromCost 通过cost获取内容 成本
func (obj *_EsDraftGoodsSkuMgr) GetFromCost(cost float64) (results []*models.EsDraftGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsSku{}).Where("`cost` = ?", cost).Find(&results).Error

	return
}

// GetBatchFromCost 批量查找 成本
func (obj *_EsDraftGoodsSkuMgr) GetBatchFromCost(costs []float64) (results []*models.EsDraftGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsSku{}).Where("`cost` IN (?)", costs).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 重量
func (obj *_EsDraftGoodsSkuMgr) GetFromWeight(weight float64) (results []*models.EsDraftGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsSku{}).Where("`weight` = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量查找 重量
func (obj *_EsDraftGoodsSkuMgr) GetBatchFromWeight(weights []float64) (results []*models.EsDraftGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsSku{}).Where("`weight` IN (?)", weights).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsDraftGoodsSkuMgr) FetchByPrimaryKey(draftSkuID int) (result models.EsDraftGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsSku{}).Where("`draft_sku_id` = ?", draftSkuID).First(&result).Error

	return
}
