package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type QuantityLogMgr struct {
	*_BaseMgr
}

// NewQuantityLogMgr open func
func NewQuantityLogMgr(db db.Repo) *QuantityLogMgr {
	if db == nil {
		panic(fmt.Errorf("NewQuantityLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &QuantityLogMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_quantity_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *QuantityLogMgr) GetTableName() string {
	return "es_quantity_log"
}

// Reset 重置gorm会话
func (obj *QuantityLogMgr) Reset() *QuantityLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *QuantityLogMgr) Get() (result models.EsQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsQuantityLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *QuantityLogMgr) Gets() (results []*models.EsQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsQuantityLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *QuantityLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsQuantityLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithLogID log_id获取 日志id
func (obj *QuantityLogMgr) WithLogID(logID int) Option {
	return optionFunc(func(o *options) { o.query["log_id"] = logID })
}

// WithOrderSn order_sn获取 订单编号
func (obj *QuantityLogMgr) WithOrderSn(orderSn string) Option {
	return optionFunc(func(o *options) { o.query["order_sn"] = orderSn })
}

// WithGoodsID goods_id获取 商品id
func (obj *QuantityLogMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithSkuID sku_id获取 sku id
func (obj *QuantityLogMgr) WithSkuID(skuID int) Option {
	return optionFunc(func(o *options) { o.query["sku_id"] = skuID })
}

// WithQuantity quantity获取 库存数
func (obj *QuantityLogMgr) WithQuantity(quantity int) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

// WithEnableQuantity enable_quantity获取 可用库存
func (obj *QuantityLogMgr) WithEnableQuantity(enableQuantity int) Option {
	return optionFunc(func(o *options) { o.query["enable_quantity"] = enableQuantity })
}

// WithOpTime op_time获取 操作时间
func (obj *QuantityLogMgr) WithOpTime(opTime int) Option {
	return optionFunc(func(o *options) { o.query["op_time"] = opTime })
}

// WithLogType log_type获取 日志类型
func (obj *QuantityLogMgr) WithLogType(logType string) Option {
	return optionFunc(func(o *options) { o.query["log_type"] = logType })
}

// WithReason reason获取 原因
func (obj *QuantityLogMgr) WithReason(reason string) Option {
	return optionFunc(func(o *options) { o.query["reason"] = reason })
}

// GetByOption 功能选项模式获取
func (obj *QuantityLogMgr) GetByOption(opts ...Option) (result models.EsQuantityLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsQuantityLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *QuantityLogMgr) GetByOptions(opts ...Option) (results []*models.EsQuantityLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsQuantityLog{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *QuantityLogMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsQuantityLog, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsQuantityLog{}).Where(options.query)
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

// GetFromLogID 通过log_id获取内容 日志id
func (obj *QuantityLogMgr) GetFromLogID(logID int) (result models.EsQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsQuantityLog{}).Where("`log_id` = ?", logID).First(&result).Error

	return
}

// GetBatchFromLogID 批量查找 日志id
func (obj *QuantityLogMgr) GetBatchFromLogID(logIDs []int) (results []*models.EsQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsQuantityLog{}).Where("`log_id` IN (?)", logIDs).Find(&results).Error

	return
}

// GetFromOrderSn 通过order_sn获取内容 订单编号
func (obj *QuantityLogMgr) GetFromOrderSn(orderSn string) (results []*models.EsQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsQuantityLog{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}

// GetBatchFromOrderSn 批量查找 订单编号
func (obj *QuantityLogMgr) GetBatchFromOrderSn(orderSns []string) (results []*models.EsQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsQuantityLog{}).Where("`order_sn` IN (?)", orderSns).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 商品id
func (obj *QuantityLogMgr) GetFromGoodsID(goodsID int) (results []*models.EsQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsQuantityLog{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 商品id
func (obj *QuantityLogMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsQuantityLog{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromSkuID 通过sku_id获取内容 sku id
func (obj *QuantityLogMgr) GetFromSkuID(skuID int) (results []*models.EsQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsQuantityLog{}).Where("`sku_id` = ?", skuID).Find(&results).Error

	return
}

// GetBatchFromSkuID 批量查找 sku id
func (obj *QuantityLogMgr) GetBatchFromSkuID(skuIDs []int) (results []*models.EsQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsQuantityLog{}).Where("`sku_id` IN (?)", skuIDs).Find(&results).Error

	return
}

// GetFromQuantity 通过quantity获取内容 库存数
func (obj *QuantityLogMgr) GetFromQuantity(quantity int) (results []*models.EsQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsQuantityLog{}).Where("`quantity` = ?", quantity).Find(&results).Error

	return
}

// GetBatchFromQuantity 批量查找 库存数
func (obj *QuantityLogMgr) GetBatchFromQuantity(quantitys []int) (results []*models.EsQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsQuantityLog{}).Where("`quantity` IN (?)", quantitys).Find(&results).Error

	return
}

// GetFromEnableQuantity 通过enable_quantity获取内容 可用库存
func (obj *QuantityLogMgr) GetFromEnableQuantity(enableQuantity int) (results []*models.EsQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsQuantityLog{}).Where("`enable_quantity` = ?", enableQuantity).Find(&results).Error

	return
}

// GetBatchFromEnableQuantity 批量查找 可用库存
func (obj *QuantityLogMgr) GetBatchFromEnableQuantity(enableQuantitys []int) (results []*models.EsQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsQuantityLog{}).Where("`enable_quantity` IN (?)", enableQuantitys).Find(&results).Error

	return
}

// GetFromOpTime 通过op_time获取内容 操作时间
func (obj *QuantityLogMgr) GetFromOpTime(opTime int) (results []*models.EsQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsQuantityLog{}).Where("`op_time` = ?", opTime).Find(&results).Error

	return
}

// GetBatchFromOpTime 批量查找 操作时间
func (obj *QuantityLogMgr) GetBatchFromOpTime(opTimes []int) (results []*models.EsQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsQuantityLog{}).Where("`op_time` IN (?)", opTimes).Find(&results).Error

	return
}

// GetFromLogType 通过log_type获取内容 日志类型
func (obj *QuantityLogMgr) GetFromLogType(logType string) (results []*models.EsQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsQuantityLog{}).Where("`log_type` = ?", logType).Find(&results).Error

	return
}

// GetBatchFromLogType 批量查找 日志类型
func (obj *QuantityLogMgr) GetBatchFromLogType(logTypes []string) (results []*models.EsQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsQuantityLog{}).Where("`log_type` IN (?)", logTypes).Find(&results).Error

	return
}

// GetFromReason 通过reason获取内容 原因
func (obj *QuantityLogMgr) GetFromReason(reason string) (results []*models.EsQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsQuantityLog{}).Where("`reason` = ?", reason).Find(&results).Error

	return
}

// GetBatchFromReason 批量查找 原因
func (obj *QuantityLogMgr) GetBatchFromReason(reasons []string) (results []*models.EsQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsQuantityLog{}).Where("`reason` IN (?)", reasons).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *QuantityLogMgr) FetchByPrimaryKey(logID int) (result models.EsQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsQuantityLog{}).Where("`log_id` = ?", logID).First(&result).Error

	return
}

// FetchIndexByIndQuantityLogGoodsid  获取多个内容
func (obj *QuantityLogMgr) FetchIndexByIndQuantityLogGoodsid(goodsID int, skuID int) (results []*models.EsQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsQuantityLog{}).Where("`goods_id` = ? AND `sku_id` = ?", goodsID, skuID).Find(&results).Error

	return
}
