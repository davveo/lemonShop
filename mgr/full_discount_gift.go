package mgr

import (
	"context"
	"fmt"

	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type FullDiscountGiftMgr struct {
	*_BaseMgr
}

// NewFullDiscountGiftMgr open func
func NewFullDiscountGiftMgr(db db.Repo) *FullDiscountGiftMgr {
	if db == nil {
		panic(fmt.Errorf("NewFullDiscountGiftMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &FullDiscountGiftMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_full_discount_gift"),
		wdb:       db.GetDbW().Table("es_full_discount_gift"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *FullDiscountGiftMgr) GetTableName() string {
	return "es_full_discount_gift"
}

// Reset 重置gorm会话
func (obj *FullDiscountGiftMgr) Reset() *FullDiscountGiftMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *FullDiscountGiftMgr) Get() (result models.EsFullDiscountGift, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscountGift{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *FullDiscountGiftMgr) Gets() (results []*models.EsFullDiscountGift, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscountGift{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *FullDiscountGiftMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscountGift{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGiftID gift_id获取 赠品id
func (obj *FullDiscountGiftMgr) WithGiftID(giftID int) Option {
	return optionFunc(func(o *options) { o.query["gift_id"] = giftID })
}

// WithGiftName gift_name获取 赠品名称
func (obj *FullDiscountGiftMgr) WithGiftName(giftName string) Option {
	return optionFunc(func(o *options) { o.query["gift_name"] = giftName })
}

// WithGiftPrice gift_price获取 赠品金额
func (obj *FullDiscountGiftMgr) WithGiftPrice(giftPrice float64) Option {
	return optionFunc(func(o *options) { o.query["gift_price"] = giftPrice })
}

// WithGiftImg gift_img获取 赠品图片
func (obj *FullDiscountGiftMgr) WithGiftImg(giftImg string) Option {
	return optionFunc(func(o *options) { o.query["gift_img"] = giftImg })
}

// WithGiftType gift_type获取 赠品类型
func (obj *FullDiscountGiftMgr) WithGiftType(giftType int) Option {
	return optionFunc(func(o *options) { o.query["gift_type"] = giftType })
}

// WithActualStore actual_store获取 库存
func (obj *FullDiscountGiftMgr) WithActualStore(actualStore int) Option {
	return optionFunc(func(o *options) { o.query["actual_store"] = actualStore })
}

// WithEnableStore enable_store获取 可用库存
func (obj *FullDiscountGiftMgr) WithEnableStore(enableStore int) Option {
	return optionFunc(func(o *options) { o.query["enable_store"] = enableStore })
}

// WithCreateTime create_time获取 活动时间
func (obj *FullDiscountGiftMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithGoodsID goods_id获取 活动商品id
func (obj *FullDiscountGiftMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithDisabled disabled获取 是否禁用
func (obj *FullDiscountGiftMgr) WithDisabled(disabled int) Option {
	return optionFunc(func(o *options) { o.query["disabled"] = disabled })
}

// WithSellerID seller_id获取 店铺id
func (obj *FullDiscountGiftMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// GetByOption 功能选项模式获取
func (obj *FullDiscountGiftMgr) GetByOption(opts ...Option) (result models.EsFullDiscountGift, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscountGift{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *FullDiscountGiftMgr) GetByOptions(opts ...Option) (results []*models.EsFullDiscountGift, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscountGift{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *FullDiscountGiftMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsFullDiscountGift, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscountGift{}).Where(options.query)
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

// GetFromGiftID 通过gift_id获取内容 赠品id
func (obj *FullDiscountGiftMgr) GetFromGiftID(giftID int) (result models.EsFullDiscountGift, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscountGift{}).Where("`gift_id` = ?", giftID).First(&result).Error

	return
}

// GetBatchFromGiftID 批量查找 赠品id
func (obj *FullDiscountGiftMgr) GetBatchFromGiftID(giftIDs []int) (results []*models.EsFullDiscountGift, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscountGift{}).Where("`gift_id` IN (?)", giftIDs).Find(&results).Error

	return
}

// GetFromGiftName 通过gift_name获取内容 赠品名称
func (obj *FullDiscountGiftMgr) GetFromGiftName(giftName string) (results []*models.EsFullDiscountGift, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscountGift{}).Where("`gift_name` = ?", giftName).Find(&results).Error

	return
}

// GetBatchFromGiftName 批量查找 赠品名称
func (obj *FullDiscountGiftMgr) GetBatchFromGiftName(giftNames []string) (results []*models.EsFullDiscountGift, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscountGift{}).Where("`gift_name` IN (?)", giftNames).Find(&results).Error

	return
}

// GetFromGiftPrice 通过gift_price获取内容 赠品金额
func (obj *FullDiscountGiftMgr) GetFromGiftPrice(giftPrice float64) (results []*models.EsFullDiscountGift, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscountGift{}).Where("`gift_price` = ?", giftPrice).Find(&results).Error

	return
}

// GetBatchFromGiftPrice 批量查找 赠品金额
func (obj *FullDiscountGiftMgr) GetBatchFromGiftPrice(giftPrices []float64) (results []*models.EsFullDiscountGift, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscountGift{}).Where("`gift_price` IN (?)", giftPrices).Find(&results).Error

	return
}

// GetFromGiftImg 通过gift_img获取内容 赠品图片
func (obj *FullDiscountGiftMgr) GetFromGiftImg(giftImg string) (results []*models.EsFullDiscountGift, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscountGift{}).Where("`gift_img` = ?", giftImg).Find(&results).Error

	return
}

// GetBatchFromGiftImg 批量查找 赠品图片
func (obj *FullDiscountGiftMgr) GetBatchFromGiftImg(giftImgs []string) (results []*models.EsFullDiscountGift, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscountGift{}).Where("`gift_img` IN (?)", giftImgs).Find(&results).Error

	return
}

// GetFromGiftType 通过gift_type获取内容 赠品类型
func (obj *FullDiscountGiftMgr) GetFromGiftType(giftType int) (results []*models.EsFullDiscountGift, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscountGift{}).Where("`gift_type` = ?", giftType).Find(&results).Error

	return
}

// GetBatchFromGiftType 批量查找 赠品类型
func (obj *FullDiscountGiftMgr) GetBatchFromGiftType(giftTypes []int) (results []*models.EsFullDiscountGift, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscountGift{}).Where("`gift_type` IN (?)", giftTypes).Find(&results).Error

	return
}

// GetFromActualStore 通过actual_store获取内容 库存
func (obj *FullDiscountGiftMgr) GetFromActualStore(actualStore int) (results []*models.EsFullDiscountGift, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscountGift{}).Where("`actual_store` = ?", actualStore).Find(&results).Error

	return
}

// GetBatchFromActualStore 批量查找 库存
func (obj *FullDiscountGiftMgr) GetBatchFromActualStore(actualStores []int) (results []*models.EsFullDiscountGift, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscountGift{}).Where("`actual_store` IN (?)", actualStores).Find(&results).Error

	return
}

// GetFromEnableStore 通过enable_store获取内容 可用库存
func (obj *FullDiscountGiftMgr) GetFromEnableStore(enableStore int) (results []*models.EsFullDiscountGift, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscountGift{}).Where("`enable_store` = ?", enableStore).Find(&results).Error

	return
}

// GetBatchFromEnableStore 批量查找 可用库存
func (obj *FullDiscountGiftMgr) GetBatchFromEnableStore(enableStores []int) (results []*models.EsFullDiscountGift, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscountGift{}).Where("`enable_store` IN (?)", enableStores).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 活动时间
func (obj *FullDiscountGiftMgr) GetFromCreateTime(createTime int64) (results []*models.EsFullDiscountGift, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscountGift{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 活动时间
func (obj *FullDiscountGiftMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsFullDiscountGift, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscountGift{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 活动商品id
func (obj *FullDiscountGiftMgr) GetFromGoodsID(goodsID int) (results []*models.EsFullDiscountGift, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscountGift{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 活动商品id
func (obj *FullDiscountGiftMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsFullDiscountGift, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscountGift{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromDisabled 通过disabled获取内容 是否禁用
func (obj *FullDiscountGiftMgr) GetFromDisabled(disabled int) (results []*models.EsFullDiscountGift, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscountGift{}).Where("`disabled` = ?", disabled).Find(&results).Error

	return
}

// GetBatchFromDisabled 批量查找 是否禁用
func (obj *FullDiscountGiftMgr) GetBatchFromDisabled(disableds []int) (results []*models.EsFullDiscountGift, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscountGift{}).Where("`disabled` IN (?)", disableds).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 店铺id
func (obj *FullDiscountGiftMgr) GetFromSellerID(sellerID int) (results []*models.EsFullDiscountGift, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscountGift{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 店铺id
func (obj *FullDiscountGiftMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsFullDiscountGift, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscountGift{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *FullDiscountGiftMgr) FetchByPrimaryKey(giftID int) (result models.EsFullDiscountGift, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFullDiscountGift{}).Where("`gift_id` = ?", giftID).First(&result).Error

	return
}
