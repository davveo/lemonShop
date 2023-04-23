package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsPintuanGoodsMgr struct {
	*_BaseMgr
}

// EsPintuanGoodsMgr open func
func EsPintuanGoodsMgr(db *gorm.DB) *_EsPintuanGoodsMgr {
	if db == nil {
		panic(fmt.Errorf("EsPintuanGoodsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsPintuanGoodsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_pintuan_goods"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsPintuanGoodsMgr) GetTableName() string {
	return "es_pintuan_goods"
}

// Reset 重置gorm会话
func (obj *_EsPintuanGoodsMgr) Reset() *_EsPintuanGoodsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsPintuanGoodsMgr) Get() (result models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsPintuanGoodsMgr) Gets() (results []*models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsPintuanGoodsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *_EsPintuanGoodsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSkuID sku_id获取 sku_id
func (obj *_EsPintuanGoodsMgr) WithSkuID(skuID int) Option {
	return optionFunc(func(o *options) { o.query["sku_id"] = skuID })
}

// WithGoodsName goods_name获取 商品名称
func (obj *_EsPintuanGoodsMgr) WithGoodsName(goodsName string) Option {
	return optionFunc(func(o *options) { o.query["goods_name"] = goodsName })
}

// WithOriginPrice origin_price获取 原价
func (obj *_EsPintuanGoodsMgr) WithOriginPrice(originPrice float64) Option {
	return optionFunc(func(o *options) { o.query["origin_price"] = originPrice })
}

// WithSalesPrice sales_price获取 活动价
func (obj *_EsPintuanGoodsMgr) WithSalesPrice(salesPrice float64) Option {
	return optionFunc(func(o *options) { o.query["sales_price"] = salesPrice })
}

// WithSn sn获取 sn
func (obj *_EsPintuanGoodsMgr) WithSn(sn string) Option {
	return optionFunc(func(o *options) { o.query["sn"] = sn })
}

// WithSoldQuantity sold_quantity获取 已售数量
func (obj *_EsPintuanGoodsMgr) WithSoldQuantity(soldQuantity int) Option {
	return optionFunc(func(o *options) { o.query["sold_quantity"] = soldQuantity })
}

// WithLockedQuantity locked_quantity获取 待发货数量
func (obj *_EsPintuanGoodsMgr) WithLockedQuantity(lockedQuantity int) Option {
	return optionFunc(func(o *options) { o.query["locked_quantity"] = lockedQuantity })
}

// WithPintuanID pintuan_id获取 拼团活动id
func (obj *_EsPintuanGoodsMgr) WithPintuanID(pintuanID int) Option {
	return optionFunc(func(o *options) { o.query["pintuan_id"] = pintuanID })
}

// WithGoodsID goods_id获取 goods_id
func (obj *_EsPintuanGoodsMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithSpecs specs获取 规格
func (obj *_EsPintuanGoodsMgr) WithSpecs(specs string) Option {
	return optionFunc(func(o *options) { o.query["specs"] = specs })
}

// WithSellerID seller_id获取 卖家id
func (obj *_EsPintuanGoodsMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithSellerName seller_name获取 卖家名称
func (obj *_EsPintuanGoodsMgr) WithSellerName(sellerName string) Option {
	return optionFunc(func(o *options) { o.query["seller_name"] = sellerName })
}

// WithThumbnail thumbnail获取 商品缩略图
func (obj *_EsPintuanGoodsMgr) WithThumbnail(thumbnail string) Option {
	return optionFunc(func(o *options) { o.query["thumbnail"] = thumbnail })
}

// GetByOption 功能选项模式获取
func (obj *_EsPintuanGoodsMgr) GetByOption(opts ...Option) (result models.EsPintuanGoods, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsPintuanGoodsMgr) GetByOptions(opts ...Option) (results []*models.EsPintuanGoods, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsPintuanGoodsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsPintuanGoods, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where(options.query)
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

// GetFromID 通过id获取内容 id
func (obj *_EsPintuanGoodsMgr) GetFromID(id int) (result models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *_EsPintuanGoodsMgr) GetBatchFromID(ids []int) (results []*models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSkuID 通过sku_id获取内容 sku_id
func (obj *_EsPintuanGoodsMgr) GetFromSkuID(skuID int) (results []*models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where("`sku_id` = ?", skuID).Find(&results).Error

	return
}

// GetBatchFromSkuID 批量查找 sku_id
func (obj *_EsPintuanGoodsMgr) GetBatchFromSkuID(skuIDs []int) (results []*models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where("`sku_id` IN (?)", skuIDs).Find(&results).Error

	return
}

// GetFromGoodsName 通过goods_name获取内容 商品名称
func (obj *_EsPintuanGoodsMgr) GetFromGoodsName(goodsName string) (results []*models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}

// GetBatchFromGoodsName 批量查找 商品名称
func (obj *_EsPintuanGoodsMgr) GetBatchFromGoodsName(goodsNames []string) (results []*models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where("`goods_name` IN (?)", goodsNames).Find(&results).Error

	return
}

// GetFromOriginPrice 通过origin_price获取内容 原价
func (obj *_EsPintuanGoodsMgr) GetFromOriginPrice(originPrice float64) (results []*models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where("`origin_price` = ?", originPrice).Find(&results).Error

	return
}

// GetBatchFromOriginPrice 批量查找 原价
func (obj *_EsPintuanGoodsMgr) GetBatchFromOriginPrice(originPrices []float64) (results []*models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where("`origin_price` IN (?)", originPrices).Find(&results).Error

	return
}

// GetFromSalesPrice 通过sales_price获取内容 活动价
func (obj *_EsPintuanGoodsMgr) GetFromSalesPrice(salesPrice float64) (results []*models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where("`sales_price` = ?", salesPrice).Find(&results).Error

	return
}

// GetBatchFromSalesPrice 批量查找 活动价
func (obj *_EsPintuanGoodsMgr) GetBatchFromSalesPrice(salesPrices []float64) (results []*models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where("`sales_price` IN (?)", salesPrices).Find(&results).Error

	return
}

// GetFromSn 通过sn获取内容 sn
func (obj *_EsPintuanGoodsMgr) GetFromSn(sn string) (results []*models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where("`sn` = ?", sn).Find(&results).Error

	return
}

// GetBatchFromSn 批量查找 sn
func (obj *_EsPintuanGoodsMgr) GetBatchFromSn(sns []string) (results []*models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where("`sn` IN (?)", sns).Find(&results).Error

	return
}

// GetFromSoldQuantity 通过sold_quantity获取内容 已售数量
func (obj *_EsPintuanGoodsMgr) GetFromSoldQuantity(soldQuantity int) (results []*models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where("`sold_quantity` = ?", soldQuantity).Find(&results).Error

	return
}

// GetBatchFromSoldQuantity 批量查找 已售数量
func (obj *_EsPintuanGoodsMgr) GetBatchFromSoldQuantity(soldQuantitys []int) (results []*models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where("`sold_quantity` IN (?)", soldQuantitys).Find(&results).Error

	return
}

// GetFromLockedQuantity 通过locked_quantity获取内容 待发货数量
func (obj *_EsPintuanGoodsMgr) GetFromLockedQuantity(lockedQuantity int) (results []*models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where("`locked_quantity` = ?", lockedQuantity).Find(&results).Error

	return
}

// GetBatchFromLockedQuantity 批量查找 待发货数量
func (obj *_EsPintuanGoodsMgr) GetBatchFromLockedQuantity(lockedQuantitys []int) (results []*models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where("`locked_quantity` IN (?)", lockedQuantitys).Find(&results).Error

	return
}

// GetFromPintuanID 通过pintuan_id获取内容 拼团活动id
func (obj *_EsPintuanGoodsMgr) GetFromPintuanID(pintuanID int) (results []*models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where("`pintuan_id` = ?", pintuanID).Find(&results).Error

	return
}

// GetBatchFromPintuanID 批量查找 拼团活动id
func (obj *_EsPintuanGoodsMgr) GetBatchFromPintuanID(pintuanIDs []int) (results []*models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where("`pintuan_id` IN (?)", pintuanIDs).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 goods_id
func (obj *_EsPintuanGoodsMgr) GetFromGoodsID(goodsID int) (results []*models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 goods_id
func (obj *_EsPintuanGoodsMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromSpecs 通过specs获取内容 规格
func (obj *_EsPintuanGoodsMgr) GetFromSpecs(specs string) (results []*models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where("`specs` = ?", specs).Find(&results).Error

	return
}

// GetBatchFromSpecs 批量查找 规格
func (obj *_EsPintuanGoodsMgr) GetBatchFromSpecs(specss []string) (results []*models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where("`specs` IN (?)", specss).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 卖家id
func (obj *_EsPintuanGoodsMgr) GetFromSellerID(sellerID int) (results []*models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 卖家id
func (obj *_EsPintuanGoodsMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromSellerName 通过seller_name获取内容 卖家名称
func (obj *_EsPintuanGoodsMgr) GetFromSellerName(sellerName string) (results []*models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where("`seller_name` = ?", sellerName).Find(&results).Error

	return
}

// GetBatchFromSellerName 批量查找 卖家名称
func (obj *_EsPintuanGoodsMgr) GetBatchFromSellerName(sellerNames []string) (results []*models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where("`seller_name` IN (?)", sellerNames).Find(&results).Error

	return
}

// GetFromThumbnail 通过thumbnail获取内容 商品缩略图
func (obj *_EsPintuanGoodsMgr) GetFromThumbnail(thumbnail string) (results []*models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where("`thumbnail` = ?", thumbnail).Find(&results).Error

	return
}

// GetBatchFromThumbnail 批量查找 商品缩略图
func (obj *_EsPintuanGoodsMgr) GetBatchFromThumbnail(thumbnails []string) (results []*models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where("`thumbnail` IN (?)", thumbnails).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsPintuanGoodsMgr) FetchByPrimaryKey(id int) (result models.EsPintuanGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanGoods{}).Where("`id` = ?", id).First(&result).Error

	return
}
