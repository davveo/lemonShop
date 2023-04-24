package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type PromotionGoodsMgr struct {
	*_BaseMgr
}

// NewPromotionGoodsMgr open func
func NewPromotionGoodsMgr(db db.Repo) *PromotionGoodsMgr {
	if db == nil {
		panic(fmt.Errorf("NewPromotionGoodsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &PromotionGoodsMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_promotion_goods"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *PromotionGoodsMgr) GetTableName() string {
	return "es_promotion_goods"
}

// Reset 重置gorm会话
func (obj *PromotionGoodsMgr) Reset() *PromotionGoodsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *PromotionGoodsMgr) Get() (result models.EsPromotionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPromotionGoods{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *PromotionGoodsMgr) Gets() (results []*models.EsPromotionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPromotionGoods{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *PromotionGoodsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsPromotionGoods{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithPgID pg_id获取 商品对照ID
func (obj *PromotionGoodsMgr) WithPgID(pgID int) Option {
	return optionFunc(func(o *options) { o.query["pg_id"] = pgID })
}

// WithGoodsID goods_id获取 商品id
func (obj *PromotionGoodsMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithProductID product_id获取 货品id
func (obj *PromotionGoodsMgr) WithProductID(productID int) Option {
	return optionFunc(func(o *options) { o.query["product_id"] = productID })
}

// WithStartTime start_time获取 活动开始时间
func (obj *PromotionGoodsMgr) WithStartTime(startTime int64) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取 活动结束时间
func (obj *PromotionGoodsMgr) WithEndTime(endTime int64) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithActivityID activity_id获取 活动id
func (obj *PromotionGoodsMgr) WithActivityID(activityID int) Option {
	return optionFunc(func(o *options) { o.query["activity_id"] = activityID })
}

// WithPromotionType promotion_type获取 促销工具类型
func (obj *PromotionGoodsMgr) WithPromotionType(promotionType string) Option {
	return optionFunc(func(o *options) { o.query["promotion_type"] = promotionType })
}

// WithTitle title获取 活动标题
func (obj *PromotionGoodsMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithNum num获取 参与活动的商品数量
func (obj *PromotionGoodsMgr) WithNum(num int) Option {
	return optionFunc(func(o *options) { o.query["num"] = num })
}

// WithPrice price获取 活动时商品的价格
func (obj *PromotionGoodsMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithSellerID seller_id获取 商家ID
func (obj *PromotionGoodsMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// GetByOption 功能选项模式获取
func (obj *PromotionGoodsMgr) GetByOption(opts ...Option) (result models.EsPromotionGoods, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPromotionGoods{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *PromotionGoodsMgr) GetByOptions(opts ...Option) (results []*models.EsPromotionGoods, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPromotionGoods{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *PromotionGoodsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsPromotionGoods, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsPromotionGoods{}).Where(options.query)
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

// GetFromPgID 通过pg_id获取内容 商品对照ID
func (obj *PromotionGoodsMgr) GetFromPgID(pgID int) (result models.EsPromotionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPromotionGoods{}).Where("`pg_id` = ?", pgID).First(&result).Error

	return
}

// GetBatchFromPgID 批量查找 商品对照ID
func (obj *PromotionGoodsMgr) GetBatchFromPgID(pgIDs []int) (results []*models.EsPromotionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPromotionGoods{}).Where("`pg_id` IN (?)", pgIDs).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 商品id
func (obj *PromotionGoodsMgr) GetFromGoodsID(goodsID int) (results []*models.EsPromotionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPromotionGoods{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 商品id
func (obj *PromotionGoodsMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsPromotionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPromotionGoods{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromProductID 通过product_id获取内容 货品id
func (obj *PromotionGoodsMgr) GetFromProductID(productID int) (results []*models.EsPromotionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPromotionGoods{}).Where("`product_id` = ?", productID).Find(&results).Error

	return
}

// GetBatchFromProductID 批量查找 货品id
func (obj *PromotionGoodsMgr) GetBatchFromProductID(productIDs []int) (results []*models.EsPromotionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPromotionGoods{}).Where("`product_id` IN (?)", productIDs).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容 活动开始时间
func (obj *PromotionGoodsMgr) GetFromStartTime(startTime int64) (results []*models.EsPromotionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPromotionGoods{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找 活动开始时间
func (obj *PromotionGoodsMgr) GetBatchFromStartTime(startTimes []int64) (results []*models.EsPromotionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPromotionGoods{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容 活动结束时间
func (obj *PromotionGoodsMgr) GetFromEndTime(endTime int64) (results []*models.EsPromotionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPromotionGoods{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找 活动结束时间
func (obj *PromotionGoodsMgr) GetBatchFromEndTime(endTimes []int64) (results []*models.EsPromotionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPromotionGoods{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromActivityID 通过activity_id获取内容 活动id
func (obj *PromotionGoodsMgr) GetFromActivityID(activityID int) (results []*models.EsPromotionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPromotionGoods{}).Where("`activity_id` = ?", activityID).Find(&results).Error

	return
}

// GetBatchFromActivityID 批量查找 活动id
func (obj *PromotionGoodsMgr) GetBatchFromActivityID(activityIDs []int) (results []*models.EsPromotionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPromotionGoods{}).Where("`activity_id` IN (?)", activityIDs).Find(&results).Error

	return
}

// GetFromPromotionType 通过promotion_type获取内容 促销工具类型
func (obj *PromotionGoodsMgr) GetFromPromotionType(promotionType string) (results []*models.EsPromotionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPromotionGoods{}).Where("`promotion_type` = ?", promotionType).Find(&results).Error

	return
}

// GetBatchFromPromotionType 批量查找 促销工具类型
func (obj *PromotionGoodsMgr) GetBatchFromPromotionType(promotionTypes []string) (results []*models.EsPromotionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPromotionGoods{}).Where("`promotion_type` IN (?)", promotionTypes).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 活动标题
func (obj *PromotionGoodsMgr) GetFromTitle(title string) (results []*models.EsPromotionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPromotionGoods{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 活动标题
func (obj *PromotionGoodsMgr) GetBatchFromTitle(titles []string) (results []*models.EsPromotionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPromotionGoods{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromNum 通过num获取内容 参与活动的商品数量
func (obj *PromotionGoodsMgr) GetFromNum(num int) (results []*models.EsPromotionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPromotionGoods{}).Where("`num` = ?", num).Find(&results).Error

	return
}

// GetBatchFromNum 批量查找 参与活动的商品数量
func (obj *PromotionGoodsMgr) GetBatchFromNum(nums []int) (results []*models.EsPromotionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPromotionGoods{}).Where("`num` IN (?)", nums).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 活动时商品的价格
func (obj *PromotionGoodsMgr) GetFromPrice(price float64) (results []*models.EsPromotionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPromotionGoods{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 活动时商品的价格
func (obj *PromotionGoodsMgr) GetBatchFromPrice(prices []float64) (results []*models.EsPromotionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPromotionGoods{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 商家ID
func (obj *PromotionGoodsMgr) GetFromSellerID(sellerID int) (results []*models.EsPromotionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPromotionGoods{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 商家ID
func (obj *PromotionGoodsMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsPromotionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPromotionGoods{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *PromotionGoodsMgr) FetchByPrimaryKey(pgID int) (result models.EsPromotionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPromotionGoods{}).Where("`pg_id` = ?", pgID).First(&result).Error

	return
}
