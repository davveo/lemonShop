package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type SeckillApplyMgr struct {
	*_BaseMgr
}

// NewSeckillApplyMgr open func
func NewSeckillApplyMgr(db db.Repo) *SeckillApplyMgr {
	if db == nil {
		panic(fmt.Errorf("NewSeckillApplyMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &SeckillApplyMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_category_spec"),
		wdb:       db.GetDbW().Table("es_category_spec"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *SeckillApplyMgr) GetTableName() string {
	return "es_seckill_apply"
}

// Reset 重置gorm会话
func (obj *SeckillApplyMgr) Reset() *SeckillApplyMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *SeckillApplyMgr) Get() (result models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *SeckillApplyMgr) Gets() (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *SeckillApplyMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithApplyID apply_id获取 主键ID
func (obj *SeckillApplyMgr) WithApplyID(applyID int) Option {
	return optionFunc(func(o *options) { o.query["apply_id"] = applyID })
}

// WithSeckillID seckill_id获取 活动id
func (obj *SeckillApplyMgr) WithSeckillID(seckillID int) Option {
	return optionFunc(func(o *options) { o.query["seckill_id"] = seckillID })
}

// WithTimeLine time_line获取 时刻
func (obj *SeckillApplyMgr) WithTimeLine(timeLine int) Option {
	return optionFunc(func(o *options) { o.query["time_line"] = timeLine })
}

// WithStartDay start_day获取 活动开始日期
func (obj *SeckillApplyMgr) WithStartDay(startDay int64) Option {
	return optionFunc(func(o *options) { o.query["start_day"] = startDay })
}

// WithGoodsID goods_id获取 商品ID
func (obj *SeckillApplyMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithGoodsName goods_name获取 商品名称
func (obj *SeckillApplyMgr) WithGoodsName(goodsName string) Option {
	return optionFunc(func(o *options) { o.query["goods_name"] = goodsName })
}

// WithSkuID sku_id获取 商品规格ID
func (obj *SeckillApplyMgr) WithSkuID(skuID int) Option {
	return optionFunc(func(o *options) { o.query["sku_id"] = skuID })
}

// WithSellerID seller_id获取 商家ID
func (obj *SeckillApplyMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithShopName shop_name获取 商家名称
func (obj *SeckillApplyMgr) WithShopName(shopName string) Option {
	return optionFunc(func(o *options) { o.query["shop_name"] = shopName })
}

// WithPrice price获取 价格
func (obj *SeckillApplyMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithSoldQuantity sold_quantity获取 售空数量
func (obj *SeckillApplyMgr) WithSoldQuantity(soldQuantity int) Option {
	return optionFunc(func(o *options) { o.query["sold_quantity"] = soldQuantity })
}

// WithStatus status获取 申请状态
func (obj *SeckillApplyMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithFailReason fail_reason获取 驳回原因
func (obj *SeckillApplyMgr) WithFailReason(failReason string) Option {
	return optionFunc(func(o *options) { o.query["fail_reason"] = failReason })
}

// WithSalesNum sales_num获取 已售数量
func (obj *SeckillApplyMgr) WithSalesNum(salesNum int) Option {
	return optionFunc(func(o *options) { o.query["sales_num"] = salesNum })
}

// WithOriginalPrice original_price获取 商品原始价格
func (obj *SeckillApplyMgr) WithOriginalPrice(originalPrice float64) Option {
	return optionFunc(func(o *options) { o.query["original_price"] = originalPrice })
}

// GetByOption 功能选项模式获取
func (obj *SeckillApplyMgr) GetByOption(opts ...Option) (result models.EsSeckillApply, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *SeckillApplyMgr) GetByOptions(opts ...Option) (results []*models.EsSeckillApply, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *SeckillApplyMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsSeckillApply, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where(options.query)
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

// GetFromApplyID 通过apply_id获取内容 主键ID
func (obj *SeckillApplyMgr) GetFromApplyID(applyID int) (result models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`apply_id` = ?", applyID).First(&result).Error

	return
}

// GetBatchFromApplyID 批量查找 主键ID
func (obj *SeckillApplyMgr) GetBatchFromApplyID(applyIDs []int) (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`apply_id` IN (?)", applyIDs).Find(&results).Error

	return
}

// GetFromSeckillID 通过seckill_id获取内容 活动id
func (obj *SeckillApplyMgr) GetFromSeckillID(seckillID int) (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`seckill_id` = ?", seckillID).Find(&results).Error

	return
}

// GetBatchFromSeckillID 批量查找 活动id
func (obj *SeckillApplyMgr) GetBatchFromSeckillID(seckillIDs []int) (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`seckill_id` IN (?)", seckillIDs).Find(&results).Error

	return
}

// GetFromTimeLine 通过time_line获取内容 时刻
func (obj *SeckillApplyMgr) GetFromTimeLine(timeLine int) (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`time_line` = ?", timeLine).Find(&results).Error

	return
}

// GetBatchFromTimeLine 批量查找 时刻
func (obj *SeckillApplyMgr) GetBatchFromTimeLine(timeLines []int) (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`time_line` IN (?)", timeLines).Find(&results).Error

	return
}

// GetFromStartDay 通过start_day获取内容 活动开始日期
func (obj *SeckillApplyMgr) GetFromStartDay(startDay int64) (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`start_day` = ?", startDay).Find(&results).Error

	return
}

// GetBatchFromStartDay 批量查找 活动开始日期
func (obj *SeckillApplyMgr) GetBatchFromStartDay(startDays []int64) (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`start_day` IN (?)", startDays).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 商品ID
func (obj *SeckillApplyMgr) GetFromGoodsID(goodsID int) (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 商品ID
func (obj *SeckillApplyMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromGoodsName 通过goods_name获取内容 商品名称
func (obj *SeckillApplyMgr) GetFromGoodsName(goodsName string) (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}

// GetBatchFromGoodsName 批量查找 商品名称
func (obj *SeckillApplyMgr) GetBatchFromGoodsName(goodsNames []string) (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`goods_name` IN (?)", goodsNames).Find(&results).Error

	return
}

// GetFromSkuID 通过sku_id获取内容 商品规格ID
func (obj *SeckillApplyMgr) GetFromSkuID(skuID int) (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`sku_id` = ?", skuID).Find(&results).Error

	return
}

// GetBatchFromSkuID 批量查找 商品规格ID
func (obj *SeckillApplyMgr) GetBatchFromSkuID(skuIDs []int) (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`sku_id` IN (?)", skuIDs).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 商家ID
func (obj *SeckillApplyMgr) GetFromSellerID(sellerID int) (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 商家ID
func (obj *SeckillApplyMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromShopName 通过shop_name获取内容 商家名称
func (obj *SeckillApplyMgr) GetFromShopName(shopName string) (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`shop_name` = ?", shopName).Find(&results).Error

	return
}

// GetBatchFromShopName 批量查找 商家名称
func (obj *SeckillApplyMgr) GetBatchFromShopName(shopNames []string) (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`shop_name` IN (?)", shopNames).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 价格
func (obj *SeckillApplyMgr) GetFromPrice(price float64) (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 价格
func (obj *SeckillApplyMgr) GetBatchFromPrice(prices []float64) (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromSoldQuantity 通过sold_quantity获取内容 售空数量
func (obj *SeckillApplyMgr) GetFromSoldQuantity(soldQuantity int) (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`sold_quantity` = ?", soldQuantity).Find(&results).Error

	return
}

// GetBatchFromSoldQuantity 批量查找 售空数量
func (obj *SeckillApplyMgr) GetBatchFromSoldQuantity(soldQuantitys []int) (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`sold_quantity` IN (?)", soldQuantitys).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 申请状态
func (obj *SeckillApplyMgr) GetFromStatus(status string) (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 申请状态
func (obj *SeckillApplyMgr) GetBatchFromStatus(statuss []string) (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromFailReason 通过fail_reason获取内容 驳回原因
func (obj *SeckillApplyMgr) GetFromFailReason(failReason string) (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`fail_reason` = ?", failReason).Find(&results).Error

	return
}

// GetBatchFromFailReason 批量查找 驳回原因
func (obj *SeckillApplyMgr) GetBatchFromFailReason(failReasons []string) (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`fail_reason` IN (?)", failReasons).Find(&results).Error

	return
}

// GetFromSalesNum 通过sales_num获取内容 已售数量
func (obj *SeckillApplyMgr) GetFromSalesNum(salesNum int) (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`sales_num` = ?", salesNum).Find(&results).Error

	return
}

// GetBatchFromSalesNum 批量查找 已售数量
func (obj *SeckillApplyMgr) GetBatchFromSalesNum(salesNums []int) (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`sales_num` IN (?)", salesNums).Find(&results).Error

	return
}

// GetFromOriginalPrice 通过original_price获取内容 商品原始价格
func (obj *SeckillApplyMgr) GetFromOriginalPrice(originalPrice float64) (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`original_price` = ?", originalPrice).Find(&results).Error

	return
}

// GetBatchFromOriginalPrice 批量查找 商品原始价格
func (obj *SeckillApplyMgr) GetBatchFromOriginalPrice(originalPrices []float64) (results []*models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`original_price` IN (?)", originalPrices).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *SeckillApplyMgr) FetchByPrimaryKey(applyID int) (result models.EsSeckillApply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillApply{}).Where("`apply_id` = ?", applyID).First(&result).Error

	return
}
