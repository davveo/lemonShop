package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type MemberShopScoreMgr struct {
	*_BaseMgr
}

// NewMemberShopScoreMgr open func
func NewMemberShopScoreMgr(db db.Repo) *MemberShopScoreMgr {
	if db == nil {
		panic(fmt.Errorf("NewMemberShopScoreMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &MemberShopScoreMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_member_shop_score"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *MemberShopScoreMgr) GetTableName() string {
	return "es_member_shop_score"
}

// Reset 重置gorm会话
func (obj *MemberShopScoreMgr) Reset() *MemberShopScoreMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *MemberShopScoreMgr) Get() (result models.EsMemberShopScore, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberShopScore{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *MemberShopScoreMgr) Gets() (results []*models.EsMemberShopScore, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberShopScore{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *MemberShopScoreMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsMemberShopScore{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithScoreID score_id获取 主键
func (obj *MemberShopScoreMgr) WithScoreID(scoreID int) Option {
	return optionFunc(func(o *options) { o.query["score_id"] = scoreID })
}

// WithMemberID member_id获取 会员id
func (obj *MemberShopScoreMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithOrderSn order_sn获取 订单编号
func (obj *MemberShopScoreMgr) WithOrderSn(orderSn string) Option {
	return optionFunc(func(o *options) { o.query["order_sn"] = orderSn })
}

// WithDeliveryScore delivery_score获取 发货速度评分
func (obj *MemberShopScoreMgr) WithDeliveryScore(deliveryScore int16) Option {
	return optionFunc(func(o *options) { o.query["delivery_score"] = deliveryScore })
}

// WithDescriptionScore description_score获取 描述相符度评分
func (obj *MemberShopScoreMgr) WithDescriptionScore(descriptionScore int16) Option {
	return optionFunc(func(o *options) { o.query["description_score"] = descriptionScore })
}

// WithServiceScore service_score获取 服务评分
func (obj *MemberShopScoreMgr) WithServiceScore(serviceScore int16) Option {
	return optionFunc(func(o *options) { o.query["service_score"] = serviceScore })
}

// WithSellerID seller_id获取 卖家
func (obj *MemberShopScoreMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// GetByOption 功能选项模式获取
func (obj *MemberShopScoreMgr) GetByOption(opts ...Option) (result models.EsMemberShopScore, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberShopScore{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *MemberShopScoreMgr) GetByOptions(opts ...Option) (results []*models.EsMemberShopScore, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberShopScore{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *MemberShopScoreMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsMemberShopScore, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsMemberShopScore{}).Where(options.query)
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

// GetFromScoreID 通过score_id获取内容 主键
func (obj *MemberShopScoreMgr) GetFromScoreID(scoreID int) (result models.EsMemberShopScore, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberShopScore{}).Where("`score_id` = ?", scoreID).First(&result).Error

	return
}

// GetBatchFromScoreID 批量查找 主键
func (obj *MemberShopScoreMgr) GetBatchFromScoreID(scoreIDs []int) (results []*models.EsMemberShopScore, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberShopScore{}).Where("`score_id` IN (?)", scoreIDs).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *MemberShopScoreMgr) GetFromMemberID(memberID int) (results []*models.EsMemberShopScore, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberShopScore{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *MemberShopScoreMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsMemberShopScore, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberShopScore{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromOrderSn 通过order_sn获取内容 订单编号
func (obj *MemberShopScoreMgr) GetFromOrderSn(orderSn string) (results []*models.EsMemberShopScore, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberShopScore{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}

// GetBatchFromOrderSn 批量查找 订单编号
func (obj *MemberShopScoreMgr) GetBatchFromOrderSn(orderSns []string) (results []*models.EsMemberShopScore, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberShopScore{}).Where("`order_sn` IN (?)", orderSns).Find(&results).Error

	return
}

// GetFromDeliveryScore 通过delivery_score获取内容 发货速度评分
func (obj *MemberShopScoreMgr) GetFromDeliveryScore(deliveryScore int16) (results []*models.EsMemberShopScore, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberShopScore{}).Where("`delivery_score` = ?", deliveryScore).Find(&results).Error

	return
}

// GetBatchFromDeliveryScore 批量查找 发货速度评分
func (obj *MemberShopScoreMgr) GetBatchFromDeliveryScore(deliveryScores []int16) (results []*models.EsMemberShopScore, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberShopScore{}).Where("`delivery_score` IN (?)", deliveryScores).Find(&results).Error

	return
}

// GetFromDescriptionScore 通过description_score获取内容 描述相符度评分
func (obj *MemberShopScoreMgr) GetFromDescriptionScore(descriptionScore int16) (results []*models.EsMemberShopScore, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberShopScore{}).Where("`description_score` = ?", descriptionScore).Find(&results).Error

	return
}

// GetBatchFromDescriptionScore 批量查找 描述相符度评分
func (obj *MemberShopScoreMgr) GetBatchFromDescriptionScore(descriptionScores []int16) (results []*models.EsMemberShopScore, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberShopScore{}).Where("`description_score` IN (?)", descriptionScores).Find(&results).Error

	return
}

// GetFromServiceScore 通过service_score获取内容 服务评分
func (obj *MemberShopScoreMgr) GetFromServiceScore(serviceScore int16) (results []*models.EsMemberShopScore, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberShopScore{}).Where("`service_score` = ?", serviceScore).Find(&results).Error

	return
}

// GetBatchFromServiceScore 批量查找 服务评分
func (obj *MemberShopScoreMgr) GetBatchFromServiceScore(serviceScores []int16) (results []*models.EsMemberShopScore, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberShopScore{}).Where("`service_score` IN (?)", serviceScores).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 卖家
func (obj *MemberShopScoreMgr) GetFromSellerID(sellerID int) (results []*models.EsMemberShopScore, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberShopScore{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 卖家
func (obj *MemberShopScoreMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsMemberShopScore, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberShopScore{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *MemberShopScoreMgr) FetchByPrimaryKey(scoreID int) (result models.EsMemberShopScore, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberShopScore{}).Where("`score_id` = ?", scoreID).First(&result).Error

	return
}
