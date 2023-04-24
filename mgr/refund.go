package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type RefundMgr struct {
	*_BaseMgr
}

// NewRefundMgr open func
func NewRefundMgr(db db.Repo) *RefundMgr {
	if db == nil {
		panic(fmt.Errorf("NewRefundMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &RefundMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_refund"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *RefundMgr) GetTableName() string {
	return "es_refund"
}

// Reset 重置gorm会话
func (obj *RefundMgr) Reset() *RefundMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *RefundMgr) Get() (result models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *RefundMgr) Gets() (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *RefundMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *RefundMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSn sn获取 售后服务单号
func (obj *RefundMgr) WithSn(sn string) Option {
	return optionFunc(func(o *options) { o.query["sn"] = sn })
}

// WithMemberID member_id获取 会员id
func (obj *RefundMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithMemberName member_name获取 会员名称
func (obj *RefundMgr) WithMemberName(memberName string) Option {
	return optionFunc(func(o *options) { o.query["member_name"] = memberName })
}

// WithSellerID seller_id获取 卖家id
func (obj *RefundMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithSellerName seller_name获取 卖家姓名
func (obj *RefundMgr) WithSellerName(sellerName string) Option {
	return optionFunc(func(o *options) { o.query["seller_name"] = sellerName })
}

// WithOrderSn order_sn获取 订单编号
func (obj *RefundMgr) WithOrderSn(orderSn string) Option {
	return optionFunc(func(o *options) { o.query["order_sn"] = orderSn })
}

// WithRefundStatus refund_status获取 退款单状态 APPLY：新创建，REFUNDING：退款中，REFUNDFAIL：退款失败，COMPLETE：完成
func (obj *RefundMgr) WithRefundStatus(refundStatus string) Option {
	return optionFunc(func(o *options) { o.query["refund_status"] = refundStatus })
}

// WithCreateTime create_time获取 创建时间
func (obj *RefundMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithRefundPrice refund_price获取 申请退款金额
func (obj *RefundMgr) WithRefundPrice(refundPrice float64) Option {
	return optionFunc(func(o *options) { o.query["refund_price"] = refundPrice })
}

// WithRefundWay refund_way获取 退款方式 ORIGINAL：原路退回，OFFLINE：线下退款，ACCOUNT：账户退款
func (obj *RefundMgr) WithRefundWay(refundWay string) Option {
	return optionFunc(func(o *options) { o.query["refund_way"] = refundWay })
}

// WithAccountType account_type获取 退款账户类型
func (obj *RefundMgr) WithAccountType(accountType string) Option {
	return optionFunc(func(o *options) { o.query["account_type"] = accountType })
}

// WithReturnAccount return_account获取 退款账户
func (obj *RefundMgr) WithReturnAccount(returnAccount string) Option {
	return optionFunc(func(o *options) { o.query["return_account"] = returnAccount })
}

// WithBankName bank_name获取 银行名称
func (obj *RefundMgr) WithBankName(bankName string) Option {
	return optionFunc(func(o *options) { o.query["bank_name"] = bankName })
}

// WithBankAccountNumber bank_account_number获取 银行账号
func (obj *RefundMgr) WithBankAccountNumber(bankAccountNumber string) Option {
	return optionFunc(func(o *options) { o.query["bank_account_number"] = bankAccountNumber })
}

// WithBankAccountName bank_account_name获取 银行开户名
func (obj *RefundMgr) WithBankAccountName(bankAccountName string) Option {
	return optionFunc(func(o *options) { o.query["bank_account_name"] = bankAccountName })
}

// WithBankDepositName bank_deposit_name获取 银行开户行
func (obj *RefundMgr) WithBankDepositName(bankDepositName string) Option {
	return optionFunc(func(o *options) { o.query["bank_deposit_name"] = bankDepositName })
}

// WithPayOrderNo pay_order_no获取 支付结果交易号
func (obj *RefundMgr) WithPayOrderNo(payOrderNo string) Option {
	return optionFunc(func(o *options) { o.query["pay_order_no"] = payOrderNo })
}

// WithPaymentType payment_type获取 订单付款类型 ONLINE：在线支付，COD：货到付款
func (obj *RefundMgr) WithPaymentType(paymentType string) Option {
	return optionFunc(func(o *options) { o.query["payment_type"] = paymentType })
}

// WithRefundFailReason refund_fail_reason获取 退款失败原因
func (obj *RefundMgr) WithRefundFailReason(refundFailReason string) Option {
	return optionFunc(func(o *options) { o.query["refund_fail_reason"] = refundFailReason })
}

// WithRefundTime refund_time获取 退款时间
func (obj *RefundMgr) WithRefundTime(refundTime int64) Option {
	return optionFunc(func(o *options) { o.query["refund_time"] = refundTime })
}

// WithMobile mobile获取 手机号
func (obj *RefundMgr) WithMobile(mobile string) Option {
	return optionFunc(func(o *options) { o.query["mobile"] = mobile })
}

// WithAgreePrice agree_price获取 商家同意的退款金额
func (obj *RefundMgr) WithAgreePrice(agreePrice float64) Option {
	return optionFunc(func(o *options) { o.query["agree_price"] = agreePrice })
}

// WithActualPrice actual_price获取 实际退款金额
func (obj *RefundMgr) WithActualPrice(actualPrice float64) Option {
	return optionFunc(func(o *options) { o.query["actual_price"] = actualPrice })
}

// WithGoodsJSON goods_json获取 售后商品信息json
func (obj *RefundMgr) WithGoodsJSON(goodsJSON string) Option {
	return optionFunc(func(o *options) { o.query["goods_json"] = goodsJSON })
}

// WithDisabled disabled获取 删除状态 DELETED：已删除 NORMAL：正常
func (obj *RefundMgr) WithDisabled(disabled string) Option {
	return optionFunc(func(o *options) { o.query["disabled"] = disabled })
}

// WithCreateChannel create_channel获取 售后服务单创建渠道 NORMAL：正常渠道创建，PINTUAN：拼团失败自动创建
func (obj *RefundMgr) WithCreateChannel(createChannel string) Option {
	return optionFunc(func(o *options) { o.query["create_channel"] = createChannel })
}

// GetByOption 功能选项模式获取
func (obj *RefundMgr) GetByOption(opts ...Option) (result models.EsRefund, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *RefundMgr) GetByOptions(opts ...Option) (results []*models.EsRefund, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *RefundMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsRefund, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where(options.query)
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
func (obj *RefundMgr) GetFromID(id int) (result models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键id
func (obj *RefundMgr) GetBatchFromID(ids []int) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSn 通过sn获取内容 售后服务单号
func (obj *RefundMgr) GetFromSn(sn string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`sn` = ?", sn).Find(&results).Error

	return
}

// GetBatchFromSn 批量查找 售后服务单号
func (obj *RefundMgr) GetBatchFromSn(sns []string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`sn` IN (?)", sns).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *RefundMgr) GetFromMemberID(memberID int) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *RefundMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromMemberName 通过member_name获取内容 会员名称
func (obj *RefundMgr) GetFromMemberName(memberName string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`member_name` = ?", memberName).Find(&results).Error

	return
}

// GetBatchFromMemberName 批量查找 会员名称
func (obj *RefundMgr) GetBatchFromMemberName(memberNames []string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`member_name` IN (?)", memberNames).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 卖家id
func (obj *RefundMgr) GetFromSellerID(sellerID int) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 卖家id
func (obj *RefundMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromSellerName 通过seller_name获取内容 卖家姓名
func (obj *RefundMgr) GetFromSellerName(sellerName string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`seller_name` = ?", sellerName).Find(&results).Error

	return
}

// GetBatchFromSellerName 批量查找 卖家姓名
func (obj *RefundMgr) GetBatchFromSellerName(sellerNames []string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`seller_name` IN (?)", sellerNames).Find(&results).Error

	return
}

// GetFromOrderSn 通过order_sn获取内容 订单编号
func (obj *RefundMgr) GetFromOrderSn(orderSn string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}

// GetBatchFromOrderSn 批量查找 订单编号
func (obj *RefundMgr) GetBatchFromOrderSn(orderSns []string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`order_sn` IN (?)", orderSns).Find(&results).Error

	return
}

// GetFromRefundStatus 通过refund_status获取内容 退款单状态 APPLY：新创建，REFUNDING：退款中，REFUNDFAIL：退款失败，COMPLETE：完成
func (obj *RefundMgr) GetFromRefundStatus(refundStatus string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`refund_status` = ?", refundStatus).Find(&results).Error

	return
}

// GetBatchFromRefundStatus 批量查找 退款单状态 APPLY：新创建，REFUNDING：退款中，REFUNDFAIL：退款失败，COMPLETE：完成
func (obj *RefundMgr) GetBatchFromRefundStatus(refundStatuss []string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`refund_status` IN (?)", refundStatuss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *RefundMgr) GetFromCreateTime(createTime int64) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *RefundMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromRefundPrice 通过refund_price获取内容 申请退款金额
func (obj *RefundMgr) GetFromRefundPrice(refundPrice float64) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`refund_price` = ?", refundPrice).Find(&results).Error

	return
}

// GetBatchFromRefundPrice 批量查找 申请退款金额
func (obj *RefundMgr) GetBatchFromRefundPrice(refundPrices []float64) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`refund_price` IN (?)", refundPrices).Find(&results).Error

	return
}

// GetFromRefundWay 通过refund_way获取内容 退款方式 ORIGINAL：原路退回，OFFLINE：线下退款，ACCOUNT：账户退款
func (obj *RefundMgr) GetFromRefundWay(refundWay string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`refund_way` = ?", refundWay).Find(&results).Error

	return
}

// GetBatchFromRefundWay 批量查找 退款方式 ORIGINAL：原路退回，OFFLINE：线下退款，ACCOUNT：账户退款
func (obj *RefundMgr) GetBatchFromRefundWay(refundWays []string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`refund_way` IN (?)", refundWays).Find(&results).Error

	return
}

// GetFromAccountType 通过account_type获取内容 退款账户类型
func (obj *RefundMgr) GetFromAccountType(accountType string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`account_type` = ?", accountType).Find(&results).Error

	return
}

// GetBatchFromAccountType 批量查找 退款账户类型
func (obj *RefundMgr) GetBatchFromAccountType(accountTypes []string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`account_type` IN (?)", accountTypes).Find(&results).Error

	return
}

// GetFromReturnAccount 通过return_account获取内容 退款账户
func (obj *RefundMgr) GetFromReturnAccount(returnAccount string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`return_account` = ?", returnAccount).Find(&results).Error

	return
}

// GetBatchFromReturnAccount 批量查找 退款账户
func (obj *RefundMgr) GetBatchFromReturnAccount(returnAccounts []string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`return_account` IN (?)", returnAccounts).Find(&results).Error

	return
}

// GetFromBankName 通过bank_name获取内容 银行名称
func (obj *RefundMgr) GetFromBankName(bankName string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`bank_name` = ?", bankName).Find(&results).Error

	return
}

// GetBatchFromBankName 批量查找 银行名称
func (obj *RefundMgr) GetBatchFromBankName(bankNames []string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`bank_name` IN (?)", bankNames).Find(&results).Error

	return
}

// GetFromBankAccountNumber 通过bank_account_number获取内容 银行账号
func (obj *RefundMgr) GetFromBankAccountNumber(bankAccountNumber string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`bank_account_number` = ?", bankAccountNumber).Find(&results).Error

	return
}

// GetBatchFromBankAccountNumber 批量查找 银行账号
func (obj *RefundMgr) GetBatchFromBankAccountNumber(bankAccountNumbers []string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`bank_account_number` IN (?)", bankAccountNumbers).Find(&results).Error

	return
}

// GetFromBankAccountName 通过bank_account_name获取内容 银行开户名
func (obj *RefundMgr) GetFromBankAccountName(bankAccountName string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`bank_account_name` = ?", bankAccountName).Find(&results).Error

	return
}

// GetBatchFromBankAccountName 批量查找 银行开户名
func (obj *RefundMgr) GetBatchFromBankAccountName(bankAccountNames []string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`bank_account_name` IN (?)", bankAccountNames).Find(&results).Error

	return
}

// GetFromBankDepositName 通过bank_deposit_name获取内容 银行开户行
func (obj *RefundMgr) GetFromBankDepositName(bankDepositName string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`bank_deposit_name` = ?", bankDepositName).Find(&results).Error

	return
}

// GetBatchFromBankDepositName 批量查找 银行开户行
func (obj *RefundMgr) GetBatchFromBankDepositName(bankDepositNames []string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`bank_deposit_name` IN (?)", bankDepositNames).Find(&results).Error

	return
}

// GetFromPayOrderNo 通过pay_order_no获取内容 支付结果交易号
func (obj *RefundMgr) GetFromPayOrderNo(payOrderNo string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`pay_order_no` = ?", payOrderNo).Find(&results).Error

	return
}

// GetBatchFromPayOrderNo 批量查找 支付结果交易号
func (obj *RefundMgr) GetBatchFromPayOrderNo(payOrderNos []string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`pay_order_no` IN (?)", payOrderNos).Find(&results).Error

	return
}

// GetFromPaymentType 通过payment_type获取内容 订单付款类型 ONLINE：在线支付，COD：货到付款
func (obj *RefundMgr) GetFromPaymentType(paymentType string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`payment_type` = ?", paymentType).Find(&results).Error

	return
}

// GetBatchFromPaymentType 批量查找 订单付款类型 ONLINE：在线支付，COD：货到付款
func (obj *RefundMgr) GetBatchFromPaymentType(paymentTypes []string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`payment_type` IN (?)", paymentTypes).Find(&results).Error

	return
}

// GetFromRefundFailReason 通过refund_fail_reason获取内容 退款失败原因
func (obj *RefundMgr) GetFromRefundFailReason(refundFailReason string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`refund_fail_reason` = ?", refundFailReason).Find(&results).Error

	return
}

// GetBatchFromRefundFailReason 批量查找 退款失败原因
func (obj *RefundMgr) GetBatchFromRefundFailReason(refundFailReasons []string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`refund_fail_reason` IN (?)", refundFailReasons).Find(&results).Error

	return
}

// GetFromRefundTime 通过refund_time获取内容 退款时间
func (obj *RefundMgr) GetFromRefundTime(refundTime int64) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`refund_time` = ?", refundTime).Find(&results).Error

	return
}

// GetBatchFromRefundTime 批量查找 退款时间
func (obj *RefundMgr) GetBatchFromRefundTime(refundTimes []int64) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`refund_time` IN (?)", refundTimes).Find(&results).Error

	return
}

// GetFromMobile 通过mobile获取内容 手机号
func (obj *RefundMgr) GetFromMobile(mobile string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`mobile` = ?", mobile).Find(&results).Error

	return
}

// GetBatchFromMobile 批量查找 手机号
func (obj *RefundMgr) GetBatchFromMobile(mobiles []string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`mobile` IN (?)", mobiles).Find(&results).Error

	return
}

// GetFromAgreePrice 通过agree_price获取内容 商家同意的退款金额
func (obj *RefundMgr) GetFromAgreePrice(agreePrice float64) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`agree_price` = ?", agreePrice).Find(&results).Error

	return
}

// GetBatchFromAgreePrice 批量查找 商家同意的退款金额
func (obj *RefundMgr) GetBatchFromAgreePrice(agreePrices []float64) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`agree_price` IN (?)", agreePrices).Find(&results).Error

	return
}

// GetFromActualPrice 通过actual_price获取内容 实际退款金额
func (obj *RefundMgr) GetFromActualPrice(actualPrice float64) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`actual_price` = ?", actualPrice).Find(&results).Error

	return
}

// GetBatchFromActualPrice 批量查找 实际退款金额
func (obj *RefundMgr) GetBatchFromActualPrice(actualPrices []float64) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`actual_price` IN (?)", actualPrices).Find(&results).Error

	return
}

// GetFromGoodsJSON 通过goods_json获取内容 售后商品信息json
func (obj *RefundMgr) GetFromGoodsJSON(goodsJSON string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`goods_json` = ?", goodsJSON).Find(&results).Error

	return
}

// GetBatchFromGoodsJSON 批量查找 售后商品信息json
func (obj *RefundMgr) GetBatchFromGoodsJSON(goodsJSONs []string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`goods_json` IN (?)", goodsJSONs).Find(&results).Error

	return
}

// GetFromDisabled 通过disabled获取内容 删除状态 DELETED：已删除 NORMAL：正常
func (obj *RefundMgr) GetFromDisabled(disabled string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`disabled` = ?", disabled).Find(&results).Error

	return
}

// GetBatchFromDisabled 批量查找 删除状态 DELETED：已删除 NORMAL：正常
func (obj *RefundMgr) GetBatchFromDisabled(disableds []string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`disabled` IN (?)", disableds).Find(&results).Error

	return
}

// GetFromCreateChannel 通过create_channel获取内容 售后服务单创建渠道 NORMAL：正常渠道创建，PINTUAN：拼团失败自动创建
func (obj *RefundMgr) GetFromCreateChannel(createChannel string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`create_channel` = ?", createChannel).Find(&results).Error

	return
}

// GetBatchFromCreateChannel 批量查找 售后服务单创建渠道 NORMAL：正常渠道创建，PINTUAN：拼团失败自动创建
func (obj *RefundMgr) GetBatchFromCreateChannel(createChannels []string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`create_channel` IN (?)", createChannels).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *RefundMgr) FetchByPrimaryKey(id int) (result models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByEsIndexID  获取多个内容
func (obj *RefundMgr) FetchIndexByEsIndexID(id int) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`id` = ?", id).Find(&results).Error

	return
}

// FetchIndexByEsIndexSn  获取多个内容
func (obj *RefundMgr) FetchIndexByEsIndexSn(sn string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`sn` = ?", sn).Find(&results).Error

	return
}

// FetchIndexByEsMemberID  获取多个内容
func (obj *RefundMgr) FetchIndexByEsMemberID(memberID int) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// FetchIndexByEsSellerID  获取多个内容
func (obj *RefundMgr) FetchIndexByEsSellerID(sellerID int) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// FetchIndexByEsIndexOrderSn  获取多个内容
func (obj *RefundMgr) FetchIndexByEsIndexOrderSn(orderSn string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}

// FetchIndexByEsIndexRefundStatus  获取多个内容
func (obj *RefundMgr) FetchIndexByEsIndexRefundStatus(refundStatus string) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`refund_status` = ?", refundStatus).Find(&results).Error

	return
}

// FetchIndexByEsIndexCreateTime  获取多个内容
func (obj *RefundMgr) FetchIndexByEsIndexCreateTime(createTime int64) (results []*models.EsRefund, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRefund{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}
