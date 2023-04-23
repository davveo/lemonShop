package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsGoodsSkuMgr struct {
	*_BaseMgr
}

// EsGoodsSkuMgr open func
func EsGoodsSkuMgr(db *gorm.DB) *_EsGoodsSkuMgr {
	if db == nil {
		panic(fmt.Errorf("EsGoodsSkuMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsGoodsSkuMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_goods_sku"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsGoodsSkuMgr) GetTableName() string {
	return "es_goods_sku"
}

// Reset 重置gorm会话
func (obj *_EsGoodsSkuMgr) Reset() *_EsGoodsSkuMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsGoodsSkuMgr) Get() (result models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsGoodsSkuMgr) Gets() (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsGoodsSkuMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSkuID sku_id获取 主键
func (obj *_EsGoodsSkuMgr) WithSkuID(skuID int) Option {
	return optionFunc(func(o *options) { o.query["sku_id"] = skuID })
}

// WithGoodsID goods_id获取 商品id
func (obj *_EsGoodsSkuMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithGoodsName goods_name获取 商品名称
func (obj *_EsGoodsSkuMgr) WithGoodsName(goodsName string) Option {
	return optionFunc(func(o *options) { o.query["goods_name"] = goodsName })
}

// WithSn sn获取 商品编号
func (obj *_EsGoodsSkuMgr) WithSn(sn string) Option {
	return optionFunc(func(o *options) { o.query["sn"] = sn })
}

// WithQuantity quantity获取 库存
func (obj *_EsGoodsSkuMgr) WithQuantity(quantity int) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

// WithEnableQuantity enable_quantity获取 可用库存
func (obj *_EsGoodsSkuMgr) WithEnableQuantity(enableQuantity int) Option {
	return optionFunc(func(o *options) { o.query["enable_quantity"] = enableQuantity })
}

// WithPrice price获取 商品价格
func (obj *_EsGoodsSkuMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithSpecs specs获取 规格信息json
func (obj *_EsGoodsSkuMgr) WithSpecs(specs string) Option {
	return optionFunc(func(o *options) { o.query["specs"] = specs })
}

// WithCost cost获取 成本价格
func (obj *_EsGoodsSkuMgr) WithCost(cost float64) Option {
	return optionFunc(func(o *options) { o.query["cost"] = cost })
}

// WithWeight weight获取 重量
func (obj *_EsGoodsSkuMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithSellerID seller_id获取 卖家id
func (obj *_EsGoodsSkuMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithSellerName seller_name获取 卖家名称
func (obj *_EsGoodsSkuMgr) WithSellerName(sellerName string) Option {
	return optionFunc(func(o *options) { o.query["seller_name"] = sellerName })
}

// WithCategoryID category_id获取 分类id
func (obj *_EsGoodsSkuMgr) WithCategoryID(categoryID int) Option {
	return optionFunc(func(o *options) { o.query["category_id"] = categoryID })
}

// WithThumbnail thumbnail获取 缩略图
func (obj *_EsGoodsSkuMgr) WithThumbnail(thumbnail string) Option {
	return optionFunc(func(o *options) { o.query["thumbnail"] = thumbnail })
}

// WithHashCode hash_code获取 标识规格唯一性
func (obj *_EsGoodsSkuMgr) WithHashCode(hashCode int) Option {
	return optionFunc(func(o *options) { o.query["hash_code"] = hashCode })
}

// GetByOption 功能选项模式获取
func (obj *_EsGoodsSkuMgr) GetByOption(opts ...Option) (result models.EsGoodsSku, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsGoodsSkuMgr) GetByOptions(opts ...Option) (results []*models.EsGoodsSku, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsGoodsSkuMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsGoodsSku, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where(options.query)
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
func (obj *_EsGoodsSkuMgr) GetFromSkuID(skuID int) (result models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`sku_id` = ?", skuID).First(&result).Error

	return
}

// GetBatchFromSkuID 批量查找 主键
func (obj *_EsGoodsSkuMgr) GetBatchFromSkuID(skuIDs []int) (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`sku_id` IN (?)", skuIDs).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 商品id
func (obj *_EsGoodsSkuMgr) GetFromGoodsID(goodsID int) (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 商品id
func (obj *_EsGoodsSkuMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromGoodsName 通过goods_name获取内容 商品名称
func (obj *_EsGoodsSkuMgr) GetFromGoodsName(goodsName string) (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}

// GetBatchFromGoodsName 批量查找 商品名称
func (obj *_EsGoodsSkuMgr) GetBatchFromGoodsName(goodsNames []string) (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`goods_name` IN (?)", goodsNames).Find(&results).Error

	return
}

// GetFromSn 通过sn获取内容 商品编号
func (obj *_EsGoodsSkuMgr) GetFromSn(sn string) (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`sn` = ?", sn).Find(&results).Error

	return
}

// GetBatchFromSn 批量查找 商品编号
func (obj *_EsGoodsSkuMgr) GetBatchFromSn(sns []string) (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`sn` IN (?)", sns).Find(&results).Error

	return
}

// GetFromQuantity 通过quantity获取内容 库存
func (obj *_EsGoodsSkuMgr) GetFromQuantity(quantity int) (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`quantity` = ?", quantity).Find(&results).Error

	return
}

// GetBatchFromQuantity 批量查找 库存
func (obj *_EsGoodsSkuMgr) GetBatchFromQuantity(quantitys []int) (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`quantity` IN (?)", quantitys).Find(&results).Error

	return
}

// GetFromEnableQuantity 通过enable_quantity获取内容 可用库存
func (obj *_EsGoodsSkuMgr) GetFromEnableQuantity(enableQuantity int) (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`enable_quantity` = ?", enableQuantity).Find(&results).Error

	return
}

// GetBatchFromEnableQuantity 批量查找 可用库存
func (obj *_EsGoodsSkuMgr) GetBatchFromEnableQuantity(enableQuantitys []int) (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`enable_quantity` IN (?)", enableQuantitys).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 商品价格
func (obj *_EsGoodsSkuMgr) GetFromPrice(price float64) (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 商品价格
func (obj *_EsGoodsSkuMgr) GetBatchFromPrice(prices []float64) (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromSpecs 通过specs获取内容 规格信息json
func (obj *_EsGoodsSkuMgr) GetFromSpecs(specs string) (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`specs` = ?", specs).Find(&results).Error

	return
}

// GetBatchFromSpecs 批量查找 规格信息json
func (obj *_EsGoodsSkuMgr) GetBatchFromSpecs(specss []string) (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`specs` IN (?)", specss).Find(&results).Error

	return
}

// GetFromCost 通过cost获取内容 成本价格
func (obj *_EsGoodsSkuMgr) GetFromCost(cost float64) (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`cost` = ?", cost).Find(&results).Error

	return
}

// GetBatchFromCost 批量查找 成本价格
func (obj *_EsGoodsSkuMgr) GetBatchFromCost(costs []float64) (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`cost` IN (?)", costs).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 重量
func (obj *_EsGoodsSkuMgr) GetFromWeight(weight float64) (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`weight` = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量查找 重量
func (obj *_EsGoodsSkuMgr) GetBatchFromWeight(weights []float64) (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`weight` IN (?)", weights).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 卖家id
func (obj *_EsGoodsSkuMgr) GetFromSellerID(sellerID int) (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 卖家id
func (obj *_EsGoodsSkuMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromSellerName 通过seller_name获取内容 卖家名称
func (obj *_EsGoodsSkuMgr) GetFromSellerName(sellerName string) (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`seller_name` = ?", sellerName).Find(&results).Error

	return
}

// GetBatchFromSellerName 批量查找 卖家名称
func (obj *_EsGoodsSkuMgr) GetBatchFromSellerName(sellerNames []string) (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`seller_name` IN (?)", sellerNames).Find(&results).Error

	return
}

// GetFromCategoryID 通过category_id获取内容 分类id
func (obj *_EsGoodsSkuMgr) GetFromCategoryID(categoryID int) (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`category_id` = ?", categoryID).Find(&results).Error

	return
}

// GetBatchFromCategoryID 批量查找 分类id
func (obj *_EsGoodsSkuMgr) GetBatchFromCategoryID(categoryIDs []int) (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`category_id` IN (?)", categoryIDs).Find(&results).Error

	return
}

// GetFromThumbnail 通过thumbnail获取内容 缩略图
func (obj *_EsGoodsSkuMgr) GetFromThumbnail(thumbnail string) (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`thumbnail` = ?", thumbnail).Find(&results).Error

	return
}

// GetBatchFromThumbnail 批量查找 缩略图
func (obj *_EsGoodsSkuMgr) GetBatchFromThumbnail(thumbnails []string) (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`thumbnail` IN (?)", thumbnails).Find(&results).Error

	return
}

// GetFromHashCode 通过hash_code获取内容 标识规格唯一性
func (obj *_EsGoodsSkuMgr) GetFromHashCode(hashCode int) (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`hash_code` = ?", hashCode).Find(&results).Error

	return
}

// GetBatchFromHashCode 批量查找 标识规格唯一性
func (obj *_EsGoodsSkuMgr) GetBatchFromHashCode(hashCodes []int) (results []*models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`hash_code` IN (?)", hashCodes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsGoodsSkuMgr) FetchByPrimaryKey(skuID int) (result models.EsGoodsSku, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsSku{}).Where("`sku_id` = ?", skuID).First(&result).Error

	return
}
