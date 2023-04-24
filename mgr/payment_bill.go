package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type PaymentBillMgr struct {
	*_BaseMgr
}

// NewPaymentBillMgr open func
func NewPaymentBillMgr(db db.Repo) *PaymentBillMgr {
	if db == nil {
		panic(fmt.Errorf("NewPaymentBillMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &PaymentBillMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_payment_bill"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *PaymentBillMgr) GetTableName() string {
	return "es_payment_bill"
}

// Reset 重置gorm会话
func (obj *PaymentBillMgr) Reset() *PaymentBillMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *PaymentBillMgr) Get() (result models.EsPaymentBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentBill{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *PaymentBillMgr) Gets() (results []*models.EsPaymentBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentBill{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *PaymentBillMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentBill{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithBillID bill_id获取 主键id
func (obj *PaymentBillMgr) WithBillID(billID int) Option {
	return optionFunc(func(o *options) { o.query["bill_id"] = billID })
}

// WithSn sn获取 单号
func (obj *PaymentBillMgr) WithSn(sn string) Option {
	return optionFunc(func(o *options) { o.query["sn"] = sn })
}

// WithOutTradeNo out_trade_no获取 提交给第三方平台单号
func (obj *PaymentBillMgr) WithOutTradeNo(outTradeNo string) Option {
	return optionFunc(func(o *options) { o.query["out_trade_no"] = outTradeNo })
}

// WithReturnTradeNo return_trade_no获取 第三方平台返回交易号
func (obj *PaymentBillMgr) WithReturnTradeNo(returnTradeNo string) Option {
	return optionFunc(func(o *options) { o.query["return_trade_no"] = returnTradeNo })
}

// WithIsPay is_pay获取 是否已支付
func (obj *PaymentBillMgr) WithIsPay(isPay int) Option {
	return optionFunc(func(o *options) { o.query["is_pay"] = isPay })
}

// WithTradeType trade_type获取 交易类型
func (obj *PaymentBillMgr) WithTradeType(tradeType string) Option {
	return optionFunc(func(o *options) { o.query["trade_type"] = tradeType })
}

// WithPaymentName payment_name获取 支付方式名称
func (obj *PaymentBillMgr) WithPaymentName(paymentName string) Option {
	return optionFunc(func(o *options) { o.query["payment_name"] = paymentName })
}

// WithPayConfig pay_config获取 支付参数
func (obj *PaymentBillMgr) WithPayConfig(payConfig string) Option {
	return optionFunc(func(o *options) { o.query["pay_config"] = payConfig })
}

// WithTradePrice trade_price获取 交易金额
func (obj *PaymentBillMgr) WithTradePrice(tradePrice float64) Option {
	return optionFunc(func(o *options) { o.query["trade_price"] = tradePrice })
}

// WithPaymentPluginID payment_plugin_id获取 支付插件
func (obj *PaymentBillMgr) WithPaymentPluginID(paymentPluginID string) Option {
	return optionFunc(func(o *options) { o.query["payment_plugin_id"] = paymentPluginID })
}

// GetByOption 功能选项模式获取
func (obj *PaymentBillMgr) GetByOption(opts ...Option) (result models.EsPaymentBill, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentBill{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *PaymentBillMgr) GetByOptions(opts ...Option) (results []*models.EsPaymentBill, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentBill{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *PaymentBillMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsPaymentBill, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentBill{}).Where(options.query)
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

// GetFromBillID 通过bill_id获取内容 主键id
func (obj *PaymentBillMgr) GetFromBillID(billID int) (result models.EsPaymentBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentBill{}).Where("`bill_id` = ?", billID).First(&result).Error

	return
}

// GetBatchFromBillID 批量查找 主键id
func (obj *PaymentBillMgr) GetBatchFromBillID(billIDs []int) (results []*models.EsPaymentBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentBill{}).Where("`bill_id` IN (?)", billIDs).Find(&results).Error

	return
}

// GetFromSn 通过sn获取内容 单号
func (obj *PaymentBillMgr) GetFromSn(sn string) (results []*models.EsPaymentBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentBill{}).Where("`sn` = ?", sn).Find(&results).Error

	return
}

// GetBatchFromSn 批量查找 单号
func (obj *PaymentBillMgr) GetBatchFromSn(sns []string) (results []*models.EsPaymentBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentBill{}).Where("`sn` IN (?)", sns).Find(&results).Error

	return
}

// GetFromOutTradeNo 通过out_trade_no获取内容 提交给第三方平台单号
func (obj *PaymentBillMgr) GetFromOutTradeNo(outTradeNo string) (results []*models.EsPaymentBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentBill{}).Where("`out_trade_no` = ?", outTradeNo).Find(&results).Error

	return
}

// GetBatchFromOutTradeNo 批量查找 提交给第三方平台单号
func (obj *PaymentBillMgr) GetBatchFromOutTradeNo(outTradeNos []string) (results []*models.EsPaymentBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentBill{}).Where("`out_trade_no` IN (?)", outTradeNos).Find(&results).Error

	return
}

// GetFromReturnTradeNo 通过return_trade_no获取内容 第三方平台返回交易号
func (obj *PaymentBillMgr) GetFromReturnTradeNo(returnTradeNo string) (results []*models.EsPaymentBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentBill{}).Where("`return_trade_no` = ?", returnTradeNo).Find(&results).Error

	return
}

// GetBatchFromReturnTradeNo 批量查找 第三方平台返回交易号
func (obj *PaymentBillMgr) GetBatchFromReturnTradeNo(returnTradeNos []string) (results []*models.EsPaymentBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentBill{}).Where("`return_trade_no` IN (?)", returnTradeNos).Find(&results).Error

	return
}

// GetFromIsPay 通过is_pay获取内容 是否已支付
func (obj *PaymentBillMgr) GetFromIsPay(isPay int) (results []*models.EsPaymentBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentBill{}).Where("`is_pay` = ?", isPay).Find(&results).Error

	return
}

// GetBatchFromIsPay 批量查找 是否已支付
func (obj *PaymentBillMgr) GetBatchFromIsPay(isPays []int) (results []*models.EsPaymentBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentBill{}).Where("`is_pay` IN (?)", isPays).Find(&results).Error

	return
}

// GetFromTradeType 通过trade_type获取内容 交易类型
func (obj *PaymentBillMgr) GetFromTradeType(tradeType string) (results []*models.EsPaymentBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentBill{}).Where("`trade_type` = ?", tradeType).Find(&results).Error

	return
}

// GetBatchFromTradeType 批量查找 交易类型
func (obj *PaymentBillMgr) GetBatchFromTradeType(tradeTypes []string) (results []*models.EsPaymentBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentBill{}).Where("`trade_type` IN (?)", tradeTypes).Find(&results).Error

	return
}

// GetFromPaymentName 通过payment_name获取内容 支付方式名称
func (obj *PaymentBillMgr) GetFromPaymentName(paymentName string) (results []*models.EsPaymentBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentBill{}).Where("`payment_name` = ?", paymentName).Find(&results).Error

	return
}

// GetBatchFromPaymentName 批量查找 支付方式名称
func (obj *PaymentBillMgr) GetBatchFromPaymentName(paymentNames []string) (results []*models.EsPaymentBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentBill{}).Where("`payment_name` IN (?)", paymentNames).Find(&results).Error

	return
}

// GetFromPayConfig 通过pay_config获取内容 支付参数
func (obj *PaymentBillMgr) GetFromPayConfig(payConfig string) (results []*models.EsPaymentBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentBill{}).Where("`pay_config` = ?", payConfig).Find(&results).Error

	return
}

// GetBatchFromPayConfig 批量查找 支付参数
func (obj *PaymentBillMgr) GetBatchFromPayConfig(payConfigs []string) (results []*models.EsPaymentBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentBill{}).Where("`pay_config` IN (?)", payConfigs).Find(&results).Error

	return
}

// GetFromTradePrice 通过trade_price获取内容 交易金额
func (obj *PaymentBillMgr) GetFromTradePrice(tradePrice float64) (results []*models.EsPaymentBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentBill{}).Where("`trade_price` = ?", tradePrice).Find(&results).Error

	return
}

// GetBatchFromTradePrice 批量查找 交易金额
func (obj *PaymentBillMgr) GetBatchFromTradePrice(tradePrices []float64) (results []*models.EsPaymentBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentBill{}).Where("`trade_price` IN (?)", tradePrices).Find(&results).Error

	return
}

// GetFromPaymentPluginID 通过payment_plugin_id获取内容 支付插件
func (obj *PaymentBillMgr) GetFromPaymentPluginID(paymentPluginID string) (results []*models.EsPaymentBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentBill{}).Where("`payment_plugin_id` = ?", paymentPluginID).Find(&results).Error

	return
}

// GetBatchFromPaymentPluginID 批量查找 支付插件
func (obj *PaymentBillMgr) GetBatchFromPaymentPluginID(paymentPluginIDs []string) (results []*models.EsPaymentBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentBill{}).Where("`payment_plugin_id` IN (?)", paymentPluginIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *PaymentBillMgr) FetchByPrimaryKey(billID int) (result models.EsPaymentBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentBill{}).Where("`bill_id` = ?", billID).First(&result).Error

	return
}
