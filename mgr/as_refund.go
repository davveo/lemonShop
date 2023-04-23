package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsAsRefundMgr struct {
	*_BaseMgr
}

// EsAsRefundMgr open func
func EsAsRefundMgr(db *gorm.DB) *_EsAsRefundMgr {
	if db == nil {
		panic(fmt.Errorf("EsAsRefundMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsAsRefundMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_as_refund"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsAsRefundMgr) GetTableName() string {
	return "es_as_refund"
}

// Reset 重置gorm会话
func (obj *_EsAsRefundMgr) Reset() *_EsAsRefundMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsAsRefundMgr) Get() (result models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsAsRefundMgr) Gets() (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsAsRefundMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键ID
func (obj *_EsAsRefundMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithServiceSn service_sn获取 售后服务单号
func (obj *_EsAsRefundMgr) WithServiceSn(serviceSn string) Option {
	return optionFunc(func(o *options) { o.query["service_sn"] = serviceSn })
}

// WithRefundPrice refund_price获取 申请的退款金额
func (obj *_EsAsRefundMgr) WithRefundPrice(refundPrice float64) Option {
	return optionFunc(func(o *options) { o.query["refund_price"] = refundPrice })
}

// WithAgreePrice agree_price获取 商家同意的退款金额
func (obj *_EsAsRefundMgr) WithAgreePrice(agreePrice float64) Option {
	return optionFunc(func(o *options) { o.query["agree_price"] = agreePrice })
}

// WithActualPrice actual_price获取 实际退款金额
func (obj *_EsAsRefundMgr) WithActualPrice(actualPrice float64) Option {
	return optionFunc(func(o *options) { o.query["actual_price"] = actualPrice })
}

// WithRefundWay refund_way获取 退款方式 ORIGINAL：原路退回，OFFLINE：线下支付
func (obj *_EsAsRefundMgr) WithRefundWay(refundWay string) Option {
	return optionFunc(func(o *options) { o.query["refund_way"] = refundWay })
}

// WithAccountType account_type获取 账号类型
func (obj *_EsAsRefundMgr) WithAccountType(accountType string) Option {
	return optionFunc(func(o *options) { o.query["account_type"] = accountType })
}

// WithReturnAccount return_account获取 退款账号
func (obj *_EsAsRefundMgr) WithReturnAccount(returnAccount string) Option {
	return optionFunc(func(o *options) { o.query["return_account"] = returnAccount })
}

// WithBankName bank_name获取 银行名称
func (obj *_EsAsRefundMgr) WithBankName(bankName string) Option {
	return optionFunc(func(o *options) { o.query["bank_name"] = bankName })
}

// WithBankAccountNumber bank_account_number获取 银行账户
func (obj *_EsAsRefundMgr) WithBankAccountNumber(bankAccountNumber string) Option {
	return optionFunc(func(o *options) { o.query["bank_account_number"] = bankAccountNumber })
}

// WithBankAccountName bank_account_name获取 银行开户名
func (obj *_EsAsRefundMgr) WithBankAccountName(bankAccountName string) Option {
	return optionFunc(func(o *options) { o.query["bank_account_name"] = bankAccountName })
}

// WithBankDepositName bank_deposit_name获取 开户行
func (obj *_EsAsRefundMgr) WithBankDepositName(bankDepositName string) Option {
	return optionFunc(func(o *options) { o.query["bank_deposit_name"] = bankDepositName })
}

// WithPayOrderNo pay_order_no获取 订单支付方式返回的交易号
func (obj *_EsAsRefundMgr) WithPayOrderNo(payOrderNo string) Option {
	return optionFunc(func(o *options) { o.query["pay_order_no"] = payOrderNo })
}

// WithRefundTime refund_time获取 退款时间
func (obj *_EsAsRefundMgr) WithRefundTime(refundTime int64) Option {
	return optionFunc(func(o *options) { o.query["refund_time"] = refundTime })
}

// GetByOption 功能选项模式获取
func (obj *_EsAsRefundMgr) GetByOption(opts ...Option) (result models.EsAsRefund, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsAsRefundMgr) GetByOptions(opts ...Option) (results []*models.EsAsRefund, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsAsRefundMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsAsRefund, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where(options.query)
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

// GetFromID 通过id获取内容 主键ID
func (obj *_EsAsRefundMgr) GetFromID(id int) (result models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键ID
func (obj *_EsAsRefundMgr) GetBatchFromID(ids []int) (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromServiceSn 通过service_sn获取内容 售后服务单号
func (obj *_EsAsRefundMgr) GetFromServiceSn(serviceSn string) (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`service_sn` = ?", serviceSn).Find(&results).Error

	return
}

// GetBatchFromServiceSn 批量查找 售后服务单号
func (obj *_EsAsRefundMgr) GetBatchFromServiceSn(serviceSns []string) (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`service_sn` IN (?)", serviceSns).Find(&results).Error

	return
}

// GetFromRefundPrice 通过refund_price获取内容 申请的退款金额
func (obj *_EsAsRefundMgr) GetFromRefundPrice(refundPrice float64) (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`refund_price` = ?", refundPrice).Find(&results).Error

	return
}

// GetBatchFromRefundPrice 批量查找 申请的退款金额
func (obj *_EsAsRefundMgr) GetBatchFromRefundPrice(refundPrices []float64) (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`refund_price` IN (?)", refundPrices).Find(&results).Error

	return
}

// GetFromAgreePrice 通过agree_price获取内容 商家同意的退款金额
func (obj *_EsAsRefundMgr) GetFromAgreePrice(agreePrice float64) (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`agree_price` = ?", agreePrice).Find(&results).Error

	return
}

// GetBatchFromAgreePrice 批量查找 商家同意的退款金额
func (obj *_EsAsRefundMgr) GetBatchFromAgreePrice(agreePrices []float64) (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`agree_price` IN (?)", agreePrices).Find(&results).Error

	return
}

// GetFromActualPrice 通过actual_price获取内容 实际退款金额
func (obj *_EsAsRefundMgr) GetFromActualPrice(actualPrice float64) (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`actual_price` = ?", actualPrice).Find(&results).Error

	return
}

// GetBatchFromActualPrice 批量查找 实际退款金额
func (obj *_EsAsRefundMgr) GetBatchFromActualPrice(actualPrices []float64) (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`actual_price` IN (?)", actualPrices).Find(&results).Error

	return
}

// GetFromRefundWay 通过refund_way获取内容 退款方式 ORIGINAL：原路退回，OFFLINE：线下支付
func (obj *_EsAsRefundMgr) GetFromRefundWay(refundWay string) (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`refund_way` = ?", refundWay).Find(&results).Error

	return
}

// GetBatchFromRefundWay 批量查找 退款方式 ORIGINAL：原路退回，OFFLINE：线下支付
func (obj *_EsAsRefundMgr) GetBatchFromRefundWay(refundWays []string) (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`refund_way` IN (?)", refundWays).Find(&results).Error

	return
}

// GetFromAccountType 通过account_type获取内容 账号类型
func (obj *_EsAsRefundMgr) GetFromAccountType(accountType string) (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`account_type` = ?", accountType).Find(&results).Error

	return
}

// GetBatchFromAccountType 批量查找 账号类型
func (obj *_EsAsRefundMgr) GetBatchFromAccountType(accountTypes []string) (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`account_type` IN (?)", accountTypes).Find(&results).Error

	return
}

// GetFromReturnAccount 通过return_account获取内容 退款账号
func (obj *_EsAsRefundMgr) GetFromReturnAccount(returnAccount string) (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`return_account` = ?", returnAccount).Find(&results).Error

	return
}

// GetBatchFromReturnAccount 批量查找 退款账号
func (obj *_EsAsRefundMgr) GetBatchFromReturnAccount(returnAccounts []string) (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`return_account` IN (?)", returnAccounts).Find(&results).Error

	return
}

// GetFromBankName 通过bank_name获取内容 银行名称
func (obj *_EsAsRefundMgr) GetFromBankName(bankName string) (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`bank_name` = ?", bankName).Find(&results).Error

	return
}

// GetBatchFromBankName 批量查找 银行名称
func (obj *_EsAsRefundMgr) GetBatchFromBankName(bankNames []string) (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`bank_name` IN (?)", bankNames).Find(&results).Error

	return
}

// GetFromBankAccountNumber 通过bank_account_number获取内容 银行账户
func (obj *_EsAsRefundMgr) GetFromBankAccountNumber(bankAccountNumber string) (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`bank_account_number` = ?", bankAccountNumber).Find(&results).Error

	return
}

// GetBatchFromBankAccountNumber 批量查找 银行账户
func (obj *_EsAsRefundMgr) GetBatchFromBankAccountNumber(bankAccountNumbers []string) (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`bank_account_number` IN (?)", bankAccountNumbers).Find(&results).Error

	return
}

// GetFromBankAccountName 通过bank_account_name获取内容 银行开户名
func (obj *_EsAsRefundMgr) GetFromBankAccountName(bankAccountName string) (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`bank_account_name` = ?", bankAccountName).Find(&results).Error

	return
}

// GetBatchFromBankAccountName 批量查找 银行开户名
func (obj *_EsAsRefundMgr) GetBatchFromBankAccountName(bankAccountNames []string) (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`bank_account_name` IN (?)", bankAccountNames).Find(&results).Error

	return
}

// GetFromBankDepositName 通过bank_deposit_name获取内容 开户行
func (obj *_EsAsRefundMgr) GetFromBankDepositName(bankDepositName string) (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`bank_deposit_name` = ?", bankDepositName).Find(&results).Error

	return
}

// GetBatchFromBankDepositName 批量查找 开户行
func (obj *_EsAsRefundMgr) GetBatchFromBankDepositName(bankDepositNames []string) (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`bank_deposit_name` IN (?)", bankDepositNames).Find(&results).Error

	return
}

// GetFromPayOrderNo 通过pay_order_no获取内容 订单支付方式返回的交易号
func (obj *_EsAsRefundMgr) GetFromPayOrderNo(payOrderNo string) (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`pay_order_no` = ?", payOrderNo).Find(&results).Error

	return
}

// GetBatchFromPayOrderNo 批量查找 订单支付方式返回的交易号
func (obj *_EsAsRefundMgr) GetBatchFromPayOrderNo(payOrderNos []string) (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`pay_order_no` IN (?)", payOrderNos).Find(&results).Error

	return
}

// GetFromRefundTime 通过refund_time获取内容 退款时间
func (obj *_EsAsRefundMgr) GetFromRefundTime(refundTime int64) (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`refund_time` = ?", refundTime).Find(&results).Error

	return
}

// GetBatchFromRefundTime 批量查找 退款时间
func (obj *_EsAsRefundMgr) GetBatchFromRefundTime(refundTimes []int64) (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`refund_time` IN (?)", refundTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsAsRefundMgr) FetchByPrimaryKey(id int) (result models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByEsIndexID  获取多个内容
func (obj *_EsAsRefundMgr) FetchIndexByEsIndexID(id int) (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`id` = ?", id).Find(&results).Error

	return
}

// FetchIndexByEsIndexServiceSn  获取多个内容
func (obj *_EsAsRefundMgr) FetchIndexByEsIndexServiceSn(serviceSn string) (results []*models.EsAsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsRefund{}).Where("`service_sn` = ?", serviceSn).Find(&results).Error

	return
}
