package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsBillMgr struct {
	*_BaseMgr
}

// EsBillMgr open func
func EsBillMgr(db *gorm.DB) *_EsBillMgr {
	if db == nil {
		panic(fmt.Errorf("EsBillMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsBillMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_bill"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsBillMgr) GetTableName() string {
	return "es_bill"
}

// Reset 重置gorm会话
func (obj *_EsBillMgr) Reset() *_EsBillMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsBillMgr) Get() (result models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsBillMgr) Gets() (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsBillMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithBillID bill_id获取 主键ID
func (obj *_EsBillMgr) WithBillID(billID int) Option {
	return optionFunc(func(o *options) { o.query["bill_id"] = billID })
}

// WithBillSn bill_sn获取 结算单编号
func (obj *_EsBillMgr) WithBillSn(billSn string) Option {
	return optionFunc(func(o *options) { o.query["bill_sn"] = billSn })
}

// WithStartTime start_time获取 结算开始时间
func (obj *_EsBillMgr) WithStartTime(startTime int64) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取 结算结束时间
func (obj *_EsBillMgr) WithEndTime(endTime int64) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithPrice price获取 结算总金额
func (obj *_EsBillMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithCommiPrice commi_price获取 佣金
func (obj *_EsBillMgr) WithCommiPrice(commiPrice float64) Option {
	return optionFunc(func(o *options) { o.query["commi_price"] = commiPrice })
}

// WithDiscountPrice discount_price获取 优惠金额
func (obj *_EsBillMgr) WithDiscountPrice(discountPrice float64) Option {
	return optionFunc(func(o *options) { o.query["discount_price"] = discountPrice })
}

// WithStatus status获取 状态
func (obj *_EsBillMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithBillType bill_type获取 账单类型
func (obj *_EsBillMgr) WithBillType(billType int) Option {
	return optionFunc(func(o *options) { o.query["bill_type"] = billType })
}

// WithSellerID seller_id获取 店铺id
func (obj *_EsBillMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithPayTime pay_time获取 付款时间
func (obj *_EsBillMgr) WithPayTime(payTime int64) Option {
	return optionFunc(func(o *options) { o.query["pay_time"] = payTime })
}

// WithCreateTime create_time获取 出账日期
func (obj *_EsBillMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithBillPrice bill_price获取 结算金额
func (obj *_EsBillMgr) WithBillPrice(billPrice float64) Option {
	return optionFunc(func(o *options) { o.query["bill_price"] = billPrice })
}

// WithRefundPrice refund_price获取 在线支付退款金额
func (obj *_EsBillMgr) WithRefundPrice(refundPrice float64) Option {
	return optionFunc(func(o *options) { o.query["refund_price"] = refundPrice })
}

// WithRefundCommiPrice refund_commi_price获取 退还佣金金额
func (obj *_EsBillMgr) WithRefundCommiPrice(refundCommiPrice float64) Option {
	return optionFunc(func(o *options) { o.query["refund_commi_price"] = refundCommiPrice })
}

// WithSn sn获取 账单号
func (obj *_EsBillMgr) WithSn(sn string) Option {
	return optionFunc(func(o *options) { o.query["sn"] = sn })
}

// WithShopName shop_name获取 店铺名称
func (obj *_EsBillMgr) WithShopName(shopName string) Option {
	return optionFunc(func(o *options) { o.query["shop_name"] = shopName })
}

// WithBankAccountName bank_account_name获取 银行开户名
func (obj *_EsBillMgr) WithBankAccountName(bankAccountName string) Option {
	return optionFunc(func(o *options) { o.query["bank_account_name"] = bankAccountName })
}

// WithBankAccountNumber bank_account_number获取 公司银行账号
func (obj *_EsBillMgr) WithBankAccountNumber(bankAccountNumber string) Option {
	return optionFunc(func(o *options) { o.query["bank_account_number"] = bankAccountNumber })
}

// WithBankName bank_name获取 开户银行支行名称
func (obj *_EsBillMgr) WithBankName(bankName string) Option {
	return optionFunc(func(o *options) { o.query["bank_name"] = bankName })
}

// WithBankCode bank_code获取 支行联行号
func (obj *_EsBillMgr) WithBankCode(bankCode string) Option {
	return optionFunc(func(o *options) { o.query["bank_code"] = bankCode })
}

// WithBankAddress bank_address获取 开户银行地址
func (obj *_EsBillMgr) WithBankAddress(bankAddress string) Option {
	return optionFunc(func(o *options) { o.query["bank_address"] = bankAddress })
}

// WithCodPrice cod_price获取 货到付款金额
func (obj *_EsBillMgr) WithCodPrice(codPrice float64) Option {
	return optionFunc(func(o *options) { o.query["cod_price"] = codPrice })
}

// WithCodRefundPrice cod_refund_price获取 货到付款退款金额
func (obj *_EsBillMgr) WithCodRefundPrice(codRefundPrice float64) Option {
	return optionFunc(func(o *options) { o.query["cod_refund_price"] = codRefundPrice })
}

// WithDistributionRebate distribution_rebate获取 分销返现支出
func (obj *_EsBillMgr) WithDistributionRebate(distributionRebate float64) Option {
	return optionFunc(func(o *options) { o.query["distribution_rebate"] = distributionRebate })
}

// WithDistributionReturnRebate distribution_return_rebate获取 分销返现支出返还
func (obj *_EsBillMgr) WithDistributionReturnRebate(distributionReturnRebate float64) Option {
	return optionFunc(func(o *options) { o.query["distribution_return_rebate"] = distributionReturnRebate })
}

// WithSiteCouponCommi site_coupon_commi获取 平台优惠券佣金
func (obj *_EsBillMgr) WithSiteCouponCommi(siteCouponCommi float64) Option {
	return optionFunc(func(o *options) { o.query["site_coupon_commi"] = siteCouponCommi })
}

// WithOrderTotalPrice order_total_price获取 结算周期内订单付款总金额
func (obj *_EsBillMgr) WithOrderTotalPrice(orderTotalPrice float64) Option {
	return optionFunc(func(o *options) { o.query["order_total_price"] = orderTotalPrice })
}

// WithRefundTotalPrice refund_total_price获取 结算周期内订单退款总金额
func (obj *_EsBillMgr) WithRefundTotalPrice(refundTotalPrice float64) Option {
	return optionFunc(func(o *options) { o.query["refund_total_price"] = refundTotalPrice })
}

// GetByOption 功能选项模式获取
func (obj *_EsBillMgr) GetByOption(opts ...Option) (result models.EsBill, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsBillMgr) GetByOptions(opts ...Option) (results []*models.EsBill, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsBillMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsBill, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where(options.query)
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

// GetFromBillID 通过bill_id获取内容 主键ID
func (obj *_EsBillMgr) GetFromBillID(billID int) (result models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`bill_id` = ?", billID).First(&result).Error

	return
}

// GetBatchFromBillID 批量查找 主键ID
func (obj *_EsBillMgr) GetBatchFromBillID(billIDs []int) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`bill_id` IN (?)", billIDs).Find(&results).Error

	return
}

// GetFromBillSn 通过bill_sn获取内容 结算单编号
func (obj *_EsBillMgr) GetFromBillSn(billSn string) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`bill_sn` = ?", billSn).Find(&results).Error

	return
}

// GetBatchFromBillSn 批量查找 结算单编号
func (obj *_EsBillMgr) GetBatchFromBillSn(billSns []string) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`bill_sn` IN (?)", billSns).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容 结算开始时间
func (obj *_EsBillMgr) GetFromStartTime(startTime int64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找 结算开始时间
func (obj *_EsBillMgr) GetBatchFromStartTime(startTimes []int64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容 结算结束时间
func (obj *_EsBillMgr) GetFromEndTime(endTime int64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找 结算结束时间
func (obj *_EsBillMgr) GetBatchFromEndTime(endTimes []int64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 结算总金额
func (obj *_EsBillMgr) GetFromPrice(price float64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 结算总金额
func (obj *_EsBillMgr) GetBatchFromPrice(prices []float64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromCommiPrice 通过commi_price获取内容 佣金
func (obj *_EsBillMgr) GetFromCommiPrice(commiPrice float64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`commi_price` = ?", commiPrice).Find(&results).Error

	return
}

// GetBatchFromCommiPrice 批量查找 佣金
func (obj *_EsBillMgr) GetBatchFromCommiPrice(commiPrices []float64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`commi_price` IN (?)", commiPrices).Find(&results).Error

	return
}

// GetFromDiscountPrice 通过discount_price获取内容 优惠金额
func (obj *_EsBillMgr) GetFromDiscountPrice(discountPrice float64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`discount_price` = ?", discountPrice).Find(&results).Error

	return
}

// GetBatchFromDiscountPrice 批量查找 优惠金额
func (obj *_EsBillMgr) GetBatchFromDiscountPrice(discountPrices []float64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`discount_price` IN (?)", discountPrices).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态
func (obj *_EsBillMgr) GetFromStatus(status string) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态
func (obj *_EsBillMgr) GetBatchFromStatus(statuss []string) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromBillType 通过bill_type获取内容 账单类型
func (obj *_EsBillMgr) GetFromBillType(billType int) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`bill_type` = ?", billType).Find(&results).Error

	return
}

// GetBatchFromBillType 批量查找 账单类型
func (obj *_EsBillMgr) GetBatchFromBillType(billTypes []int) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`bill_type` IN (?)", billTypes).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 店铺id
func (obj *_EsBillMgr) GetFromSellerID(sellerID int) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 店铺id
func (obj *_EsBillMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromPayTime 通过pay_time获取内容 付款时间
func (obj *_EsBillMgr) GetFromPayTime(payTime int64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`pay_time` = ?", payTime).Find(&results).Error

	return
}

// GetBatchFromPayTime 批量查找 付款时间
func (obj *_EsBillMgr) GetBatchFromPayTime(payTimes []int64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`pay_time` IN (?)", payTimes).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 出账日期
func (obj *_EsBillMgr) GetFromCreateTime(createTime int64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 出账日期
func (obj *_EsBillMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromBillPrice 通过bill_price获取内容 结算金额
func (obj *_EsBillMgr) GetFromBillPrice(billPrice float64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`bill_price` = ?", billPrice).Find(&results).Error

	return
}

// GetBatchFromBillPrice 批量查找 结算金额
func (obj *_EsBillMgr) GetBatchFromBillPrice(billPrices []float64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`bill_price` IN (?)", billPrices).Find(&results).Error

	return
}

// GetFromRefundPrice 通过refund_price获取内容 在线支付退款金额
func (obj *_EsBillMgr) GetFromRefundPrice(refundPrice float64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`refund_price` = ?", refundPrice).Find(&results).Error

	return
}

// GetBatchFromRefundPrice 批量查找 在线支付退款金额
func (obj *_EsBillMgr) GetBatchFromRefundPrice(refundPrices []float64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`refund_price` IN (?)", refundPrices).Find(&results).Error

	return
}

// GetFromRefundCommiPrice 通过refund_commi_price获取内容 退还佣金金额
func (obj *_EsBillMgr) GetFromRefundCommiPrice(refundCommiPrice float64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`refund_commi_price` = ?", refundCommiPrice).Find(&results).Error

	return
}

// GetBatchFromRefundCommiPrice 批量查找 退还佣金金额
func (obj *_EsBillMgr) GetBatchFromRefundCommiPrice(refundCommiPrices []float64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`refund_commi_price` IN (?)", refundCommiPrices).Find(&results).Error

	return
}

// GetFromSn 通过sn获取内容 账单号
func (obj *_EsBillMgr) GetFromSn(sn string) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`sn` = ?", sn).Find(&results).Error

	return
}

// GetBatchFromSn 批量查找 账单号
func (obj *_EsBillMgr) GetBatchFromSn(sns []string) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`sn` IN (?)", sns).Find(&results).Error

	return
}

// GetFromShopName 通过shop_name获取内容 店铺名称
func (obj *_EsBillMgr) GetFromShopName(shopName string) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`shop_name` = ?", shopName).Find(&results).Error

	return
}

// GetBatchFromShopName 批量查找 店铺名称
func (obj *_EsBillMgr) GetBatchFromShopName(shopNames []string) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`shop_name` IN (?)", shopNames).Find(&results).Error

	return
}

// GetFromBankAccountName 通过bank_account_name获取内容 银行开户名
func (obj *_EsBillMgr) GetFromBankAccountName(bankAccountName string) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`bank_account_name` = ?", bankAccountName).Find(&results).Error

	return
}

// GetBatchFromBankAccountName 批量查找 银行开户名
func (obj *_EsBillMgr) GetBatchFromBankAccountName(bankAccountNames []string) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`bank_account_name` IN (?)", bankAccountNames).Find(&results).Error

	return
}

// GetFromBankAccountNumber 通过bank_account_number获取内容 公司银行账号
func (obj *_EsBillMgr) GetFromBankAccountNumber(bankAccountNumber string) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`bank_account_number` = ?", bankAccountNumber).Find(&results).Error

	return
}

// GetBatchFromBankAccountNumber 批量查找 公司银行账号
func (obj *_EsBillMgr) GetBatchFromBankAccountNumber(bankAccountNumbers []string) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`bank_account_number` IN (?)", bankAccountNumbers).Find(&results).Error

	return
}

// GetFromBankName 通过bank_name获取内容 开户银行支行名称
func (obj *_EsBillMgr) GetFromBankName(bankName string) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`bank_name` = ?", bankName).Find(&results).Error

	return
}

// GetBatchFromBankName 批量查找 开户银行支行名称
func (obj *_EsBillMgr) GetBatchFromBankName(bankNames []string) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`bank_name` IN (?)", bankNames).Find(&results).Error

	return
}

// GetFromBankCode 通过bank_code获取内容 支行联行号
func (obj *_EsBillMgr) GetFromBankCode(bankCode string) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`bank_code` = ?", bankCode).Find(&results).Error

	return
}

// GetBatchFromBankCode 批量查找 支行联行号
func (obj *_EsBillMgr) GetBatchFromBankCode(bankCodes []string) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`bank_code` IN (?)", bankCodes).Find(&results).Error

	return
}

// GetFromBankAddress 通过bank_address获取内容 开户银行地址
func (obj *_EsBillMgr) GetFromBankAddress(bankAddress string) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`bank_address` = ?", bankAddress).Find(&results).Error

	return
}

// GetBatchFromBankAddress 批量查找 开户银行地址
func (obj *_EsBillMgr) GetBatchFromBankAddress(bankAddresss []string) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`bank_address` IN (?)", bankAddresss).Find(&results).Error

	return
}

// GetFromCodPrice 通过cod_price获取内容 货到付款金额
func (obj *_EsBillMgr) GetFromCodPrice(codPrice float64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`cod_price` = ?", codPrice).Find(&results).Error

	return
}

// GetBatchFromCodPrice 批量查找 货到付款金额
func (obj *_EsBillMgr) GetBatchFromCodPrice(codPrices []float64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`cod_price` IN (?)", codPrices).Find(&results).Error

	return
}

// GetFromCodRefundPrice 通过cod_refund_price获取内容 货到付款退款金额
func (obj *_EsBillMgr) GetFromCodRefundPrice(codRefundPrice float64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`cod_refund_price` = ?", codRefundPrice).Find(&results).Error

	return
}

// GetBatchFromCodRefundPrice 批量查找 货到付款退款金额
func (obj *_EsBillMgr) GetBatchFromCodRefundPrice(codRefundPrices []float64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`cod_refund_price` IN (?)", codRefundPrices).Find(&results).Error

	return
}

// GetFromDistributionRebate 通过distribution_rebate获取内容 分销返现支出
func (obj *_EsBillMgr) GetFromDistributionRebate(distributionRebate float64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`distribution_rebate` = ?", distributionRebate).Find(&results).Error

	return
}

// GetBatchFromDistributionRebate 批量查找 分销返现支出
func (obj *_EsBillMgr) GetBatchFromDistributionRebate(distributionRebates []float64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`distribution_rebate` IN (?)", distributionRebates).Find(&results).Error

	return
}

// GetFromDistributionReturnRebate 通过distribution_return_rebate获取内容 分销返现支出返还
func (obj *_EsBillMgr) GetFromDistributionReturnRebate(distributionReturnRebate float64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`distribution_return_rebate` = ?", distributionReturnRebate).Find(&results).Error

	return
}

// GetBatchFromDistributionReturnRebate 批量查找 分销返现支出返还
func (obj *_EsBillMgr) GetBatchFromDistributionReturnRebate(distributionReturnRebates []float64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`distribution_return_rebate` IN (?)", distributionReturnRebates).Find(&results).Error

	return
}

// GetFromSiteCouponCommi 通过site_coupon_commi获取内容 平台优惠券佣金
func (obj *_EsBillMgr) GetFromSiteCouponCommi(siteCouponCommi float64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`site_coupon_commi` = ?", siteCouponCommi).Find(&results).Error

	return
}

// GetBatchFromSiteCouponCommi 批量查找 平台优惠券佣金
func (obj *_EsBillMgr) GetBatchFromSiteCouponCommi(siteCouponCommis []float64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`site_coupon_commi` IN (?)", siteCouponCommis).Find(&results).Error

	return
}

// GetFromOrderTotalPrice 通过order_total_price获取内容 结算周期内订单付款总金额
func (obj *_EsBillMgr) GetFromOrderTotalPrice(orderTotalPrice float64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`order_total_price` = ?", orderTotalPrice).Find(&results).Error

	return
}

// GetBatchFromOrderTotalPrice 批量查找 结算周期内订单付款总金额
func (obj *_EsBillMgr) GetBatchFromOrderTotalPrice(orderTotalPrices []float64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`order_total_price` IN (?)", orderTotalPrices).Find(&results).Error

	return
}

// GetFromRefundTotalPrice 通过refund_total_price获取内容 结算周期内订单退款总金额
func (obj *_EsBillMgr) GetFromRefundTotalPrice(refundTotalPrice float64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`refund_total_price` = ?", refundTotalPrice).Find(&results).Error

	return
}

// GetBatchFromRefundTotalPrice 批量查找 结算周期内订单退款总金额
func (obj *_EsBillMgr) GetBatchFromRefundTotalPrice(refundTotalPrices []float64) (results []*models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`refund_total_price` IN (?)", refundTotalPrices).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsBillMgr) FetchByPrimaryKey(billID int) (result models.EsBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBill{}).Where("`bill_id` = ?", billID).First(&result).Error

	return
}
