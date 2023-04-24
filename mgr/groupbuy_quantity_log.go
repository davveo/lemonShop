package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type GroupbuyQuantityLogMgr struct {
	*_BaseMgr
}

// NewGroupbuyQuantityLogMgr open func
func NewGroupbuyQuantityLogMgr(db db.Repo) *GroupbuyQuantityLogMgr {
	if db == nil {
		panic(fmt.Errorf("NewGroupbuyQuantityLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &GroupbuyQuantityLogMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_groupbuy_quantity_log"), wdb: db.GetDbW().Table("es_groupbuy_quantity_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *GroupbuyQuantityLogMgr) GetTableName() string {
	return "es_groupbuy_quantity_log"
}

// Reset 重置gorm会话
func (obj *GroupbuyQuantityLogMgr) Reset() *GroupbuyQuantityLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *GroupbuyQuantityLogMgr) Get() (result models.EsGroupbuyQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyQuantityLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *GroupbuyQuantityLogMgr) Gets() (results []*models.EsGroupbuyQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyQuantityLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *GroupbuyQuantityLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyQuantityLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithLogID log_id获取 日志id
func (obj *GroupbuyQuantityLogMgr) WithLogID(logID int) Option {
	return optionFunc(func(o *options) { o.query["log_id"] = logID })
}

// WithOrderSn order_sn获取 订单编号
func (obj *GroupbuyQuantityLogMgr) WithOrderSn(orderSn string) Option {
	return optionFunc(func(o *options) { o.query["order_sn"] = orderSn })
}

// WithGoodsID goods_id获取 商品ID
func (obj *GroupbuyQuantityLogMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithQuantity quantity获取 团购售空数量
func (obj *GroupbuyQuantityLogMgr) WithQuantity(quantity int) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

// WithOpTime op_time获取 操作时间
func (obj *GroupbuyQuantityLogMgr) WithOpTime(opTime int64) Option {
	return optionFunc(func(o *options) { o.query["op_time"] = opTime })
}

// WithLogType log_type获取 日志类型
func (obj *GroupbuyQuantityLogMgr) WithLogType(logType string) Option {
	return optionFunc(func(o *options) { o.query["log_type"] = logType })
}

// WithReason reason获取 操作原因
func (obj *GroupbuyQuantityLogMgr) WithReason(reason string) Option {
	return optionFunc(func(o *options) { o.query["reason"] = reason })
}

// WithGbID gb_id获取 团购商品id
func (obj *GroupbuyQuantityLogMgr) WithGbID(gbID int) Option {
	return optionFunc(func(o *options) { o.query["gb_id"] = gbID })
}

// GetByOption 功能选项模式获取
func (obj *GroupbuyQuantityLogMgr) GetByOption(opts ...Option) (result models.EsGroupbuyQuantityLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyQuantityLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *GroupbuyQuantityLogMgr) GetByOptions(opts ...Option) (results []*models.EsGroupbuyQuantityLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyQuantityLog{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *GroupbuyQuantityLogMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsGroupbuyQuantityLog, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyQuantityLog{}).Where(options.query)
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
func (obj *GroupbuyQuantityLogMgr) GetFromLogID(logID int) (result models.EsGroupbuyQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyQuantityLog{}).Where("`log_id` = ?", logID).First(&result).Error

	return
}

// GetBatchFromLogID 批量查找 日志id
func (obj *GroupbuyQuantityLogMgr) GetBatchFromLogID(logIDs []int) (results []*models.EsGroupbuyQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyQuantityLog{}).Where("`log_id` IN (?)", logIDs).Find(&results).Error

	return
}

// GetFromOrderSn 通过order_sn获取内容 订单编号
func (obj *GroupbuyQuantityLogMgr) GetFromOrderSn(orderSn string) (results []*models.EsGroupbuyQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyQuantityLog{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}

// GetBatchFromOrderSn 批量查找 订单编号
func (obj *GroupbuyQuantityLogMgr) GetBatchFromOrderSn(orderSns []string) (results []*models.EsGroupbuyQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyQuantityLog{}).Where("`order_sn` IN (?)", orderSns).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 商品ID
func (obj *GroupbuyQuantityLogMgr) GetFromGoodsID(goodsID int) (results []*models.EsGroupbuyQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyQuantityLog{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 商品ID
func (obj *GroupbuyQuantityLogMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsGroupbuyQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyQuantityLog{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromQuantity 通过quantity获取内容 团购售空数量
func (obj *GroupbuyQuantityLogMgr) GetFromQuantity(quantity int) (results []*models.EsGroupbuyQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyQuantityLog{}).Where("`quantity` = ?", quantity).Find(&results).Error

	return
}

// GetBatchFromQuantity 批量查找 团购售空数量
func (obj *GroupbuyQuantityLogMgr) GetBatchFromQuantity(quantitys []int) (results []*models.EsGroupbuyQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyQuantityLog{}).Where("`quantity` IN (?)", quantitys).Find(&results).Error

	return
}

// GetFromOpTime 通过op_time获取内容 操作时间
func (obj *GroupbuyQuantityLogMgr) GetFromOpTime(opTime int64) (results []*models.EsGroupbuyQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyQuantityLog{}).Where("`op_time` = ?", opTime).Find(&results).Error

	return
}

// GetBatchFromOpTime 批量查找 操作时间
func (obj *GroupbuyQuantityLogMgr) GetBatchFromOpTime(opTimes []int64) (results []*models.EsGroupbuyQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyQuantityLog{}).Where("`op_time` IN (?)", opTimes).Find(&results).Error

	return
}

// GetFromLogType 通过log_type获取内容 日志类型
func (obj *GroupbuyQuantityLogMgr) GetFromLogType(logType string) (results []*models.EsGroupbuyQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyQuantityLog{}).Where("`log_type` = ?", logType).Find(&results).Error

	return
}

// GetBatchFromLogType 批量查找 日志类型
func (obj *GroupbuyQuantityLogMgr) GetBatchFromLogType(logTypes []string) (results []*models.EsGroupbuyQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyQuantityLog{}).Where("`log_type` IN (?)", logTypes).Find(&results).Error

	return
}

// GetFromReason 通过reason获取内容 操作原因
func (obj *GroupbuyQuantityLogMgr) GetFromReason(reason string) (results []*models.EsGroupbuyQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyQuantityLog{}).Where("`reason` = ?", reason).Find(&results).Error

	return
}

// GetBatchFromReason 批量查找 操作原因
func (obj *GroupbuyQuantityLogMgr) GetBatchFromReason(reasons []string) (results []*models.EsGroupbuyQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyQuantityLog{}).Where("`reason` IN (?)", reasons).Find(&results).Error

	return
}

// GetFromGbID 通过gb_id获取内容 团购商品id
func (obj *GroupbuyQuantityLogMgr) GetFromGbID(gbID int) (results []*models.EsGroupbuyQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyQuantityLog{}).Where("`gb_id` = ?", gbID).Find(&results).Error

	return
}

// GetBatchFromGbID 批量查找 团购商品id
func (obj *GroupbuyQuantityLogMgr) GetBatchFromGbID(gbIDs []int) (results []*models.EsGroupbuyQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyQuantityLog{}).Where("`gb_id` IN (?)", gbIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *GroupbuyQuantityLogMgr) FetchByPrimaryKey(logID int) (result models.EsGroupbuyQuantityLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyQuantityLog{}).Where("`log_id` = ?", logID).First(&result).Error

	return
}
