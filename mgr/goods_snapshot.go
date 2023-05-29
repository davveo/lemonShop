package mgr

import (
	"context"
	"fmt"

	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type GoodsSnapshotMgr struct {
	*_BaseMgr
}

// NewGoodsSnapshotMgr open func
func NewGoodsSnapshotMgr(db db.Repo) *GoodsSnapshotMgr {
	if db == nil {
		panic(fmt.Errorf("NewGoodsSnapshotMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &GoodsSnapshotMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_goods_snapshot"),
		wdb:       db.GetDbW().Table("es_goods_snapshot"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *GoodsSnapshotMgr) GetTableName() string {
	return "es_goods_snapshot"
}

// Reset 重置gorm会话
func (obj *GoodsSnapshotMgr) Reset() *GoodsSnapshotMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *GoodsSnapshotMgr) Get() (result models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *GoodsSnapshotMgr) Gets() (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *GoodsSnapshotMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSnapshotID snapshot_id获取 主键
func (obj *GoodsSnapshotMgr) WithSnapshotID(snapshotID int) Option {
	return optionFunc(func(o *options) { o.query["snapshot_id"] = snapshotID })
}

// WithGoodsID goods_id获取 商品id
func (obj *GoodsSnapshotMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithName name获取 商品名称
func (obj *GoodsSnapshotMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithSn sn获取 商品编号
func (obj *GoodsSnapshotMgr) WithSn(sn string) Option {
	return optionFunc(func(o *options) { o.query["sn"] = sn })
}

// WithBrandName brand_name获取 品牌名称
func (obj *GoodsSnapshotMgr) WithBrandName(brandName string) Option {
	return optionFunc(func(o *options) { o.query["brand_name"] = brandName })
}

// WithCategoryName category_name获取 分类名称
func (obj *GoodsSnapshotMgr) WithCategoryName(categoryName string) Option {
	return optionFunc(func(o *options) { o.query["category_name"] = categoryName })
}

// WithGoodsType goods_type获取 商品类型
func (obj *GoodsSnapshotMgr) WithGoodsType(goodsType string) Option {
	return optionFunc(func(o *options) { o.query["goods_type"] = goodsType })
}

// WithWeight weight获取 重量
func (obj *GoodsSnapshotMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithIntro intro获取 商品详情
func (obj *GoodsSnapshotMgr) WithIntro(intro string) Option {
	return optionFunc(func(o *options) { o.query["intro"] = intro })
}

// WithPrice price获取 商品价格
func (obj *GoodsSnapshotMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithCost cost获取 商品成本价
func (obj *GoodsSnapshotMgr) WithCost(cost float64) Option {
	return optionFunc(func(o *options) { o.query["cost"] = cost })
}

// WithMktprice mktprice获取 商品市场价
func (obj *GoodsSnapshotMgr) WithMktprice(mktprice float64) Option {
	return optionFunc(func(o *options) { o.query["mktprice"] = mktprice })
}

// WithHaveSpec have_spec获取 商品是否开启规格1 开启 0 未开启
func (obj *GoodsSnapshotMgr) WithHaveSpec(haveSpec int16) Option {
	return optionFunc(func(o *options) { o.query["have_spec"] = haveSpec })
}

// WithParamsJSON params_json获取 参数json
func (obj *GoodsSnapshotMgr) WithParamsJSON(paramsJSON string) Option {
	return optionFunc(func(o *options) { o.query["params_json"] = paramsJSON })
}

// WithImgJSON img_json获取 图片json
func (obj *GoodsSnapshotMgr) WithImgJSON(imgJSON string) Option {
	return optionFunc(func(o *options) { o.query["img_json"] = imgJSON })
}

// WithCreateTime create_time获取 快照时间
func (obj *GoodsSnapshotMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithPoint point获取 商品使用积分
func (obj *GoodsSnapshotMgr) WithPoint(point int) Option {
	return optionFunc(func(o *options) { o.query["point"] = point })
}

// WithSellerID seller_id获取 所属卖家
func (obj *GoodsSnapshotMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithPromotionJSON promotion_json获取 促销json值
func (obj *GoodsSnapshotMgr) WithPromotionJSON(promotionJSON string) Option {
	return optionFunc(func(o *options) { o.query["promotion_json"] = promotionJSON })
}

// WithCouponJSON coupon_json获取 优惠券json值
func (obj *GoodsSnapshotMgr) WithCouponJSON(couponJSON string) Option {
	return optionFunc(func(o *options) { o.query["coupon_json"] = couponJSON })
}

// WithMemberID member_id获取 会员id
func (obj *GoodsSnapshotMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// GetByOption 功能选项模式获取
func (obj *GoodsSnapshotMgr) GetByOption(opts ...Option) (result models.EsGoodsSnapshot, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *GoodsSnapshotMgr) GetByOptions(opts ...Option) (results []*models.EsGoodsSnapshot, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *GoodsSnapshotMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsGoodsSnapshot, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where(options.query)
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

// GetFromSnapshotID 通过snapshot_id获取内容 主键
func (obj *GoodsSnapshotMgr) GetFromSnapshotID(snapshotID int) (result models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`snapshot_id` = ?", snapshotID).First(&result).Error

	return
}

// GetBatchFromSnapshotID 批量查找 主键
func (obj *GoodsSnapshotMgr) GetBatchFromSnapshotID(snapshotIDs []int) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`snapshot_id` IN (?)", snapshotIDs).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 商品id
func (obj *GoodsSnapshotMgr) GetFromGoodsID(goodsID int) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 商品id
func (obj *GoodsSnapshotMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 商品名称
func (obj *GoodsSnapshotMgr) GetFromName(name string) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 商品名称
func (obj *GoodsSnapshotMgr) GetBatchFromName(names []string) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromSn 通过sn获取内容 商品编号
func (obj *GoodsSnapshotMgr) GetFromSn(sn string) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`sn` = ?", sn).Find(&results).Error

	return
}

// GetBatchFromSn 批量查找 商品编号
func (obj *GoodsSnapshotMgr) GetBatchFromSn(sns []string) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`sn` IN (?)", sns).Find(&results).Error

	return
}

// GetFromBrandName 通过brand_name获取内容 品牌名称
func (obj *GoodsSnapshotMgr) GetFromBrandName(brandName string) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`brand_name` = ?", brandName).Find(&results).Error

	return
}

// GetBatchFromBrandName 批量查找 品牌名称
func (obj *GoodsSnapshotMgr) GetBatchFromBrandName(brandNames []string) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`brand_name` IN (?)", brandNames).Find(&results).Error

	return
}

// GetFromCategoryName 通过category_name获取内容 分类名称
func (obj *GoodsSnapshotMgr) GetFromCategoryName(categoryName string) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`category_name` = ?", categoryName).Find(&results).Error

	return
}

// GetBatchFromCategoryName 批量查找 分类名称
func (obj *GoodsSnapshotMgr) GetBatchFromCategoryName(categoryNames []string) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`category_name` IN (?)", categoryNames).Find(&results).Error

	return
}

// GetFromGoodsType 通过goods_type获取内容 商品类型
func (obj *GoodsSnapshotMgr) GetFromGoodsType(goodsType string) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`goods_type` = ?", goodsType).Find(&results).Error

	return
}

// GetBatchFromGoodsType 批量查找 商品类型
func (obj *GoodsSnapshotMgr) GetBatchFromGoodsType(goodsTypes []string) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`goods_type` IN (?)", goodsTypes).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 重量
func (obj *GoodsSnapshotMgr) GetFromWeight(weight float64) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`weight` = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量查找 重量
func (obj *GoodsSnapshotMgr) GetBatchFromWeight(weights []float64) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`weight` IN (?)", weights).Find(&results).Error

	return
}

// GetFromIntro 通过intro获取内容 商品详情
func (obj *GoodsSnapshotMgr) GetFromIntro(intro string) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`intro` = ?", intro).Find(&results).Error

	return
}

// GetBatchFromIntro 批量查找 商品详情
func (obj *GoodsSnapshotMgr) GetBatchFromIntro(intros []string) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`intro` IN (?)", intros).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 商品价格
func (obj *GoodsSnapshotMgr) GetFromPrice(price float64) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 商品价格
func (obj *GoodsSnapshotMgr) GetBatchFromPrice(prices []float64) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromCost 通过cost获取内容 商品成本价
func (obj *GoodsSnapshotMgr) GetFromCost(cost float64) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`cost` = ?", cost).Find(&results).Error

	return
}

// GetBatchFromCost 批量查找 商品成本价
func (obj *GoodsSnapshotMgr) GetBatchFromCost(costs []float64) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`cost` IN (?)", costs).Find(&results).Error

	return
}

// GetFromMktprice 通过mktprice获取内容 商品市场价
func (obj *GoodsSnapshotMgr) GetFromMktprice(mktprice float64) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`mktprice` = ?", mktprice).Find(&results).Error

	return
}

// GetBatchFromMktprice 批量查找 商品市场价
func (obj *GoodsSnapshotMgr) GetBatchFromMktprice(mktprices []float64) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`mktprice` IN (?)", mktprices).Find(&results).Error

	return
}

// GetFromHaveSpec 通过have_spec获取内容 商品是否开启规格1 开启 0 未开启
func (obj *GoodsSnapshotMgr) GetFromHaveSpec(haveSpec int16) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`have_spec` = ?", haveSpec).Find(&results).Error

	return
}

// GetBatchFromHaveSpec 批量查找 商品是否开启规格1 开启 0 未开启
func (obj *GoodsSnapshotMgr) GetBatchFromHaveSpec(haveSpecs []int16) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`have_spec` IN (?)", haveSpecs).Find(&results).Error

	return
}

// GetFromParamsJSON 通过params_json获取内容 参数json
func (obj *GoodsSnapshotMgr) GetFromParamsJSON(paramsJSON string) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`params_json` = ?", paramsJSON).Find(&results).Error

	return
}

// GetBatchFromParamsJSON 批量查找 参数json
func (obj *GoodsSnapshotMgr) GetBatchFromParamsJSON(paramsJSONs []string) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`params_json` IN (?)", paramsJSONs).Find(&results).Error

	return
}

// GetFromImgJSON 通过img_json获取内容 图片json
func (obj *GoodsSnapshotMgr) GetFromImgJSON(imgJSON string) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`img_json` = ?", imgJSON).Find(&results).Error

	return
}

// GetBatchFromImgJSON 批量查找 图片json
func (obj *GoodsSnapshotMgr) GetBatchFromImgJSON(imgJSONs []string) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`img_json` IN (?)", imgJSONs).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 快照时间
func (obj *GoodsSnapshotMgr) GetFromCreateTime(createTime int64) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 快照时间
func (obj *GoodsSnapshotMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromPoint 通过point获取内容 商品使用积分
func (obj *GoodsSnapshotMgr) GetFromPoint(point int) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`point` = ?", point).Find(&results).Error

	return
}

// GetBatchFromPoint 批量查找 商品使用积分
func (obj *GoodsSnapshotMgr) GetBatchFromPoint(points []int) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`point` IN (?)", points).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 所属卖家
func (obj *GoodsSnapshotMgr) GetFromSellerID(sellerID int) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 所属卖家
func (obj *GoodsSnapshotMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromPromotionJSON 通过promotion_json获取内容 促销json值
func (obj *GoodsSnapshotMgr) GetFromPromotionJSON(promotionJSON string) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`promotion_json` = ?", promotionJSON).Find(&results).Error

	return
}

// GetBatchFromPromotionJSON 批量查找 促销json值
func (obj *GoodsSnapshotMgr) GetBatchFromPromotionJSON(promotionJSONs []string) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`promotion_json` IN (?)", promotionJSONs).Find(&results).Error

	return
}

// GetFromCouponJSON 通过coupon_json获取内容 优惠券json值
func (obj *GoodsSnapshotMgr) GetFromCouponJSON(couponJSON string) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`coupon_json` = ?", couponJSON).Find(&results).Error

	return
}

// GetBatchFromCouponJSON 批量查找 优惠券json值
func (obj *GoodsSnapshotMgr) GetBatchFromCouponJSON(couponJSONs []string) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`coupon_json` IN (?)", couponJSONs).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *GoodsSnapshotMgr) GetFromMemberID(memberID int) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *GoodsSnapshotMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *GoodsSnapshotMgr) FetchByPrimaryKey(snapshotID int) (result models.EsGoodsSnapshot, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsSnapshot{}).Where("`snapshot_id` = ?", snapshotID).First(&result).Error

	return
}
