package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsCouponMgr struct {
	*_BaseMgr
}

// EsCouponMgr open func
func EsCouponMgr(db *gorm.DB) *_EsCouponMgr {
	if db == nil {
		panic(fmt.Errorf("EsCouponMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsCouponMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_coupon"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsCouponMgr) GetTableName() string {
	return "es_coupon"
}

// Reset 重置gorm会话
func (obj *_EsCouponMgr) Reset() *_EsCouponMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsCouponMgr) Get() (result models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsCouponMgr) Gets() (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsCouponMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTitle title获取 优惠券标题
func (obj *_EsCouponMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithCouponID coupon_id获取 主键
func (obj *_EsCouponMgr) WithCouponID(couponID int) Option {
	return optionFunc(func(o *options) { o.query["coupon_id"] = couponID })
}

// WithCouponPrice coupon_price获取 优惠券面额
func (obj *_EsCouponMgr) WithCouponPrice(couponPrice float64) Option {
	return optionFunc(func(o *options) { o.query["coupon_price"] = couponPrice })
}

// WithCouponThresholdPrice coupon_threshold_price获取 优惠券门槛价格
func (obj *_EsCouponMgr) WithCouponThresholdPrice(couponThresholdPrice float64) Option {
	return optionFunc(func(o *options) { o.query["coupon_threshold_price"] = couponThresholdPrice })
}

// WithStartTime start_time获取 使用起始时间
func (obj *_EsCouponMgr) WithStartTime(startTime int64) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取 使用截止时间
func (obj *_EsCouponMgr) WithEndTime(endTime int64) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithCreateNum create_num获取 发行量
func (obj *_EsCouponMgr) WithCreateNum(createNum int) Option {
	return optionFunc(func(o *options) { o.query["create_num"] = createNum })
}

// WithLimitNum limit_num获取 每人限领数量
func (obj *_EsCouponMgr) WithLimitNum(limitNum int) Option {
	return optionFunc(func(o *options) { o.query["limit_num"] = limitNum })
}

// WithUsedNum used_num获取 已被使用的数量
func (obj *_EsCouponMgr) WithUsedNum(usedNum int) Option {
	return optionFunc(func(o *options) { o.query["used_num"] = usedNum })
}

// WithSellerID seller_id获取 店铺ID
func (obj *_EsCouponMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithReceivedNum received_num获取 已被领取的数量
func (obj *_EsCouponMgr) WithReceivedNum(receivedNum int) Option {
	return optionFunc(func(o *options) { o.query["received_num"] = receivedNum })
}

// WithSellerName seller_name获取 店铺名称
func (obj *_EsCouponMgr) WithSellerName(sellerName string) Option {
	return optionFunc(func(o *options) { o.query["seller_name"] = sellerName })
}

// WithType type获取 优惠券类型，分为免费领取和活动赠送
func (obj *_EsCouponMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithUseScope use_scope获取 使用范围，全品，分类，部分商品
func (obj *_EsCouponMgr) WithUseScope(useScope string) Option {
	return optionFunc(func(o *options) { o.query["use_scope"] = useScope })
}

// WithScopeID scope_id获取 全品或者商家优惠券时为0<br />分类时为分类id<br />部分商品时为商品ID集合
func (obj *_EsCouponMgr) WithScopeID(scopeID string) Option {
	return optionFunc(func(o *options) { o.query["scope_id"] = scopeID })
}

// WithShopCommission shop_commission获取 店铺承担比例，正整数
func (obj *_EsCouponMgr) WithShopCommission(shopCommission int) Option {
	return optionFunc(func(o *options) { o.query["shop_commission"] = shopCommission })
}

// WithScopeDescription scope_description获取 范围描述
func (obj *_EsCouponMgr) WithScopeDescription(scopeDescription string) Option {
	return optionFunc(func(o *options) { o.query["scope_description"] = scopeDescription })
}

// WithActivityDescription activity_description获取 活动说明
func (obj *_EsCouponMgr) WithActivityDescription(activityDescription string) Option {
	return optionFunc(func(o *options) { o.query["activity_description"] = activityDescription })
}

// GetByOption 功能选项模式获取
func (obj *_EsCouponMgr) GetByOption(opts ...Option) (result models.EsCoupon, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsCouponMgr) GetByOptions(opts ...Option) (results []*models.EsCoupon, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsCouponMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsCoupon, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where(options.query)
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

// GetFromTitle 通过title获取内容 优惠券标题
func (obj *_EsCouponMgr) GetFromTitle(title string) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 优惠券标题
func (obj *_EsCouponMgr) GetBatchFromTitle(titles []string) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromCouponID 通过coupon_id获取内容 主键
func (obj *_EsCouponMgr) GetFromCouponID(couponID int) (result models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`coupon_id` = ?", couponID).First(&result).Error

	return
}

// GetBatchFromCouponID 批量查找 主键
func (obj *_EsCouponMgr) GetBatchFromCouponID(couponIDs []int) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`coupon_id` IN (?)", couponIDs).Find(&results).Error

	return
}

// GetFromCouponPrice 通过coupon_price获取内容 优惠券面额
func (obj *_EsCouponMgr) GetFromCouponPrice(couponPrice float64) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`coupon_price` = ?", couponPrice).Find(&results).Error

	return
}

// GetBatchFromCouponPrice 批量查找 优惠券面额
func (obj *_EsCouponMgr) GetBatchFromCouponPrice(couponPrices []float64) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`coupon_price` IN (?)", couponPrices).Find(&results).Error

	return
}

// GetFromCouponThresholdPrice 通过coupon_threshold_price获取内容 优惠券门槛价格
func (obj *_EsCouponMgr) GetFromCouponThresholdPrice(couponThresholdPrice float64) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`coupon_threshold_price` = ?", couponThresholdPrice).Find(&results).Error

	return
}

// GetBatchFromCouponThresholdPrice 批量查找 优惠券门槛价格
func (obj *_EsCouponMgr) GetBatchFromCouponThresholdPrice(couponThresholdPrices []float64) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`coupon_threshold_price` IN (?)", couponThresholdPrices).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容 使用起始时间
func (obj *_EsCouponMgr) GetFromStartTime(startTime int64) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找 使用起始时间
func (obj *_EsCouponMgr) GetBatchFromStartTime(startTimes []int64) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容 使用截止时间
func (obj *_EsCouponMgr) GetFromEndTime(endTime int64) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找 使用截止时间
func (obj *_EsCouponMgr) GetBatchFromEndTime(endTimes []int64) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromCreateNum 通过create_num获取内容 发行量
func (obj *_EsCouponMgr) GetFromCreateNum(createNum int) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`create_num` = ?", createNum).Find(&results).Error

	return
}

// GetBatchFromCreateNum 批量查找 发行量
func (obj *_EsCouponMgr) GetBatchFromCreateNum(createNums []int) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`create_num` IN (?)", createNums).Find(&results).Error

	return
}

// GetFromLimitNum 通过limit_num获取内容 每人限领数量
func (obj *_EsCouponMgr) GetFromLimitNum(limitNum int) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`limit_num` = ?", limitNum).Find(&results).Error

	return
}

// GetBatchFromLimitNum 批量查找 每人限领数量
func (obj *_EsCouponMgr) GetBatchFromLimitNum(limitNums []int) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`limit_num` IN (?)", limitNums).Find(&results).Error

	return
}

// GetFromUsedNum 通过used_num获取内容 已被使用的数量
func (obj *_EsCouponMgr) GetFromUsedNum(usedNum int) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`used_num` = ?", usedNum).Find(&results).Error

	return
}

// GetBatchFromUsedNum 批量查找 已被使用的数量
func (obj *_EsCouponMgr) GetBatchFromUsedNum(usedNums []int) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`used_num` IN (?)", usedNums).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 店铺ID
func (obj *_EsCouponMgr) GetFromSellerID(sellerID int) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 店铺ID
func (obj *_EsCouponMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromReceivedNum 通过received_num获取内容 已被领取的数量
func (obj *_EsCouponMgr) GetFromReceivedNum(receivedNum int) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`received_num` = ?", receivedNum).Find(&results).Error

	return
}

// GetBatchFromReceivedNum 批量查找 已被领取的数量
func (obj *_EsCouponMgr) GetBatchFromReceivedNum(receivedNums []int) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`received_num` IN (?)", receivedNums).Find(&results).Error

	return
}

// GetFromSellerName 通过seller_name获取内容 店铺名称
func (obj *_EsCouponMgr) GetFromSellerName(sellerName string) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`seller_name` = ?", sellerName).Find(&results).Error

	return
}

// GetBatchFromSellerName 批量查找 店铺名称
func (obj *_EsCouponMgr) GetBatchFromSellerName(sellerNames []string) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`seller_name` IN (?)", sellerNames).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 优惠券类型，分为免费领取和活动赠送
func (obj *_EsCouponMgr) GetFromType(_type string) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 优惠券类型，分为免费领取和活动赠送
func (obj *_EsCouponMgr) GetBatchFromType(_types []string) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromUseScope 通过use_scope获取内容 使用范围，全品，分类，部分商品
func (obj *_EsCouponMgr) GetFromUseScope(useScope string) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`use_scope` = ?", useScope).Find(&results).Error

	return
}

// GetBatchFromUseScope 批量查找 使用范围，全品，分类，部分商品
func (obj *_EsCouponMgr) GetBatchFromUseScope(useScopes []string) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`use_scope` IN (?)", useScopes).Find(&results).Error

	return
}

// GetFromScopeID 通过scope_id获取内容 全品或者商家优惠券时为0<br />分类时为分类id<br />部分商品时为商品ID集合
func (obj *_EsCouponMgr) GetFromScopeID(scopeID string) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`scope_id` = ?", scopeID).Find(&results).Error

	return
}

// GetBatchFromScopeID 批量查找 全品或者商家优惠券时为0<br />分类时为分类id<br />部分商品时为商品ID集合
func (obj *_EsCouponMgr) GetBatchFromScopeID(scopeIDs []string) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`scope_id` IN (?)", scopeIDs).Find(&results).Error

	return
}

// GetFromShopCommission 通过shop_commission获取内容 店铺承担比例，正整数
func (obj *_EsCouponMgr) GetFromShopCommission(shopCommission int) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`shop_commission` = ?", shopCommission).Find(&results).Error

	return
}

// GetBatchFromShopCommission 批量查找 店铺承担比例，正整数
func (obj *_EsCouponMgr) GetBatchFromShopCommission(shopCommissions []int) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`shop_commission` IN (?)", shopCommissions).Find(&results).Error

	return
}

// GetFromScopeDescription 通过scope_description获取内容 范围描述
func (obj *_EsCouponMgr) GetFromScopeDescription(scopeDescription string) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`scope_description` = ?", scopeDescription).Find(&results).Error

	return
}

// GetBatchFromScopeDescription 批量查找 范围描述
func (obj *_EsCouponMgr) GetBatchFromScopeDescription(scopeDescriptions []string) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`scope_description` IN (?)", scopeDescriptions).Find(&results).Error

	return
}

// GetFromActivityDescription 通过activity_description获取内容 活动说明
func (obj *_EsCouponMgr) GetFromActivityDescription(activityDescription string) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`activity_description` = ?", activityDescription).Find(&results).Error

	return
}

// GetBatchFromActivityDescription 批量查找 活动说明
func (obj *_EsCouponMgr) GetBatchFromActivityDescription(activityDescriptions []string) (results []*models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`activity_description` IN (?)", activityDescriptions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsCouponMgr) FetchByPrimaryKey(couponID int) (result models.EsCoupon, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCoupon{}).Where("`coupon_id` = ?", couponID).First(&result).Error

	return
}
