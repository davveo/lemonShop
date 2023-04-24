package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type OrderLogMgr struct {
	*_BaseMgr
}

// NewOrderLogMgr open func
func NewOrderLogMgr(db db.Repo) *OrderLogMgr {
	if db == nil {
		panic(fmt.Errorf("NewOrderLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &OrderLogMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_order_log"), wdb: db.GetDbW().Table("es_order_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *OrderLogMgr) GetTableName() string {
	return "es_order_log"
}

// Reset 重置gorm会话
func (obj *OrderLogMgr) Reset() *OrderLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *OrderLogMgr) Get() (result models.EsOrderLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *OrderLogMgr) Gets() (results []*models.EsOrderLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *OrderLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsOrderLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithLogID log_id获取 主键ID
func (obj *OrderLogMgr) WithLogID(logID int) Option {
	return optionFunc(func(o *options) { o.query["log_id"] = logID })
}

// WithOrderSn order_sn获取 订单编号
func (obj *OrderLogMgr) WithOrderSn(orderSn string) Option {
	return optionFunc(func(o *options) { o.query["order_sn"] = orderSn })
}

// WithOpID op_id获取 操作者id
func (obj *OrderLogMgr) WithOpID(opID int) Option {
	return optionFunc(func(o *options) { o.query["op_id"] = opID })
}

// WithOpName op_name获取 操作者名称
func (obj *OrderLogMgr) WithOpName(opName string) Option {
	return optionFunc(func(o *options) { o.query["op_name"] = opName })
}

// WithMessage message获取 日志信息
func (obj *OrderLogMgr) WithMessage(message string) Option {
	return optionFunc(func(o *options) { o.query["message"] = message })
}

// WithOpTime op_time获取 操作时间
func (obj *OrderLogMgr) WithOpTime(opTime int64) Option {
	return optionFunc(func(o *options) { o.query["op_time"] = opTime })
}

// GetByOption 功能选项模式获取
func (obj *OrderLogMgr) GetByOption(opts ...Option) (result models.EsOrderLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *OrderLogMgr) GetByOptions(opts ...Option) (results []*models.EsOrderLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderLog{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *OrderLogMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsOrderLog, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsOrderLog{}).Where(options.query)
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

// GetFromLogID 通过log_id获取内容 主键ID
func (obj *OrderLogMgr) GetFromLogID(logID int) (result models.EsOrderLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderLog{}).Where("`log_id` = ?", logID).First(&result).Error

	return
}

// GetBatchFromLogID 批量查找 主键ID
func (obj *OrderLogMgr) GetBatchFromLogID(logIDs []int) (results []*models.EsOrderLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderLog{}).Where("`log_id` IN (?)", logIDs).Find(&results).Error

	return
}

// GetFromOrderSn 通过order_sn获取内容 订单编号
func (obj *OrderLogMgr) GetFromOrderSn(orderSn string) (results []*models.EsOrderLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderLog{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}

// GetBatchFromOrderSn 批量查找 订单编号
func (obj *OrderLogMgr) GetBatchFromOrderSn(orderSns []string) (results []*models.EsOrderLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderLog{}).Where("`order_sn` IN (?)", orderSns).Find(&results).Error

	return
}

// GetFromOpID 通过op_id获取内容 操作者id
func (obj *OrderLogMgr) GetFromOpID(opID int) (results []*models.EsOrderLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderLog{}).Where("`op_id` = ?", opID).Find(&results).Error

	return
}

// GetBatchFromOpID 批量查找 操作者id
func (obj *OrderLogMgr) GetBatchFromOpID(opIDs []int) (results []*models.EsOrderLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderLog{}).Where("`op_id` IN (?)", opIDs).Find(&results).Error

	return
}

// GetFromOpName 通过op_name获取内容 操作者名称
func (obj *OrderLogMgr) GetFromOpName(opName string) (results []*models.EsOrderLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderLog{}).Where("`op_name` = ?", opName).Find(&results).Error

	return
}

// GetBatchFromOpName 批量查找 操作者名称
func (obj *OrderLogMgr) GetBatchFromOpName(opNames []string) (results []*models.EsOrderLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderLog{}).Where("`op_name` IN (?)", opNames).Find(&results).Error

	return
}

// GetFromMessage 通过message获取内容 日志信息
func (obj *OrderLogMgr) GetFromMessage(message string) (results []*models.EsOrderLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderLog{}).Where("`message` = ?", message).Find(&results).Error

	return
}

// GetBatchFromMessage 批量查找 日志信息
func (obj *OrderLogMgr) GetBatchFromMessage(messages []string) (results []*models.EsOrderLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderLog{}).Where("`message` IN (?)", messages).Find(&results).Error

	return
}

// GetFromOpTime 通过op_time获取内容 操作时间
func (obj *OrderLogMgr) GetFromOpTime(opTime int64) (results []*models.EsOrderLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderLog{}).Where("`op_time` = ?", opTime).Find(&results).Error

	return
}

// GetBatchFromOpTime 批量查找 操作时间
func (obj *OrderLogMgr) GetBatchFromOpTime(opTimes []int64) (results []*models.EsOrderLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderLog{}).Where("`op_time` IN (?)", opTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *OrderLogMgr) FetchByPrimaryKey(logID int) (result models.EsOrderLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderLog{}).Where("`log_id` = ?", logID).First(&result).Error

	return
}

// FetchIndexByIndOrderLog  获取多个内容
func (obj *OrderLogMgr) FetchIndexByIndOrderLog(orderSn string) (results []*models.EsOrderLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderLog{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}
