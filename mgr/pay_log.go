package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type PayLogMgr struct {
	*_BaseMgr
}

// NewPayLogMgr open func
func NewPayLogMgr(db db.Repo) *PayLogMgr {
	if db == nil {
		panic(fmt.Errorf("NewPayLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &PayLogMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_pay_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *PayLogMgr) GetTableName() string {
	return "es_pay_log"
}

// Reset 重置gorm会话
func (obj *PayLogMgr) Reset() *PayLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *PayLogMgr) Get() (result models.EsPayLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPayLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *PayLogMgr) Gets() (results []*models.EsPayLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPayLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *PayLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsPayLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithPayLogID pay_log_id获取 收款单ID
func (obj *PayLogMgr) WithPayLogID(payLogID int) Option {
	return optionFunc(func(o *options) { o.query["pay_log_id"] = payLogID })
}

// WithOrderSn order_sn获取 订单编号
func (obj *PayLogMgr) WithOrderSn(orderSn string) Option {
	return optionFunc(func(o *options) { o.query["order_sn"] = orderSn })
}

// WithPayWay pay_way获取 支付方式
func (obj *PayLogMgr) WithPayWay(payWay string) Option {
	return optionFunc(func(o *options) { o.query["pay_way"] = payWay })
}

// WithPayType pay_type获取 支付类型
func (obj *PayLogMgr) WithPayType(payType string) Option {
	return optionFunc(func(o *options) { o.query["pay_type"] = payType })
}

// WithPayTime pay_time获取 付款时间
func (obj *PayLogMgr) WithPayTime(payTime int64) Option {
	return optionFunc(func(o *options) { o.query["pay_time"] = payTime })
}

// WithPayMoney pay_money获取 付款金额
func (obj *PayLogMgr) WithPayMoney(payMoney float64) Option {
	return optionFunc(func(o *options) { o.query["pay_money"] = payMoney })
}

// WithPayMemberName pay_member_name获取 付款会员名
func (obj *PayLogMgr) WithPayMemberName(payMemberName string) Option {
	return optionFunc(func(o *options) { o.query["pay_member_name"] = payMemberName })
}

// WithPayStatus pay_status获取 付款状态
func (obj *PayLogMgr) WithPayStatus(payStatus string) Option {
	return optionFunc(func(o *options) { o.query["pay_status"] = payStatus })
}

// WithPayLogSn pay_log_sn获取 流水号
func (obj *PayLogMgr) WithPayLogSn(payLogSn string) Option {
	return optionFunc(func(o *options) { o.query["pay_log_sn"] = payLogSn })
}

// WithPayOrderNo pay_order_no获取 支付方式返回流水号
func (obj *PayLogMgr) WithPayOrderNo(payOrderNo string) Option {
	return optionFunc(func(o *options) { o.query["pay_order_no"] = payOrderNo })
}

// GetByOption 功能选项模式获取
func (obj *PayLogMgr) GetByOption(opts ...Option) (result models.EsPayLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPayLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *PayLogMgr) GetByOptions(opts ...Option) (results []*models.EsPayLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPayLog{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *PayLogMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsPayLog, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsPayLog{}).Where(options.query)
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

// GetFromPayLogID 通过pay_log_id获取内容 收款单ID
func (obj *PayLogMgr) GetFromPayLogID(payLogID int) (result models.EsPayLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPayLog{}).Where("`pay_log_id` = ?", payLogID).First(&result).Error

	return
}

// GetBatchFromPayLogID 批量查找 收款单ID
func (obj *PayLogMgr) GetBatchFromPayLogID(payLogIDs []int) (results []*models.EsPayLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPayLog{}).Where("`pay_log_id` IN (?)", payLogIDs).Find(&results).Error

	return
}

// GetFromOrderSn 通过order_sn获取内容 订单编号
func (obj *PayLogMgr) GetFromOrderSn(orderSn string) (results []*models.EsPayLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPayLog{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}

// GetBatchFromOrderSn 批量查找 订单编号
func (obj *PayLogMgr) GetBatchFromOrderSn(orderSns []string) (results []*models.EsPayLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPayLog{}).Where("`order_sn` IN (?)", orderSns).Find(&results).Error

	return
}

// GetFromPayWay 通过pay_way获取内容 支付方式
func (obj *PayLogMgr) GetFromPayWay(payWay string) (results []*models.EsPayLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPayLog{}).Where("`pay_way` = ?", payWay).Find(&results).Error

	return
}

// GetBatchFromPayWay 批量查找 支付方式
func (obj *PayLogMgr) GetBatchFromPayWay(payWays []string) (results []*models.EsPayLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPayLog{}).Where("`pay_way` IN (?)", payWays).Find(&results).Error

	return
}

// GetFromPayType 通过pay_type获取内容 支付类型
func (obj *PayLogMgr) GetFromPayType(payType string) (results []*models.EsPayLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPayLog{}).Where("`pay_type` = ?", payType).Find(&results).Error

	return
}

// GetBatchFromPayType 批量查找 支付类型
func (obj *PayLogMgr) GetBatchFromPayType(payTypes []string) (results []*models.EsPayLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPayLog{}).Where("`pay_type` IN (?)", payTypes).Find(&results).Error

	return
}

// GetFromPayTime 通过pay_time获取内容 付款时间
func (obj *PayLogMgr) GetFromPayTime(payTime int64) (results []*models.EsPayLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPayLog{}).Where("`pay_time` = ?", payTime).Find(&results).Error

	return
}

// GetBatchFromPayTime 批量查找 付款时间
func (obj *PayLogMgr) GetBatchFromPayTime(payTimes []int64) (results []*models.EsPayLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPayLog{}).Where("`pay_time` IN (?)", payTimes).Find(&results).Error

	return
}

// GetFromPayMoney 通过pay_money获取内容 付款金额
func (obj *PayLogMgr) GetFromPayMoney(payMoney float64) (results []*models.EsPayLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPayLog{}).Where("`pay_money` = ?", payMoney).Find(&results).Error

	return
}

// GetBatchFromPayMoney 批量查找 付款金额
func (obj *PayLogMgr) GetBatchFromPayMoney(payMoneys []float64) (results []*models.EsPayLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPayLog{}).Where("`pay_money` IN (?)", payMoneys).Find(&results).Error

	return
}

// GetFromPayMemberName 通过pay_member_name获取内容 付款会员名
func (obj *PayLogMgr) GetFromPayMemberName(payMemberName string) (results []*models.EsPayLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPayLog{}).Where("`pay_member_name` = ?", payMemberName).Find(&results).Error

	return
}

// GetBatchFromPayMemberName 批量查找 付款会员名
func (obj *PayLogMgr) GetBatchFromPayMemberName(payMemberNames []string) (results []*models.EsPayLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPayLog{}).Where("`pay_member_name` IN (?)", payMemberNames).Find(&results).Error

	return
}

// GetFromPayStatus 通过pay_status获取内容 付款状态
func (obj *PayLogMgr) GetFromPayStatus(payStatus string) (results []*models.EsPayLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPayLog{}).Where("`pay_status` = ?", payStatus).Find(&results).Error

	return
}

// GetBatchFromPayStatus 批量查找 付款状态
func (obj *PayLogMgr) GetBatchFromPayStatus(payStatuss []string) (results []*models.EsPayLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPayLog{}).Where("`pay_status` IN (?)", payStatuss).Find(&results).Error

	return
}

// GetFromPayLogSn 通过pay_log_sn获取内容 流水号
func (obj *PayLogMgr) GetFromPayLogSn(payLogSn string) (results []*models.EsPayLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPayLog{}).Where("`pay_log_sn` = ?", payLogSn).Find(&results).Error

	return
}

// GetBatchFromPayLogSn 批量查找 流水号
func (obj *PayLogMgr) GetBatchFromPayLogSn(payLogSns []string) (results []*models.EsPayLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPayLog{}).Where("`pay_log_sn` IN (?)", payLogSns).Find(&results).Error

	return
}

// GetFromPayOrderNo 通过pay_order_no获取内容 支付方式返回流水号
func (obj *PayLogMgr) GetFromPayOrderNo(payOrderNo string) (results []*models.EsPayLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPayLog{}).Where("`pay_order_no` = ?", payOrderNo).Find(&results).Error

	return
}

// GetBatchFromPayOrderNo 批量查找 支付方式返回流水号
func (obj *PayLogMgr) GetBatchFromPayOrderNo(payOrderNos []string) (results []*models.EsPayLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPayLog{}).Where("`pay_order_no` IN (?)", payOrderNos).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *PayLogMgr) FetchByPrimaryKey(payLogID int) (result models.EsPayLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPayLog{}).Where("`pay_log_id` = ?", payLogID).First(&result).Error

	return
}

// FetchIndexByIndPayLog  获取多个内容
func (obj *PayLogMgr) FetchIndexByIndPayLog(orderSn string, payStatus string) (results []*models.EsPayLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPayLog{}).Where("`order_sn` = ? AND `pay_status` = ?", orderSn, payStatus).Find(&results).Error

	return
}
