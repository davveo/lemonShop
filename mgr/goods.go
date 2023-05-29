package mgr

import (
	"context"
	"fmt"

	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type GoodsMgr struct {
	*_BaseMgr
}

// EsGoodsMgr open func
func EsGoodsMgr(db db.Repo) *GoodsMgr {
	if db == nil {
		panic(fmt.Errorf("EsGoodsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &GoodsMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_goods"),
		wdb:       db.GetDbW().Table("es_goods"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *GoodsMgr) GetTableName() string {
	return "es_goods"
}

// Reset 重置gorm会话
func (obj *GoodsMgr) Reset() *GoodsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *GoodsMgr) Get() (result models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *GoodsMgr) Gets() (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *GoodsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGoodsID goods_id获取 主键
func (obj *GoodsMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithGoodsName goods_name获取 商品名称
func (obj *GoodsMgr) WithGoodsName(goodsName string) Option {
	return optionFunc(func(o *options) { o.query["goods_name"] = goodsName })
}

// WithSn sn获取 商品编号
func (obj *GoodsMgr) WithSn(sn string) Option {
	return optionFunc(func(o *options) { o.query["sn"] = sn })
}

// WithBrandID brand_id获取 品牌id
func (obj *GoodsMgr) WithBrandID(brandID int) Option {
	return optionFunc(func(o *options) { o.query["brand_id"] = brandID })
}

// WithCategoryID category_id获取 分类id
func (obj *GoodsMgr) WithCategoryID(categoryID int) Option {
	return optionFunc(func(o *options) { o.query["category_id"] = categoryID })
}

// WithGoodsType goods_type获取 商品类型normal普通point积分
func (obj *GoodsMgr) WithGoodsType(goodsType string) Option {
	return optionFunc(func(o *options) { o.query["goods_type"] = goodsType })
}

// WithWeight weight获取 重量
func (obj *GoodsMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithMarketEnable market_enable获取 上架状态 1上架  0下架
func (obj *GoodsMgr) WithMarketEnable(marketEnable int) Option {
	return optionFunc(func(o *options) { o.query["market_enable"] = marketEnable })
}

// WithIntro intro获取 详情
func (obj *GoodsMgr) WithIntro(intro string) Option {
	return optionFunc(func(o *options) { o.query["intro"] = intro })
}

// WithPrice price获取 商品价格
func (obj *GoodsMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithCost cost获取 成本价格
func (obj *GoodsMgr) WithCost(cost float64) Option {
	return optionFunc(func(o *options) { o.query["cost"] = cost })
}

// WithMktprice mktprice获取 市场价格
func (obj *GoodsMgr) WithMktprice(mktprice float64) Option {
	return optionFunc(func(o *options) { o.query["mktprice"] = mktprice })
}

// WithHaveSpec have_spec获取 是否有规格0没有 1有
func (obj *GoodsMgr) WithHaveSpec(haveSpec int) Option {
	return optionFunc(func(o *options) { o.query["have_spec"] = haveSpec })
}

// WithCreateTime create_time获取 创建时间
func (obj *GoodsMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithLastModify last_modify获取 最后修改时间
func (obj *GoodsMgr) WithLastModify(lastModify int64) Option {
	return optionFunc(func(o *options) { o.query["last_modify"] = lastModify })
}

// WithViewCount view_count获取 浏览数量
func (obj *GoodsMgr) WithViewCount(viewCount int) Option {
	return optionFunc(func(o *options) { o.query["view_count"] = viewCount })
}

// WithBuyCount buy_count获取 购买数量
func (obj *GoodsMgr) WithBuyCount(buyCount int) Option {
	return optionFunc(func(o *options) { o.query["buy_count"] = buyCount })
}

// WithDisabled disabled获取 是否被删除0 删除 1未删除
func (obj *GoodsMgr) WithDisabled(disabled int) Option {
	return optionFunc(func(o *options) { o.query["disabled"] = disabled })
}

// WithQuantity quantity获取 库存
func (obj *GoodsMgr) WithQuantity(quantity int) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

// WithEnableQuantity enable_quantity获取 可用库存
func (obj *GoodsMgr) WithEnableQuantity(enableQuantity int) Option {
	return optionFunc(func(o *options) { o.query["enable_quantity"] = enableQuantity })
}

// WithPoint point获取 如果是积分商品需要使用的积分
func (obj *GoodsMgr) WithPoint(point int) Option {
	return optionFunc(func(o *options) { o.query["point"] = point })
}

// WithPageTitle page_title获取 seo标题
func (obj *GoodsMgr) WithPageTitle(pageTitle string) Option {
	return optionFunc(func(o *options) { o.query["page_title"] = pageTitle })
}

// WithMetaKeywords meta_keywords获取 seo关键字
func (obj *GoodsMgr) WithMetaKeywords(metaKeywords string) Option {
	return optionFunc(func(o *options) { o.query["meta_keywords"] = metaKeywords })
}

// WithMetaDescription meta_description获取 seo描述
func (obj *GoodsMgr) WithMetaDescription(metaDescription string) Option {
	return optionFunc(func(o *options) { o.query["meta_description"] = metaDescription })
}

// WithGrade grade获取 商品好评率
func (obj *GoodsMgr) WithGrade(grade float64) Option {
	return optionFunc(func(o *options) { o.query["grade"] = grade })
}

// WithThumbnail thumbnail获取 缩略图路径
func (obj *GoodsMgr) WithThumbnail(thumbnail string) Option {
	return optionFunc(func(o *options) { o.query["thumbnail"] = thumbnail })
}

// WithBig big获取 大图路径
func (obj *GoodsMgr) WithBig(big string) Option {
	return optionFunc(func(o *options) { o.query["big"] = big })
}

// WithSmall small获取 小图路径
func (obj *GoodsMgr) WithSmall(small string) Option {
	return optionFunc(func(o *options) { o.query["small"] = small })
}

// WithOriginal original获取 原图路径
func (obj *GoodsMgr) WithOriginal(original string) Option {
	return optionFunc(func(o *options) { o.query["original"] = original })
}

// WithSellerID seller_id获取 卖家id
func (obj *GoodsMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithShopCatID shop_cat_id获取 店铺分类id
func (obj *GoodsMgr) WithShopCatID(shopCatID int) Option {
	return optionFunc(func(o *options) { o.query["shop_cat_id"] = shopCatID })
}

// WithCommentNum comment_num获取 评论数量
func (obj *GoodsMgr) WithCommentNum(commentNum int) Option {
	return optionFunc(func(o *options) { o.query["comment_num"] = commentNum })
}

// WithTemplateID template_id获取 运费模板id
func (obj *GoodsMgr) WithTemplateID(templateID int) Option {
	return optionFunc(func(o *options) { o.query["template_id"] = templateID })
}

// WithGoodsTransfeeCharge goods_transfee_charge获取 谁承担运费0：买家承担，1：卖家承担
func (obj *GoodsMgr) WithGoodsTransfeeCharge(goodsTransfeeCharge int) Option {
	return optionFunc(func(o *options) { o.query["goods_transfee_charge"] = goodsTransfeeCharge })
}

// WithSellerName seller_name获取 卖家名字
func (obj *GoodsMgr) WithSellerName(sellerName string) Option {
	return optionFunc(func(o *options) { o.query["seller_name"] = sellerName })
}

// WithIsAuth is_auth获取 0 需要审核 并且待审核，1 不需要审核 2需要审核 且审核通过 3 需要审核 且审核未通过
func (obj *GoodsMgr) WithIsAuth(isAuth int) Option {
	return optionFunc(func(o *options) { o.query["is_auth"] = isAuth })
}

// WithAuthMessage auth_message获取 审核信息
func (obj *GoodsMgr) WithAuthMessage(authMessage string) Option {
	return optionFunc(func(o *options) { o.query["auth_message"] = authMessage })
}

// WithSelfOperated self_operated获取 是否是自营商品 0 不是 1是
func (obj *GoodsMgr) WithSelfOperated(selfOperated int) Option {
	return optionFunc(func(o *options) { o.query["self_operated"] = selfOperated })
}

// WithUnderMessage under_message获取 下架原因
func (obj *GoodsMgr) WithUnderMessage(underMessage string) Option {
	return optionFunc(func(o *options) { o.query["under_message"] = underMessage })
}

// WithPriority priority获取 优先级:高(3)、中(2)、低(1)
func (obj *GoodsMgr) WithPriority(priority int) Option {
	return optionFunc(func(o *options) { o.query["priority"] = priority })
}

// GetByOption 功能选项模式获取
func (obj *GoodsMgr) GetByOption(opts ...Option) (result models.EsGoods, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *GoodsMgr) GetByOptions(opts ...Option) (results []*models.EsGoods, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *GoodsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsGoods, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where(options.query)
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
func (obj *GoodsMgr) GetFromGoodsID(goodsID int) (result models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`goods_id` = ?", goodsID).First(&result).Error

	return
}

// GetBatchFromGoodsID 批量查找 主键
func (obj *GoodsMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromGoodsName 通过goods_name获取内容 商品名称
func (obj *GoodsMgr) GetFromGoodsName(goodsName string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}

// GetBatchFromGoodsName 批量查找 商品名称
func (obj *GoodsMgr) GetBatchFromGoodsName(goodsNames []string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`goods_name` IN (?)", goodsNames).Find(&results).Error

	return
}

// GetFromSn 通过sn获取内容 商品编号
func (obj *GoodsMgr) GetFromSn(sn string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`sn` = ?", sn).Find(&results).Error

	return
}

// GetBatchFromSn 批量查找 商品编号
func (obj *GoodsMgr) GetBatchFromSn(sns []string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`sn` IN (?)", sns).Find(&results).Error

	return
}

// GetFromBrandID 通过brand_id获取内容 品牌id
func (obj *GoodsMgr) GetFromBrandID(brandID int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`brand_id` = ?", brandID).Find(&results).Error

	return
}

// GetBatchFromBrandID 批量查找 品牌id
func (obj *GoodsMgr) GetBatchFromBrandID(brandIDs []int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`brand_id` IN (?)", brandIDs).Find(&results).Error

	return
}

// GetFromCategoryID 通过category_id获取内容 分类id
func (obj *GoodsMgr) GetFromCategoryID(categoryID int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`category_id` = ?", categoryID).Find(&results).Error

	return
}

// GetBatchFromCategoryID 批量查找 分类id
func (obj *GoodsMgr) GetBatchFromCategoryID(categoryIDs []int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`category_id` IN (?)", categoryIDs).Find(&results).Error

	return
}

// GetFromGoodsType 通过goods_type获取内容 商品类型normal普通point积分
func (obj *GoodsMgr) GetFromGoodsType(goodsType string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`goods_type` = ?", goodsType).Find(&results).Error

	return
}

// GetBatchFromGoodsType 批量查找 商品类型normal普通point积分
func (obj *GoodsMgr) GetBatchFromGoodsType(goodsTypes []string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`goods_type` IN (?)", goodsTypes).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 重量
func (obj *GoodsMgr) GetFromWeight(weight float64) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`weight` = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量查找 重量
func (obj *GoodsMgr) GetBatchFromWeight(weights []float64) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`weight` IN (?)", weights).Find(&results).Error

	return
}

// GetFromMarketEnable 通过market_enable获取内容 上架状态 1上架  0下架
func (obj *GoodsMgr) GetFromMarketEnable(marketEnable int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`market_enable` = ?", marketEnable).Find(&results).Error

	return
}

// GetBatchFromMarketEnable 批量查找 上架状态 1上架  0下架
func (obj *GoodsMgr) GetBatchFromMarketEnable(marketEnables []int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`market_enable` IN (?)", marketEnables).Find(&results).Error

	return
}

// GetFromIntro 通过intro获取内容 详情
func (obj *GoodsMgr) GetFromIntro(intro string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`intro` = ?", intro).Find(&results).Error

	return
}

// GetBatchFromIntro 批量查找 详情
func (obj *GoodsMgr) GetBatchFromIntro(intros []string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`intro` IN (?)", intros).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 商品价格
func (obj *GoodsMgr) GetFromPrice(price float64) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 商品价格
func (obj *GoodsMgr) GetBatchFromPrice(prices []float64) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromCost 通过cost获取内容 成本价格
func (obj *GoodsMgr) GetFromCost(cost float64) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`cost` = ?", cost).Find(&results).Error

	return
}

// GetBatchFromCost 批量查找 成本价格
func (obj *GoodsMgr) GetBatchFromCost(costs []float64) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`cost` IN (?)", costs).Find(&results).Error

	return
}

// GetFromMktprice 通过mktprice获取内容 市场价格
func (obj *GoodsMgr) GetFromMktprice(mktprice float64) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`mktprice` = ?", mktprice).Find(&results).Error

	return
}

// GetBatchFromMktprice 批量查找 市场价格
func (obj *GoodsMgr) GetBatchFromMktprice(mktprices []float64) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`mktprice` IN (?)", mktprices).Find(&results).Error

	return
}

// GetFromHaveSpec 通过have_spec获取内容 是否有规格0没有 1有
func (obj *GoodsMgr) GetFromHaveSpec(haveSpec int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`have_spec` = ?", haveSpec).Find(&results).Error

	return
}

// GetBatchFromHaveSpec 批量查找 是否有规格0没有 1有
func (obj *GoodsMgr) GetBatchFromHaveSpec(haveSpecs []int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`have_spec` IN (?)", haveSpecs).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *GoodsMgr) GetFromCreateTime(createTime int64) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *GoodsMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromLastModify 通过last_modify获取内容 最后修改时间
func (obj *GoodsMgr) GetFromLastModify(lastModify int64) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`last_modify` = ?", lastModify).Find(&results).Error

	return
}

// GetBatchFromLastModify 批量查找 最后修改时间
func (obj *GoodsMgr) GetBatchFromLastModify(lastModifys []int64) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`last_modify` IN (?)", lastModifys).Find(&results).Error

	return
}

// GetFromViewCount 通过view_count获取内容 浏览数量
func (obj *GoodsMgr) GetFromViewCount(viewCount int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`view_count` = ?", viewCount).Find(&results).Error

	return
}

// GetBatchFromViewCount 批量查找 浏览数量
func (obj *GoodsMgr) GetBatchFromViewCount(viewCounts []int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`view_count` IN (?)", viewCounts).Find(&results).Error

	return
}

// GetFromBuyCount 通过buy_count获取内容 购买数量
func (obj *GoodsMgr) GetFromBuyCount(buyCount int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`buy_count` = ?", buyCount).Find(&results).Error

	return
}

// GetBatchFromBuyCount 批量查找 购买数量
func (obj *GoodsMgr) GetBatchFromBuyCount(buyCounts []int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`buy_count` IN (?)", buyCounts).Find(&results).Error

	return
}

// GetFromDisabled 通过disabled获取内容 是否被删除0 删除 1未删除
func (obj *GoodsMgr) GetFromDisabled(disabled int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`disabled` = ?", disabled).Find(&results).Error

	return
}

// GetBatchFromDisabled 批量查找 是否被删除0 删除 1未删除
func (obj *GoodsMgr) GetBatchFromDisabled(disableds []int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`disabled` IN (?)", disableds).Find(&results).Error

	return
}

// GetFromQuantity 通过quantity获取内容 库存
func (obj *GoodsMgr) GetFromQuantity(quantity int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`quantity` = ?", quantity).Find(&results).Error

	return
}

// GetBatchFromQuantity 批量查找 库存
func (obj *GoodsMgr) GetBatchFromQuantity(quantitys []int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`quantity` IN (?)", quantitys).Find(&results).Error

	return
}

// GetFromEnableQuantity 通过enable_quantity获取内容 可用库存
func (obj *GoodsMgr) GetFromEnableQuantity(enableQuantity int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`enable_quantity` = ?", enableQuantity).Find(&results).Error

	return
}

// GetBatchFromEnableQuantity 批量查找 可用库存
func (obj *GoodsMgr) GetBatchFromEnableQuantity(enableQuantitys []int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`enable_quantity` IN (?)", enableQuantitys).Find(&results).Error

	return
}

// GetFromPoint 通过point获取内容 如果是积分商品需要使用的积分
func (obj *GoodsMgr) GetFromPoint(point int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`point` = ?", point).Find(&results).Error

	return
}

// GetBatchFromPoint 批量查找 如果是积分商品需要使用的积分
func (obj *GoodsMgr) GetBatchFromPoint(points []int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`point` IN (?)", points).Find(&results).Error

	return
}

// GetFromPageTitle 通过page_title获取内容 seo标题
func (obj *GoodsMgr) GetFromPageTitle(pageTitle string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`page_title` = ?", pageTitle).Find(&results).Error

	return
}

// GetBatchFromPageTitle 批量查找 seo标题
func (obj *GoodsMgr) GetBatchFromPageTitle(pageTitles []string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`page_title` IN (?)", pageTitles).Find(&results).Error

	return
}

// GetFromMetaKeywords 通过meta_keywords获取内容 seo关键字
func (obj *GoodsMgr) GetFromMetaKeywords(metaKeywords string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`meta_keywords` = ?", metaKeywords).Find(&results).Error

	return
}

// GetBatchFromMetaKeywords 批量查找 seo关键字
func (obj *GoodsMgr) GetBatchFromMetaKeywords(metaKeywordss []string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`meta_keywords` IN (?)", metaKeywordss).Find(&results).Error

	return
}

// GetFromMetaDescription 通过meta_description获取内容 seo描述
func (obj *GoodsMgr) GetFromMetaDescription(metaDescription string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`meta_description` = ?", metaDescription).Find(&results).Error

	return
}

// GetBatchFromMetaDescription 批量查找 seo描述
func (obj *GoodsMgr) GetBatchFromMetaDescription(metaDescriptions []string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`meta_description` IN (?)", metaDescriptions).Find(&results).Error

	return
}

// GetFromGrade 通过grade获取内容 商品好评率
func (obj *GoodsMgr) GetFromGrade(grade float64) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`grade` = ?", grade).Find(&results).Error

	return
}

// GetBatchFromGrade 批量查找 商品好评率
func (obj *GoodsMgr) GetBatchFromGrade(grades []float64) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`grade` IN (?)", grades).Find(&results).Error

	return
}

// GetFromThumbnail 通过thumbnail获取内容 缩略图路径
func (obj *GoodsMgr) GetFromThumbnail(thumbnail string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`thumbnail` = ?", thumbnail).Find(&results).Error

	return
}

// GetBatchFromThumbnail 批量查找 缩略图路径
func (obj *GoodsMgr) GetBatchFromThumbnail(thumbnails []string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`thumbnail` IN (?)", thumbnails).Find(&results).Error

	return
}

// GetFromBig 通过big获取内容 大图路径
func (obj *GoodsMgr) GetFromBig(big string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`big` = ?", big).Find(&results).Error

	return
}

// GetBatchFromBig 批量查找 大图路径
func (obj *GoodsMgr) GetBatchFromBig(bigs []string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`big` IN (?)", bigs).Find(&results).Error

	return
}

// GetFromSmall 通过small获取内容 小图路径
func (obj *GoodsMgr) GetFromSmall(small string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`small` = ?", small).Find(&results).Error

	return
}

// GetBatchFromSmall 批量查找 小图路径
func (obj *GoodsMgr) GetBatchFromSmall(smalls []string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`small` IN (?)", smalls).Find(&results).Error

	return
}

// GetFromOriginal 通过original获取内容 原图路径
func (obj *GoodsMgr) GetFromOriginal(original string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`original` = ?", original).Find(&results).Error

	return
}

// GetBatchFromOriginal 批量查找 原图路径
func (obj *GoodsMgr) GetBatchFromOriginal(originals []string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`original` IN (?)", originals).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 卖家id
func (obj *GoodsMgr) GetFromSellerID(sellerID int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 卖家id
func (obj *GoodsMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromShopCatID 通过shop_cat_id获取内容 店铺分类id
func (obj *GoodsMgr) GetFromShopCatID(shopCatID int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`shop_cat_id` = ?", shopCatID).Find(&results).Error

	return
}

// GetBatchFromShopCatID 批量查找 店铺分类id
func (obj *GoodsMgr) GetBatchFromShopCatID(shopCatIDs []int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`shop_cat_id` IN (?)", shopCatIDs).Find(&results).Error

	return
}

// GetFromCommentNum 通过comment_num获取内容 评论数量
func (obj *GoodsMgr) GetFromCommentNum(commentNum int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`comment_num` = ?", commentNum).Find(&results).Error

	return
}

// GetBatchFromCommentNum 批量查找 评论数量
func (obj *GoodsMgr) GetBatchFromCommentNum(commentNums []int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`comment_num` IN (?)", commentNums).Find(&results).Error

	return
}

// GetFromTemplateID 通过template_id获取内容 运费模板id
func (obj *GoodsMgr) GetFromTemplateID(templateID int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`template_id` = ?", templateID).Find(&results).Error

	return
}

// GetBatchFromTemplateID 批量查找 运费模板id
func (obj *GoodsMgr) GetBatchFromTemplateID(templateIDs []int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`template_id` IN (?)", templateIDs).Find(&results).Error

	return
}

// GetFromGoodsTransfeeCharge 通过goods_transfee_charge获取内容 谁承担运费0：买家承担，1：卖家承担
func (obj *GoodsMgr) GetFromGoodsTransfeeCharge(goodsTransfeeCharge int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`goods_transfee_charge` = ?", goodsTransfeeCharge).Find(&results).Error

	return
}

// GetBatchFromGoodsTransfeeCharge 批量查找 谁承担运费0：买家承担，1：卖家承担
func (obj *GoodsMgr) GetBatchFromGoodsTransfeeCharge(goodsTransfeeCharges []int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`goods_transfee_charge` IN (?)", goodsTransfeeCharges).Find(&results).Error

	return
}

// GetFromSellerName 通过seller_name获取内容 卖家名字
func (obj *GoodsMgr) GetFromSellerName(sellerName string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`seller_name` = ?", sellerName).Find(&results).Error

	return
}

// GetBatchFromSellerName 批量查找 卖家名字
func (obj *GoodsMgr) GetBatchFromSellerName(sellerNames []string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`seller_name` IN (?)", sellerNames).Find(&results).Error

	return
}

// GetFromIsAuth 通过is_auth获取内容 0 需要审核 并且待审核，1 不需要审核 2需要审核 且审核通过 3 需要审核 且审核未通过
func (obj *GoodsMgr) GetFromIsAuth(isAuth int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`is_auth` = ?", isAuth).Find(&results).Error

	return
}

// GetBatchFromIsAuth 批量查找 0 需要审核 并且待审核，1 不需要审核 2需要审核 且审核通过 3 需要审核 且审核未通过
func (obj *GoodsMgr) GetBatchFromIsAuth(isAuths []int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`is_auth` IN (?)", isAuths).Find(&results).Error

	return
}

// GetFromAuthMessage 通过auth_message获取内容 审核信息
func (obj *GoodsMgr) GetFromAuthMessage(authMessage string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`auth_message` = ?", authMessage).Find(&results).Error

	return
}

// GetBatchFromAuthMessage 批量查找 审核信息
func (obj *GoodsMgr) GetBatchFromAuthMessage(authMessages []string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`auth_message` IN (?)", authMessages).Find(&results).Error

	return
}

// GetFromSelfOperated 通过self_operated获取内容 是否是自营商品 0 不是 1是
func (obj *GoodsMgr) GetFromSelfOperated(selfOperated int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`self_operated` = ?", selfOperated).Find(&results).Error

	return
}

// GetBatchFromSelfOperated 批量查找 是否是自营商品 0 不是 1是
func (obj *GoodsMgr) GetBatchFromSelfOperated(selfOperateds []int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`self_operated` IN (?)", selfOperateds).Find(&results).Error

	return
}

// GetFromUnderMessage 通过under_message获取内容 下架原因
func (obj *GoodsMgr) GetFromUnderMessage(underMessage string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`under_message` = ?", underMessage).Find(&results).Error

	return
}

// GetBatchFromUnderMessage 批量查找 下架原因
func (obj *GoodsMgr) GetBatchFromUnderMessage(underMessages []string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`under_message` IN (?)", underMessages).Find(&results).Error

	return
}

// GetFromPriority 通过priority获取内容 优先级:高(3)、中(2)、低(1)
func (obj *GoodsMgr) GetFromPriority(priority int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`priority` = ?", priority).Find(&results).Error

	return
}

// GetBatchFromPriority 批量查找 优先级:高(3)、中(2)、低(1)
func (obj *GoodsMgr) GetBatchFromPriority(prioritys []int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`priority` IN (?)", prioritys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *GoodsMgr) FetchByPrimaryKey(goodsID int) (result models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`goods_id` = ?", goodsID).First(&result).Error

	return
}

// FetchIndexByIndGoodsName  获取多个内容
func (obj *GoodsMgr) FetchIndexByIndGoodsName(goodsName string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}

// FetchIndexByIndGoodsSn  获取多个内容
func (obj *GoodsMgr) FetchIndexByIndGoodsSn(sn string) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`sn` = ?", sn).Find(&results).Error

	return
}

// FetchIndexByIndGoodsBrandID  获取多个内容
func (obj *GoodsMgr) FetchIndexByIndGoodsBrandID(brandID int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`brand_id` = ?", brandID).Find(&results).Error

	return
}

// FetchIndexByIndGoodsCategoryID  获取多个内容
func (obj *GoodsMgr) FetchIndexByIndGoodsCategoryID(categoryID int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`category_id` = ?", categoryID).Find(&results).Error

	return
}

// FetchIndexByIndGoodsOther  获取多个内容
func (obj *GoodsMgr) FetchIndexByIndGoodsOther(goodsType string, marketEnable int, disabled int) (results []*models.EsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoods{}).Where("`goods_type` = ? AND `market_enable` = ? AND `disabled` = ?", goodsType, marketEnable, disabled).Find(&results).Error

	return
}
