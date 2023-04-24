package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type OrderComplainMgr struct {
	*_BaseMgr
}

// NewOrderComplainMgr open func
func NewOrderComplainMgr(db db.Repo) *OrderComplainMgr {
	if db == nil {
		panic(fmt.Errorf("NewOrderComplainMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &OrderComplainMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_order_complain"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *OrderComplainMgr) GetTableName() string {
	return "es_order_complain"
}

// Reset 重置gorm会话
func (obj *OrderComplainMgr) Reset() *OrderComplainMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *OrderComplainMgr) Get() (result models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *OrderComplainMgr) Gets() (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *OrderComplainMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithComplainID complain_id获取 主键
func (obj *OrderComplainMgr) WithComplainID(complainID int) Option {
	return optionFunc(func(o *options) { o.query["complain_id"] = complainID })
}

// WithComplainTopic complain_topic获取 投诉主题
func (obj *OrderComplainMgr) WithComplainTopic(complainTopic string) Option {
	return optionFunc(func(o *options) { o.query["complain_topic"] = complainTopic })
}

// WithContent content获取 投诉内容
func (obj *OrderComplainMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithCreateTime create_time获取 投诉时间
func (obj *OrderComplainMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithImages images获取 投诉凭证图片
func (obj *OrderComplainMgr) WithImages(images string) Option {
	return optionFunc(func(o *options) { o.query["images"] = images })
}

// WithStatus status获取 状态，见ComplainStatusEnum.java
func (obj *OrderComplainMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithAppealContent appeal_content获取 商家申诉内容
func (obj *OrderComplainMgr) WithAppealContent(appealContent string) Option {
	return optionFunc(func(o *options) { o.query["appeal_content"] = appealContent })
}

// WithAppealTime appeal_time获取 商家申诉时间
func (obj *OrderComplainMgr) WithAppealTime(appealTime int64) Option {
	return optionFunc(func(o *options) { o.query["appeal_time"] = appealTime })
}

// WithAppealImages appeal_images获取 商家申诉上传的图片
func (obj *OrderComplainMgr) WithAppealImages(appealImages string) Option {
	return optionFunc(func(o *options) { o.query["appeal_images"] = appealImages })
}

// WithOrderSn order_sn获取 订单号
func (obj *OrderComplainMgr) WithOrderSn(orderSn string) Option {
	return optionFunc(func(o *options) { o.query["order_sn"] = orderSn })
}

// WithOrderTime order_time获取 下单时间
func (obj *OrderComplainMgr) WithOrderTime(orderTime int64) Option {
	return optionFunc(func(o *options) { o.query["order_time"] = orderTime })
}

// WithGoodsName goods_name获取 商品名称
func (obj *OrderComplainMgr) WithGoodsName(goodsName string) Option {
	return optionFunc(func(o *options) { o.query["goods_name"] = goodsName })
}

// WithGoodsID goods_id获取 商品id
func (obj *OrderComplainMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithGoodsPrice goods_price获取 商品价格
func (obj *OrderComplainMgr) WithGoodsPrice(goodsPrice float64) Option {
	return optionFunc(func(o *options) { o.query["goods_price"] = goodsPrice })
}

// WithNum num获取 购买的商品数量
func (obj *OrderComplainMgr) WithNum(num int) Option {
	return optionFunc(func(o *options) { o.query["num"] = num })
}

// WithShippingPrice shipping_price获取 运费
func (obj *OrderComplainMgr) WithShippingPrice(shippingPrice float64) Option {
	return optionFunc(func(o *options) { o.query["shipping_price"] = shippingPrice })
}

// WithOrderPrice order_price获取 订单金额
func (obj *OrderComplainMgr) WithOrderPrice(orderPrice float64) Option {
	return optionFunc(func(o *options) { o.query["order_price"] = orderPrice })
}

// WithShipNo ship_no获取 物流单号
func (obj *OrderComplainMgr) WithShipNo(shipNo string) Option {
	return optionFunc(func(o *options) { o.query["ship_no"] = shipNo })
}

// WithSellerID seller_id获取 商家id
func (obj *OrderComplainMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithSellerName seller_name获取 商家名称
func (obj *OrderComplainMgr) WithSellerName(sellerName string) Option {
	return optionFunc(func(o *options) { o.query["seller_name"] = sellerName })
}

// WithMemberID member_id获取 会员id
func (obj *OrderComplainMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithMemberName member_name获取 会员名称
func (obj *OrderComplainMgr) WithMemberName(memberName string) Option {
	return optionFunc(func(o *options) { o.query["member_name"] = memberName })
}

// WithShipName ship_name获取 收货人
func (obj *OrderComplainMgr) WithShipName(shipName string) Option {
	return optionFunc(func(o *options) { o.query["ship_name"] = shipName })
}

// WithShipAddr ship_addr获取 收货地址
func (obj *OrderComplainMgr) WithShipAddr(shipAddr string) Option {
	return optionFunc(func(o *options) { o.query["ship_addr"] = shipAddr })
}

// WithShipMobile ship_mobile获取 收货人手机
func (obj *OrderComplainMgr) WithShipMobile(shipMobile string) Option {
	return optionFunc(func(o *options) { o.query["ship_mobile"] = shipMobile })
}

// WithArbitrationResult arbitration_result获取 仲裁结果
func (obj *OrderComplainMgr) WithArbitrationResult(arbitrationResult string) Option {
	return optionFunc(func(o *options) { o.query["arbitration_result"] = arbitrationResult })
}

// WithSkuID sku_id获取 sku 主键
func (obj *OrderComplainMgr) WithSkuID(skuID int) Option {
	return optionFunc(func(o *options) { o.query["sku_id"] = skuID })
}

// WithGoodsImage goods_image获取 商品图片
func (obj *OrderComplainMgr) WithGoodsImage(goodsImage string) Option {
	return optionFunc(func(o *options) { o.query["goods_image"] = goodsImage })
}

// GetByOption 功能选项模式获取
func (obj *OrderComplainMgr) GetByOption(opts ...Option) (result models.EsOrderComplain, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *OrderComplainMgr) GetByOptions(opts ...Option) (results []*models.EsOrderComplain, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *OrderComplainMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsOrderComplain, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where(options.query)
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

// GetFromComplainID 通过complain_id获取内容 主键
func (obj *OrderComplainMgr) GetFromComplainID(complainID int) (result models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`complain_id` = ?", complainID).First(&result).Error

	return
}

// GetBatchFromComplainID 批量查找 主键
func (obj *OrderComplainMgr) GetBatchFromComplainID(complainIDs []int) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`complain_id` IN (?)", complainIDs).Find(&results).Error

	return
}

// GetFromComplainTopic 通过complain_topic获取内容 投诉主题
func (obj *OrderComplainMgr) GetFromComplainTopic(complainTopic string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`complain_topic` = ?", complainTopic).Find(&results).Error

	return
}

// GetBatchFromComplainTopic 批量查找 投诉主题
func (obj *OrderComplainMgr) GetBatchFromComplainTopic(complainTopics []string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`complain_topic` IN (?)", complainTopics).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 投诉内容
func (obj *OrderComplainMgr) GetFromContent(content string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 投诉内容
func (obj *OrderComplainMgr) GetBatchFromContent(contents []string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 投诉时间
func (obj *OrderComplainMgr) GetFromCreateTime(createTime int64) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 投诉时间
func (obj *OrderComplainMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromImages 通过images获取内容 投诉凭证图片
func (obj *OrderComplainMgr) GetFromImages(images string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`images` = ?", images).Find(&results).Error

	return
}

// GetBatchFromImages 批量查找 投诉凭证图片
func (obj *OrderComplainMgr) GetBatchFromImages(imagess []string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`images` IN (?)", imagess).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态，见ComplainStatusEnum.java
func (obj *OrderComplainMgr) GetFromStatus(status string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态，见ComplainStatusEnum.java
func (obj *OrderComplainMgr) GetBatchFromStatus(statuss []string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromAppealContent 通过appeal_content获取内容 商家申诉内容
func (obj *OrderComplainMgr) GetFromAppealContent(appealContent string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`appeal_content` = ?", appealContent).Find(&results).Error

	return
}

// GetBatchFromAppealContent 批量查找 商家申诉内容
func (obj *OrderComplainMgr) GetBatchFromAppealContent(appealContents []string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`appeal_content` IN (?)", appealContents).Find(&results).Error

	return
}

// GetFromAppealTime 通过appeal_time获取内容 商家申诉时间
func (obj *OrderComplainMgr) GetFromAppealTime(appealTime int64) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`appeal_time` = ?", appealTime).Find(&results).Error

	return
}

// GetBatchFromAppealTime 批量查找 商家申诉时间
func (obj *OrderComplainMgr) GetBatchFromAppealTime(appealTimes []int64) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`appeal_time` IN (?)", appealTimes).Find(&results).Error

	return
}

// GetFromAppealImages 通过appeal_images获取内容 商家申诉上传的图片
func (obj *OrderComplainMgr) GetFromAppealImages(appealImages string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`appeal_images` = ?", appealImages).Find(&results).Error

	return
}

// GetBatchFromAppealImages 批量查找 商家申诉上传的图片
func (obj *OrderComplainMgr) GetBatchFromAppealImages(appealImagess []string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`appeal_images` IN (?)", appealImagess).Find(&results).Error

	return
}

// GetFromOrderSn 通过order_sn获取内容 订单号
func (obj *OrderComplainMgr) GetFromOrderSn(orderSn string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}

// GetBatchFromOrderSn 批量查找 订单号
func (obj *OrderComplainMgr) GetBatchFromOrderSn(orderSns []string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`order_sn` IN (?)", orderSns).Find(&results).Error

	return
}

// GetFromOrderTime 通过order_time获取内容 下单时间
func (obj *OrderComplainMgr) GetFromOrderTime(orderTime int64) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`order_time` = ?", orderTime).Find(&results).Error

	return
}

// GetBatchFromOrderTime 批量查找 下单时间
func (obj *OrderComplainMgr) GetBatchFromOrderTime(orderTimes []int64) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`order_time` IN (?)", orderTimes).Find(&results).Error

	return
}

// GetFromGoodsName 通过goods_name获取内容 商品名称
func (obj *OrderComplainMgr) GetFromGoodsName(goodsName string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}

// GetBatchFromGoodsName 批量查找 商品名称
func (obj *OrderComplainMgr) GetBatchFromGoodsName(goodsNames []string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`goods_name` IN (?)", goodsNames).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 商品id
func (obj *OrderComplainMgr) GetFromGoodsID(goodsID int) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 商品id
func (obj *OrderComplainMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromGoodsPrice 通过goods_price获取内容 商品价格
func (obj *OrderComplainMgr) GetFromGoodsPrice(goodsPrice float64) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`goods_price` = ?", goodsPrice).Find(&results).Error

	return
}

// GetBatchFromGoodsPrice 批量查找 商品价格
func (obj *OrderComplainMgr) GetBatchFromGoodsPrice(goodsPrices []float64) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`goods_price` IN (?)", goodsPrices).Find(&results).Error

	return
}

// GetFromNum 通过num获取内容 购买的商品数量
func (obj *OrderComplainMgr) GetFromNum(num int) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`num` = ?", num).Find(&results).Error

	return
}

// GetBatchFromNum 批量查找 购买的商品数量
func (obj *OrderComplainMgr) GetBatchFromNum(nums []int) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`num` IN (?)", nums).Find(&results).Error

	return
}

// GetFromShippingPrice 通过shipping_price获取内容 运费
func (obj *OrderComplainMgr) GetFromShippingPrice(shippingPrice float64) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`shipping_price` = ?", shippingPrice).Find(&results).Error

	return
}

// GetBatchFromShippingPrice 批量查找 运费
func (obj *OrderComplainMgr) GetBatchFromShippingPrice(shippingPrices []float64) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`shipping_price` IN (?)", shippingPrices).Find(&results).Error

	return
}

// GetFromOrderPrice 通过order_price获取内容 订单金额
func (obj *OrderComplainMgr) GetFromOrderPrice(orderPrice float64) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`order_price` = ?", orderPrice).Find(&results).Error

	return
}

// GetBatchFromOrderPrice 批量查找 订单金额
func (obj *OrderComplainMgr) GetBatchFromOrderPrice(orderPrices []float64) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`order_price` IN (?)", orderPrices).Find(&results).Error

	return
}

// GetFromShipNo 通过ship_no获取内容 物流单号
func (obj *OrderComplainMgr) GetFromShipNo(shipNo string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`ship_no` = ?", shipNo).Find(&results).Error

	return
}

// GetBatchFromShipNo 批量查找 物流单号
func (obj *OrderComplainMgr) GetBatchFromShipNo(shipNos []string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`ship_no` IN (?)", shipNos).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 商家id
func (obj *OrderComplainMgr) GetFromSellerID(sellerID int) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 商家id
func (obj *OrderComplainMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromSellerName 通过seller_name获取内容 商家名称
func (obj *OrderComplainMgr) GetFromSellerName(sellerName string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`seller_name` = ?", sellerName).Find(&results).Error

	return
}

// GetBatchFromSellerName 批量查找 商家名称
func (obj *OrderComplainMgr) GetBatchFromSellerName(sellerNames []string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`seller_name` IN (?)", sellerNames).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *OrderComplainMgr) GetFromMemberID(memberID int) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *OrderComplainMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromMemberName 通过member_name获取内容 会员名称
func (obj *OrderComplainMgr) GetFromMemberName(memberName string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`member_name` = ?", memberName).Find(&results).Error

	return
}

// GetBatchFromMemberName 批量查找 会员名称
func (obj *OrderComplainMgr) GetBatchFromMemberName(memberNames []string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`member_name` IN (?)", memberNames).Find(&results).Error

	return
}

// GetFromShipName 通过ship_name获取内容 收货人
func (obj *OrderComplainMgr) GetFromShipName(shipName string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`ship_name` = ?", shipName).Find(&results).Error

	return
}

// GetBatchFromShipName 批量查找 收货人
func (obj *OrderComplainMgr) GetBatchFromShipName(shipNames []string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`ship_name` IN (?)", shipNames).Find(&results).Error

	return
}

// GetFromShipAddr 通过ship_addr获取内容 收货地址
func (obj *OrderComplainMgr) GetFromShipAddr(shipAddr string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`ship_addr` = ?", shipAddr).Find(&results).Error

	return
}

// GetBatchFromShipAddr 批量查找 收货地址
func (obj *OrderComplainMgr) GetBatchFromShipAddr(shipAddrs []string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`ship_addr` IN (?)", shipAddrs).Find(&results).Error

	return
}

// GetFromShipMobile 通过ship_mobile获取内容 收货人手机
func (obj *OrderComplainMgr) GetFromShipMobile(shipMobile string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`ship_mobile` = ?", shipMobile).Find(&results).Error

	return
}

// GetBatchFromShipMobile 批量查找 收货人手机
func (obj *OrderComplainMgr) GetBatchFromShipMobile(shipMobiles []string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`ship_mobile` IN (?)", shipMobiles).Find(&results).Error

	return
}

// GetFromArbitrationResult 通过arbitration_result获取内容 仲裁结果
func (obj *OrderComplainMgr) GetFromArbitrationResult(arbitrationResult string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`arbitration_result` = ?", arbitrationResult).Find(&results).Error

	return
}

// GetBatchFromArbitrationResult 批量查找 仲裁结果
func (obj *OrderComplainMgr) GetBatchFromArbitrationResult(arbitrationResults []string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`arbitration_result` IN (?)", arbitrationResults).Find(&results).Error

	return
}

// GetFromSkuID 通过sku_id获取内容 sku 主键
func (obj *OrderComplainMgr) GetFromSkuID(skuID int) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`sku_id` = ?", skuID).Find(&results).Error

	return
}

// GetBatchFromSkuID 批量查找 sku 主键
func (obj *OrderComplainMgr) GetBatchFromSkuID(skuIDs []int) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`sku_id` IN (?)", skuIDs).Find(&results).Error

	return
}

// GetFromGoodsImage 通过goods_image获取内容 商品图片
func (obj *OrderComplainMgr) GetFromGoodsImage(goodsImage string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`goods_image` = ?", goodsImage).Find(&results).Error

	return
}

// GetBatchFromGoodsImage 批量查找 商品图片
func (obj *OrderComplainMgr) GetBatchFromGoodsImage(goodsImages []string) (results []*models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`goods_image` IN (?)", goodsImages).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *OrderComplainMgr) FetchByPrimaryKey(complainID int) (result models.EsOrderComplain, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplain{}).Where("`complain_id` = ?", complainID).First(&result).Error

	return
}
