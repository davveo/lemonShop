package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsOrderMetaMgr struct {
	*_BaseMgr
}

// EsOrderMetaMgr open func
func EsOrderMetaMgr(db *gorm.DB) *_EsOrderMetaMgr {
	if db == nil {
		panic(fmt.Errorf("EsOrderMetaMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsOrderMetaMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_order_meta"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsOrderMetaMgr) GetTableName() string {
	return "es_order_meta"
}

// Reset 重置gorm会话
func (obj *_EsOrderMetaMgr) Reset() *_EsOrderMetaMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsOrderMetaMgr) Get() (result models.EsOrderMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsOrderMeta{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsOrderMetaMgr) Gets() (results []*models.EsOrderMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsOrderMeta{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsOrderMetaMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsOrderMeta{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithMetaID meta_id获取 主键ID
func (obj *_EsOrderMetaMgr) WithMetaID(metaID int) Option {
	return optionFunc(func(o *options) { o.query["meta_id"] = metaID })
}

// WithOrderSn order_sn获取 订单编号
func (obj *_EsOrderMetaMgr) WithOrderSn(orderSn string) Option {
	return optionFunc(func(o *options) { o.query["order_sn"] = orderSn })
}

// WithMetaKey meta_key获取 扩展-键
func (obj *_EsOrderMetaMgr) WithMetaKey(metaKey string) Option {
	return optionFunc(func(o *options) { o.query["meta_key"] = metaKey })
}

// WithMetaValue meta_value获取 扩展-值
func (obj *_EsOrderMetaMgr) WithMetaValue(metaValue string) Option {
	return optionFunc(func(o *options) { o.query["meta_value"] = metaValue })
}

// WithStatus status获取 售后状态
func (obj *_EsOrderMetaMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_EsOrderMetaMgr) GetByOption(opts ...Option) (result models.EsOrderMeta, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsOrderMeta{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsOrderMetaMgr) GetByOptions(opts ...Option) (results []*models.EsOrderMeta, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsOrderMeta{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsOrderMetaMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsOrderMeta, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsOrderMeta{}).Where(options.query)
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

// GetFromMetaID 通过meta_id获取内容 主键ID
func (obj *_EsOrderMetaMgr) GetFromMetaID(metaID int) (result models.EsOrderMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsOrderMeta{}).Where("`meta_id` = ?", metaID).First(&result).Error

	return
}

// GetBatchFromMetaID 批量查找 主键ID
func (obj *_EsOrderMetaMgr) GetBatchFromMetaID(metaIDs []int) (results []*models.EsOrderMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsOrderMeta{}).Where("`meta_id` IN (?)", metaIDs).Find(&results).Error

	return
}

// GetFromOrderSn 通过order_sn获取内容 订单编号
func (obj *_EsOrderMetaMgr) GetFromOrderSn(orderSn string) (results []*models.EsOrderMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsOrderMeta{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}

// GetBatchFromOrderSn 批量查找 订单编号
func (obj *_EsOrderMetaMgr) GetBatchFromOrderSn(orderSns []string) (results []*models.EsOrderMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsOrderMeta{}).Where("`order_sn` IN (?)", orderSns).Find(&results).Error

	return
}

// GetFromMetaKey 通过meta_key获取内容 扩展-键
func (obj *_EsOrderMetaMgr) GetFromMetaKey(metaKey string) (results []*models.EsOrderMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsOrderMeta{}).Where("`meta_key` = ?", metaKey).Find(&results).Error

	return
}

// GetBatchFromMetaKey 批量查找 扩展-键
func (obj *_EsOrderMetaMgr) GetBatchFromMetaKey(metaKeys []string) (results []*models.EsOrderMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsOrderMeta{}).Where("`meta_key` IN (?)", metaKeys).Find(&results).Error

	return
}

// GetFromMetaValue 通过meta_value获取内容 扩展-值
func (obj *_EsOrderMetaMgr) GetFromMetaValue(metaValue string) (results []*models.EsOrderMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsOrderMeta{}).Where("`meta_value` = ?", metaValue).Find(&results).Error

	return
}

// GetBatchFromMetaValue 批量查找 扩展-值
func (obj *_EsOrderMetaMgr) GetBatchFromMetaValue(metaValues []string) (results []*models.EsOrderMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsOrderMeta{}).Where("`meta_value` IN (?)", metaValues).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 售后状态
func (obj *_EsOrderMetaMgr) GetFromStatus(status string) (results []*models.EsOrderMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsOrderMeta{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 售后状态
func (obj *_EsOrderMetaMgr) GetBatchFromStatus(statuss []string) (results []*models.EsOrderMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsOrderMeta{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsOrderMetaMgr) FetchByPrimaryKey(metaID int) (result models.EsOrderMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsOrderMeta{}).Where("`meta_id` = ?", metaID).First(&result).Error

	return
}

// FetchIndexByEsIndOrderexMetaid  获取多个内容
func (obj *_EsOrderMetaMgr) FetchIndexByEsIndOrderexMetaid(metaID int) (results []*models.EsOrderMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsOrderMeta{}).Where("`meta_id` = ?", metaID).Find(&results).Error

	return
}
