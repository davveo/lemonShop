package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type FullDiscountMgr struct {
	*_BaseMgr
}

// NewFullDiscountMgr open func
func NewFullDiscountMgr(db db.Repo) *FullDiscountMgr {
	if db == nil {
		panic(fmt.Errorf("NewFullDiscountMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &FullDiscountMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_category_spec"),
		wdb:       db.GetDbW().Table("es_category_spec"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *FullDiscountMgr) GetTableName() string {
	return "es_full_discount"
}

// Reset 重置gorm会话
func (obj *FullDiscountMgr) Reset() *FullDiscountMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *FullDiscountMgr) Get() (result models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *FullDiscountMgr) Gets() (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *FullDiscountMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithFdID fd_id获取 活动id
func (obj *FullDiscountMgr) WithFdID(fdID int) Option {
	return optionFunc(func(o *options) { o.query["fd_id"] = fdID })
}

// WithFullMoney full_money获取 优惠门槛金额
func (obj *FullDiscountMgr) WithFullMoney(fullMoney float64) Option {
	return optionFunc(func(o *options) { o.query["full_money"] = fullMoney })
}

// WithMinusValue minus_value获取 减现金
func (obj *FullDiscountMgr) WithMinusValue(minusValue float64) Option {
	return optionFunc(func(o *options) { o.query["minus_value"] = minusValue })
}

// WithPointValue point_value获取 送积分
func (obj *FullDiscountMgr) WithPointValue(pointValue int) Option {
	return optionFunc(func(o *options) { o.query["point_value"] = pointValue })
}

// WithIsFullMinus is_full_minus获取 活动是否减现金
func (obj *FullDiscountMgr) WithIsFullMinus(isFullMinus int) Option {
	return optionFunc(func(o *options) { o.query["is_full_minus"] = isFullMinus })
}

// WithIsFreeShip is_free_ship获取 是否免邮
func (obj *FullDiscountMgr) WithIsFreeShip(isFreeShip int) Option {
	return optionFunc(func(o *options) { o.query["is_free_ship"] = isFreeShip })
}

// WithIsSendPoint is_send_point获取 是否送积分
func (obj *FullDiscountMgr) WithIsSendPoint(isSendPoint int) Option {
	return optionFunc(func(o *options) { o.query["is_send_point"] = isSendPoint })
}

// WithIsSendGift is_send_gift获取 是否有赠品
func (obj *FullDiscountMgr) WithIsSendGift(isSendGift int) Option {
	return optionFunc(func(o *options) { o.query["is_send_gift"] = isSendGift })
}

// WithIsSendBonus is_send_bonus获取 是否增优惠券
func (obj *FullDiscountMgr) WithIsSendBonus(isSendBonus int) Option {
	return optionFunc(func(o *options) { o.query["is_send_bonus"] = isSendBonus })
}

// WithGiftID gift_id获取 赠品id
func (obj *FullDiscountMgr) WithGiftID(giftID int) Option {
	return optionFunc(func(o *options) { o.query["gift_id"] = giftID })
}

// WithBonusID bonus_id获取 优惠券id
func (obj *FullDiscountMgr) WithBonusID(bonusID int) Option {
	return optionFunc(func(o *options) { o.query["bonus_id"] = bonusID })
}

// WithIsDiscount is_discount获取 是否打折
func (obj *FullDiscountMgr) WithIsDiscount(isDiscount int) Option {
	return optionFunc(func(o *options) { o.query["is_discount"] = isDiscount })
}

// WithDiscountValue discount_value获取 打多少折
func (obj *FullDiscountMgr) WithDiscountValue(discountValue float64) Option {
	return optionFunc(func(o *options) { o.query["discount_value"] = discountValue })
}

// WithStartTime start_time获取 活动开始时间
func (obj *FullDiscountMgr) WithStartTime(startTime int64) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取 活动结束时间
func (obj *FullDiscountMgr) WithEndTime(endTime int64) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithTitle title获取 活动标题
func (obj *FullDiscountMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithRangeType range_type获取 是否全部商品参与
func (obj *FullDiscountMgr) WithRangeType(rangeType int) Option {
	return optionFunc(func(o *options) { o.query["range_type"] = rangeType })
}

// WithDisabled disabled获取 是否停用
func (obj *FullDiscountMgr) WithDisabled(disabled int) Option {
	return optionFunc(func(o *options) { o.query["disabled"] = disabled })
}

// WithDescription description获取 活动说明
func (obj *FullDiscountMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithSellerID seller_id获取 商家id
func (obj *FullDiscountMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// GetByOption 功能选项模式获取
func (obj *FullDiscountMgr) GetByOption(opts ...Option) (result models.EsFullDiscount, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *FullDiscountMgr) GetByOptions(opts ...Option) (results []*models.EsFullDiscount, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *FullDiscountMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsFullDiscount, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where(options.query)
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

// GetFromFdID 通过fd_id获取内容 活动id
func (obj *FullDiscountMgr) GetFromFdID(fdID int) (result models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`fd_id` = ?", fdID).First(&result).Error

	return
}

// GetBatchFromFdID 批量查找 活动id
func (obj *FullDiscountMgr) GetBatchFromFdID(fdIDs []int) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`fd_id` IN (?)", fdIDs).Find(&results).Error

	return
}

// GetFromFullMoney 通过full_money获取内容 优惠门槛金额
func (obj *FullDiscountMgr) GetFromFullMoney(fullMoney float64) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`full_money` = ?", fullMoney).Find(&results).Error

	return
}

// GetBatchFromFullMoney 批量查找 优惠门槛金额
func (obj *FullDiscountMgr) GetBatchFromFullMoney(fullMoneys []float64) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`full_money` IN (?)", fullMoneys).Find(&results).Error

	return
}

// GetFromMinusValue 通过minus_value获取内容 减现金
func (obj *FullDiscountMgr) GetFromMinusValue(minusValue float64) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`minus_value` = ?", minusValue).Find(&results).Error

	return
}

// GetBatchFromMinusValue 批量查找 减现金
func (obj *FullDiscountMgr) GetBatchFromMinusValue(minusValues []float64) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`minus_value` IN (?)", minusValues).Find(&results).Error

	return
}

// GetFromPointValue 通过point_value获取内容 送积分
func (obj *FullDiscountMgr) GetFromPointValue(pointValue int) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`point_value` = ?", pointValue).Find(&results).Error

	return
}

// GetBatchFromPointValue 批量查找 送积分
func (obj *FullDiscountMgr) GetBatchFromPointValue(pointValues []int) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`point_value` IN (?)", pointValues).Find(&results).Error

	return
}

// GetFromIsFullMinus 通过is_full_minus获取内容 活动是否减现金
func (obj *FullDiscountMgr) GetFromIsFullMinus(isFullMinus int) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`is_full_minus` = ?", isFullMinus).Find(&results).Error

	return
}

// GetBatchFromIsFullMinus 批量查找 活动是否减现金
func (obj *FullDiscountMgr) GetBatchFromIsFullMinus(isFullMinuss []int) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`is_full_minus` IN (?)", isFullMinuss).Find(&results).Error

	return
}

// GetFromIsFreeShip 通过is_free_ship获取内容 是否免邮
func (obj *FullDiscountMgr) GetFromIsFreeShip(isFreeShip int) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`is_free_ship` = ?", isFreeShip).Find(&results).Error

	return
}

// GetBatchFromIsFreeShip 批量查找 是否免邮
func (obj *FullDiscountMgr) GetBatchFromIsFreeShip(isFreeShips []int) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`is_free_ship` IN (?)", isFreeShips).Find(&results).Error

	return
}

// GetFromIsSendPoint 通过is_send_point获取内容 是否送积分
func (obj *FullDiscountMgr) GetFromIsSendPoint(isSendPoint int) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`is_send_point` = ?", isSendPoint).Find(&results).Error

	return
}

// GetBatchFromIsSendPoint 批量查找 是否送积分
func (obj *FullDiscountMgr) GetBatchFromIsSendPoint(isSendPoints []int) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`is_send_point` IN (?)", isSendPoints).Find(&results).Error

	return
}

// GetFromIsSendGift 通过is_send_gift获取内容 是否有赠品
func (obj *FullDiscountMgr) GetFromIsSendGift(isSendGift int) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`is_send_gift` = ?", isSendGift).Find(&results).Error

	return
}

// GetBatchFromIsSendGift 批量查找 是否有赠品
func (obj *FullDiscountMgr) GetBatchFromIsSendGift(isSendGifts []int) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`is_send_gift` IN (?)", isSendGifts).Find(&results).Error

	return
}

// GetFromIsSendBonus 通过is_send_bonus获取内容 是否增优惠券
func (obj *FullDiscountMgr) GetFromIsSendBonus(isSendBonus int) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`is_send_bonus` = ?", isSendBonus).Find(&results).Error

	return
}

// GetBatchFromIsSendBonus 批量查找 是否增优惠券
func (obj *FullDiscountMgr) GetBatchFromIsSendBonus(isSendBonuss []int) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`is_send_bonus` IN (?)", isSendBonuss).Find(&results).Error

	return
}

// GetFromGiftID 通过gift_id获取内容 赠品id
func (obj *FullDiscountMgr) GetFromGiftID(giftID int) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`gift_id` = ?", giftID).Find(&results).Error

	return
}

// GetBatchFromGiftID 批量查找 赠品id
func (obj *FullDiscountMgr) GetBatchFromGiftID(giftIDs []int) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`gift_id` IN (?)", giftIDs).Find(&results).Error

	return
}

// GetFromBonusID 通过bonus_id获取内容 优惠券id
func (obj *FullDiscountMgr) GetFromBonusID(bonusID int) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`bonus_id` = ?", bonusID).Find(&results).Error

	return
}

// GetBatchFromBonusID 批量查找 优惠券id
func (obj *FullDiscountMgr) GetBatchFromBonusID(bonusIDs []int) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`bonus_id` IN (?)", bonusIDs).Find(&results).Error

	return
}

// GetFromIsDiscount 通过is_discount获取内容 是否打折
func (obj *FullDiscountMgr) GetFromIsDiscount(isDiscount int) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`is_discount` = ?", isDiscount).Find(&results).Error

	return
}

// GetBatchFromIsDiscount 批量查找 是否打折
func (obj *FullDiscountMgr) GetBatchFromIsDiscount(isDiscounts []int) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`is_discount` IN (?)", isDiscounts).Find(&results).Error

	return
}

// GetFromDiscountValue 通过discount_value获取内容 打多少折
func (obj *FullDiscountMgr) GetFromDiscountValue(discountValue float64) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`discount_value` = ?", discountValue).Find(&results).Error

	return
}

// GetBatchFromDiscountValue 批量查找 打多少折
func (obj *FullDiscountMgr) GetBatchFromDiscountValue(discountValues []float64) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`discount_value` IN (?)", discountValues).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容 活动开始时间
func (obj *FullDiscountMgr) GetFromStartTime(startTime int64) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找 活动开始时间
func (obj *FullDiscountMgr) GetBatchFromStartTime(startTimes []int64) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容 活动结束时间
func (obj *FullDiscountMgr) GetFromEndTime(endTime int64) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找 活动结束时间
func (obj *FullDiscountMgr) GetBatchFromEndTime(endTimes []int64) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 活动标题
func (obj *FullDiscountMgr) GetFromTitle(title string) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 活动标题
func (obj *FullDiscountMgr) GetBatchFromTitle(titles []string) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromRangeType 通过range_type获取内容 是否全部商品参与
func (obj *FullDiscountMgr) GetFromRangeType(rangeType int) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`range_type` = ?", rangeType).Find(&results).Error

	return
}

// GetBatchFromRangeType 批量查找 是否全部商品参与
func (obj *FullDiscountMgr) GetBatchFromRangeType(rangeTypes []int) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`range_type` IN (?)", rangeTypes).Find(&results).Error

	return
}

// GetFromDisabled 通过disabled获取内容 是否停用
func (obj *FullDiscountMgr) GetFromDisabled(disabled int) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`disabled` = ?", disabled).Find(&results).Error

	return
}

// GetBatchFromDisabled 批量查找 是否停用
func (obj *FullDiscountMgr) GetBatchFromDisabled(disableds []int) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`disabled` IN (?)", disableds).Find(&results).Error

	return
}

// GetFromDescription 通过description获取内容 活动说明
func (obj *FullDiscountMgr) GetFromDescription(description string) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`description` = ?", description).Find(&results).Error

	return
}

// GetBatchFromDescription 批量查找 活动说明
func (obj *FullDiscountMgr) GetBatchFromDescription(descriptions []string) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`description` IN (?)", descriptions).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 商家id
func (obj *FullDiscountMgr) GetFromSellerID(sellerID int) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 商家id
func (obj *FullDiscountMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *FullDiscountMgr) FetchByPrimaryKey(fdID int) (result models.EsFullDiscount, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscount{}).Where("`fd_id` = ?", fdID).First(&result).Error

	return
}
