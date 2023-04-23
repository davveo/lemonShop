package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsBillItemMgr struct {
	*_BaseMgr
}

// EsBillItemMgr open func
func EsBillItemMgr(db *gorm.DB) *_EsBillItemMgr {
	if db == nil {
		panic(fmt.Errorf("EsBillItemMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsBillItemMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_bill_item"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsBillItemMgr) GetTableName() string {
	return "es_bill_item"
}

// Reset 重置gorm会话
func (obj *_EsBillItemMgr) Reset() *_EsBillItemMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsBillItemMgr) Get() (result models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsBillItemMgr) Gets() (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsBillItemMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_EsBillItemMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderSn order_sn获取 订单编号
func (obj *_EsBillItemMgr) WithOrderSn(orderSn string) Option {
	return optionFunc(func(o *options) { o.query["order_sn"] = orderSn })
}

// WithPrice price获取 订单价格
func (obj *_EsBillItemMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithDiscountPrice discount_price获取 优惠价格
func (obj *_EsBillItemMgr) WithDiscountPrice(discountPrice float64) Option {
	return optionFunc(func(o *options) { o.query["discount_price"] = discountPrice })
}

// WithItemType item_type获取 单项类型 收款/退款
func (obj *_EsBillItemMgr) WithItemType(itemType string) Option {
	return optionFunc(func(o *options) { o.query["item_type"] = itemType })
}

// WithAddTime add_time获取 加入时间
func (obj *_EsBillItemMgr) WithAddTime(addTime int) Option {
	return optionFunc(func(o *options) { o.query["add_time"] = addTime })
}

// WithBillID bill_id获取 所属账单id
func (obj *_EsBillItemMgr) WithBillID(billID int) Option {
	return optionFunc(func(o *options) { o.query["bill_id"] = billID })
}

// WithMemberID member_id获取 会员id
func (obj *_EsBillItemMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithStatus status获取 状态
func (obj *_EsBillItemMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithSellerID seller_id获取 店铺id
func (obj *_EsBillItemMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithOrderTime order_time获取 下单时间
func (obj *_EsBillItemMgr) WithOrderTime(orderTime int64) Option {
	return optionFunc(func(o *options) { o.query["order_time"] = orderTime })
}

// WithRefundSn refund_sn获取 退款单号
func (obj *_EsBillItemMgr) WithRefundSn(refundSn string) Option {
	return optionFunc(func(o *options) { o.query["refund_sn"] = refundSn })
}

// WithMemberName member_name获取 会员名称
func (obj *_EsBillItemMgr) WithMemberName(memberName string) Option {
	return optionFunc(func(o *options) { o.query["member_name"] = memberName })
}

// WithShipName ship_name获取 收货人
func (obj *_EsBillItemMgr) WithShipName(shipName string) Option {
	return optionFunc(func(o *options) { o.query["ship_name"] = shipName })
}

// WithPaymentType payment_type获取 支付方式
func (obj *_EsBillItemMgr) WithPaymentType(paymentType string) Option {
	return optionFunc(func(o *options) { o.query["payment_type"] = paymentType })
}

// WithRefundTime refund_time获取 退货时间
func (obj *_EsBillItemMgr) WithRefundTime(refundTime int64) Option {
	return optionFunc(func(o *options) { o.query["refund_time"] = refundTime })
}

// WithSiteCouponPrice site_coupon_price获取 使用平台优惠券金额
func (obj *_EsBillItemMgr) WithSiteCouponPrice(siteCouponPrice float64) Option {
	return optionFunc(func(o *options) { o.query["site_coupon_price"] = siteCouponPrice })
}

// WithCouponCommission coupon_commission获取 优惠券佣金比例
func (obj *_EsBillItemMgr) WithCouponCommission(couponCommission int) Option {
	return optionFunc(func(o *options) { o.query["coupon_commission"] = couponCommission })
}

// GetByOption 功能选项模式获取
func (obj *_EsBillItemMgr) GetByOption(opts ...Option) (result models.EsBillItem, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsBillItemMgr) GetByOptions(opts ...Option) (results []*models.EsBillItem, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsBillItemMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsBillItem, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where(options.query)
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
func (obj *_EsBillItemMgr) GetFromID(id int) (result models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键id
func (obj *_EsBillItemMgr) GetBatchFromID(ids []int) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderSn 通过order_sn获取内容 订单编号
func (obj *_EsBillItemMgr) GetFromOrderSn(orderSn string) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}

// GetBatchFromOrderSn 批量查找 订单编号
func (obj *_EsBillItemMgr) GetBatchFromOrderSn(orderSns []string) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`order_sn` IN (?)", orderSns).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 订单价格
func (obj *_EsBillItemMgr) GetFromPrice(price float64) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 订单价格
func (obj *_EsBillItemMgr) GetBatchFromPrice(prices []float64) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromDiscountPrice 通过discount_price获取内容 优惠价格
func (obj *_EsBillItemMgr) GetFromDiscountPrice(discountPrice float64) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`discount_price` = ?", discountPrice).Find(&results).Error

	return
}

// GetBatchFromDiscountPrice 批量查找 优惠价格
func (obj *_EsBillItemMgr) GetBatchFromDiscountPrice(discountPrices []float64) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`discount_price` IN (?)", discountPrices).Find(&results).Error

	return
}

// GetFromItemType 通过item_type获取内容 单项类型 收款/退款
func (obj *_EsBillItemMgr) GetFromItemType(itemType string) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`item_type` = ?", itemType).Find(&results).Error

	return
}

// GetBatchFromItemType 批量查找 单项类型 收款/退款
func (obj *_EsBillItemMgr) GetBatchFromItemType(itemTypes []string) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`item_type` IN (?)", itemTypes).Find(&results).Error

	return
}

// GetFromAddTime 通过add_time获取内容 加入时间
func (obj *_EsBillItemMgr) GetFromAddTime(addTime int) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`add_time` = ?", addTime).Find(&results).Error

	return
}

// GetBatchFromAddTime 批量查找 加入时间
func (obj *_EsBillItemMgr) GetBatchFromAddTime(addTimes []int) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`add_time` IN (?)", addTimes).Find(&results).Error

	return
}

// GetFromBillID 通过bill_id获取内容 所属账单id
func (obj *_EsBillItemMgr) GetFromBillID(billID int) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`bill_id` = ?", billID).Find(&results).Error

	return
}

// GetBatchFromBillID 批量查找 所属账单id
func (obj *_EsBillItemMgr) GetBatchFromBillID(billIDs []int) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`bill_id` IN (?)", billIDs).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *_EsBillItemMgr) GetFromMemberID(memberID int) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *_EsBillItemMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态
func (obj *_EsBillItemMgr) GetFromStatus(status int) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态
func (obj *_EsBillItemMgr) GetBatchFromStatus(statuss []int) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 店铺id
func (obj *_EsBillItemMgr) GetFromSellerID(sellerID int) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 店铺id
func (obj *_EsBillItemMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromOrderTime 通过order_time获取内容 下单时间
func (obj *_EsBillItemMgr) GetFromOrderTime(orderTime int64) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`order_time` = ?", orderTime).Find(&results).Error

	return
}

// GetBatchFromOrderTime 批量查找 下单时间
func (obj *_EsBillItemMgr) GetBatchFromOrderTime(orderTimes []int64) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`order_time` IN (?)", orderTimes).Find(&results).Error

	return
}

// GetFromRefundSn 通过refund_sn获取内容 退款单号
func (obj *_EsBillItemMgr) GetFromRefundSn(refundSn string) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`refund_sn` = ?", refundSn).Find(&results).Error

	return
}

// GetBatchFromRefundSn 批量查找 退款单号
func (obj *_EsBillItemMgr) GetBatchFromRefundSn(refundSns []string) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`refund_sn` IN (?)", refundSns).Find(&results).Error

	return
}

// GetFromMemberName 通过member_name获取内容 会员名称
func (obj *_EsBillItemMgr) GetFromMemberName(memberName string) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`member_name` = ?", memberName).Find(&results).Error

	return
}

// GetBatchFromMemberName 批量查找 会员名称
func (obj *_EsBillItemMgr) GetBatchFromMemberName(memberNames []string) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`member_name` IN (?)", memberNames).Find(&results).Error

	return
}

// GetFromShipName 通过ship_name获取内容 收货人
func (obj *_EsBillItemMgr) GetFromShipName(shipName string) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`ship_name` = ?", shipName).Find(&results).Error

	return
}

// GetBatchFromShipName 批量查找 收货人
func (obj *_EsBillItemMgr) GetBatchFromShipName(shipNames []string) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`ship_name` IN (?)", shipNames).Find(&results).Error

	return
}

// GetFromPaymentType 通过payment_type获取内容 支付方式
func (obj *_EsBillItemMgr) GetFromPaymentType(paymentType string) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`payment_type` = ?", paymentType).Find(&results).Error

	return
}

// GetBatchFromPaymentType 批量查找 支付方式
func (obj *_EsBillItemMgr) GetBatchFromPaymentType(paymentTypes []string) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`payment_type` IN (?)", paymentTypes).Find(&results).Error

	return
}

// GetFromRefundTime 通过refund_time获取内容 退货时间
func (obj *_EsBillItemMgr) GetFromRefundTime(refundTime int64) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`refund_time` = ?", refundTime).Find(&results).Error

	return
}

// GetBatchFromRefundTime 批量查找 退货时间
func (obj *_EsBillItemMgr) GetBatchFromRefundTime(refundTimes []int64) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`refund_time` IN (?)", refundTimes).Find(&results).Error

	return
}

// GetFromSiteCouponPrice 通过site_coupon_price获取内容 使用平台优惠券金额
func (obj *_EsBillItemMgr) GetFromSiteCouponPrice(siteCouponPrice float64) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`site_coupon_price` = ?", siteCouponPrice).Find(&results).Error

	return
}

// GetBatchFromSiteCouponPrice 批量查找 使用平台优惠券金额
func (obj *_EsBillItemMgr) GetBatchFromSiteCouponPrice(siteCouponPrices []float64) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`site_coupon_price` IN (?)", siteCouponPrices).Find(&results).Error

	return
}

// GetFromCouponCommission 通过coupon_commission获取内容 优惠券佣金比例
func (obj *_EsBillItemMgr) GetFromCouponCommission(couponCommission int) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`coupon_commission` = ?", couponCommission).Find(&results).Error

	return
}

// GetBatchFromCouponCommission 批量查找 优惠券佣金比例
func (obj *_EsBillItemMgr) GetBatchFromCouponCommission(couponCommissions []int) (results []*models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`coupon_commission` IN (?)", couponCommissions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsBillItemMgr) FetchByPrimaryKey(id int) (result models.EsBillItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillItem{}).Where("`id` = ?", id).First(&result).Error

	return
}
