package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsDraftGoodsMgr struct {
	*_BaseMgr
}

// EsDraftGoodsMgr open func
func EsDraftGoodsMgr(db *gorm.DB) *_EsDraftGoodsMgr {
	if db == nil {
		panic(fmt.Errorf("EsDraftGoodsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsDraftGoodsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_draft_goods"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsDraftGoodsMgr) GetTableName() string {
	return "es_draft_goods"
}

// Reset 重置gorm会话
func (obj *_EsDraftGoodsMgr) Reset() *_EsDraftGoodsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsDraftGoodsMgr) Get() (result models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsDraftGoodsMgr) Gets() (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsDraftGoodsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithDraftGoodsID draft_goods_id获取 草稿商品id
func (obj *_EsDraftGoodsMgr) WithDraftGoodsID(draftGoodsID int) Option {
	return optionFunc(func(o *options) { o.query["draft_goods_id"] = draftGoodsID })
}

// WithQuantity quantity获取 商品总库存
func (obj *_EsDraftGoodsMgr) WithQuantity(quantity int) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

// WithOriginal original获取 商品原始图片
func (obj *_EsDraftGoodsMgr) WithOriginal(original string) Option {
	return optionFunc(func(o *options) { o.query["original"] = original })
}

// WithSellerID seller_id获取 商品所属卖家ID
func (obj *_EsDraftGoodsMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithShopCatID shop_cat_id获取 商品所属店铺类目ID
func (obj *_EsDraftGoodsMgr) WithShopCatID(shopCatID int) Option {
	return optionFunc(func(o *options) { o.query["shop_cat_id"] = shopCatID })
}

// WithTemplateID template_id获取 商品运费模板ID
func (obj *_EsDraftGoodsMgr) WithTemplateID(templateID int) Option {
	return optionFunc(func(o *options) { o.query["template_id"] = templateID })
}

// WithGoodsTransfeeCharge goods_transfee_charge获取 是否为买家承担运费
func (obj *_EsDraftGoodsMgr) WithGoodsTransfeeCharge(goodsTransfeeCharge int) Option {
	return optionFunc(func(o *options) { o.query["goods_transfee_charge"] = goodsTransfeeCharge })
}

// WithSellerName seller_name获取 商品所属店铺名称
func (obj *_EsDraftGoodsMgr) WithSellerName(sellerName string) Option {
	return optionFunc(func(o *options) { o.query["seller_name"] = sellerName })
}

// WithPageTitle page_title获取 seo 标题
func (obj *_EsDraftGoodsMgr) WithPageTitle(pageTitle string) Option {
	return optionFunc(func(o *options) { o.query["page_title"] = pageTitle })
}

// WithMetaKeywords meta_keywords获取 seo关键字
func (obj *_EsDraftGoodsMgr) WithMetaKeywords(metaKeywords string) Option {
	return optionFunc(func(o *options) { o.query["meta_keywords"] = metaKeywords })
}

// WithMetaDescription meta_description获取 seo描述
func (obj *_EsDraftGoodsMgr) WithMetaDescription(metaDescription string) Option {
	return optionFunc(func(o *options) { o.query["meta_description"] = metaDescription })
}

// WithCreateTime create_time获取 商品添加时间
func (obj *_EsDraftGoodsMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithHaveSpec have_spec获取 是否开启规格
func (obj *_EsDraftGoodsMgr) WithHaveSpec(haveSpec int) Option {
	return optionFunc(func(o *options) { o.query["have_spec"] = haveSpec })
}

// WithSn sn获取 商品编号
func (obj *_EsDraftGoodsMgr) WithSn(sn string) Option {
	return optionFunc(func(o *options) { o.query["sn"] = sn })
}

// WithGoodsName goods_name获取 商品名称
func (obj *_EsDraftGoodsMgr) WithGoodsName(goodsName string) Option {
	return optionFunc(func(o *options) { o.query["goods_name"] = goodsName })
}

// WithBrandID brand_id获取 商品品牌ID
func (obj *_EsDraftGoodsMgr) WithBrandID(brandID int) Option {
	return optionFunc(func(o *options) { o.query["brand_id"] = brandID })
}

// WithCategoryID category_id获取 商品分类ID
func (obj *_EsDraftGoodsMgr) WithCategoryID(categoryID int) Option {
	return optionFunc(func(o *options) { o.query["category_id"] = categoryID })
}

// WithWeight weight获取 商品重量
func (obj *_EsDraftGoodsMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithIntro intro获取 商品详情
func (obj *_EsDraftGoodsMgr) WithIntro(intro string) Option {
	return optionFunc(func(o *options) { o.query["intro"] = intro })
}

// WithPrice price获取 商品价格
func (obj *_EsDraftGoodsMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithCost cost获取 商品成本价
func (obj *_EsDraftGoodsMgr) WithCost(cost float64) Option {
	return optionFunc(func(o *options) { o.query["cost"] = cost })
}

// WithMktprice mktprice获取 商品市场价
func (obj *_EsDraftGoodsMgr) WithMktprice(mktprice float64) Option {
	return optionFunc(func(o *options) { o.query["mktprice"] = mktprice })
}

// WithGoodsType goods_type获取 商品类型
func (obj *_EsDraftGoodsMgr) WithGoodsType(goodsType string) Option {
	return optionFunc(func(o *options) { o.query["goods_type"] = goodsType })
}

// WithExchangeMoney exchange_money获取 积分商品需要的金额
func (obj *_EsDraftGoodsMgr) WithExchangeMoney(exchangeMoney float64) Option {
	return optionFunc(func(o *options) { o.query["exchange_money"] = exchangeMoney })
}

// WithExchangePoint exchange_point获取 积分商品需要的积分
func (obj *_EsDraftGoodsMgr) WithExchangePoint(exchangePoint int) Option {
	return optionFunc(func(o *options) { o.query["exchange_point"] = exchangePoint })
}

// WithExchangeCategoryID exchange_category_id获取 积分商品的分类id
func (obj *_EsDraftGoodsMgr) WithExchangeCategoryID(exchangeCategoryID int) Option {
	return optionFunc(func(o *options) { o.query["exchange_category_id"] = exchangeCategoryID })
}

// GetByOption 功能选项模式获取
func (obj *_EsDraftGoodsMgr) GetByOption(opts ...Option) (result models.EsDraftGoods, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsDraftGoodsMgr) GetByOptions(opts ...Option) (results []*models.EsDraftGoods, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsDraftGoodsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsDraftGoods, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where(options.query)
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

// GetFromDraftGoodsID 通过draft_goods_id获取内容 草稿商品id
func (obj *_EsDraftGoodsMgr) GetFromDraftGoodsID(draftGoodsID int) (result models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`draft_goods_id` = ?", draftGoodsID).First(&result).Error

	return
}

// GetBatchFromDraftGoodsID 批量查找 草稿商品id
func (obj *_EsDraftGoodsMgr) GetBatchFromDraftGoodsID(draftGoodsIDs []int) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`draft_goods_id` IN (?)", draftGoodsIDs).Find(&results).Error

	return
}

// GetFromQuantity 通过quantity获取内容 商品总库存
func (obj *_EsDraftGoodsMgr) GetFromQuantity(quantity int) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`quantity` = ?", quantity).Find(&results).Error

	return
}

// GetBatchFromQuantity 批量查找 商品总库存
func (obj *_EsDraftGoodsMgr) GetBatchFromQuantity(quantitys []int) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`quantity` IN (?)", quantitys).Find(&results).Error

	return
}

// GetFromOriginal 通过original获取内容 商品原始图片
func (obj *_EsDraftGoodsMgr) GetFromOriginal(original string) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`original` = ?", original).Find(&results).Error

	return
}

// GetBatchFromOriginal 批量查找 商品原始图片
func (obj *_EsDraftGoodsMgr) GetBatchFromOriginal(originals []string) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`original` IN (?)", originals).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 商品所属卖家ID
func (obj *_EsDraftGoodsMgr) GetFromSellerID(sellerID int) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 商品所属卖家ID
func (obj *_EsDraftGoodsMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromShopCatID 通过shop_cat_id获取内容 商品所属店铺类目ID
func (obj *_EsDraftGoodsMgr) GetFromShopCatID(shopCatID int) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`shop_cat_id` = ?", shopCatID).Find(&results).Error

	return
}

// GetBatchFromShopCatID 批量查找 商品所属店铺类目ID
func (obj *_EsDraftGoodsMgr) GetBatchFromShopCatID(shopCatIDs []int) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`shop_cat_id` IN (?)", shopCatIDs).Find(&results).Error

	return
}

// GetFromTemplateID 通过template_id获取内容 商品运费模板ID
func (obj *_EsDraftGoodsMgr) GetFromTemplateID(templateID int) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`template_id` = ?", templateID).Find(&results).Error

	return
}

// GetBatchFromTemplateID 批量查找 商品运费模板ID
func (obj *_EsDraftGoodsMgr) GetBatchFromTemplateID(templateIDs []int) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`template_id` IN (?)", templateIDs).Find(&results).Error

	return
}

// GetFromGoodsTransfeeCharge 通过goods_transfee_charge获取内容 是否为买家承担运费
func (obj *_EsDraftGoodsMgr) GetFromGoodsTransfeeCharge(goodsTransfeeCharge int) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`goods_transfee_charge` = ?", goodsTransfeeCharge).Find(&results).Error

	return
}

// GetBatchFromGoodsTransfeeCharge 批量查找 是否为买家承担运费
func (obj *_EsDraftGoodsMgr) GetBatchFromGoodsTransfeeCharge(goodsTransfeeCharges []int) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`goods_transfee_charge` IN (?)", goodsTransfeeCharges).Find(&results).Error

	return
}

// GetFromSellerName 通过seller_name获取内容 商品所属店铺名称
func (obj *_EsDraftGoodsMgr) GetFromSellerName(sellerName string) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`seller_name` = ?", sellerName).Find(&results).Error

	return
}

// GetBatchFromSellerName 批量查找 商品所属店铺名称
func (obj *_EsDraftGoodsMgr) GetBatchFromSellerName(sellerNames []string) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`seller_name` IN (?)", sellerNames).Find(&results).Error

	return
}

// GetFromPageTitle 通过page_title获取内容 seo 标题
func (obj *_EsDraftGoodsMgr) GetFromPageTitle(pageTitle string) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`page_title` = ?", pageTitle).Find(&results).Error

	return
}

// GetBatchFromPageTitle 批量查找 seo 标题
func (obj *_EsDraftGoodsMgr) GetBatchFromPageTitle(pageTitles []string) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`page_title` IN (?)", pageTitles).Find(&results).Error

	return
}

// GetFromMetaKeywords 通过meta_keywords获取内容 seo关键字
func (obj *_EsDraftGoodsMgr) GetFromMetaKeywords(metaKeywords string) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`meta_keywords` = ?", metaKeywords).Find(&results).Error

	return
}

// GetBatchFromMetaKeywords 批量查找 seo关键字
func (obj *_EsDraftGoodsMgr) GetBatchFromMetaKeywords(metaKeywordss []string) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`meta_keywords` IN (?)", metaKeywordss).Find(&results).Error

	return
}

// GetFromMetaDescription 通过meta_description获取内容 seo描述
func (obj *_EsDraftGoodsMgr) GetFromMetaDescription(metaDescription string) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`meta_description` = ?", metaDescription).Find(&results).Error

	return
}

// GetBatchFromMetaDescription 批量查找 seo描述
func (obj *_EsDraftGoodsMgr) GetBatchFromMetaDescription(metaDescriptions []string) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`meta_description` IN (?)", metaDescriptions).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 商品添加时间
func (obj *_EsDraftGoodsMgr) GetFromCreateTime(createTime int64) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 商品添加时间
func (obj *_EsDraftGoodsMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromHaveSpec 通过have_spec获取内容 是否开启规格
func (obj *_EsDraftGoodsMgr) GetFromHaveSpec(haveSpec int) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`have_spec` = ?", haveSpec).Find(&results).Error

	return
}

// GetBatchFromHaveSpec 批量查找 是否开启规格
func (obj *_EsDraftGoodsMgr) GetBatchFromHaveSpec(haveSpecs []int) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`have_spec` IN (?)", haveSpecs).Find(&results).Error

	return
}

// GetFromSn 通过sn获取内容 商品编号
func (obj *_EsDraftGoodsMgr) GetFromSn(sn string) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`sn` = ?", sn).Find(&results).Error

	return
}

// GetBatchFromSn 批量查找 商品编号
func (obj *_EsDraftGoodsMgr) GetBatchFromSn(sns []string) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`sn` IN (?)", sns).Find(&results).Error

	return
}

// GetFromGoodsName 通过goods_name获取内容 商品名称
func (obj *_EsDraftGoodsMgr) GetFromGoodsName(goodsName string) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}

// GetBatchFromGoodsName 批量查找 商品名称
func (obj *_EsDraftGoodsMgr) GetBatchFromGoodsName(goodsNames []string) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`goods_name` IN (?)", goodsNames).Find(&results).Error

	return
}

// GetFromBrandID 通过brand_id获取内容 商品品牌ID
func (obj *_EsDraftGoodsMgr) GetFromBrandID(brandID int) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`brand_id` = ?", brandID).Find(&results).Error

	return
}

// GetBatchFromBrandID 批量查找 商品品牌ID
func (obj *_EsDraftGoodsMgr) GetBatchFromBrandID(brandIDs []int) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`brand_id` IN (?)", brandIDs).Find(&results).Error

	return
}

// GetFromCategoryID 通过category_id获取内容 商品分类ID
func (obj *_EsDraftGoodsMgr) GetFromCategoryID(categoryID int) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`category_id` = ?", categoryID).Find(&results).Error

	return
}

// GetBatchFromCategoryID 批量查找 商品分类ID
func (obj *_EsDraftGoodsMgr) GetBatchFromCategoryID(categoryIDs []int) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`category_id` IN (?)", categoryIDs).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 商品重量
func (obj *_EsDraftGoodsMgr) GetFromWeight(weight float64) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`weight` = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量查找 商品重量
func (obj *_EsDraftGoodsMgr) GetBatchFromWeight(weights []float64) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`weight` IN (?)", weights).Find(&results).Error

	return
}

// GetFromIntro 通过intro获取内容 商品详情
func (obj *_EsDraftGoodsMgr) GetFromIntro(intro string) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`intro` = ?", intro).Find(&results).Error

	return
}

// GetBatchFromIntro 批量查找 商品详情
func (obj *_EsDraftGoodsMgr) GetBatchFromIntro(intros []string) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`intro` IN (?)", intros).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 商品价格
func (obj *_EsDraftGoodsMgr) GetFromPrice(price float64) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 商品价格
func (obj *_EsDraftGoodsMgr) GetBatchFromPrice(prices []float64) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromCost 通过cost获取内容 商品成本价
func (obj *_EsDraftGoodsMgr) GetFromCost(cost float64) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`cost` = ?", cost).Find(&results).Error

	return
}

// GetBatchFromCost 批量查找 商品成本价
func (obj *_EsDraftGoodsMgr) GetBatchFromCost(costs []float64) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`cost` IN (?)", costs).Find(&results).Error

	return
}

// GetFromMktprice 通过mktprice获取内容 商品市场价
func (obj *_EsDraftGoodsMgr) GetFromMktprice(mktprice float64) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`mktprice` = ?", mktprice).Find(&results).Error

	return
}

// GetBatchFromMktprice 批量查找 商品市场价
func (obj *_EsDraftGoodsMgr) GetBatchFromMktprice(mktprices []float64) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`mktprice` IN (?)", mktprices).Find(&results).Error

	return
}

// GetFromGoodsType 通过goods_type获取内容 商品类型
func (obj *_EsDraftGoodsMgr) GetFromGoodsType(goodsType string) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`goods_type` = ?", goodsType).Find(&results).Error

	return
}

// GetBatchFromGoodsType 批量查找 商品类型
func (obj *_EsDraftGoodsMgr) GetBatchFromGoodsType(goodsTypes []string) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`goods_type` IN (?)", goodsTypes).Find(&results).Error

	return
}

// GetFromExchangeMoney 通过exchange_money获取内容 积分商品需要的金额
func (obj *_EsDraftGoodsMgr) GetFromExchangeMoney(exchangeMoney float64) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`exchange_money` = ?", exchangeMoney).Find(&results).Error

	return
}

// GetBatchFromExchangeMoney 批量查找 积分商品需要的金额
func (obj *_EsDraftGoodsMgr) GetBatchFromExchangeMoney(exchangeMoneys []float64) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`exchange_money` IN (?)", exchangeMoneys).Find(&results).Error

	return
}

// GetFromExchangePoint 通过exchange_point获取内容 积分商品需要的积分
func (obj *_EsDraftGoodsMgr) GetFromExchangePoint(exchangePoint int) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`exchange_point` = ?", exchangePoint).Find(&results).Error

	return
}

// GetBatchFromExchangePoint 批量查找 积分商品需要的积分
func (obj *_EsDraftGoodsMgr) GetBatchFromExchangePoint(exchangePoints []int) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`exchange_point` IN (?)", exchangePoints).Find(&results).Error

	return
}

// GetFromExchangeCategoryID 通过exchange_category_id获取内容 积分商品的分类id
func (obj *_EsDraftGoodsMgr) GetFromExchangeCategoryID(exchangeCategoryID int) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`exchange_category_id` = ?", exchangeCategoryID).Find(&results).Error

	return
}

// GetBatchFromExchangeCategoryID 批量查找 积分商品的分类id
func (obj *_EsDraftGoodsMgr) GetBatchFromExchangeCategoryID(exchangeCategoryIDs []int) (results []*models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`exchange_category_id` IN (?)", exchangeCategoryIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsDraftGoodsMgr) FetchByPrimaryKey(draftGoodsID int) (result models.EsDraftGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`draft_goods_id` = ?", draftGoodsID).First(&result).Error

	return
}
