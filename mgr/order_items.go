package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type OrderItemsMgr struct {
	*_BaseMgr
}

// NewOrderItemsMgr open func
func NewOrderItemsMgr(db db.Repo) *OrderItemsMgr {
	if db == nil {
		panic(fmt.Errorf("NewOrderItemsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &OrderItemsMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_order_items"), wdb: db.GetDbW().Table("es_order_items"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *OrderItemsMgr) GetTableName() string {
	return "es_order_items"
}

// Reset 重置gorm会话
func (obj *OrderItemsMgr) Reset() *OrderItemsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *OrderItemsMgr) Get() (result models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *OrderItemsMgr) Gets() (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *OrderItemsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithItemID item_id获取 主键ID
func (obj *OrderItemsMgr) WithItemID(itemID int) Option {
	return optionFunc(func(o *options) { o.query["item_id"] = itemID })
}

// WithGoodsID goods_id获取 商品ID
func (obj *OrderItemsMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithProductID product_id获取 货品ID
func (obj *OrderItemsMgr) WithProductID(productID int) Option {
	return optionFunc(func(o *options) { o.query["product_id"] = productID })
}

// WithNum num获取 销售量
func (obj *OrderItemsMgr) WithNum(num int) Option {
	return optionFunc(func(o *options) { o.query["num"] = num })
}

// WithShipNum ship_num获取 发货量
func (obj *OrderItemsMgr) WithShipNum(shipNum int) Option {
	return optionFunc(func(o *options) { o.query["ship_num"] = shipNum })
}

// WithTradeSn trade_sn获取 交易编号
func (obj *OrderItemsMgr) WithTradeSn(tradeSn string) Option {
	return optionFunc(func(o *options) { o.query["trade_sn"] = tradeSn })
}

// WithOrderSn order_sn获取 订单编号
func (obj *OrderItemsMgr) WithOrderSn(orderSn string) Option {
	return optionFunc(func(o *options) { o.query["order_sn"] = orderSn })
}

// WithImage image获取 图片
func (obj *OrderItemsMgr) WithImage(image string) Option {
	return optionFunc(func(o *options) { o.query["image"] = image })
}

// WithName name获取 商品名称
func (obj *OrderItemsMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithPrice price获取 销售金额
func (obj *OrderItemsMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithCatID cat_id获取 分类ID
func (obj *OrderItemsMgr) WithCatID(catID int) Option {
	return optionFunc(func(o *options) { o.query["cat_id"] = catID })
}

// WithState state获取 状态
func (obj *OrderItemsMgr) WithState(state int) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithSnapshotID snapshot_id获取 快照id
func (obj *OrderItemsMgr) WithSnapshotID(snapshotID int) Option {
	return optionFunc(func(o *options) { o.query["snapshot_id"] = snapshotID })
}

// WithSpecJSON spec_json获取 规格json
func (obj *OrderItemsMgr) WithSpecJSON(specJSON string) Option {
	return optionFunc(func(o *options) { o.query["spec_json"] = specJSON })
}

// WithPromotionType promotion_type获取 促销类型
func (obj *OrderItemsMgr) WithPromotionType(promotionType string) Option {
	return optionFunc(func(o *options) { o.query["promotion_type"] = promotionType })
}

// WithPromotionID promotion_id获取 促销id
func (obj *OrderItemsMgr) WithPromotionID(promotionID int) Option {
	return optionFunc(func(o *options) { o.query["promotion_id"] = promotionID })
}

// WithRefundPrice refund_price获取 订单项可退款金额
func (obj *OrderItemsMgr) WithRefundPrice(refundPrice float64) Option {
	return optionFunc(func(o *options) { o.query["refund_price"] = refundPrice })
}

// WithCommentStatus comment_status获取 评论状态
func (obj *OrderItemsMgr) WithCommentStatus(commentStatus string) Option {
	return optionFunc(func(o *options) { o.query["comment_status"] = commentStatus })
}

// GetByOption 功能选项模式获取
func (obj *OrderItemsMgr) GetByOption(opts ...Option) (result models.EsOrderItems, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *OrderItemsMgr) GetByOptions(opts ...Option) (results []*models.EsOrderItems, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *OrderItemsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsOrderItems, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where(options.query)
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

// GetFromItemID 通过item_id获取内容 主键ID
func (obj *OrderItemsMgr) GetFromItemID(itemID int) (result models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`item_id` = ?", itemID).First(&result).Error

	return
}

// GetBatchFromItemID 批量查找 主键ID
func (obj *OrderItemsMgr) GetBatchFromItemID(itemIDs []int) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`item_id` IN (?)", itemIDs).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 商品ID
func (obj *OrderItemsMgr) GetFromGoodsID(goodsID int) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 商品ID
func (obj *OrderItemsMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromProductID 通过product_id获取内容 货品ID
func (obj *OrderItemsMgr) GetFromProductID(productID int) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`product_id` = ?", productID).Find(&results).Error

	return
}

// GetBatchFromProductID 批量查找 货品ID
func (obj *OrderItemsMgr) GetBatchFromProductID(productIDs []int) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`product_id` IN (?)", productIDs).Find(&results).Error

	return
}

// GetFromNum 通过num获取内容 销售量
func (obj *OrderItemsMgr) GetFromNum(num int) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`num` = ?", num).Find(&results).Error

	return
}

// GetBatchFromNum 批量查找 销售量
func (obj *OrderItemsMgr) GetBatchFromNum(nums []int) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`num` IN (?)", nums).Find(&results).Error

	return
}

// GetFromShipNum 通过ship_num获取内容 发货量
func (obj *OrderItemsMgr) GetFromShipNum(shipNum int) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`ship_num` = ?", shipNum).Find(&results).Error

	return
}

// GetBatchFromShipNum 批量查找 发货量
func (obj *OrderItemsMgr) GetBatchFromShipNum(shipNums []int) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`ship_num` IN (?)", shipNums).Find(&results).Error

	return
}

// GetFromTradeSn 通过trade_sn获取内容 交易编号
func (obj *OrderItemsMgr) GetFromTradeSn(tradeSn string) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`trade_sn` = ?", tradeSn).Find(&results).Error

	return
}

// GetBatchFromTradeSn 批量查找 交易编号
func (obj *OrderItemsMgr) GetBatchFromTradeSn(tradeSns []string) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`trade_sn` IN (?)", tradeSns).Find(&results).Error

	return
}

// GetFromOrderSn 通过order_sn获取内容 订单编号
func (obj *OrderItemsMgr) GetFromOrderSn(orderSn string) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}

// GetBatchFromOrderSn 批量查找 订单编号
func (obj *OrderItemsMgr) GetBatchFromOrderSn(orderSns []string) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`order_sn` IN (?)", orderSns).Find(&results).Error

	return
}

// GetFromImage 通过image获取内容 图片
func (obj *OrderItemsMgr) GetFromImage(image string) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`image` = ?", image).Find(&results).Error

	return
}

// GetBatchFromImage 批量查找 图片
func (obj *OrderItemsMgr) GetBatchFromImage(images []string) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`image` IN (?)", images).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 商品名称
func (obj *OrderItemsMgr) GetFromName(name string) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 商品名称
func (obj *OrderItemsMgr) GetBatchFromName(names []string) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 销售金额
func (obj *OrderItemsMgr) GetFromPrice(price float64) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 销售金额
func (obj *OrderItemsMgr) GetBatchFromPrice(prices []float64) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromCatID 通过cat_id获取内容 分类ID
func (obj *OrderItemsMgr) GetFromCatID(catID int) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`cat_id` = ?", catID).Find(&results).Error

	return
}

// GetBatchFromCatID 批量查找 分类ID
func (obj *OrderItemsMgr) GetBatchFromCatID(catIDs []int) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`cat_id` IN (?)", catIDs).Find(&results).Error

	return
}

// GetFromState 通过state获取内容 状态
func (obj *OrderItemsMgr) GetFromState(state int) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`state` = ?", state).Find(&results).Error

	return
}

// GetBatchFromState 批量查找 状态
func (obj *OrderItemsMgr) GetBatchFromState(states []int) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`state` IN (?)", states).Find(&results).Error

	return
}

// GetFromSnapshotID 通过snapshot_id获取内容 快照id
func (obj *OrderItemsMgr) GetFromSnapshotID(snapshotID int) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`snapshot_id` = ?", snapshotID).Find(&results).Error

	return
}

// GetBatchFromSnapshotID 批量查找 快照id
func (obj *OrderItemsMgr) GetBatchFromSnapshotID(snapshotIDs []int) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`snapshot_id` IN (?)", snapshotIDs).Find(&results).Error

	return
}

// GetFromSpecJSON 通过spec_json获取内容 规格json
func (obj *OrderItemsMgr) GetFromSpecJSON(specJSON string) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`spec_json` = ?", specJSON).Find(&results).Error

	return
}

// GetBatchFromSpecJSON 批量查找 规格json
func (obj *OrderItemsMgr) GetBatchFromSpecJSON(specJSONs []string) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`spec_json` IN (?)", specJSONs).Find(&results).Error

	return
}

// GetFromPromotionType 通过promotion_type获取内容 促销类型
func (obj *OrderItemsMgr) GetFromPromotionType(promotionType string) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`promotion_type` = ?", promotionType).Find(&results).Error

	return
}

// GetBatchFromPromotionType 批量查找 促销类型
func (obj *OrderItemsMgr) GetBatchFromPromotionType(promotionTypes []string) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`promotion_type` IN (?)", promotionTypes).Find(&results).Error

	return
}

// GetFromPromotionID 通过promotion_id获取内容 促销id
func (obj *OrderItemsMgr) GetFromPromotionID(promotionID int) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`promotion_id` = ?", promotionID).Find(&results).Error

	return
}

// GetBatchFromPromotionID 批量查找 促销id
func (obj *OrderItemsMgr) GetBatchFromPromotionID(promotionIDs []int) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`promotion_id` IN (?)", promotionIDs).Find(&results).Error

	return
}

// GetFromRefundPrice 通过refund_price获取内容 订单项可退款金额
func (obj *OrderItemsMgr) GetFromRefundPrice(refundPrice float64) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`refund_price` = ?", refundPrice).Find(&results).Error

	return
}

// GetBatchFromRefundPrice 批量查找 订单项可退款金额
func (obj *OrderItemsMgr) GetBatchFromRefundPrice(refundPrices []float64) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`refund_price` IN (?)", refundPrices).Find(&results).Error

	return
}

// GetFromCommentStatus 通过comment_status获取内容 评论状态
func (obj *OrderItemsMgr) GetFromCommentStatus(commentStatus string) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`comment_status` = ?", commentStatus).Find(&results).Error

	return
}

// GetBatchFromCommentStatus 批量查找 评论状态
func (obj *OrderItemsMgr) GetBatchFromCommentStatus(commentStatuss []string) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`comment_status` IN (?)", commentStatuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *OrderItemsMgr) FetchByPrimaryKey(itemID int) (result models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`item_id` = ?", itemID).First(&result).Error

	return
}

// FetchIndexByEsOrderItem  获取多个内容
func (obj *OrderItemsMgr) FetchIndexByEsOrderItem(orderSn string) (results []*models.EsOrderItems, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderItems{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}
