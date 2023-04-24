package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type GoodsSkuMgr struct {
	*_BaseMgr
}

// NewGoodsSkuMgr open func
func NewGoodsSkuMgr(db db.Repo) *GoodsSkuMgr {
	if db == nil {
		panic(fmt.Errorf("NewGoodsSkuMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &GoodsSkuMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_category_spec"),
		wdb:       db.GetDbW().Table("es_category_spec"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *GoodsSkuMgr) GetTableName() string {
	return "es_goods_sku"
}

// Reset 重置gorm会话
func (obj *GoodsSkuMgr) Reset() *GoodsSkuMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *GoodsSkuMgr) Get() (result models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *GoodsSkuMgr) Gets() (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *GoodsSkuMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSkuID sku_id获取 主键
func (obj *GoodsSkuMgr) WithSkuID(skuID int) Option {
	return optionFunc(func(o *options) { o.query["sku_id"] = skuID })
}

// WithGoodsID goods_id获取 商品id
func (obj *GoodsSkuMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithGoodsName goods_name获取 商品名称
func (obj *GoodsSkuMgr) WithGoodsName(goodsName string) Option {
	return optionFunc(func(o *options) { o.query["goods_name"] = goodsName })
}

// WithSn sn获取 商品编号
func (obj *GoodsSkuMgr) WithSn(sn string) Option {
	return optionFunc(func(o *options) { o.query["sn"] = sn })
}

// WithQuantity quantity获取 库存
func (obj *GoodsSkuMgr) WithQuantity(quantity int) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

// WithEnableQuantity enable_quantity获取 可用库存
func (obj *GoodsSkuMgr) WithEnableQuantity(enableQuantity int) Option {
	return optionFunc(func(o *options) { o.query["enable_quantity"] = enableQuantity })
}

// WithPrice price获取 商品价格
func (obj *GoodsSkuMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithSpecs specs获取 规格信息json
func (obj *GoodsSkuMgr) WithSpecs(specs string) Option {
	return optionFunc(func(o *options) { o.query["specs"] = specs })
}

// WithCost cost获取 成本价格
func (obj *GoodsSkuMgr) WithCost(cost float64) Option {
	return optionFunc(func(o *options) { o.query["cost"] = cost })
}

// WithWeight weight获取 重量
func (obj *GoodsSkuMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithSellerID seller_id获取 卖家id
func (obj *GoodsSkuMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithSellerName seller_name获取 卖家名称
func (obj *GoodsSkuMgr) WithSellerName(sellerName string) Option {
	return optionFunc(func(o *options) { o.query["seller_name"] = sellerName })
}

// WithCategoryID category_id获取 分类id
func (obj *GoodsSkuMgr) WithCategoryID(categoryID int) Option {
	return optionFunc(func(o *options) { o.query["category_id"] = categoryID })
}

// WithThumbnail thumbnail获取 缩略图
func (obj *GoodsSkuMgr) WithThumbnail(thumbnail string) Option {
	return optionFunc(func(o *options) { o.query["thumbnail"] = thumbnail })
}

// WithHashCode hash_code获取 标识规格唯一性
func (obj *GoodsSkuMgr) WithHashCode(hashCode int) Option {
	return optionFunc(func(o *options) { o.query["hash_code"] = hashCode })
}

// GetByOption 功能选项模式获取
func (obj *GoodsSkuMgr) GetByOption(opts ...Option) (result models.EsGoodsSku, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *GoodsSkuMgr) GetByOptions(opts ...Option) (results []*models.EsGoodsSku, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *GoodsSkuMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsGoodsSku, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where(options.query)
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

// GetFromSkuID 通过sku_id获取内容 主键
func (obj *GoodsSkuMgr) GetFromSkuID(skuID int) (result models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`sku_id` = ?", skuID).First(&result).Error

	return
}

// GetBatchFromSkuID 批量查找 主键
func (obj *GoodsSkuMgr) GetBatchFromSkuID(skuIDs []int) (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`sku_id` IN (?)", skuIDs).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 商品id
func (obj *GoodsSkuMgr) GetFromGoodsID(goodsID int) (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 商品id
func (obj *GoodsSkuMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromGoodsName 通过goods_name获取内容 商品名称
func (obj *GoodsSkuMgr) GetFromGoodsName(goodsName string) (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}

// GetBatchFromGoodsName 批量查找 商品名称
func (obj *GoodsSkuMgr) GetBatchFromGoodsName(goodsNames []string) (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`goods_name` IN (?)", goodsNames).Find(&results).Error

	return
}

// GetFromSn 通过sn获取内容 商品编号
func (obj *GoodsSkuMgr) GetFromSn(sn string) (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`sn` = ?", sn).Find(&results).Error

	return
}

// GetBatchFromSn 批量查找 商品编号
func (obj *GoodsSkuMgr) GetBatchFromSn(sns []string) (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`sn` IN (?)", sns).Find(&results).Error

	return
}

// GetFromQuantity 通过quantity获取内容 库存
func (obj *GoodsSkuMgr) GetFromQuantity(quantity int) (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`quantity` = ?", quantity).Find(&results).Error

	return
}

// GetBatchFromQuantity 批量查找 库存
func (obj *GoodsSkuMgr) GetBatchFromQuantity(quantitys []int) (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`quantity` IN (?)", quantitys).Find(&results).Error

	return
}

// GetFromEnableQuantity 通过enable_quantity获取内容 可用库存
func (obj *GoodsSkuMgr) GetFromEnableQuantity(enableQuantity int) (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`enable_quantity` = ?", enableQuantity).Find(&results).Error

	return
}

// GetBatchFromEnableQuantity 批量查找 可用库存
func (obj *GoodsSkuMgr) GetBatchFromEnableQuantity(enableQuantitys []int) (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`enable_quantity` IN (?)", enableQuantitys).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 商品价格
func (obj *GoodsSkuMgr) GetFromPrice(price float64) (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 商品价格
func (obj *GoodsSkuMgr) GetBatchFromPrice(prices []float64) (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromSpecs 通过specs获取内容 规格信息json
func (obj *GoodsSkuMgr) GetFromSpecs(specs string) (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`specs` = ?", specs).Find(&results).Error

	return
}

// GetBatchFromSpecs 批量查找 规格信息json
func (obj *GoodsSkuMgr) GetBatchFromSpecs(specss []string) (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`specs` IN (?)", specss).Find(&results).Error

	return
}

// GetFromCost 通过cost获取内容 成本价格
func (obj *GoodsSkuMgr) GetFromCost(cost float64) (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`cost` = ?", cost).Find(&results).Error

	return
}

// GetBatchFromCost 批量查找 成本价格
func (obj *GoodsSkuMgr) GetBatchFromCost(costs []float64) (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`cost` IN (?)", costs).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 重量
func (obj *GoodsSkuMgr) GetFromWeight(weight float64) (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`weight` = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量查找 重量
func (obj *GoodsSkuMgr) GetBatchFromWeight(weights []float64) (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`weight` IN (?)", weights).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 卖家id
func (obj *GoodsSkuMgr) GetFromSellerID(sellerID int) (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 卖家id
func (obj *GoodsSkuMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromSellerName 通过seller_name获取内容 卖家名称
func (obj *GoodsSkuMgr) GetFromSellerName(sellerName string) (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`seller_name` = ?", sellerName).Find(&results).Error

	return
}

// GetBatchFromSellerName 批量查找 卖家名称
func (obj *GoodsSkuMgr) GetBatchFromSellerName(sellerNames []string) (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`seller_name` IN (?)", sellerNames).Find(&results).Error

	return
}

// GetFromCategoryID 通过category_id获取内容 分类id
func (obj *GoodsSkuMgr) GetFromCategoryID(categoryID int) (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`category_id` = ?", categoryID).Find(&results).Error

	return
}

// GetBatchFromCategoryID 批量查找 分类id
func (obj *GoodsSkuMgr) GetBatchFromCategoryID(categoryIDs []int) (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`category_id` IN (?)", categoryIDs).Find(&results).Error

	return
}

// GetFromThumbnail 通过thumbnail获取内容 缩略图
func (obj *GoodsSkuMgr) GetFromThumbnail(thumbnail string) (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`thumbnail` = ?", thumbnail).Find(&results).Error

	return
}

// GetBatchFromThumbnail 批量查找 缩略图
func (obj *GoodsSkuMgr) GetBatchFromThumbnail(thumbnails []string) (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`thumbnail` IN (?)", thumbnails).Find(&results).Error

	return
}

// GetFromHashCode 通过hash_code获取内容 标识规格唯一性
func (obj *GoodsSkuMgr) GetFromHashCode(hashCode int) (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`hash_code` = ?", hashCode).Find(&results).Error

	return
}

// GetBatchFromHashCode 批量查找 标识规格唯一性
func (obj *GoodsSkuMgr) GetBatchFromHashCode(hashCodes []int) (results []*models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`hash_code` IN (?)", hashCodes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *GoodsSkuMgr) FetchByPrimaryKey(skuID int) (result models.EsGoodsSku, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`sku_id` = ?", skuID).First(&result).Error

	return
}
