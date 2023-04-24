package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type DraftGoodsMgr struct {
	*_BaseMgr
}

// NewDraftGoodsMgr open func
func NewDraftGoodsMgr(db db.Repo) *DraftGoodsMgr {
	if db == nil {
		panic(fmt.Errorf("NewDraftGoodsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &DraftGoodsMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_category_spec"),
		wdb:       db.GetDbW().Table("es_category_spec"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *DraftGoodsMgr) GetTableName() string {
	return "es_draft_goods"
}

// Reset 重置gorm会话
func (obj *DraftGoodsMgr) Reset() *DraftGoodsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *DraftGoodsMgr) Get() (result models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *DraftGoodsMgr) Gets() (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *DraftGoodsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithDraftGoodsID draft_goods_id获取 草稿商品id
func (obj *DraftGoodsMgr) WithDraftGoodsID(draftGoodsID int) Option {
	return optionFunc(func(o *options) { o.query["draft_goods_id"] = draftGoodsID })
}

// WithQuantity quantity获取 商品总库存
func (obj *DraftGoodsMgr) WithQuantity(quantity int) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

// WithOriginal original获取 商品原始图片
func (obj *DraftGoodsMgr) WithOriginal(original string) Option {
	return optionFunc(func(o *options) { o.query["original"] = original })
}

// WithSellerID seller_id获取 商品所属卖家ID
func (obj *DraftGoodsMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithShopCatID shop_cat_id获取 商品所属店铺类目ID
func (obj *DraftGoodsMgr) WithShopCatID(shopCatID int) Option {
	return optionFunc(func(o *options) { o.query["shop_cat_id"] = shopCatID })
}

// WithTemplateID template_id获取 商品运费模板ID
func (obj *DraftGoodsMgr) WithTemplateID(templateID int) Option {
	return optionFunc(func(o *options) { o.query["template_id"] = templateID })
}

// WithGoodsTransfeeCharge goods_transfee_charge获取 是否为买家承担运费
func (obj *DraftGoodsMgr) WithGoodsTransfeeCharge(goodsTransfeeCharge int) Option {
	return optionFunc(func(o *options) { o.query["goods_transfee_charge"] = goodsTransfeeCharge })
}

// WithSellerName seller_name获取 商品所属店铺名称
func (obj *DraftGoodsMgr) WithSellerName(sellerName string) Option {
	return optionFunc(func(o *options) { o.query["seller_name"] = sellerName })
}

// WithPageTitle page_title获取 seo 标题
func (obj *DraftGoodsMgr) WithPageTitle(pageTitle string) Option {
	return optionFunc(func(o *options) { o.query["page_title"] = pageTitle })
}

// WithMetaKeywords meta_keywords获取 seo关键字
func (obj *DraftGoodsMgr) WithMetaKeywords(metaKeywords string) Option {
	return optionFunc(func(o *options) { o.query["meta_keywords"] = metaKeywords })
}

// WithMetaDescription meta_description获取 seo描述
func (obj *DraftGoodsMgr) WithMetaDescription(metaDescription string) Option {
	return optionFunc(func(o *options) { o.query["meta_description"] = metaDescription })
}

// WithCreateTime create_time获取 商品添加时间
func (obj *DraftGoodsMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithHaveSpec have_spec获取 是否开启规格
func (obj *DraftGoodsMgr) WithHaveSpec(haveSpec int) Option {
	return optionFunc(func(o *options) { o.query["have_spec"] = haveSpec })
}

// WithSn sn获取 商品编号
func (obj *DraftGoodsMgr) WithSn(sn string) Option {
	return optionFunc(func(o *options) { o.query["sn"] = sn })
}

// WithGoodsName goods_name获取 商品名称
func (obj *DraftGoodsMgr) WithGoodsName(goodsName string) Option {
	return optionFunc(func(o *options) { o.query["goods_name"] = goodsName })
}

// WithBrandID brand_id获取 商品品牌ID
func (obj *DraftGoodsMgr) WithBrandID(brandID int) Option {
	return optionFunc(func(o *options) { o.query["brand_id"] = brandID })
}

// WithCategoryID category_id获取 商品分类ID
func (obj *DraftGoodsMgr) WithCategoryID(categoryID int) Option {
	return optionFunc(func(o *options) { o.query["category_id"] = categoryID })
}

// WithWeight weight获取 商品重量
func (obj *DraftGoodsMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithIntro intro获取 商品详情
func (obj *DraftGoodsMgr) WithIntro(intro string) Option {
	return optionFunc(func(o *options) { o.query["intro"] = intro })
}

// WithPrice price获取 商品价格
func (obj *DraftGoodsMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithCost cost获取 商品成本价
func (obj *DraftGoodsMgr) WithCost(cost float64) Option {
	return optionFunc(func(o *options) { o.query["cost"] = cost })
}

// WithMktprice mktprice获取 商品市场价
func (obj *DraftGoodsMgr) WithMktprice(mktprice float64) Option {
	return optionFunc(func(o *options) { o.query["mktprice"] = mktprice })
}

// WithGoodsType goods_type获取 商品类型
func (obj *DraftGoodsMgr) WithGoodsType(goodsType string) Option {
	return optionFunc(func(o *options) { o.query["goods_type"] = goodsType })
}

// WithExchangeMoney exchange_money获取 积分商品需要的金额
func (obj *DraftGoodsMgr) WithExchangeMoney(exchangeMoney float64) Option {
	return optionFunc(func(o *options) { o.query["exchange_money"] = exchangeMoney })
}

// WithExchangePoint exchange_point获取 积分商品需要的积分
func (obj *DraftGoodsMgr) WithExchangePoint(exchangePoint int) Option {
	return optionFunc(func(o *options) { o.query["exchange_point"] = exchangePoint })
}

// WithExchangeCategoryID exchange_category_id获取 积分商品的分类id
func (obj *DraftGoodsMgr) WithExchangeCategoryID(exchangeCategoryID int) Option {
	return optionFunc(func(o *options) { o.query["exchange_category_id"] = exchangeCategoryID })
}

// GetByOption 功能选项模式获取
func (obj *DraftGoodsMgr) GetByOption(opts ...Option) (result models.EsDraftGoods, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *DraftGoodsMgr) GetByOptions(opts ...Option) (results []*models.EsDraftGoods, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *DraftGoodsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsDraftGoods, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where(options.query)
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
func (obj *DraftGoodsMgr) GetFromDraftGoodsID(draftGoodsID int) (result models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`draft_goods_id` = ?", draftGoodsID).First(&result).Error

	return
}

// GetBatchFromDraftGoodsID 批量查找 草稿商品id
func (obj *DraftGoodsMgr) GetBatchFromDraftGoodsID(draftGoodsIDs []int) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`draft_goods_id` IN (?)", draftGoodsIDs).Find(&results).Error

	return
}

// GetFromQuantity 通过quantity获取内容 商品总库存
func (obj *DraftGoodsMgr) GetFromQuantity(quantity int) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`quantity` = ?", quantity).Find(&results).Error

	return
}

// GetBatchFromQuantity 批量查找 商品总库存
func (obj *DraftGoodsMgr) GetBatchFromQuantity(quantitys []int) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`quantity` IN (?)", quantitys).Find(&results).Error

	return
}

// GetFromOriginal 通过original获取内容 商品原始图片
func (obj *DraftGoodsMgr) GetFromOriginal(original string) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`original` = ?", original).Find(&results).Error

	return
}

// GetBatchFromOriginal 批量查找 商品原始图片
func (obj *DraftGoodsMgr) GetBatchFromOriginal(originals []string) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`original` IN (?)", originals).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 商品所属卖家ID
func (obj *DraftGoodsMgr) GetFromSellerID(sellerID int) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 商品所属卖家ID
func (obj *DraftGoodsMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromShopCatID 通过shop_cat_id获取内容 商品所属店铺类目ID
func (obj *DraftGoodsMgr) GetFromShopCatID(shopCatID int) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`shop_cat_id` = ?", shopCatID).Find(&results).Error

	return
}

// GetBatchFromShopCatID 批量查找 商品所属店铺类目ID
func (obj *DraftGoodsMgr) GetBatchFromShopCatID(shopCatIDs []int) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`shop_cat_id` IN (?)", shopCatIDs).Find(&results).Error

	return
}

// GetFromTemplateID 通过template_id获取内容 商品运费模板ID
func (obj *DraftGoodsMgr) GetFromTemplateID(templateID int) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`template_id` = ?", templateID).Find(&results).Error

	return
}

// GetBatchFromTemplateID 批量查找 商品运费模板ID
func (obj *DraftGoodsMgr) GetBatchFromTemplateID(templateIDs []int) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`template_id` IN (?)", templateIDs).Find(&results).Error

	return
}

// GetFromGoodsTransfeeCharge 通过goods_transfee_charge获取内容 是否为买家承担运费
func (obj *DraftGoodsMgr) GetFromGoodsTransfeeCharge(goodsTransfeeCharge int) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`goods_transfee_charge` = ?", goodsTransfeeCharge).Find(&results).Error

	return
}

// GetBatchFromGoodsTransfeeCharge 批量查找 是否为买家承担运费
func (obj *DraftGoodsMgr) GetBatchFromGoodsTransfeeCharge(goodsTransfeeCharges []int) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`goods_transfee_charge` IN (?)", goodsTransfeeCharges).Find(&results).Error

	return
}

// GetFromSellerName 通过seller_name获取内容 商品所属店铺名称
func (obj *DraftGoodsMgr) GetFromSellerName(sellerName string) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`seller_name` = ?", sellerName).Find(&results).Error

	return
}

// GetBatchFromSellerName 批量查找 商品所属店铺名称
func (obj *DraftGoodsMgr) GetBatchFromSellerName(sellerNames []string) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`seller_name` IN (?)", sellerNames).Find(&results).Error

	return
}

// GetFromPageTitle 通过page_title获取内容 seo 标题
func (obj *DraftGoodsMgr) GetFromPageTitle(pageTitle string) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`page_title` = ?", pageTitle).Find(&results).Error

	return
}

// GetBatchFromPageTitle 批量查找 seo 标题
func (obj *DraftGoodsMgr) GetBatchFromPageTitle(pageTitles []string) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`page_title` IN (?)", pageTitles).Find(&results).Error

	return
}

// GetFromMetaKeywords 通过meta_keywords获取内容 seo关键字
func (obj *DraftGoodsMgr) GetFromMetaKeywords(metaKeywords string) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`meta_keywords` = ?", metaKeywords).Find(&results).Error

	return
}

// GetBatchFromMetaKeywords 批量查找 seo关键字
func (obj *DraftGoodsMgr) GetBatchFromMetaKeywords(metaKeywordss []string) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`meta_keywords` IN (?)", metaKeywordss).Find(&results).Error

	return
}

// GetFromMetaDescription 通过meta_description获取内容 seo描述
func (obj *DraftGoodsMgr) GetFromMetaDescription(metaDescription string) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`meta_description` = ?", metaDescription).Find(&results).Error

	return
}

// GetBatchFromMetaDescription 批量查找 seo描述
func (obj *DraftGoodsMgr) GetBatchFromMetaDescription(metaDescriptions []string) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`meta_description` IN (?)", metaDescriptions).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 商品添加时间
func (obj *DraftGoodsMgr) GetFromCreateTime(createTime int64) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 商品添加时间
func (obj *DraftGoodsMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromHaveSpec 通过have_spec获取内容 是否开启规格
func (obj *DraftGoodsMgr) GetFromHaveSpec(haveSpec int) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`have_spec` = ?", haveSpec).Find(&results).Error

	return
}

// GetBatchFromHaveSpec 批量查找 是否开启规格
func (obj *DraftGoodsMgr) GetBatchFromHaveSpec(haveSpecs []int) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`have_spec` IN (?)", haveSpecs).Find(&results).Error

	return
}

// GetFromSn 通过sn获取内容 商品编号
func (obj *DraftGoodsMgr) GetFromSn(sn string) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`sn` = ?", sn).Find(&results).Error

	return
}

// GetBatchFromSn 批量查找 商品编号
func (obj *DraftGoodsMgr) GetBatchFromSn(sns []string) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`sn` IN (?)", sns).Find(&results).Error

	return
}

// GetFromGoodsName 通过goods_name获取内容 商品名称
func (obj *DraftGoodsMgr) GetFromGoodsName(goodsName string) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}

// GetBatchFromGoodsName 批量查找 商品名称
func (obj *DraftGoodsMgr) GetBatchFromGoodsName(goodsNames []string) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`goods_name` IN (?)", goodsNames).Find(&results).Error

	return
}

// GetFromBrandID 通过brand_id获取内容 商品品牌ID
func (obj *DraftGoodsMgr) GetFromBrandID(brandID int) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`brand_id` = ?", brandID).Find(&results).Error

	return
}

// GetBatchFromBrandID 批量查找 商品品牌ID
func (obj *DraftGoodsMgr) GetBatchFromBrandID(brandIDs []int) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`brand_id` IN (?)", brandIDs).Find(&results).Error

	return
}

// GetFromCategoryID 通过category_id获取内容 商品分类ID
func (obj *DraftGoodsMgr) GetFromCategoryID(categoryID int) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`category_id` = ?", categoryID).Find(&results).Error

	return
}

// GetBatchFromCategoryID 批量查找 商品分类ID
func (obj *DraftGoodsMgr) GetBatchFromCategoryID(categoryIDs []int) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`category_id` IN (?)", categoryIDs).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 商品重量
func (obj *DraftGoodsMgr) GetFromWeight(weight float64) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`weight` = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量查找 商品重量
func (obj *DraftGoodsMgr) GetBatchFromWeight(weights []float64) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`weight` IN (?)", weights).Find(&results).Error

	return
}

// GetFromIntro 通过intro获取内容 商品详情
func (obj *DraftGoodsMgr) GetFromIntro(intro string) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`intro` = ?", intro).Find(&results).Error

	return
}

// GetBatchFromIntro 批量查找 商品详情
func (obj *DraftGoodsMgr) GetBatchFromIntro(intros []string) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`intro` IN (?)", intros).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 商品价格
func (obj *DraftGoodsMgr) GetFromPrice(price float64) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 商品价格
func (obj *DraftGoodsMgr) GetBatchFromPrice(prices []float64) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromCost 通过cost获取内容 商品成本价
func (obj *DraftGoodsMgr) GetFromCost(cost float64) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`cost` = ?", cost).Find(&results).Error

	return
}

// GetBatchFromCost 批量查找 商品成本价
func (obj *DraftGoodsMgr) GetBatchFromCost(costs []float64) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`cost` IN (?)", costs).Find(&results).Error

	return
}

// GetFromMktprice 通过mktprice获取内容 商品市场价
func (obj *DraftGoodsMgr) GetFromMktprice(mktprice float64) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`mktprice` = ?", mktprice).Find(&results).Error

	return
}

// GetBatchFromMktprice 批量查找 商品市场价
func (obj *DraftGoodsMgr) GetBatchFromMktprice(mktprices []float64) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`mktprice` IN (?)", mktprices).Find(&results).Error

	return
}

// GetFromGoodsType 通过goods_type获取内容 商品类型
func (obj *DraftGoodsMgr) GetFromGoodsType(goodsType string) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`goods_type` = ?", goodsType).Find(&results).Error

	return
}

// GetBatchFromGoodsType 批量查找 商品类型
func (obj *DraftGoodsMgr) GetBatchFromGoodsType(goodsTypes []string) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`goods_type` IN (?)", goodsTypes).Find(&results).Error

	return
}

// GetFromExchangeMoney 通过exchange_money获取内容 积分商品需要的金额
func (obj *DraftGoodsMgr) GetFromExchangeMoney(exchangeMoney float64) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`exchange_money` = ?", exchangeMoney).Find(&results).Error

	return
}

// GetBatchFromExchangeMoney 批量查找 积分商品需要的金额
func (obj *DraftGoodsMgr) GetBatchFromExchangeMoney(exchangeMoneys []float64) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`exchange_money` IN (?)", exchangeMoneys).Find(&results).Error

	return
}

// GetFromExchangePoint 通过exchange_point获取内容 积分商品需要的积分
func (obj *DraftGoodsMgr) GetFromExchangePoint(exchangePoint int) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`exchange_point` = ?", exchangePoint).Find(&results).Error

	return
}

// GetBatchFromExchangePoint 批量查找 积分商品需要的积分
func (obj *DraftGoodsMgr) GetBatchFromExchangePoint(exchangePoints []int) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`exchange_point` IN (?)", exchangePoints).Find(&results).Error

	return
}

// GetFromExchangeCategoryID 通过exchange_category_id获取内容 积分商品的分类id
func (obj *DraftGoodsMgr) GetFromExchangeCategoryID(exchangeCategoryID int) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`exchange_category_id` = ?", exchangeCategoryID).Find(&results).Error

	return
}

// GetBatchFromExchangeCategoryID 批量查找 积分商品的分类id
func (obj *DraftGoodsMgr) GetBatchFromExchangeCategoryID(exchangeCategoryIDs []int) (results []*models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`exchange_category_id` IN (?)", exchangeCategoryIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *DraftGoodsMgr) FetchByPrimaryKey(draftGoodsID int) (result models.EsDraftGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoods{}).Where("`draft_goods_id` = ?", draftGoodsID).First(&result).Error

	return
}
