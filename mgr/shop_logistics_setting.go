package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type ShopLogisticsSettingMgr struct {
	*_BaseMgr
}

// NewShopLogisticsSettingMgr open func
func NewShopLogisticsSettingMgr(db db.Repo) *ShopLogisticsSettingMgr {
	if db == nil {
		panic(fmt.Errorf("NewShopLogisticsSettingMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &ShopLogisticsSettingMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_shop_logistics_setting"), wdb: db.GetDbW().Table("es_shop_logistics_setting"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *ShopLogisticsSettingMgr) GetTableName() string {
	return "es_shop_logistics_setting"
}

// Reset 重置gorm会话
func (obj *ShopLogisticsSettingMgr) Reset() *ShopLogisticsSettingMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *ShopLogisticsSettingMgr) Get() (result models.EsShopLogisticsSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopLogisticsSetting{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *ShopLogisticsSettingMgr) Gets() (results []*models.EsShopLogisticsSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopLogisticsSetting{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *ShopLogisticsSettingMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsShopLogisticsSetting{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *ShopLogisticsSettingMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithShopID shop_id获取
func (obj *ShopLogisticsSettingMgr) WithShopID(shopID int) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithLogisticsID logistics_id获取
func (obj *ShopLogisticsSettingMgr) WithLogisticsID(logisticsID int) Option {
	return optionFunc(func(o *options) { o.query["logistics_id"] = logisticsID })
}

// WithConfig config获取
func (obj *ShopLogisticsSettingMgr) WithConfig(config string) Option {
	return optionFunc(func(o *options) { o.query["config"] = config })
}

// GetByOption 功能选项模式获取
func (obj *ShopLogisticsSettingMgr) GetByOption(opts ...Option) (result models.EsShopLogisticsSetting, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopLogisticsSetting{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *ShopLogisticsSettingMgr) GetByOptions(opts ...Option) (results []*models.EsShopLogisticsSetting, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopLogisticsSetting{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *ShopLogisticsSettingMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsShopLogisticsSetting, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsShopLogisticsSetting{}).Where(options.query)
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

// GetFromID 通过id获取内容
func (obj *ShopLogisticsSettingMgr) GetFromID(id int) (result models.EsShopLogisticsSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopLogisticsSetting{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *ShopLogisticsSettingMgr) GetBatchFromID(ids []int) (results []*models.EsShopLogisticsSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopLogisticsSetting{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容
func (obj *ShopLogisticsSettingMgr) GetFromShopID(shopID int) (results []*models.EsShopLogisticsSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopLogisticsSetting{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找
func (obj *ShopLogisticsSettingMgr) GetBatchFromShopID(shopIDs []int) (results []*models.EsShopLogisticsSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopLogisticsSetting{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromLogisticsID 通过logistics_id获取内容
func (obj *ShopLogisticsSettingMgr) GetFromLogisticsID(logisticsID int) (results []*models.EsShopLogisticsSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopLogisticsSetting{}).Where("`logistics_id` = ?", logisticsID).Find(&results).Error

	return
}

// GetBatchFromLogisticsID 批量查找
func (obj *ShopLogisticsSettingMgr) GetBatchFromLogisticsID(logisticsIDs []int) (results []*models.EsShopLogisticsSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopLogisticsSetting{}).Where("`logistics_id` IN (?)", logisticsIDs).Find(&results).Error

	return
}

// GetFromConfig 通过config获取内容
func (obj *ShopLogisticsSettingMgr) GetFromConfig(config string) (results []*models.EsShopLogisticsSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopLogisticsSetting{}).Where("`config` = ?", config).Find(&results).Error

	return
}

// GetBatchFromConfig 批量查找
func (obj *ShopLogisticsSettingMgr) GetBatchFromConfig(configs []string) (results []*models.EsShopLogisticsSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopLogisticsSetting{}).Where("`config` IN (?)", configs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *ShopLogisticsSettingMgr) FetchByPrimaryKey(id int) (result models.EsShopLogisticsSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopLogisticsSetting{}).Where("`id` = ?", id).First(&result).Error

	return
}
