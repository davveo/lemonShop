package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsRefundMgr struct {
	*_BaseMgr
}

// EsRefundMgr open func
func EsRefundMgr(db *gorm.DB) *_EsRefundMgr {
	if db == nil {
		panic(fmt.Errorf("EsRefundMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsRefundMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_refund"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsRefundMgr) GetTableName() string {
	return "es_refund"
}

// Reset 重置gorm会话
func (obj *_EsRefundMgr) Reset() *_EsRefundMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsRefundMgr) Get() (result models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsRefundMgr) Gets() (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsRefundMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_EsRefundMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSn sn获取 售后服务单号
func (obj *_EsRefundMgr) WithSn(sn string) Option {
	return optionFunc(func(o *options) { o.query["sn"] = sn })
}

// WithMemberID member_id获取 会员id
func (obj *_EsRefundMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithMemberName member_name获取 会员名称
func (obj *_EsRefundMgr) WithMemberName(memberName string) Option {
	return optionFunc(func(o *options) { o.query["member_name"] = memberName })
}

// WithSellerID seller_id获取 卖家id
func (obj *_EsRefundMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithSellerName seller_name获取 卖家姓名
func (obj *_EsRefundMgr) WithSellerName(sellerName string) Option {
	return optionFunc(func(o *options) { o.query["seller_name"] = sellerName })
}

// WithOrderSn order_sn获取 订单编号
func (obj *_EsRefundMgr) WithOrderSn(orderSn string) Option {
	return optionFunc(func(o *options) { o.query["order_sn"] = orderSn })
}

// WithRefundStatus refund_status获取 退款单状态 APPLY：新创建，REFUNDING：退款中，REFUNDFAIL：退款失败，COMPLETE：完成
func (obj *_EsRefundMgr) WithRefundStatus(refundStatus string) Option {
	return optionFunc(func(o *options) { o.query["refund_status"] = refundStatus })
}

// WithCreateTime create_time获取 创建时间
func (obj *_EsRefundMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithRefundPrice refund_price获取 申请退款金额
func (obj *_EsRefundMgr) WithRefundPrice(refundPrice float64) Option {
	return optionFunc(func(o *options) { o.query["refund_price"] = refundPrice })
}

// WithRefundWay refund_way获取 退款方式 ORIGINAL：原路退回，OFFLINE：线下退款，ACCOUNT：账户退款
func (obj *_EsRefundMgr) WithRefundWay(refundWay string) Option {
	return optionFunc(func(o *options) { o.query["refund_way"] = refundWay })
}

// WithAccountType account_type获取 退款账户类型
func (obj *_EsRefundMgr) WithAccountType(accountType string) Option {
	return optionFunc(func(o *options) { o.query["account_type"] = accountType })
}

// WithReturnAccount return_account获取 退款账户
func (obj *_EsRefundMgr) WithReturnAccount(returnAccount string) Option {
	return optionFunc(func(o *options) { o.query["return_account"] = returnAccount })
}

// WithBankName bank_name获取 银行名称
func (obj *_EsRefundMgr) WithBankName(bankName string) Option {
	return optionFunc(func(o *options) { o.query["bank_name"] = bankName })
}

// WithBankAccountNumber bank_account_number获取 银行账号
func (obj *_EsRefundMgr) WithBankAccountNumber(bankAccountNumber string) Option {
	return optionFunc(func(o *options) { o.query["bank_account_number"] = bankAccountNumber })
}

// WithBankAccountName bank_account_name获取 银行开户名
func (obj *_EsRefundMgr) WithBankAccountName(bankAccountName string) Option {
	return optionFunc(func(o *options) { o.query["bank_account_name"] = bankAccountName })
}

// WithBankDepositName bank_deposit_name获取 银行开户行
func (obj *_EsRefundMgr) WithBankDepositName(bankDepositName string) Option {
	return optionFunc(func(o *options) { o.query["bank_deposit_name"] = bankDepositName })
}

// WithPayOrderNo pay_order_no获取 支付结果交易号
func (obj *_EsRefundMgr) WithPayOrderNo(payOrderNo string) Option {
	return optionFunc(func(o *options) { o.query["pay_order_no"] = payOrderNo })
}

// WithPaymentType payment_type获取 订单付款类型 ONLINE：在线支付，COD：货到付款
func (obj *_EsRefundMgr) WithPaymentType(paymentType string) Option {
	return optionFunc(func(o *options) { o.query["payment_type"] = paymentType })
}

// WithRefundFailReason refund_fail_reason获取 退款失败原因
func (obj *_EsRefundMgr) WithRefundFailReason(refundFailReason string) Option {
	return optionFunc(func(o *options) { o.query["refund_fail_reason"] = refundFailReason })
}

// WithRefundTime refund_time获取 退款时间
func (obj *_EsRefundMgr) WithRefundTime(refundTime int64) Option {
	return optionFunc(func(o *options) { o.query["refund_time"] = refundTime })
}

// WithMobile mobile获取 手机号
func (obj *_EsRefundMgr) WithMobile(mobile string) Option {
	return optionFunc(func(o *options) { o.query["mobile"] = mobile })
}

// WithAgreePrice agree_price获取 商家同意的退款金额
func (obj *_EsRefundMgr) WithAgreePrice(agreePrice float64) Option {
	return optionFunc(func(o *options) { o.query["agree_price"] = agreePrice })
}

// WithActualPrice actual_price获取 实际退款金额
func (obj *_EsRefundMgr) WithActualPrice(actualPrice float64) Option {
	return optionFunc(func(o *options) { o.query["actual_price"] = actualPrice })
}

// WithGoodsJSON goods_json获取 售后商品信息json
func (obj *_EsRefundMgr) WithGoodsJSON(goodsJSON string) Option {
	return optionFunc(func(o *options) { o.query["goods_json"] = goodsJSON })
}

// WithDisabled disabled获取 删除状态 DELETED：已删除 NORMAL：正常
func (obj *_EsRefundMgr) WithDisabled(disabled string) Option {
	return optionFunc(func(o *options) { o.query["disabled"] = disabled })
}

// WithCreateChannel create_channel获取 售后服务单创建渠道 NORMAL：正常渠道创建，PINTUAN：拼团失败自动创建
func (obj *_EsRefundMgr) WithCreateChannel(createChannel string) Option {
	return optionFunc(func(o *options) { o.query["create_channel"] = createChannel })
}

// GetByOption 功能选项模式获取
func (obj *_EsRefundMgr) GetByOption(opts ...Option) (result models.EsRefund, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsRefundMgr) GetByOptions(opts ...Option) (results []*models.EsRefund, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsRefundMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsRefund, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where(options.query)
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

// GetFromID 通过id获取内容 主键id
func (obj *_EsRefundMgr) GetFromID(id int) (result models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键id
func (obj *_EsRefundMgr) GetBatchFromID(ids []int) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSn 通过sn获取内容 售后服务单号
func (obj *_EsRefundMgr) GetFromSn(sn string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`sn` = ?", sn).Find(&results).Error

	return
}

// GetBatchFromSn 批量查找 售后服务单号
func (obj *_EsRefundMgr) GetBatchFromSn(sns []string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`sn` IN (?)", sns).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *_EsRefundMgr) GetFromMemberID(memberID int) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *_EsRefundMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromMemberName 通过member_name获取内容 会员名称
func (obj *_EsRefundMgr) GetFromMemberName(memberName string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`member_name` = ?", memberName).Find(&results).Error

	return
}

// GetBatchFromMemberName 批量查找 会员名称
func (obj *_EsRefundMgr) GetBatchFromMemberName(memberNames []string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`member_name` IN (?)", memberNames).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 卖家id
func (obj *_EsRefundMgr) GetFromSellerID(sellerID int) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 卖家id
func (obj *_EsRefundMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromSellerName 通过seller_name获取内容 卖家姓名
func (obj *_EsRefundMgr) GetFromSellerName(sellerName string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`seller_name` = ?", sellerName).Find(&results).Error

	return
}

// GetBatchFromSellerName 批量查找 卖家姓名
func (obj *_EsRefundMgr) GetBatchFromSellerName(sellerNames []string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`seller_name` IN (?)", sellerNames).Find(&results).Error

	return
}

// GetFromOrderSn 通过order_sn获取内容 订单编号
func (obj *_EsRefundMgr) GetFromOrderSn(orderSn string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}

// GetBatchFromOrderSn 批量查找 订单编号
func (obj *_EsRefundMgr) GetBatchFromOrderSn(orderSns []string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`order_sn` IN (?)", orderSns).Find(&results).Error

	return
}

// GetFromRefundStatus 通过refund_status获取内容 退款单状态 APPLY：新创建，REFUNDING：退款中，REFUNDFAIL：退款失败，COMPLETE：完成
func (obj *_EsRefundMgr) GetFromRefundStatus(refundStatus string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`refund_status` = ?", refundStatus).Find(&results).Error

	return
}

// GetBatchFromRefundStatus 批量查找 退款单状态 APPLY：新创建，REFUNDING：退款中，REFUNDFAIL：退款失败，COMPLETE：完成
func (obj *_EsRefundMgr) GetBatchFromRefundStatus(refundStatuss []string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`refund_status` IN (?)", refundStatuss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_EsRefundMgr) GetFromCreateTime(createTime int64) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_EsRefundMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromRefundPrice 通过refund_price获取内容 申请退款金额
func (obj *_EsRefundMgr) GetFromRefundPrice(refundPrice float64) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`refund_price` = ?", refundPrice).Find(&results).Error

	return
}

// GetBatchFromRefundPrice 批量查找 申请退款金额
func (obj *_EsRefundMgr) GetBatchFromRefundPrice(refundPrices []float64) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`refund_price` IN (?)", refundPrices).Find(&results).Error

	return
}

// GetFromRefundWay 通过refund_way获取内容 退款方式 ORIGINAL：原路退回，OFFLINE：线下退款，ACCOUNT：账户退款
func (obj *_EsRefundMgr) GetFromRefundWay(refundWay string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`refund_way` = ?", refundWay).Find(&results).Error

	return
}

// GetBatchFromRefundWay 批量查找 退款方式 ORIGINAL：原路退回，OFFLINE：线下退款，ACCOUNT：账户退款
func (obj *_EsRefundMgr) GetBatchFromRefundWay(refundWays []string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`refund_way` IN (?)", refundWays).Find(&results).Error

	return
}

// GetFromAccountType 通过account_type获取内容 退款账户类型
func (obj *_EsRefundMgr) GetFromAccountType(accountType string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`account_type` = ?", accountType).Find(&results).Error

	return
}

// GetBatchFromAccountType 批量查找 退款账户类型
func (obj *_EsRefundMgr) GetBatchFromAccountType(accountTypes []string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`account_type` IN (?)", accountTypes).Find(&results).Error

	return
}

// GetFromReturnAccount 通过return_account获取内容 退款账户
func (obj *_EsRefundMgr) GetFromReturnAccount(returnAccount string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`return_account` = ?", returnAccount).Find(&results).Error

	return
}

// GetBatchFromReturnAccount 批量查找 退款账户
func (obj *_EsRefundMgr) GetBatchFromReturnAccount(returnAccounts []string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`return_account` IN (?)", returnAccounts).Find(&results).Error

	return
}

// GetFromBankName 通过bank_name获取内容 银行名称
func (obj *_EsRefundMgr) GetFromBankName(bankName string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`bank_name` = ?", bankName).Find(&results).Error

	return
}

// GetBatchFromBankName 批量查找 银行名称
func (obj *_EsRefundMgr) GetBatchFromBankName(bankNames []string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`bank_name` IN (?)", bankNames).Find(&results).Error

	return
}

// GetFromBankAccountNumber 通过bank_account_number获取内容 银行账号
func (obj *_EsRefundMgr) GetFromBankAccountNumber(bankAccountNumber string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`bank_account_number` = ?", bankAccountNumber).Find(&results).Error

	return
}

// GetBatchFromBankAccountNumber 批量查找 银行账号
func (obj *_EsRefundMgr) GetBatchFromBankAccountNumber(bankAccountNumbers []string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`bank_account_number` IN (?)", bankAccountNumbers).Find(&results).Error

	return
}

// GetFromBankAccountName 通过bank_account_name获取内容 银行开户名
func (obj *_EsRefundMgr) GetFromBankAccountName(bankAccountName string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`bank_account_name` = ?", bankAccountName).Find(&results).Error

	return
}

// GetBatchFromBankAccountName 批量查找 银行开户名
func (obj *_EsRefundMgr) GetBatchFromBankAccountName(bankAccountNames []string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`bank_account_name` IN (?)", bankAccountNames).Find(&results).Error

	return
}

// GetFromBankDepositName 通过bank_deposit_name获取内容 银行开户行
func (obj *_EsRefundMgr) GetFromBankDepositName(bankDepositName string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`bank_deposit_name` = ?", bankDepositName).Find(&results).Error

	return
}

// GetBatchFromBankDepositName 批量查找 银行开户行
func (obj *_EsRefundMgr) GetBatchFromBankDepositName(bankDepositNames []string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`bank_deposit_name` IN (?)", bankDepositNames).Find(&results).Error

	return
}

// GetFromPayOrderNo 通过pay_order_no获取内容 支付结果交易号
func (obj *_EsRefundMgr) GetFromPayOrderNo(payOrderNo string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`pay_order_no` = ?", payOrderNo).Find(&results).Error

	return
}

// GetBatchFromPayOrderNo 批量查找 支付结果交易号
func (obj *_EsRefundMgr) GetBatchFromPayOrderNo(payOrderNos []string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`pay_order_no` IN (?)", payOrderNos).Find(&results).Error

	return
}

// GetFromPaymentType 通过payment_type获取内容 订单付款类型 ONLINE：在线支付，COD：货到付款
func (obj *_EsRefundMgr) GetFromPaymentType(paymentType string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`payment_type` = ?", paymentType).Find(&results).Error

	return
}

// GetBatchFromPaymentType 批量查找 订单付款类型 ONLINE：在线支付，COD：货到付款
func (obj *_EsRefundMgr) GetBatchFromPaymentType(paymentTypes []string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`payment_type` IN (?)", paymentTypes).Find(&results).Error

	return
}

// GetFromRefundFailReason 通过refund_fail_reason获取内容 退款失败原因
func (obj *_EsRefundMgr) GetFromRefundFailReason(refundFailReason string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`refund_fail_reason` = ?", refundFailReason).Find(&results).Error

	return
}

// GetBatchFromRefundFailReason 批量查找 退款失败原因
func (obj *_EsRefundMgr) GetBatchFromRefundFailReason(refundFailReasons []string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`refund_fail_reason` IN (?)", refundFailReasons).Find(&results).Error

	return
}

// GetFromRefundTime 通过refund_time获取内容 退款时间
func (obj *_EsRefundMgr) GetFromRefundTime(refundTime int64) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`refund_time` = ?", refundTime).Find(&results).Error

	return
}

// GetBatchFromRefundTime 批量查找 退款时间
func (obj *_EsRefundMgr) GetBatchFromRefundTime(refundTimes []int64) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`refund_time` IN (?)", refundTimes).Find(&results).Error

	return
}

// GetFromMobile 通过mobile获取内容 手机号
func (obj *_EsRefundMgr) GetFromMobile(mobile string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`mobile` = ?", mobile).Find(&results).Error

	return
}

// GetBatchFromMobile 批量查找 手机号
func (obj *_EsRefundMgr) GetBatchFromMobile(mobiles []string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`mobile` IN (?)", mobiles).Find(&results).Error

	return
}

// GetFromAgreePrice 通过agree_price获取内容 商家同意的退款金额
func (obj *_EsRefundMgr) GetFromAgreePrice(agreePrice float64) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`agree_price` = ?", agreePrice).Find(&results).Error

	return
}

// GetBatchFromAgreePrice 批量查找 商家同意的退款金额
func (obj *_EsRefundMgr) GetBatchFromAgreePrice(agreePrices []float64) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`agree_price` IN (?)", agreePrices).Find(&results).Error

	return
}

// GetFromActualPrice 通过actual_price获取内容 实际退款金额
func (obj *_EsRefundMgr) GetFromActualPrice(actualPrice float64) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`actual_price` = ?", actualPrice).Find(&results).Error

	return
}

// GetBatchFromActualPrice 批量查找 实际退款金额
func (obj *_EsRefundMgr) GetBatchFromActualPrice(actualPrices []float64) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`actual_price` IN (?)", actualPrices).Find(&results).Error

	return
}

// GetFromGoodsJSON 通过goods_json获取内容 售后商品信息json
func (obj *_EsRefundMgr) GetFromGoodsJSON(goodsJSON string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`goods_json` = ?", goodsJSON).Find(&results).Error

	return
}

// GetBatchFromGoodsJSON 批量查找 售后商品信息json
func (obj *_EsRefundMgr) GetBatchFromGoodsJSON(goodsJSONs []string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`goods_json` IN (?)", goodsJSONs).Find(&results).Error

	return
}

// GetFromDisabled 通过disabled获取内容 删除状态 DELETED：已删除 NORMAL：正常
func (obj *_EsRefundMgr) GetFromDisabled(disabled string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`disabled` = ?", disabled).Find(&results).Error

	return
}

// GetBatchFromDisabled 批量查找 删除状态 DELETED：已删除 NORMAL：正常
func (obj *_EsRefundMgr) GetBatchFromDisabled(disableds []string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`disabled` IN (?)", disableds).Find(&results).Error

	return
}

// GetFromCreateChannel 通过create_channel获取内容 售后服务单创建渠道 NORMAL：正常渠道创建，PINTUAN：拼团失败自动创建
func (obj *_EsRefundMgr) GetFromCreateChannel(createChannel string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`create_channel` = ?", createChannel).Find(&results).Error

	return
}

// GetBatchFromCreateChannel 批量查找 售后服务单创建渠道 NORMAL：正常渠道创建，PINTUAN：拼团失败自动创建
func (obj *_EsRefundMgr) GetBatchFromCreateChannel(createChannels []string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`create_channel` IN (?)", createChannels).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsRefundMgr) FetchByPrimaryKey(id int) (result models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByEsIndexID  获取多个内容
func (obj *_EsRefundMgr) FetchIndexByEsIndexID(id int) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`id` = ?", id).Find(&results).Error

	return
}

// FetchIndexByEsIndexSn  获取多个内容
func (obj *_EsRefundMgr) FetchIndexByEsIndexSn(sn string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`sn` = ?", sn).Find(&results).Error

	return
}

// FetchIndexByEsMemberID  获取多个内容
func (obj *_EsRefundMgr) FetchIndexByEsMemberID(memberID int) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// FetchIndexByEsSellerID  获取多个内容
func (obj *_EsRefundMgr) FetchIndexByEsSellerID(sellerID int) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// FetchIndexByEsIndexOrderSn  获取多个内容
func (obj *_EsRefundMgr) FetchIndexByEsIndexOrderSn(orderSn string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}

// FetchIndexByEsIndexRefundStatus  获取多个内容
func (obj *_EsRefundMgr) FetchIndexByEsIndexRefundStatus(refundStatus string) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`refund_status` = ?", refundStatus).Find(&results).Error

	return
}

// FetchIndexByEsIndexCreateTime  获取多个内容
func (obj *_EsRefundMgr) FetchIndexByEsIndexCreateTime(createTime int64) (results []*models.EsRefund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}
