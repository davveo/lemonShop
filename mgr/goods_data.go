package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsSssGoodsDataMgr struct {
	*_BaseMgr
}

// EsSssGoodsDataMgr open func
func EsSssGoodsDataMgr(db *gorm.DB) *_EsSssGoodsDataMgr {
	if db == nil {
		panic(fmt.Errorf("EsSssGoodsDataMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsSssGoodsDataMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_sss_goods_data"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsSssGoodsDataMgr) GetTableName() string {
	return "es_sss_goods_data"
}

// Reset 重置gorm会话
func (obj *_EsSssGoodsDataMgr) Reset() *_EsSssGoodsDataMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsSssGoodsDataMgr) Get() (result models.EsSssGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsData{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsSssGoodsDataMgr) Gets() (results []*models.EsSssGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsData{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsSssGoodsDataMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsData{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *_EsSssGoodsDataMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithGoodsID goods_id获取 商品id
func (obj *_EsSssGoodsDataMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithGoodsName goods_name获取 商品名称
func (obj *_EsSssGoodsDataMgr) WithGoodsName(goodsName string) Option {
	return optionFunc(func(o *options) { o.query["goods_name"] = goodsName })
}

// WithBrandID brand_id获取 品牌id
func (obj *_EsSssGoodsDataMgr) WithBrandID(brandID int) Option {
	return optionFunc(func(o *options) { o.query["brand_id"] = brandID })
}

// WithCategoryID category_id获取 分类id
func (obj *_EsSssGoodsDataMgr) WithCategoryID(categoryID int) Option {
	return optionFunc(func(o *options) { o.query["category_id"] = categoryID })
}

// WithCategoryPath category_path获取 分类路径
func (obj *_EsSssGoodsDataMgr) WithCategoryPath(categoryPath string) Option {
	return optionFunc(func(o *options) { o.query["category_path"] = categoryPath })
}

// WithSellerID seller_id获取 商家id
func (obj *_EsSssGoodsDataMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithShopCatID shop_cat_id获取 商家分类id
func (obj *_EsSssGoodsDataMgr) WithShopCatID(shopCatID int) Option {
	return optionFunc(func(o *options) { o.query["shop_cat_id"] = shopCatID })
}

// WithPrice price获取 商品价格
func (obj *_EsSssGoodsDataMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithFavoriteNum favorite_num获取 收藏数量
func (obj *_EsSssGoodsDataMgr) WithFavoriteNum(favoriteNum int) Option {
	return optionFunc(func(o *options) { o.query["favorite_num"] = favoriteNum })
}

// WithMarketEnable market_enable获取 是否上架
func (obj *_EsSssGoodsDataMgr) WithMarketEnable(marketEnable int16) Option {
	return optionFunc(func(o *options) { o.query["market_enable"] = marketEnable })
}

// GetByOption 功能选项模式获取
func (obj *_EsSssGoodsDataMgr) GetByOption(opts ...Option) (result models.EsSssGoodsData, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsData{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsSssGoodsDataMgr) GetByOptions(opts ...Option) (results []*models.EsSssGoodsData, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsData{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsSssGoodsDataMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsSssGoodsData, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsData{}).Where(options.query)
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
func (obj *_EsSssGoodsDataMgr) GetFromID(id int) (result models.EsSssGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsData{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *_EsSssGoodsDataMgr) GetBatchFromID(ids []int) (results []*models.EsSssGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsData{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 商品id
func (obj *_EsSssGoodsDataMgr) GetFromGoodsID(goodsID int) (results []*models.EsSssGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsData{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 商品id
func (obj *_EsSssGoodsDataMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsSssGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsData{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromGoodsName 通过goods_name获取内容 商品名称
func (obj *_EsSssGoodsDataMgr) GetFromGoodsName(goodsName string) (results []*models.EsSssGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsData{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}

// GetBatchFromGoodsName 批量查找 商品名称
func (obj *_EsSssGoodsDataMgr) GetBatchFromGoodsName(goodsNames []string) (results []*models.EsSssGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsData{}).Where("`goods_name` IN (?)", goodsNames).Find(&results).Error

	return
}

// GetFromBrandID 通过brand_id获取内容 品牌id
func (obj *_EsSssGoodsDataMgr) GetFromBrandID(brandID int) (results []*models.EsSssGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsData{}).Where("`brand_id` = ?", brandID).Find(&results).Error

	return
}

// GetBatchFromBrandID 批量查找 品牌id
func (obj *_EsSssGoodsDataMgr) GetBatchFromBrandID(brandIDs []int) (results []*models.EsSssGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsData{}).Where("`brand_id` IN (?)", brandIDs).Find(&results).Error

	return
}

// GetFromCategoryID 通过category_id获取内容 分类id
func (obj *_EsSssGoodsDataMgr) GetFromCategoryID(categoryID int) (results []*models.EsSssGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsData{}).Where("`category_id` = ?", categoryID).Find(&results).Error

	return
}

// GetBatchFromCategoryID 批量查找 分类id
func (obj *_EsSssGoodsDataMgr) GetBatchFromCategoryID(categoryIDs []int) (results []*models.EsSssGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsData{}).Where("`category_id` IN (?)", categoryIDs).Find(&results).Error

	return
}

// GetFromCategoryPath 通过category_path获取内容 分类路径
func (obj *_EsSssGoodsDataMgr) GetFromCategoryPath(categoryPath string) (results []*models.EsSssGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsData{}).Where("`category_path` = ?", categoryPath).Find(&results).Error

	return
}

// GetBatchFromCategoryPath 批量查找 分类路径
func (obj *_EsSssGoodsDataMgr) GetBatchFromCategoryPath(categoryPaths []string) (results []*models.EsSssGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsData{}).Where("`category_path` IN (?)", categoryPaths).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 商家id
func (obj *_EsSssGoodsDataMgr) GetFromSellerID(sellerID int) (results []*models.EsSssGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsData{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 商家id
func (obj *_EsSssGoodsDataMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsSssGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsData{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromShopCatID 通过shop_cat_id获取内容 商家分类id
func (obj *_EsSssGoodsDataMgr) GetFromShopCatID(shopCatID int) (results []*models.EsSssGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsData{}).Where("`shop_cat_id` = ?", shopCatID).Find(&results).Error

	return
}

// GetBatchFromShopCatID 批量查找 商家分类id
func (obj *_EsSssGoodsDataMgr) GetBatchFromShopCatID(shopCatIDs []int) (results []*models.EsSssGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsData{}).Where("`shop_cat_id` IN (?)", shopCatIDs).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 商品价格
func (obj *_EsSssGoodsDataMgr) GetFromPrice(price float64) (results []*models.EsSssGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsData{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 商品价格
func (obj *_EsSssGoodsDataMgr) GetBatchFromPrice(prices []float64) (results []*models.EsSssGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsData{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromFavoriteNum 通过favorite_num获取内容 收藏数量
func (obj *_EsSssGoodsDataMgr) GetFromFavoriteNum(favoriteNum int) (results []*models.EsSssGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsData{}).Where("`favorite_num` = ?", favoriteNum).Find(&results).Error

	return
}

// GetBatchFromFavoriteNum 批量查找 收藏数量
func (obj *_EsSssGoodsDataMgr) GetBatchFromFavoriteNum(favoriteNums []int) (results []*models.EsSssGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsData{}).Where("`favorite_num` IN (?)", favoriteNums).Find(&results).Error

	return
}

// GetFromMarketEnable 通过market_enable获取内容 是否上架
func (obj *_EsSssGoodsDataMgr) GetFromMarketEnable(marketEnable int16) (results []*models.EsSssGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsData{}).Where("`market_enable` = ?", marketEnable).Find(&results).Error

	return
}

// GetBatchFromMarketEnable 批量查找 是否上架
func (obj *_EsSssGoodsDataMgr) GetBatchFromMarketEnable(marketEnables []int16) (results []*models.EsSssGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsData{}).Where("`market_enable` IN (?)", marketEnables).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsSssGoodsDataMgr) FetchByPrimaryKey(id int) (result models.EsSssGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsData{}).Where("`id` = ?", id).First(&result).Error

	return
}
