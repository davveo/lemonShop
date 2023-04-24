package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type GroupbuyGoodsMgr struct {
	*_BaseMgr
}

// NewGroupbuyGoodsMgr open func
func NewGroupbuyGoodsMgr(db db.Repo) *GroupbuyGoodsMgr {
	if db == nil {
		panic(fmt.Errorf("NewGroupbuyGoodsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &GroupbuyGoodsMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_groupbuy_goods"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *GroupbuyGoodsMgr) GetTableName() string {
	return "es_groupbuy_goods"
}

// Reset 重置gorm会话
func (obj *GroupbuyGoodsMgr) Reset() *GroupbuyGoodsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *GroupbuyGoodsMgr) Get() (result models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *GroupbuyGoodsMgr) Gets() (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *GroupbuyGoodsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGbID gb_id获取 团购商品Id
func (obj *GroupbuyGoodsMgr) WithGbID(gbID int) Option {
	return optionFunc(func(o *options) { o.query["gb_id"] = gbID })
}

// WithSkuID sku_id获取 货品Id
func (obj *GroupbuyGoodsMgr) WithSkuID(skuID int) Option {
	return optionFunc(func(o *options) { o.query["sku_id"] = skuID })
}

// WithActID act_id获取 活动Id
func (obj *GroupbuyGoodsMgr) WithActID(actID int) Option {
	return optionFunc(func(o *options) { o.query["act_id"] = actID })
}

// WithCatID cat_id获取 cat_id
func (obj *GroupbuyGoodsMgr) WithCatID(catID int) Option {
	return optionFunc(func(o *options) { o.query["cat_id"] = catID })
}

// WithAreaID area_id获取 地区Id
func (obj *GroupbuyGoodsMgr) WithAreaID(areaID int) Option {
	return optionFunc(func(o *options) { o.query["area_id"] = areaID })
}

// WithGbName gb_name获取 团购名称
func (obj *GroupbuyGoodsMgr) WithGbName(gbName string) Option {
	return optionFunc(func(o *options) { o.query["gb_name"] = gbName })
}

// WithGbTitle gb_title获取 副标题
func (obj *GroupbuyGoodsMgr) WithGbTitle(gbTitle string) Option {
	return optionFunc(func(o *options) { o.query["gb_title"] = gbTitle })
}

// WithGoodsName goods_name获取 商品名称
func (obj *GroupbuyGoodsMgr) WithGoodsName(goodsName string) Option {
	return optionFunc(func(o *options) { o.query["goods_name"] = goodsName })
}

// WithGoodsID goods_id获取 商品Id
func (obj *GroupbuyGoodsMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithOriginalPrice original_price获取 原始价格
func (obj *GroupbuyGoodsMgr) WithOriginalPrice(originalPrice float64) Option {
	return optionFunc(func(o *options) { o.query["original_price"] = originalPrice })
}

// WithPrice price获取 团购价格
func (obj *GroupbuyGoodsMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithImgURL img_url获取 图片地址
func (obj *GroupbuyGoodsMgr) WithImgURL(imgURL string) Option {
	return optionFunc(func(o *options) { o.query["img_url"] = imgURL })
}

// WithGoodsNum goods_num获取 商品总数
func (obj *GroupbuyGoodsMgr) WithGoodsNum(goodsNum int) Option {
	return optionFunc(func(o *options) { o.query["goods_num"] = goodsNum })
}

// WithVisualNum visual_num获取 虚拟数量
func (obj *GroupbuyGoodsMgr) WithVisualNum(visualNum int) Option {
	return optionFunc(func(o *options) { o.query["visual_num"] = visualNum })
}

// WithLimitNum limit_num获取 限购数量
func (obj *GroupbuyGoodsMgr) WithLimitNum(limitNum int) Option {
	return optionFunc(func(o *options) { o.query["limit_num"] = limitNum })
}

// WithBuyNum buy_num获取 已团购数量
func (obj *GroupbuyGoodsMgr) WithBuyNum(buyNum int) Option {
	return optionFunc(func(o *options) { o.query["buy_num"] = buyNum })
}

// WithViewNum view_num获取 浏览数量
func (obj *GroupbuyGoodsMgr) WithViewNum(viewNum int) Option {
	return optionFunc(func(o *options) { o.query["view_num"] = viewNum })
}

// WithRemark remark获取 介绍
func (obj *GroupbuyGoodsMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithGbStatus gb_status获取 状态
func (obj *GroupbuyGoodsMgr) WithGbStatus(gbStatus int) Option {
	return optionFunc(func(o *options) { o.query["gb_status"] = gbStatus })
}

// WithAddTime add_time获取 添加时间
func (obj *GroupbuyGoodsMgr) WithAddTime(addTime int64) Option {
	return optionFunc(func(o *options) { o.query["add_time"] = addTime })
}

// WithWapThumbnail wap_thumbnail获取 wap缩略图
func (obj *GroupbuyGoodsMgr) WithWapThumbnail(wapThumbnail string) Option {
	return optionFunc(func(o *options) { o.query["wap_thumbnail"] = wapThumbnail })
}

// WithThumbnail thumbnail获取 pc缩略图
func (obj *GroupbuyGoodsMgr) WithThumbnail(thumbnail string) Option {
	return optionFunc(func(o *options) { o.query["thumbnail"] = thumbnail })
}

// WithSellerID seller_id获取 商家ID
func (obj *GroupbuyGoodsMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithSellerName seller_name获取 店铺名称
func (obj *GroupbuyGoodsMgr) WithSellerName(sellerName string) Option {
	return optionFunc(func(o *options) { o.query["seller_name"] = sellerName })
}

// GetByOption 功能选项模式获取
func (obj *GroupbuyGoodsMgr) GetByOption(opts ...Option) (result models.EsGroupbuyGoods, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *GroupbuyGoodsMgr) GetByOptions(opts ...Option) (results []*models.EsGroupbuyGoods, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *GroupbuyGoodsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsGroupbuyGoods, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where(options.query)
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

// GetFromGbID 通过gb_id获取内容 团购商品Id
func (obj *GroupbuyGoodsMgr) GetFromGbID(gbID int) (result models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`gb_id` = ?", gbID).First(&result).Error

	return
}

// GetBatchFromGbID 批量查找 团购商品Id
func (obj *GroupbuyGoodsMgr) GetBatchFromGbID(gbIDs []int) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`gb_id` IN (?)", gbIDs).Find(&results).Error

	return
}

// GetFromSkuID 通过sku_id获取内容 货品Id
func (obj *GroupbuyGoodsMgr) GetFromSkuID(skuID int) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`sku_id` = ?", skuID).Find(&results).Error

	return
}

// GetBatchFromSkuID 批量查找 货品Id
func (obj *GroupbuyGoodsMgr) GetBatchFromSkuID(skuIDs []int) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`sku_id` IN (?)", skuIDs).Find(&results).Error

	return
}

// GetFromActID 通过act_id获取内容 活动Id
func (obj *GroupbuyGoodsMgr) GetFromActID(actID int) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`act_id` = ?", actID).Find(&results).Error

	return
}

// GetBatchFromActID 批量查找 活动Id
func (obj *GroupbuyGoodsMgr) GetBatchFromActID(actIDs []int) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`act_id` IN (?)", actIDs).Find(&results).Error

	return
}

// GetFromCatID 通过cat_id获取内容 cat_id
func (obj *GroupbuyGoodsMgr) GetFromCatID(catID int) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`cat_id` = ?", catID).Find(&results).Error

	return
}

// GetBatchFromCatID 批量查找 cat_id
func (obj *GroupbuyGoodsMgr) GetBatchFromCatID(catIDs []int) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`cat_id` IN (?)", catIDs).Find(&results).Error

	return
}

// GetFromAreaID 通过area_id获取内容 地区Id
func (obj *GroupbuyGoodsMgr) GetFromAreaID(areaID int) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`area_id` = ?", areaID).Find(&results).Error

	return
}

// GetBatchFromAreaID 批量查找 地区Id
func (obj *GroupbuyGoodsMgr) GetBatchFromAreaID(areaIDs []int) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`area_id` IN (?)", areaIDs).Find(&results).Error

	return
}

// GetFromGbName 通过gb_name获取内容 团购名称
func (obj *GroupbuyGoodsMgr) GetFromGbName(gbName string) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`gb_name` = ?", gbName).Find(&results).Error

	return
}

// GetBatchFromGbName 批量查找 团购名称
func (obj *GroupbuyGoodsMgr) GetBatchFromGbName(gbNames []string) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`gb_name` IN (?)", gbNames).Find(&results).Error

	return
}

// GetFromGbTitle 通过gb_title获取内容 副标题
func (obj *GroupbuyGoodsMgr) GetFromGbTitle(gbTitle string) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`gb_title` = ?", gbTitle).Find(&results).Error

	return
}

// GetBatchFromGbTitle 批量查找 副标题
func (obj *GroupbuyGoodsMgr) GetBatchFromGbTitle(gbTitles []string) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`gb_title` IN (?)", gbTitles).Find(&results).Error

	return
}

// GetFromGoodsName 通过goods_name获取内容 商品名称
func (obj *GroupbuyGoodsMgr) GetFromGoodsName(goodsName string) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}

// GetBatchFromGoodsName 批量查找 商品名称
func (obj *GroupbuyGoodsMgr) GetBatchFromGoodsName(goodsNames []string) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`goods_name` IN (?)", goodsNames).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 商品Id
func (obj *GroupbuyGoodsMgr) GetFromGoodsID(goodsID int) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 商品Id
func (obj *GroupbuyGoodsMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromOriginalPrice 通过original_price获取内容 原始价格
func (obj *GroupbuyGoodsMgr) GetFromOriginalPrice(originalPrice float64) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`original_price` = ?", originalPrice).Find(&results).Error

	return
}

// GetBatchFromOriginalPrice 批量查找 原始价格
func (obj *GroupbuyGoodsMgr) GetBatchFromOriginalPrice(originalPrices []float64) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`original_price` IN (?)", originalPrices).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 团购价格
func (obj *GroupbuyGoodsMgr) GetFromPrice(price float64) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 团购价格
func (obj *GroupbuyGoodsMgr) GetBatchFromPrice(prices []float64) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromImgURL 通过img_url获取内容 图片地址
func (obj *GroupbuyGoodsMgr) GetFromImgURL(imgURL string) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`img_url` = ?", imgURL).Find(&results).Error

	return
}

// GetBatchFromImgURL 批量查找 图片地址
func (obj *GroupbuyGoodsMgr) GetBatchFromImgURL(imgURLs []string) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`img_url` IN (?)", imgURLs).Find(&results).Error

	return
}

// GetFromGoodsNum 通过goods_num获取内容 商品总数
func (obj *GroupbuyGoodsMgr) GetFromGoodsNum(goodsNum int) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`goods_num` = ?", goodsNum).Find(&results).Error

	return
}

// GetBatchFromGoodsNum 批量查找 商品总数
func (obj *GroupbuyGoodsMgr) GetBatchFromGoodsNum(goodsNums []int) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`goods_num` IN (?)", goodsNums).Find(&results).Error

	return
}

// GetFromVisualNum 通过visual_num获取内容 虚拟数量
func (obj *GroupbuyGoodsMgr) GetFromVisualNum(visualNum int) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`visual_num` = ?", visualNum).Find(&results).Error

	return
}

// GetBatchFromVisualNum 批量查找 虚拟数量
func (obj *GroupbuyGoodsMgr) GetBatchFromVisualNum(visualNums []int) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`visual_num` IN (?)", visualNums).Find(&results).Error

	return
}

// GetFromLimitNum 通过limit_num获取内容 限购数量
func (obj *GroupbuyGoodsMgr) GetFromLimitNum(limitNum int) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`limit_num` = ?", limitNum).Find(&results).Error

	return
}

// GetBatchFromLimitNum 批量查找 限购数量
func (obj *GroupbuyGoodsMgr) GetBatchFromLimitNum(limitNums []int) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`limit_num` IN (?)", limitNums).Find(&results).Error

	return
}

// GetFromBuyNum 通过buy_num获取内容 已团购数量
func (obj *GroupbuyGoodsMgr) GetFromBuyNum(buyNum int) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`buy_num` = ?", buyNum).Find(&results).Error

	return
}

// GetBatchFromBuyNum 批量查找 已团购数量
func (obj *GroupbuyGoodsMgr) GetBatchFromBuyNum(buyNums []int) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`buy_num` IN (?)", buyNums).Find(&results).Error

	return
}

// GetFromViewNum 通过view_num获取内容 浏览数量
func (obj *GroupbuyGoodsMgr) GetFromViewNum(viewNum int) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`view_num` = ?", viewNum).Find(&results).Error

	return
}

// GetBatchFromViewNum 批量查找 浏览数量
func (obj *GroupbuyGoodsMgr) GetBatchFromViewNum(viewNums []int) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`view_num` IN (?)", viewNums).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 介绍
func (obj *GroupbuyGoodsMgr) GetFromRemark(remark string) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 介绍
func (obj *GroupbuyGoodsMgr) GetBatchFromRemark(remarks []string) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromGbStatus 通过gb_status获取内容 状态
func (obj *GroupbuyGoodsMgr) GetFromGbStatus(gbStatus int) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`gb_status` = ?", gbStatus).Find(&results).Error

	return
}

// GetBatchFromGbStatus 批量查找 状态
func (obj *GroupbuyGoodsMgr) GetBatchFromGbStatus(gbStatuss []int) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`gb_status` IN (?)", gbStatuss).Find(&results).Error

	return
}

// GetFromAddTime 通过add_time获取内容 添加时间
func (obj *GroupbuyGoodsMgr) GetFromAddTime(addTime int64) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`add_time` = ?", addTime).Find(&results).Error

	return
}

// GetBatchFromAddTime 批量查找 添加时间
func (obj *GroupbuyGoodsMgr) GetBatchFromAddTime(addTimes []int64) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`add_time` IN (?)", addTimes).Find(&results).Error

	return
}

// GetFromWapThumbnail 通过wap_thumbnail获取内容 wap缩略图
func (obj *GroupbuyGoodsMgr) GetFromWapThumbnail(wapThumbnail string) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`wap_thumbnail` = ?", wapThumbnail).Find(&results).Error

	return
}

// GetBatchFromWapThumbnail 批量查找 wap缩略图
func (obj *GroupbuyGoodsMgr) GetBatchFromWapThumbnail(wapThumbnails []string) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`wap_thumbnail` IN (?)", wapThumbnails).Find(&results).Error

	return
}

// GetFromThumbnail 通过thumbnail获取内容 pc缩略图
func (obj *GroupbuyGoodsMgr) GetFromThumbnail(thumbnail string) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`thumbnail` = ?", thumbnail).Find(&results).Error

	return
}

// GetBatchFromThumbnail 批量查找 pc缩略图
func (obj *GroupbuyGoodsMgr) GetBatchFromThumbnail(thumbnails []string) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`thumbnail` IN (?)", thumbnails).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 商家ID
func (obj *GroupbuyGoodsMgr) GetFromSellerID(sellerID int) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 商家ID
func (obj *GroupbuyGoodsMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromSellerName 通过seller_name获取内容 店铺名称
func (obj *GroupbuyGoodsMgr) GetFromSellerName(sellerName string) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`seller_name` = ?", sellerName).Find(&results).Error

	return
}

// GetBatchFromSellerName 批量查找 店铺名称
func (obj *GroupbuyGoodsMgr) GetBatchFromSellerName(sellerNames []string) (results []*models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`seller_name` IN (?)", sellerNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *GroupbuyGoodsMgr) FetchByPrimaryKey(gbID int) (result models.EsGroupbuyGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyGoods{}).Where("`gb_id` = ?", gbID).First(&result).Error

	return
}
