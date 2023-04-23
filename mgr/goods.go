package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsGoodsMgr struct {
	*_BaseMgr
}

// EsGoodsMgr open func
func EsGoodsMgr(db *gorm.DB) *_EsGoodsMgr {
	if db == nil {
		panic(fmt.Errorf("EsGoodsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsGoodsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_goods"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsGoodsMgr) GetTableName() string {
	return "es_goods"
}

// Reset 重置gorm会话
func (obj *_EsGoodsMgr) Reset() *_EsGoodsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsGoodsMgr) Get() (result models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsGoodsMgr) Gets() (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsGoodsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGoodsID goods_id获取 主键
func (obj *_EsGoodsMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithGoodsName goods_name获取 商品名称
func (obj *_EsGoodsMgr) WithGoodsName(goodsName string) Option {
	return optionFunc(func(o *options) { o.query["goods_name"] = goodsName })
}

// WithSn sn获取 商品编号
func (obj *_EsGoodsMgr) WithSn(sn string) Option {
	return optionFunc(func(o *options) { o.query["sn"] = sn })
}

// WithBrandID brand_id获取 品牌id
func (obj *_EsGoodsMgr) WithBrandID(brandID int) Option {
	return optionFunc(func(o *options) { o.query["brand_id"] = brandID })
}

// WithCategoryID category_id获取 分类id
func (obj *_EsGoodsMgr) WithCategoryID(categoryID int) Option {
	return optionFunc(func(o *options) { o.query["category_id"] = categoryID })
}

// WithGoodsType goods_type获取 商品类型normal普通point积分
func (obj *_EsGoodsMgr) WithGoodsType(goodsType string) Option {
	return optionFunc(func(o *options) { o.query["goods_type"] = goodsType })
}

// WithWeight weight获取 重量
func (obj *_EsGoodsMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithMarketEnable market_enable获取 上架状态 1上架  0下架
func (obj *_EsGoodsMgr) WithMarketEnable(marketEnable int) Option {
	return optionFunc(func(o *options) { o.query["market_enable"] = marketEnable })
}

// WithIntro intro获取 详情
func (obj *_EsGoodsMgr) WithIntro(intro string) Option {
	return optionFunc(func(o *options) { o.query["intro"] = intro })
}

// WithPrice price获取 商品价格
func (obj *_EsGoodsMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithCost cost获取 成本价格
func (obj *_EsGoodsMgr) WithCost(cost float64) Option {
	return optionFunc(func(o *options) { o.query["cost"] = cost })
}

// WithMktprice mktprice获取 市场价格
func (obj *_EsGoodsMgr) WithMktprice(mktprice float64) Option {
	return optionFunc(func(o *options) { o.query["mktprice"] = mktprice })
}

// WithHaveSpec have_spec获取 是否有规格0没有 1有
func (obj *_EsGoodsMgr) WithHaveSpec(haveSpec int) Option {
	return optionFunc(func(o *options) { o.query["have_spec"] = haveSpec })
}

// WithCreateTime create_time获取 创建时间
func (obj *_EsGoodsMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithLastModify last_modify获取 最后修改时间
func (obj *_EsGoodsMgr) WithLastModify(lastModify int64) Option {
	return optionFunc(func(o *options) { o.query["last_modify"] = lastModify })
}

// WithViewCount view_count获取 浏览数量
func (obj *_EsGoodsMgr) WithViewCount(viewCount int) Option {
	return optionFunc(func(o *options) { o.query["view_count"] = viewCount })
}

// WithBuyCount buy_count获取 购买数量
func (obj *_EsGoodsMgr) WithBuyCount(buyCount int) Option {
	return optionFunc(func(o *options) { o.query["buy_count"] = buyCount })
}

// WithDisabled disabled获取 是否被删除0 删除 1未删除
func (obj *_EsGoodsMgr) WithDisabled(disabled int) Option {
	return optionFunc(func(o *options) { o.query["disabled"] = disabled })
}

// WithQuantity quantity获取 库存
func (obj *_EsGoodsMgr) WithQuantity(quantity int) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

// WithEnableQuantity enable_quantity获取 可用库存
func (obj *_EsGoodsMgr) WithEnableQuantity(enableQuantity int) Option {
	return optionFunc(func(o *options) { o.query["enable_quantity"] = enableQuantity })
}

// WithPoint point获取 如果是积分商品需要使用的积分
func (obj *_EsGoodsMgr) WithPoint(point int) Option {
	return optionFunc(func(o *options) { o.query["point"] = point })
}

// WithPageTitle page_title获取 seo标题
func (obj *_EsGoodsMgr) WithPageTitle(pageTitle string) Option {
	return optionFunc(func(o *options) { o.query["page_title"] = pageTitle })
}

// WithMetaKeywords meta_keywords获取 seo关键字
func (obj *_EsGoodsMgr) WithMetaKeywords(metaKeywords string) Option {
	return optionFunc(func(o *options) { o.query["meta_keywords"] = metaKeywords })
}

// WithMetaDescription meta_description获取 seo描述
func (obj *_EsGoodsMgr) WithMetaDescription(metaDescription string) Option {
	return optionFunc(func(o *options) { o.query["meta_description"] = metaDescription })
}

// WithGrade grade获取 商品好评率
func (obj *_EsGoodsMgr) WithGrade(grade float64) Option {
	return optionFunc(func(o *options) { o.query["grade"] = grade })
}

// WithThumbnail thumbnail获取 缩略图路径
func (obj *_EsGoodsMgr) WithThumbnail(thumbnail string) Option {
	return optionFunc(func(o *options) { o.query["thumbnail"] = thumbnail })
}

// WithBig big获取 大图路径
func (obj *_EsGoodsMgr) WithBig(big string) Option {
	return optionFunc(func(o *options) { o.query["big"] = big })
}

// WithSmall small获取 小图路径
func (obj *_EsGoodsMgr) WithSmall(small string) Option {
	return optionFunc(func(o *options) { o.query["small"] = small })
}

// WithOriginal original获取 原图路径
func (obj *_EsGoodsMgr) WithOriginal(original string) Option {
	return optionFunc(func(o *options) { o.query["original"] = original })
}

// WithSellerID seller_id获取 卖家id
func (obj *_EsGoodsMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithShopCatID shop_cat_id获取 店铺分类id
func (obj *_EsGoodsMgr) WithShopCatID(shopCatID int) Option {
	return optionFunc(func(o *options) { o.query["shop_cat_id"] = shopCatID })
}

// WithCommentNum comment_num获取 评论数量
func (obj *_EsGoodsMgr) WithCommentNum(commentNum int) Option {
	return optionFunc(func(o *options) { o.query["comment_num"] = commentNum })
}

// WithTemplateID template_id获取 运费模板id
func (obj *_EsGoodsMgr) WithTemplateID(templateID int) Option {
	return optionFunc(func(o *options) { o.query["template_id"] = templateID })
}

// WithGoodsTransfeeCharge goods_transfee_charge获取 谁承担运费0：买家承担，1：卖家承担
func (obj *_EsGoodsMgr) WithGoodsTransfeeCharge(goodsTransfeeCharge int) Option {
	return optionFunc(func(o *options) { o.query["goods_transfee_charge"] = goodsTransfeeCharge })
}

// WithSellerName seller_name获取 卖家名字
func (obj *_EsGoodsMgr) WithSellerName(sellerName string) Option {
	return optionFunc(func(o *options) { o.query["seller_name"] = sellerName })
}

// WithIsAuth is_auth获取 0 需要审核 并且待审核，1 不需要审核 2需要审核 且审核通过 3 需要审核 且审核未通过
func (obj *_EsGoodsMgr) WithIsAuth(isAuth int) Option {
	return optionFunc(func(o *options) { o.query["is_auth"] = isAuth })
}

// WithAuthMessage auth_message获取 审核信息
func (obj *_EsGoodsMgr) WithAuthMessage(authMessage string) Option {
	return optionFunc(func(o *options) { o.query["auth_message"] = authMessage })
}

// WithSelfOperated self_operated获取 是否是自营商品 0 不是 1是
func (obj *_EsGoodsMgr) WithSelfOperated(selfOperated int) Option {
	return optionFunc(func(o *options) { o.query["self_operated"] = selfOperated })
}

// WithUnderMessage under_message获取 下架原因
func (obj *_EsGoodsMgr) WithUnderMessage(underMessage string) Option {
	return optionFunc(func(o *options) { o.query["under_message"] = underMessage })
}

// WithPriority priority获取 优先级:高(3)、中(2)、低(1)
func (obj *_EsGoodsMgr) WithPriority(priority int) Option {
	return optionFunc(func(o *options) { o.query["priority"] = priority })
}

// GetByOption 功能选项模式获取
func (obj *_EsGoodsMgr) GetByOption(opts ...Option) (result models.EsGoods, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsGoodsMgr) GetByOptions(opts ...Option) (results []*models.EsGoods, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsGoodsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsGoods, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where(options.query)
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

// GetFromGoodsID 通过goods_id获取内容 主键
func (obj *_EsGoodsMgr) GetFromGoodsID(goodsID int) (result models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`goods_id` = ?", goodsID).First(&result).Error

	return
}

// GetBatchFromGoodsID 批量查找 主键
func (obj *_EsGoodsMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromGoodsName 通过goods_name获取内容 商品名称
func (obj *_EsGoodsMgr) GetFromGoodsName(goodsName string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}

// GetBatchFromGoodsName 批量查找 商品名称
func (obj *_EsGoodsMgr) GetBatchFromGoodsName(goodsNames []string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`goods_name` IN (?)", goodsNames).Find(&results).Error

	return
}

// GetFromSn 通过sn获取内容 商品编号
func (obj *_EsGoodsMgr) GetFromSn(sn string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`sn` = ?", sn).Find(&results).Error

	return
}

// GetBatchFromSn 批量查找 商品编号
func (obj *_EsGoodsMgr) GetBatchFromSn(sns []string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`sn` IN (?)", sns).Find(&results).Error

	return
}

// GetFromBrandID 通过brand_id获取内容 品牌id
func (obj *_EsGoodsMgr) GetFromBrandID(brandID int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`brand_id` = ?", brandID).Find(&results).Error

	return
}

// GetBatchFromBrandID 批量查找 品牌id
func (obj *_EsGoodsMgr) GetBatchFromBrandID(brandIDs []int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`brand_id` IN (?)", brandIDs).Find(&results).Error

	return
}

// GetFromCategoryID 通过category_id获取内容 分类id
func (obj *_EsGoodsMgr) GetFromCategoryID(categoryID int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`category_id` = ?", categoryID).Find(&results).Error

	return
}

// GetBatchFromCategoryID 批量查找 分类id
func (obj *_EsGoodsMgr) GetBatchFromCategoryID(categoryIDs []int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`category_id` IN (?)", categoryIDs).Find(&results).Error

	return
}

// GetFromGoodsType 通过goods_type获取内容 商品类型normal普通point积分
func (obj *_EsGoodsMgr) GetFromGoodsType(goodsType string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`goods_type` = ?", goodsType).Find(&results).Error

	return
}

// GetBatchFromGoodsType 批量查找 商品类型normal普通point积分
func (obj *_EsGoodsMgr) GetBatchFromGoodsType(goodsTypes []string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`goods_type` IN (?)", goodsTypes).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 重量
func (obj *_EsGoodsMgr) GetFromWeight(weight float64) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`weight` = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量查找 重量
func (obj *_EsGoodsMgr) GetBatchFromWeight(weights []float64) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`weight` IN (?)", weights).Find(&results).Error

	return
}

// GetFromMarketEnable 通过market_enable获取内容 上架状态 1上架  0下架
func (obj *_EsGoodsMgr) GetFromMarketEnable(marketEnable int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`market_enable` = ?", marketEnable).Find(&results).Error

	return
}

// GetBatchFromMarketEnable 批量查找 上架状态 1上架  0下架
func (obj *_EsGoodsMgr) GetBatchFromMarketEnable(marketEnables []int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`market_enable` IN (?)", marketEnables).Find(&results).Error

	return
}

// GetFromIntro 通过intro获取内容 详情
func (obj *_EsGoodsMgr) GetFromIntro(intro string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`intro` = ?", intro).Find(&results).Error

	return
}

// GetBatchFromIntro 批量查找 详情
func (obj *_EsGoodsMgr) GetBatchFromIntro(intros []string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`intro` IN (?)", intros).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 商品价格
func (obj *_EsGoodsMgr) GetFromPrice(price float64) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 商品价格
func (obj *_EsGoodsMgr) GetBatchFromPrice(prices []float64) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromCost 通过cost获取内容 成本价格
func (obj *_EsGoodsMgr) GetFromCost(cost float64) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`cost` = ?", cost).Find(&results).Error

	return
}

// GetBatchFromCost 批量查找 成本价格
func (obj *_EsGoodsMgr) GetBatchFromCost(costs []float64) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`cost` IN (?)", costs).Find(&results).Error

	return
}

// GetFromMktprice 通过mktprice获取内容 市场价格
func (obj *_EsGoodsMgr) GetFromMktprice(mktprice float64) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`mktprice` = ?", mktprice).Find(&results).Error

	return
}

// GetBatchFromMktprice 批量查找 市场价格
func (obj *_EsGoodsMgr) GetBatchFromMktprice(mktprices []float64) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`mktprice` IN (?)", mktprices).Find(&results).Error

	return
}

// GetFromHaveSpec 通过have_spec获取内容 是否有规格0没有 1有
func (obj *_EsGoodsMgr) GetFromHaveSpec(haveSpec int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`have_spec` = ?", haveSpec).Find(&results).Error

	return
}

// GetBatchFromHaveSpec 批量查找 是否有规格0没有 1有
func (obj *_EsGoodsMgr) GetBatchFromHaveSpec(haveSpecs []int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`have_spec` IN (?)", haveSpecs).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_EsGoodsMgr) GetFromCreateTime(createTime int64) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_EsGoodsMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromLastModify 通过last_modify获取内容 最后修改时间
func (obj *_EsGoodsMgr) GetFromLastModify(lastModify int64) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`last_modify` = ?", lastModify).Find(&results).Error

	return
}

// GetBatchFromLastModify 批量查找 最后修改时间
func (obj *_EsGoodsMgr) GetBatchFromLastModify(lastModifys []int64) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`last_modify` IN (?)", lastModifys).Find(&results).Error

	return
}

// GetFromViewCount 通过view_count获取内容 浏览数量
func (obj *_EsGoodsMgr) GetFromViewCount(viewCount int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`view_count` = ?", viewCount).Find(&results).Error

	return
}

// GetBatchFromViewCount 批量查找 浏览数量
func (obj *_EsGoodsMgr) GetBatchFromViewCount(viewCounts []int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`view_count` IN (?)", viewCounts).Find(&results).Error

	return
}

// GetFromBuyCount 通过buy_count获取内容 购买数量
func (obj *_EsGoodsMgr) GetFromBuyCount(buyCount int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`buy_count` = ?", buyCount).Find(&results).Error

	return
}

// GetBatchFromBuyCount 批量查找 购买数量
func (obj *_EsGoodsMgr) GetBatchFromBuyCount(buyCounts []int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`buy_count` IN (?)", buyCounts).Find(&results).Error

	return
}

// GetFromDisabled 通过disabled获取内容 是否被删除0 删除 1未删除
func (obj *_EsGoodsMgr) GetFromDisabled(disabled int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`disabled` = ?", disabled).Find(&results).Error

	return
}

// GetBatchFromDisabled 批量查找 是否被删除0 删除 1未删除
func (obj *_EsGoodsMgr) GetBatchFromDisabled(disableds []int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`disabled` IN (?)", disableds).Find(&results).Error

	return
}

// GetFromQuantity 通过quantity获取内容 库存
func (obj *_EsGoodsMgr) GetFromQuantity(quantity int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`quantity` = ?", quantity).Find(&results).Error

	return
}

// GetBatchFromQuantity 批量查找 库存
func (obj *_EsGoodsMgr) GetBatchFromQuantity(quantitys []int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`quantity` IN (?)", quantitys).Find(&results).Error

	return
}

// GetFromEnableQuantity 通过enable_quantity获取内容 可用库存
func (obj *_EsGoodsMgr) GetFromEnableQuantity(enableQuantity int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`enable_quantity` = ?", enableQuantity).Find(&results).Error

	return
}

// GetBatchFromEnableQuantity 批量查找 可用库存
func (obj *_EsGoodsMgr) GetBatchFromEnableQuantity(enableQuantitys []int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`enable_quantity` IN (?)", enableQuantitys).Find(&results).Error

	return
}

// GetFromPoint 通过point获取内容 如果是积分商品需要使用的积分
func (obj *_EsGoodsMgr) GetFromPoint(point int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`point` = ?", point).Find(&results).Error

	return
}

// GetBatchFromPoint 批量查找 如果是积分商品需要使用的积分
func (obj *_EsGoodsMgr) GetBatchFromPoint(points []int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`point` IN (?)", points).Find(&results).Error

	return
}

// GetFromPageTitle 通过page_title获取内容 seo标题
func (obj *_EsGoodsMgr) GetFromPageTitle(pageTitle string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`page_title` = ?", pageTitle).Find(&results).Error

	return
}

// GetBatchFromPageTitle 批量查找 seo标题
func (obj *_EsGoodsMgr) GetBatchFromPageTitle(pageTitles []string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`page_title` IN (?)", pageTitles).Find(&results).Error

	return
}

// GetFromMetaKeywords 通过meta_keywords获取内容 seo关键字
func (obj *_EsGoodsMgr) GetFromMetaKeywords(metaKeywords string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`meta_keywords` = ?", metaKeywords).Find(&results).Error

	return
}

// GetBatchFromMetaKeywords 批量查找 seo关键字
func (obj *_EsGoodsMgr) GetBatchFromMetaKeywords(metaKeywordss []string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`meta_keywords` IN (?)", metaKeywordss).Find(&results).Error

	return
}

// GetFromMetaDescription 通过meta_description获取内容 seo描述
func (obj *_EsGoodsMgr) GetFromMetaDescription(metaDescription string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`meta_description` = ?", metaDescription).Find(&results).Error

	return
}

// GetBatchFromMetaDescription 批量查找 seo描述
func (obj *_EsGoodsMgr) GetBatchFromMetaDescription(metaDescriptions []string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`meta_description` IN (?)", metaDescriptions).Find(&results).Error

	return
}

// GetFromGrade 通过grade获取内容 商品好评率
func (obj *_EsGoodsMgr) GetFromGrade(grade float64) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`grade` = ?", grade).Find(&results).Error

	return
}

// GetBatchFromGrade 批量查找 商品好评率
func (obj *_EsGoodsMgr) GetBatchFromGrade(grades []float64) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`grade` IN (?)", grades).Find(&results).Error

	return
}

// GetFromThumbnail 通过thumbnail获取内容 缩略图路径
func (obj *_EsGoodsMgr) GetFromThumbnail(thumbnail string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`thumbnail` = ?", thumbnail).Find(&results).Error

	return
}

// GetBatchFromThumbnail 批量查找 缩略图路径
func (obj *_EsGoodsMgr) GetBatchFromThumbnail(thumbnails []string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`thumbnail` IN (?)", thumbnails).Find(&results).Error

	return
}

// GetFromBig 通过big获取内容 大图路径
func (obj *_EsGoodsMgr) GetFromBig(big string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`big` = ?", big).Find(&results).Error

	return
}

// GetBatchFromBig 批量查找 大图路径
func (obj *_EsGoodsMgr) GetBatchFromBig(bigs []string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`big` IN (?)", bigs).Find(&results).Error

	return
}

// GetFromSmall 通过small获取内容 小图路径
func (obj *_EsGoodsMgr) GetFromSmall(small string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`small` = ?", small).Find(&results).Error

	return
}

// GetBatchFromSmall 批量查找 小图路径
func (obj *_EsGoodsMgr) GetBatchFromSmall(smalls []string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`small` IN (?)", smalls).Find(&results).Error

	return
}

// GetFromOriginal 通过original获取内容 原图路径
func (obj *_EsGoodsMgr) GetFromOriginal(original string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`original` = ?", original).Find(&results).Error

	return
}

// GetBatchFromOriginal 批量查找 原图路径
func (obj *_EsGoodsMgr) GetBatchFromOriginal(originals []string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`original` IN (?)", originals).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 卖家id
func (obj *_EsGoodsMgr) GetFromSellerID(sellerID int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 卖家id
func (obj *_EsGoodsMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromShopCatID 通过shop_cat_id获取内容 店铺分类id
func (obj *_EsGoodsMgr) GetFromShopCatID(shopCatID int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`shop_cat_id` = ?", shopCatID).Find(&results).Error

	return
}

// GetBatchFromShopCatID 批量查找 店铺分类id
func (obj *_EsGoodsMgr) GetBatchFromShopCatID(shopCatIDs []int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`shop_cat_id` IN (?)", shopCatIDs).Find(&results).Error

	return
}

// GetFromCommentNum 通过comment_num获取内容 评论数量
func (obj *_EsGoodsMgr) GetFromCommentNum(commentNum int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`comment_num` = ?", commentNum).Find(&results).Error

	return
}

// GetBatchFromCommentNum 批量查找 评论数量
func (obj *_EsGoodsMgr) GetBatchFromCommentNum(commentNums []int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`comment_num` IN (?)", commentNums).Find(&results).Error

	return
}

// GetFromTemplateID 通过template_id获取内容 运费模板id
func (obj *_EsGoodsMgr) GetFromTemplateID(templateID int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`template_id` = ?", templateID).Find(&results).Error

	return
}

// GetBatchFromTemplateID 批量查找 运费模板id
func (obj *_EsGoodsMgr) GetBatchFromTemplateID(templateIDs []int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`template_id` IN (?)", templateIDs).Find(&results).Error

	return
}

// GetFromGoodsTransfeeCharge 通过goods_transfee_charge获取内容 谁承担运费0：买家承担，1：卖家承担
func (obj *_EsGoodsMgr) GetFromGoodsTransfeeCharge(goodsTransfeeCharge int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`goods_transfee_charge` = ?", goodsTransfeeCharge).Find(&results).Error

	return
}

// GetBatchFromGoodsTransfeeCharge 批量查找 谁承担运费0：买家承担，1：卖家承担
func (obj *_EsGoodsMgr) GetBatchFromGoodsTransfeeCharge(goodsTransfeeCharges []int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`goods_transfee_charge` IN (?)", goodsTransfeeCharges).Find(&results).Error

	return
}

// GetFromSellerName 通过seller_name获取内容 卖家名字
func (obj *_EsGoodsMgr) GetFromSellerName(sellerName string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`seller_name` = ?", sellerName).Find(&results).Error

	return
}

// GetBatchFromSellerName 批量查找 卖家名字
func (obj *_EsGoodsMgr) GetBatchFromSellerName(sellerNames []string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`seller_name` IN (?)", sellerNames).Find(&results).Error

	return
}

// GetFromIsAuth 通过is_auth获取内容 0 需要审核 并且待审核，1 不需要审核 2需要审核 且审核通过 3 需要审核 且审核未通过
func (obj *_EsGoodsMgr) GetFromIsAuth(isAuth int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`is_auth` = ?", isAuth).Find(&results).Error

	return
}

// GetBatchFromIsAuth 批量查找 0 需要审核 并且待审核，1 不需要审核 2需要审核 且审核通过 3 需要审核 且审核未通过
func (obj *_EsGoodsMgr) GetBatchFromIsAuth(isAuths []int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`is_auth` IN (?)", isAuths).Find(&results).Error

	return
}

// GetFromAuthMessage 通过auth_message获取内容 审核信息
func (obj *_EsGoodsMgr) GetFromAuthMessage(authMessage string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`auth_message` = ?", authMessage).Find(&results).Error

	return
}

// GetBatchFromAuthMessage 批量查找 审核信息
func (obj *_EsGoodsMgr) GetBatchFromAuthMessage(authMessages []string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`auth_message` IN (?)", authMessages).Find(&results).Error

	return
}

// GetFromSelfOperated 通过self_operated获取内容 是否是自营商品 0 不是 1是
func (obj *_EsGoodsMgr) GetFromSelfOperated(selfOperated int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`self_operated` = ?", selfOperated).Find(&results).Error

	return
}

// GetBatchFromSelfOperated 批量查找 是否是自营商品 0 不是 1是
func (obj *_EsGoodsMgr) GetBatchFromSelfOperated(selfOperateds []int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`self_operated` IN (?)", selfOperateds).Find(&results).Error

	return
}

// GetFromUnderMessage 通过under_message获取内容 下架原因
func (obj *_EsGoodsMgr) GetFromUnderMessage(underMessage string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`under_message` = ?", underMessage).Find(&results).Error

	return
}

// GetBatchFromUnderMessage 批量查找 下架原因
func (obj *_EsGoodsMgr) GetBatchFromUnderMessage(underMessages []string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`under_message` IN (?)", underMessages).Find(&results).Error

	return
}

// GetFromPriority 通过priority获取内容 优先级:高(3)、中(2)、低(1)
func (obj *_EsGoodsMgr) GetFromPriority(priority int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`priority` = ?", priority).Find(&results).Error

	return
}

// GetBatchFromPriority 批量查找 优先级:高(3)、中(2)、低(1)
func (obj *_EsGoodsMgr) GetBatchFromPriority(prioritys []int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`priority` IN (?)", prioritys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsGoodsMgr) FetchByPrimaryKey(goodsID int) (result models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`goods_id` = ?", goodsID).First(&result).Error

	return
}

// FetchIndexByIndGoodsName  获取多个内容
func (obj *_EsGoodsMgr) FetchIndexByIndGoodsName(goodsName string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}

// FetchIndexByIndGoodsSn  获取多个内容
func (obj *_EsGoodsMgr) FetchIndexByIndGoodsSn(sn string) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`sn` = ?", sn).Find(&results).Error

	return
}

// FetchIndexByIndGoodsBrandID  获取多个内容
func (obj *_EsGoodsMgr) FetchIndexByIndGoodsBrandID(brandID int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`brand_id` = ?", brandID).Find(&results).Error

	return
}

// FetchIndexByIndGoodsCategoryID  获取多个内容
func (obj *_EsGoodsMgr) FetchIndexByIndGoodsCategoryID(categoryID int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`category_id` = ?", categoryID).Find(&results).Error

	return
}

// FetchIndexByIndGoodsOther  获取多个内容
func (obj *_EsGoodsMgr) FetchIndexByIndGoodsOther(goodsType string, marketEnable int, disabled int) (results []*models.EsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`goods_type` = ? AND `market_enable` = ? AND `disabled` = ?", goodsType, marketEnable, disabled).Find(&results).Error

	return
}
