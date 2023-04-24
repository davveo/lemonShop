package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type OrderMgr struct {
	*_BaseMgr
}

// NewOrderMgr open func
func NewOrderMgr(db db.Repo) *OrderMgr {
	if db == nil {
		panic(fmt.Errorf("NewOrderMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &OrderMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_order"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *OrderMgr) GetTableName() string {
	return "es_order"
}

// Reset 重置gorm会话
func (obj *OrderMgr) Reset() *OrderMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *OrderMgr) Get() (result models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *OrderMgr) Gets() (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *OrderMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithOrderID order_id获取 主键ID
func (obj *OrderMgr) WithOrderID(orderID int) Option {
	return optionFunc(func(o *options) { o.query["order_id"] = orderID })
}

// WithTradeSn trade_sn获取 交易编号
func (obj *OrderMgr) WithTradeSn(tradeSn string) Option {
	return optionFunc(func(o *options) { o.query["trade_sn"] = tradeSn })
}

// WithSellerID seller_id获取 店铺ID
func (obj *OrderMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithSellerName seller_name获取 店铺名称
func (obj *OrderMgr) WithSellerName(sellerName string) Option {
	return optionFunc(func(o *options) { o.query["seller_name"] = sellerName })
}

// WithMemberID member_id获取 会员ID
func (obj *OrderMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithMemberName member_name获取 买家账号
func (obj *OrderMgr) WithMemberName(memberName string) Option {
	return optionFunc(func(o *options) { o.query["member_name"] = memberName })
}

// WithOrderStatus order_status获取 订单状态
func (obj *OrderMgr) WithOrderStatus(orderStatus string) Option {
	return optionFunc(func(o *options) { o.query["order_status"] = orderStatus })
}

// WithPayStatus pay_status获取 付款状态
func (obj *OrderMgr) WithPayStatus(payStatus string) Option {
	return optionFunc(func(o *options) { o.query["pay_status"] = payStatus })
}

// WithShipStatus ship_status获取 货运状态
func (obj *OrderMgr) WithShipStatus(shipStatus string) Option {
	return optionFunc(func(o *options) { o.query["ship_status"] = shipStatus })
}

// WithShippingID shipping_id获取 配送方式ID
func (obj *OrderMgr) WithShippingID(shippingID int) Option {
	return optionFunc(func(o *options) { o.query["shipping_id"] = shippingID })
}

// WithCommentStatus comment_status获取 评论是否完成
func (obj *OrderMgr) WithCommentStatus(commentStatus string) Option {
	return optionFunc(func(o *options) { o.query["comment_status"] = commentStatus })
}

// WithShippingType shipping_type获取 配送方式
func (obj *OrderMgr) WithShippingType(shippingType string) Option {
	return optionFunc(func(o *options) { o.query["shipping_type"] = shippingType })
}

// WithPaymentMethodID payment_method_id获取 支付方式id
func (obj *OrderMgr) WithPaymentMethodID(paymentMethodID string) Option {
	return optionFunc(func(o *options) { o.query["payment_method_id"] = paymentMethodID })
}

// WithPaymentPluginID payment_plugin_id获取 支付插件id
func (obj *OrderMgr) WithPaymentPluginID(paymentPluginID string) Option {
	return optionFunc(func(o *options) { o.query["payment_plugin_id"] = paymentPluginID })
}

// WithPaymentMethodName payment_method_name获取 支付方式名称
func (obj *OrderMgr) WithPaymentMethodName(paymentMethodName string) Option {
	return optionFunc(func(o *options) { o.query["payment_method_name"] = paymentMethodName })
}

// WithPaymentType payment_type获取 支付方式类型
func (obj *OrderMgr) WithPaymentType(paymentType string) Option {
	return optionFunc(func(o *options) { o.query["payment_type"] = paymentType })
}

// WithPaymentTime payment_time获取 支付时间
func (obj *OrderMgr) WithPaymentTime(paymentTime int64) Option {
	return optionFunc(func(o *options) { o.query["payment_time"] = paymentTime })
}

// WithPayMoney pay_money获取 已支付金额
func (obj *OrderMgr) WithPayMoney(payMoney float64) Option {
	return optionFunc(func(o *options) { o.query["pay_money"] = payMoney })
}

// WithShipName ship_name获取 收货人姓名
func (obj *OrderMgr) WithShipName(shipName string) Option {
	return optionFunc(func(o *options) { o.query["ship_name"] = shipName })
}

// WithShipAddr ship_addr获取 收货地址
func (obj *OrderMgr) WithShipAddr(shipAddr string) Option {
	return optionFunc(func(o *options) { o.query["ship_addr"] = shipAddr })
}

// WithShipZip ship_zip获取 收货人邮编
func (obj *OrderMgr) WithShipZip(shipZip string) Option {
	return optionFunc(func(o *options) { o.query["ship_zip"] = shipZip })
}

// WithShipMobile ship_mobile获取 收货人手机
func (obj *OrderMgr) WithShipMobile(shipMobile string) Option {
	return optionFunc(func(o *options) { o.query["ship_mobile"] = shipMobile })
}

// WithShipTel ship_tel获取 收货人电话
func (obj *OrderMgr) WithShipTel(shipTel string) Option {
	return optionFunc(func(o *options) { o.query["ship_tel"] = shipTel })
}

// WithReceiveTime receive_time获取 收货时间
func (obj *OrderMgr) WithReceiveTime(receiveTime string) Option {
	return optionFunc(func(o *options) { o.query["receive_time"] = receiveTime })
}

// WithShipProvinceID ship_province_id获取 配送地区-省份ID
func (obj *OrderMgr) WithShipProvinceID(shipProvinceID int) Option {
	return optionFunc(func(o *options) { o.query["ship_province_id"] = shipProvinceID })
}

// WithShipCityID ship_city_id获取 配送地区-城市ID
func (obj *OrderMgr) WithShipCityID(shipCityID int) Option {
	return optionFunc(func(o *options) { o.query["ship_city_id"] = shipCityID })
}

// WithShipCountyID ship_county_id获取 配送地区-区(县)ID
func (obj *OrderMgr) WithShipCountyID(shipCountyID int) Option {
	return optionFunc(func(o *options) { o.query["ship_county_id"] = shipCountyID })
}

// WithShipTownID ship_town_id获取 配送街道id
func (obj *OrderMgr) WithShipTownID(shipTownID int) Option {
	return optionFunc(func(o *options) { o.query["ship_town_id"] = shipTownID })
}

// WithShipProvince ship_province获取 配送地区-省份
func (obj *OrderMgr) WithShipProvince(shipProvince string) Option {
	return optionFunc(func(o *options) { o.query["ship_province"] = shipProvince })
}

// WithShipCity ship_city获取 配送地区-城市
func (obj *OrderMgr) WithShipCity(shipCity string) Option {
	return optionFunc(func(o *options) { o.query["ship_city"] = shipCity })
}

// WithShipCounty ship_county获取 配送地区-区(县)
func (obj *OrderMgr) WithShipCounty(shipCounty string) Option {
	return optionFunc(func(o *options) { o.query["ship_county"] = shipCounty })
}

// WithShipTown ship_town获取 配送街道
func (obj *OrderMgr) WithShipTown(shipTown string) Option {
	return optionFunc(func(o *options) { o.query["ship_town"] = shipTown })
}

// WithOrderPrice order_price获取 订单总额
func (obj *OrderMgr) WithOrderPrice(orderPrice float64) Option {
	return optionFunc(func(o *options) { o.query["order_price"] = orderPrice })
}

// WithGoodsPrice goods_price获取 商品总额
func (obj *OrderMgr) WithGoodsPrice(goodsPrice float64) Option {
	return optionFunc(func(o *options) { o.query["goods_price"] = goodsPrice })
}

// WithShippingPrice shipping_price获取 配送费用
func (obj *OrderMgr) WithShippingPrice(shippingPrice float64) Option {
	return optionFunc(func(o *options) { o.query["shipping_price"] = shippingPrice })
}

// WithDiscountPrice discount_price获取 优惠金额
func (obj *OrderMgr) WithDiscountPrice(discountPrice float64) Option {
	return optionFunc(func(o *options) { o.query["discount_price"] = discountPrice })
}

// WithDisabled disabled获取 是否被删除
func (obj *OrderMgr) WithDisabled(disabled int) Option {
	return optionFunc(func(o *options) { o.query["disabled"] = disabled })
}

// WithWeight weight获取 订单商品总重量
func (obj *OrderMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithGoodsNum goods_num获取 商品数量
func (obj *OrderMgr) WithGoodsNum(goodsNum int) Option {
	return optionFunc(func(o *options) { o.query["goods_num"] = goodsNum })
}

// WithRemark remark获取 订单备注
func (obj *OrderMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCancelReason cancel_reason获取 订单取消原因
func (obj *OrderMgr) WithCancelReason(cancelReason string) Option {
	return optionFunc(func(o *options) { o.query["cancel_reason"] = cancelReason })
}

// WithTheSign the_sign获取 签收人
func (obj *OrderMgr) WithTheSign(theSign string) Option {
	return optionFunc(func(o *options) { o.query["the_sign"] = theSign })
}

// WithItemsJSON items_json获取 货物列表json
func (obj *OrderMgr) WithItemsJSON(itemsJSON string) Option {
	return optionFunc(func(o *options) { o.query["items_json"] = itemsJSON })
}

// WithWarehouseID warehouse_id获取 发货仓库ID
func (obj *OrderMgr) WithWarehouseID(warehouseID int) Option {
	return optionFunc(func(o *options) { o.query["warehouse_id"] = warehouseID })
}

// WithNeedPayMoney need_pay_money获取 应付金额
func (obj *OrderMgr) WithNeedPayMoney(needPayMoney float64) Option {
	return optionFunc(func(o *options) { o.query["need_pay_money"] = needPayMoney })
}

// WithShipNo ship_no获取 发货单号
func (obj *OrderMgr) WithShipNo(shipNo string) Option {
	return optionFunc(func(o *options) { o.query["ship_no"] = shipNo })
}

// WithAddressID address_id获取 收货地址ID
func (obj *OrderMgr) WithAddressID(addressID int) Option {
	return optionFunc(func(o *options) { o.query["address_id"] = addressID })
}

// WithAdminRemark admin_remark获取 管理员备注
func (obj *OrderMgr) WithAdminRemark(adminRemark int) Option {
	return optionFunc(func(o *options) { o.query["admin_remark"] = adminRemark })
}

// WithLogiID logi_id获取 物流公司ID
func (obj *OrderMgr) WithLogiID(logiID int) Option {
	return optionFunc(func(o *options) { o.query["logi_id"] = logiID })
}

// WithLogiName logi_name获取 物流公司名称
func (obj *OrderMgr) WithLogiName(logiName string) Option {
	return optionFunc(func(o *options) { o.query["logi_name"] = logiName })
}

// WithCompleteTime complete_time获取 完成时间
func (obj *OrderMgr) WithCompleteTime(completeTime int64) Option {
	return optionFunc(func(o *options) { o.query["complete_time"] = completeTime })
}

// WithCreateTime create_time获取 订单创建时间
func (obj *OrderMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithSigningTime signing_time获取 签收时间
func (obj *OrderMgr) WithSigningTime(signingTime int64) Option {
	return optionFunc(func(o *options) { o.query["signing_time"] = signingTime })
}

// WithShipTime ship_time获取 送货时间
func (obj *OrderMgr) WithShipTime(shipTime int64) Option {
	return optionFunc(func(o *options) { o.query["ship_time"] = shipTime })
}

// WithPayOrderNo pay_order_no获取 支付方式返回的交易号
func (obj *OrderMgr) WithPayOrderNo(payOrderNo string) Option {
	return optionFunc(func(o *options) { o.query["pay_order_no"] = payOrderNo })
}

// WithServiceStatus service_status获取 售后状态
func (obj *OrderMgr) WithServiceStatus(serviceStatus string) Option {
	return optionFunc(func(o *options) { o.query["service_status"] = serviceStatus })
}

// WithBillStatus bill_status获取 结算状态
func (obj *OrderMgr) WithBillStatus(billStatus int) Option {
	return optionFunc(func(o *options) { o.query["bill_status"] = billStatus })
}

// WithBillSn bill_sn获取 结算单号
func (obj *OrderMgr) WithBillSn(billSn string) Option {
	return optionFunc(func(o *options) { o.query["bill_sn"] = billSn })
}

// WithClientType client_type获取 订单来源
func (obj *OrderMgr) WithClientType(clientType string) Option {
	return optionFunc(func(o *options) { o.query["client_type"] = clientType })
}

// WithSn sn获取 订单编号
func (obj *OrderMgr) WithSn(sn string) Option {
	return optionFunc(func(o *options) { o.query["sn"] = sn })
}

// WithNeedReceipt need_receipt获取 是否需要发票
func (obj *OrderMgr) WithNeedReceipt(needReceipt int) Option {
	return optionFunc(func(o *options) { o.query["need_receipt"] = needReceipt })
}

// WithOrderType order_type获取 订单类型
func (obj *OrderMgr) WithOrderType(orderType string) Option {
	return optionFunc(func(o *options) { o.query["order_type"] = orderType })
}

// WithOrderData order_data获取 订单数据
func (obj *OrderMgr) WithOrderData(orderData string) Option {
	return optionFunc(func(o *options) { o.query["order_data"] = orderData })
}

// GetByOption 功能选项模式获取
func (obj *OrderMgr) GetByOption(opts ...Option) (result models.EsOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *OrderMgr) GetByOptions(opts ...Option) (results []*models.EsOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *OrderMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsOrder, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where(options.query)
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

// GetFromOrderID 通过order_id获取内容 主键ID
func (obj *OrderMgr) GetFromOrderID(orderID int) (result models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`order_id` = ?", orderID).First(&result).Error

	return
}

// GetBatchFromOrderID 批量查找 主键ID
func (obj *OrderMgr) GetBatchFromOrderID(orderIDs []int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`order_id` IN (?)", orderIDs).Find(&results).Error

	return
}

// GetFromTradeSn 通过trade_sn获取内容 交易编号
func (obj *OrderMgr) GetFromTradeSn(tradeSn string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`trade_sn` = ?", tradeSn).Find(&results).Error

	return
}

// GetBatchFromTradeSn 批量查找 交易编号
func (obj *OrderMgr) GetBatchFromTradeSn(tradeSns []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`trade_sn` IN (?)", tradeSns).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 店铺ID
func (obj *OrderMgr) GetFromSellerID(sellerID int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 店铺ID
func (obj *OrderMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromSellerName 通过seller_name获取内容 店铺名称
func (obj *OrderMgr) GetFromSellerName(sellerName string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`seller_name` = ?", sellerName).Find(&results).Error

	return
}

// GetBatchFromSellerName 批量查找 店铺名称
func (obj *OrderMgr) GetBatchFromSellerName(sellerNames []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`seller_name` IN (?)", sellerNames).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员ID
func (obj *OrderMgr) GetFromMemberID(memberID int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员ID
func (obj *OrderMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromMemberName 通过member_name获取内容 买家账号
func (obj *OrderMgr) GetFromMemberName(memberName string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`member_name` = ?", memberName).Find(&results).Error

	return
}

// GetBatchFromMemberName 批量查找 买家账号
func (obj *OrderMgr) GetBatchFromMemberName(memberNames []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`member_name` IN (?)", memberNames).Find(&results).Error

	return
}

// GetFromOrderStatus 通过order_status获取内容 订单状态
func (obj *OrderMgr) GetFromOrderStatus(orderStatus string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`order_status` = ?", orderStatus).Find(&results).Error

	return
}

// GetBatchFromOrderStatus 批量查找 订单状态
func (obj *OrderMgr) GetBatchFromOrderStatus(orderStatuss []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`order_status` IN (?)", orderStatuss).Find(&results).Error

	return
}

// GetFromPayStatus 通过pay_status获取内容 付款状态
func (obj *OrderMgr) GetFromPayStatus(payStatus string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`pay_status` = ?", payStatus).Find(&results).Error

	return
}

// GetBatchFromPayStatus 批量查找 付款状态
func (obj *OrderMgr) GetBatchFromPayStatus(payStatuss []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`pay_status` IN (?)", payStatuss).Find(&results).Error

	return
}

// GetFromShipStatus 通过ship_status获取内容 货运状态
func (obj *OrderMgr) GetFromShipStatus(shipStatus string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_status` = ?", shipStatus).Find(&results).Error

	return
}

// GetBatchFromShipStatus 批量查找 货运状态
func (obj *OrderMgr) GetBatchFromShipStatus(shipStatuss []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_status` IN (?)", shipStatuss).Find(&results).Error

	return
}

// GetFromShippingID 通过shipping_id获取内容 配送方式ID
func (obj *OrderMgr) GetFromShippingID(shippingID int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`shipping_id` = ?", shippingID).Find(&results).Error

	return
}

// GetBatchFromShippingID 批量查找 配送方式ID
func (obj *OrderMgr) GetBatchFromShippingID(shippingIDs []int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`shipping_id` IN (?)", shippingIDs).Find(&results).Error

	return
}

// GetFromCommentStatus 通过comment_status获取内容 评论是否完成
func (obj *OrderMgr) GetFromCommentStatus(commentStatus string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`comment_status` = ?", commentStatus).Find(&results).Error

	return
}

// GetBatchFromCommentStatus 批量查找 评论是否完成
func (obj *OrderMgr) GetBatchFromCommentStatus(commentStatuss []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`comment_status` IN (?)", commentStatuss).Find(&results).Error

	return
}

// GetFromShippingType 通过shipping_type获取内容 配送方式
func (obj *OrderMgr) GetFromShippingType(shippingType string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`shipping_type` = ?", shippingType).Find(&results).Error

	return
}

// GetBatchFromShippingType 批量查找 配送方式
func (obj *OrderMgr) GetBatchFromShippingType(shippingTypes []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`shipping_type` IN (?)", shippingTypes).Find(&results).Error

	return
}

// GetFromPaymentMethodID 通过payment_method_id获取内容 支付方式id
func (obj *OrderMgr) GetFromPaymentMethodID(paymentMethodID string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`payment_method_id` = ?", paymentMethodID).Find(&results).Error

	return
}

// GetBatchFromPaymentMethodID 批量查找 支付方式id
func (obj *OrderMgr) GetBatchFromPaymentMethodID(paymentMethodIDs []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`payment_method_id` IN (?)", paymentMethodIDs).Find(&results).Error

	return
}

// GetFromPaymentPluginID 通过payment_plugin_id获取内容 支付插件id
func (obj *OrderMgr) GetFromPaymentPluginID(paymentPluginID string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`payment_plugin_id` = ?", paymentPluginID).Find(&results).Error

	return
}

// GetBatchFromPaymentPluginID 批量查找 支付插件id
func (obj *OrderMgr) GetBatchFromPaymentPluginID(paymentPluginIDs []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`payment_plugin_id` IN (?)", paymentPluginIDs).Find(&results).Error

	return
}

// GetFromPaymentMethodName 通过payment_method_name获取内容 支付方式名称
func (obj *OrderMgr) GetFromPaymentMethodName(paymentMethodName string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`payment_method_name` = ?", paymentMethodName).Find(&results).Error

	return
}

// GetBatchFromPaymentMethodName 批量查找 支付方式名称
func (obj *OrderMgr) GetBatchFromPaymentMethodName(paymentMethodNames []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`payment_method_name` IN (?)", paymentMethodNames).Find(&results).Error

	return
}

// GetFromPaymentType 通过payment_type获取内容 支付方式类型
func (obj *OrderMgr) GetFromPaymentType(paymentType string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`payment_type` = ?", paymentType).Find(&results).Error

	return
}

// GetBatchFromPaymentType 批量查找 支付方式类型
func (obj *OrderMgr) GetBatchFromPaymentType(paymentTypes []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`payment_type` IN (?)", paymentTypes).Find(&results).Error

	return
}

// GetFromPaymentTime 通过payment_time获取内容 支付时间
func (obj *OrderMgr) GetFromPaymentTime(paymentTime int64) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`payment_time` = ?", paymentTime).Find(&results).Error

	return
}

// GetBatchFromPaymentTime 批量查找 支付时间
func (obj *OrderMgr) GetBatchFromPaymentTime(paymentTimes []int64) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`payment_time` IN (?)", paymentTimes).Find(&results).Error

	return
}

// GetFromPayMoney 通过pay_money获取内容 已支付金额
func (obj *OrderMgr) GetFromPayMoney(payMoney float64) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`pay_money` = ?", payMoney).Find(&results).Error

	return
}

// GetBatchFromPayMoney 批量查找 已支付金额
func (obj *OrderMgr) GetBatchFromPayMoney(payMoneys []float64) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`pay_money` IN (?)", payMoneys).Find(&results).Error

	return
}

// GetFromShipName 通过ship_name获取内容 收货人姓名
func (obj *OrderMgr) GetFromShipName(shipName string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_name` = ?", shipName).Find(&results).Error

	return
}

// GetBatchFromShipName 批量查找 收货人姓名
func (obj *OrderMgr) GetBatchFromShipName(shipNames []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_name` IN (?)", shipNames).Find(&results).Error

	return
}

// GetFromShipAddr 通过ship_addr获取内容 收货地址
func (obj *OrderMgr) GetFromShipAddr(shipAddr string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_addr` = ?", shipAddr).Find(&results).Error

	return
}

// GetBatchFromShipAddr 批量查找 收货地址
func (obj *OrderMgr) GetBatchFromShipAddr(shipAddrs []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_addr` IN (?)", shipAddrs).Find(&results).Error

	return
}

// GetFromShipZip 通过ship_zip获取内容 收货人邮编
func (obj *OrderMgr) GetFromShipZip(shipZip string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_zip` = ?", shipZip).Find(&results).Error

	return
}

// GetBatchFromShipZip 批量查找 收货人邮编
func (obj *OrderMgr) GetBatchFromShipZip(shipZips []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_zip` IN (?)", shipZips).Find(&results).Error

	return
}

// GetFromShipMobile 通过ship_mobile获取内容 收货人手机
func (obj *OrderMgr) GetFromShipMobile(shipMobile string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_mobile` = ?", shipMobile).Find(&results).Error

	return
}

// GetBatchFromShipMobile 批量查找 收货人手机
func (obj *OrderMgr) GetBatchFromShipMobile(shipMobiles []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_mobile` IN (?)", shipMobiles).Find(&results).Error

	return
}

// GetFromShipTel 通过ship_tel获取内容 收货人电话
func (obj *OrderMgr) GetFromShipTel(shipTel string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_tel` = ?", shipTel).Find(&results).Error

	return
}

// GetBatchFromShipTel 批量查找 收货人电话
func (obj *OrderMgr) GetBatchFromShipTel(shipTels []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_tel` IN (?)", shipTels).Find(&results).Error

	return
}

// GetFromReceiveTime 通过receive_time获取内容 收货时间
func (obj *OrderMgr) GetFromReceiveTime(receiveTime string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`receive_time` = ?", receiveTime).Find(&results).Error

	return
}

// GetBatchFromReceiveTime 批量查找 收货时间
func (obj *OrderMgr) GetBatchFromReceiveTime(receiveTimes []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`receive_time` IN (?)", receiveTimes).Find(&results).Error

	return
}

// GetFromShipProvinceID 通过ship_province_id获取内容 配送地区-省份ID
func (obj *OrderMgr) GetFromShipProvinceID(shipProvinceID int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_province_id` = ?", shipProvinceID).Find(&results).Error

	return
}

// GetBatchFromShipProvinceID 批量查找 配送地区-省份ID
func (obj *OrderMgr) GetBatchFromShipProvinceID(shipProvinceIDs []int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_province_id` IN (?)", shipProvinceIDs).Find(&results).Error

	return
}

// GetFromShipCityID 通过ship_city_id获取内容 配送地区-城市ID
func (obj *OrderMgr) GetFromShipCityID(shipCityID int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_city_id` = ?", shipCityID).Find(&results).Error

	return
}

// GetBatchFromShipCityID 批量查找 配送地区-城市ID
func (obj *OrderMgr) GetBatchFromShipCityID(shipCityIDs []int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_city_id` IN (?)", shipCityIDs).Find(&results).Error

	return
}

// GetFromShipCountyID 通过ship_county_id获取内容 配送地区-区(县)ID
func (obj *OrderMgr) GetFromShipCountyID(shipCountyID int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_county_id` = ?", shipCountyID).Find(&results).Error

	return
}

// GetBatchFromShipCountyID 批量查找 配送地区-区(县)ID
func (obj *OrderMgr) GetBatchFromShipCountyID(shipCountyIDs []int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_county_id` IN (?)", shipCountyIDs).Find(&results).Error

	return
}

// GetFromShipTownID 通过ship_town_id获取内容 配送街道id
func (obj *OrderMgr) GetFromShipTownID(shipTownID int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_town_id` = ?", shipTownID).Find(&results).Error

	return
}

// GetBatchFromShipTownID 批量查找 配送街道id
func (obj *OrderMgr) GetBatchFromShipTownID(shipTownIDs []int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_town_id` IN (?)", shipTownIDs).Find(&results).Error

	return
}

// GetFromShipProvince 通过ship_province获取内容 配送地区-省份
func (obj *OrderMgr) GetFromShipProvince(shipProvince string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_province` = ?", shipProvince).Find(&results).Error

	return
}

// GetBatchFromShipProvince 批量查找 配送地区-省份
func (obj *OrderMgr) GetBatchFromShipProvince(shipProvinces []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_province` IN (?)", shipProvinces).Find(&results).Error

	return
}

// GetFromShipCity 通过ship_city获取内容 配送地区-城市
func (obj *OrderMgr) GetFromShipCity(shipCity string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_city` = ?", shipCity).Find(&results).Error

	return
}

// GetBatchFromShipCity 批量查找 配送地区-城市
func (obj *OrderMgr) GetBatchFromShipCity(shipCitys []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_city` IN (?)", shipCitys).Find(&results).Error

	return
}

// GetFromShipCounty 通过ship_county获取内容 配送地区-区(县)
func (obj *OrderMgr) GetFromShipCounty(shipCounty string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_county` = ?", shipCounty).Find(&results).Error

	return
}

// GetBatchFromShipCounty 批量查找 配送地区-区(县)
func (obj *OrderMgr) GetBatchFromShipCounty(shipCountys []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_county` IN (?)", shipCountys).Find(&results).Error

	return
}

// GetFromShipTown 通过ship_town获取内容 配送街道
func (obj *OrderMgr) GetFromShipTown(shipTown string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_town` = ?", shipTown).Find(&results).Error

	return
}

// GetBatchFromShipTown 批量查找 配送街道
func (obj *OrderMgr) GetBatchFromShipTown(shipTowns []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_town` IN (?)", shipTowns).Find(&results).Error

	return
}

// GetFromOrderPrice 通过order_price获取内容 订单总额
func (obj *OrderMgr) GetFromOrderPrice(orderPrice float64) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`order_price` = ?", orderPrice).Find(&results).Error

	return
}

// GetBatchFromOrderPrice 批量查找 订单总额
func (obj *OrderMgr) GetBatchFromOrderPrice(orderPrices []float64) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`order_price` IN (?)", orderPrices).Find(&results).Error

	return
}

// GetFromGoodsPrice 通过goods_price获取内容 商品总额
func (obj *OrderMgr) GetFromGoodsPrice(goodsPrice float64) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`goods_price` = ?", goodsPrice).Find(&results).Error

	return
}

// GetBatchFromGoodsPrice 批量查找 商品总额
func (obj *OrderMgr) GetBatchFromGoodsPrice(goodsPrices []float64) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`goods_price` IN (?)", goodsPrices).Find(&results).Error

	return
}

// GetFromShippingPrice 通过shipping_price获取内容 配送费用
func (obj *OrderMgr) GetFromShippingPrice(shippingPrice float64) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`shipping_price` = ?", shippingPrice).Find(&results).Error

	return
}

// GetBatchFromShippingPrice 批量查找 配送费用
func (obj *OrderMgr) GetBatchFromShippingPrice(shippingPrices []float64) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`shipping_price` IN (?)", shippingPrices).Find(&results).Error

	return
}

// GetFromDiscountPrice 通过discount_price获取内容 优惠金额
func (obj *OrderMgr) GetFromDiscountPrice(discountPrice float64) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`discount_price` = ?", discountPrice).Find(&results).Error

	return
}

// GetBatchFromDiscountPrice 批量查找 优惠金额
func (obj *OrderMgr) GetBatchFromDiscountPrice(discountPrices []float64) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`discount_price` IN (?)", discountPrices).Find(&results).Error

	return
}

// GetFromDisabled 通过disabled获取内容 是否被删除
func (obj *OrderMgr) GetFromDisabled(disabled int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`disabled` = ?", disabled).Find(&results).Error

	return
}

// GetBatchFromDisabled 批量查找 是否被删除
func (obj *OrderMgr) GetBatchFromDisabled(disableds []int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`disabled` IN (?)", disableds).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 订单商品总重量
func (obj *OrderMgr) GetFromWeight(weight float64) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`weight` = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量查找 订单商品总重量
func (obj *OrderMgr) GetBatchFromWeight(weights []float64) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`weight` IN (?)", weights).Find(&results).Error

	return
}

// GetFromGoodsNum 通过goods_num获取内容 商品数量
func (obj *OrderMgr) GetFromGoodsNum(goodsNum int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`goods_num` = ?", goodsNum).Find(&results).Error

	return
}

// GetBatchFromGoodsNum 批量查找 商品数量
func (obj *OrderMgr) GetBatchFromGoodsNum(goodsNums []int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`goods_num` IN (?)", goodsNums).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 订单备注
func (obj *OrderMgr) GetFromRemark(remark string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 订单备注
func (obj *OrderMgr) GetBatchFromRemark(remarks []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCancelReason 通过cancel_reason获取内容 订单取消原因
func (obj *OrderMgr) GetFromCancelReason(cancelReason string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`cancel_reason` = ?", cancelReason).Find(&results).Error

	return
}

// GetBatchFromCancelReason 批量查找 订单取消原因
func (obj *OrderMgr) GetBatchFromCancelReason(cancelReasons []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`cancel_reason` IN (?)", cancelReasons).Find(&results).Error

	return
}

// GetFromTheSign 通过the_sign获取内容 签收人
func (obj *OrderMgr) GetFromTheSign(theSign string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`the_sign` = ?", theSign).Find(&results).Error

	return
}

// GetBatchFromTheSign 批量查找 签收人
func (obj *OrderMgr) GetBatchFromTheSign(theSigns []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`the_sign` IN (?)", theSigns).Find(&results).Error

	return
}

// GetFromItemsJSON 通过items_json获取内容 货物列表json
func (obj *OrderMgr) GetFromItemsJSON(itemsJSON string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`items_json` = ?", itemsJSON).Find(&results).Error

	return
}

// GetBatchFromItemsJSON 批量查找 货物列表json
func (obj *OrderMgr) GetBatchFromItemsJSON(itemsJSONs []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`items_json` IN (?)", itemsJSONs).Find(&results).Error

	return
}

// GetFromWarehouseID 通过warehouse_id获取内容 发货仓库ID
func (obj *OrderMgr) GetFromWarehouseID(warehouseID int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`warehouse_id` = ?", warehouseID).Find(&results).Error

	return
}

// GetBatchFromWarehouseID 批量查找 发货仓库ID
func (obj *OrderMgr) GetBatchFromWarehouseID(warehouseIDs []int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`warehouse_id` IN (?)", warehouseIDs).Find(&results).Error

	return
}

// GetFromNeedPayMoney 通过need_pay_money获取内容 应付金额
func (obj *OrderMgr) GetFromNeedPayMoney(needPayMoney float64) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`need_pay_money` = ?", needPayMoney).Find(&results).Error

	return
}

// GetBatchFromNeedPayMoney 批量查找 应付金额
func (obj *OrderMgr) GetBatchFromNeedPayMoney(needPayMoneys []float64) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`need_pay_money` IN (?)", needPayMoneys).Find(&results).Error

	return
}

// GetFromShipNo 通过ship_no获取内容 发货单号
func (obj *OrderMgr) GetFromShipNo(shipNo string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_no` = ?", shipNo).Find(&results).Error

	return
}

// GetBatchFromShipNo 批量查找 发货单号
func (obj *OrderMgr) GetBatchFromShipNo(shipNos []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_no` IN (?)", shipNos).Find(&results).Error

	return
}

// GetFromAddressID 通过address_id获取内容 收货地址ID
func (obj *OrderMgr) GetFromAddressID(addressID int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`address_id` = ?", addressID).Find(&results).Error

	return
}

// GetBatchFromAddressID 批量查找 收货地址ID
func (obj *OrderMgr) GetBatchFromAddressID(addressIDs []int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`address_id` IN (?)", addressIDs).Find(&results).Error

	return
}

// GetFromAdminRemark 通过admin_remark获取内容 管理员备注
func (obj *OrderMgr) GetFromAdminRemark(adminRemark int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`admin_remark` = ?", adminRemark).Find(&results).Error

	return
}

// GetBatchFromAdminRemark 批量查找 管理员备注
func (obj *OrderMgr) GetBatchFromAdminRemark(adminRemarks []int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`admin_remark` IN (?)", adminRemarks).Find(&results).Error

	return
}

// GetFromLogiID 通过logi_id获取内容 物流公司ID
func (obj *OrderMgr) GetFromLogiID(logiID int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`logi_id` = ?", logiID).Find(&results).Error

	return
}

// GetBatchFromLogiID 批量查找 物流公司ID
func (obj *OrderMgr) GetBatchFromLogiID(logiIDs []int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`logi_id` IN (?)", logiIDs).Find(&results).Error

	return
}

// GetFromLogiName 通过logi_name获取内容 物流公司名称
func (obj *OrderMgr) GetFromLogiName(logiName string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`logi_name` = ?", logiName).Find(&results).Error

	return
}

// GetBatchFromLogiName 批量查找 物流公司名称
func (obj *OrderMgr) GetBatchFromLogiName(logiNames []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`logi_name` IN (?)", logiNames).Find(&results).Error

	return
}

// GetFromCompleteTime 通过complete_time获取内容 完成时间
func (obj *OrderMgr) GetFromCompleteTime(completeTime int64) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`complete_time` = ?", completeTime).Find(&results).Error

	return
}

// GetBatchFromCompleteTime 批量查找 完成时间
func (obj *OrderMgr) GetBatchFromCompleteTime(completeTimes []int64) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`complete_time` IN (?)", completeTimes).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 订单创建时间
func (obj *OrderMgr) GetFromCreateTime(createTime int64) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 订单创建时间
func (obj *OrderMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromSigningTime 通过signing_time获取内容 签收时间
func (obj *OrderMgr) GetFromSigningTime(signingTime int64) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`signing_time` = ?", signingTime).Find(&results).Error

	return
}

// GetBatchFromSigningTime 批量查找 签收时间
func (obj *OrderMgr) GetBatchFromSigningTime(signingTimes []int64) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`signing_time` IN (?)", signingTimes).Find(&results).Error

	return
}

// GetFromShipTime 通过ship_time获取内容 送货时间
func (obj *OrderMgr) GetFromShipTime(shipTime int64) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_time` = ?", shipTime).Find(&results).Error

	return
}

// GetBatchFromShipTime 批量查找 送货时间
func (obj *OrderMgr) GetBatchFromShipTime(shipTimes []int64) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`ship_time` IN (?)", shipTimes).Find(&results).Error

	return
}

// GetFromPayOrderNo 通过pay_order_no获取内容 支付方式返回的交易号
func (obj *OrderMgr) GetFromPayOrderNo(payOrderNo string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`pay_order_no` = ?", payOrderNo).Find(&results).Error

	return
}

// GetBatchFromPayOrderNo 批量查找 支付方式返回的交易号
func (obj *OrderMgr) GetBatchFromPayOrderNo(payOrderNos []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`pay_order_no` IN (?)", payOrderNos).Find(&results).Error

	return
}

// GetFromServiceStatus 通过service_status获取内容 售后状态
func (obj *OrderMgr) GetFromServiceStatus(serviceStatus string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`service_status` = ?", serviceStatus).Find(&results).Error

	return
}

// GetBatchFromServiceStatus 批量查找 售后状态
func (obj *OrderMgr) GetBatchFromServiceStatus(serviceStatuss []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`service_status` IN (?)", serviceStatuss).Find(&results).Error

	return
}

// GetFromBillStatus 通过bill_status获取内容 结算状态
func (obj *OrderMgr) GetFromBillStatus(billStatus int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`bill_status` = ?", billStatus).Find(&results).Error

	return
}

// GetBatchFromBillStatus 批量查找 结算状态
func (obj *OrderMgr) GetBatchFromBillStatus(billStatuss []int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`bill_status` IN (?)", billStatuss).Find(&results).Error

	return
}

// GetFromBillSn 通过bill_sn获取内容 结算单号
func (obj *OrderMgr) GetFromBillSn(billSn string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`bill_sn` = ?", billSn).Find(&results).Error

	return
}

// GetBatchFromBillSn 批量查找 结算单号
func (obj *OrderMgr) GetBatchFromBillSn(billSns []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`bill_sn` IN (?)", billSns).Find(&results).Error

	return
}

// GetFromClientType 通过client_type获取内容 订单来源
func (obj *OrderMgr) GetFromClientType(clientType string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`client_type` = ?", clientType).Find(&results).Error

	return
}

// GetBatchFromClientType 批量查找 订单来源
func (obj *OrderMgr) GetBatchFromClientType(clientTypes []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`client_type` IN (?)", clientTypes).Find(&results).Error

	return
}

// GetFromSn 通过sn获取内容 订单编号
func (obj *OrderMgr) GetFromSn(sn string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`sn` = ?", sn).Find(&results).Error

	return
}

// GetBatchFromSn 批量查找 订单编号
func (obj *OrderMgr) GetBatchFromSn(sns []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`sn` IN (?)", sns).Find(&results).Error

	return
}

// GetFromNeedReceipt 通过need_receipt获取内容 是否需要发票
func (obj *OrderMgr) GetFromNeedReceipt(needReceipt int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`need_receipt` = ?", needReceipt).Find(&results).Error

	return
}

// GetBatchFromNeedReceipt 批量查找 是否需要发票
func (obj *OrderMgr) GetBatchFromNeedReceipt(needReceipts []int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`need_receipt` IN (?)", needReceipts).Find(&results).Error

	return
}

// GetFromOrderType 通过order_type获取内容 订单类型
func (obj *OrderMgr) GetFromOrderType(orderType string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`order_type` = ?", orderType).Find(&results).Error

	return
}

// GetBatchFromOrderType 批量查找 订单类型
func (obj *OrderMgr) GetBatchFromOrderType(orderTypes []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`order_type` IN (?)", orderTypes).Find(&results).Error

	return
}

// GetFromOrderData 通过order_data获取内容 订单数据
func (obj *OrderMgr) GetFromOrderData(orderData string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`order_data` = ?", orderData).Find(&results).Error

	return
}

// GetBatchFromOrderData 批量查找 订单数据
func (obj *OrderMgr) GetBatchFromOrderData(orderDatas []string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`order_data` IN (?)", orderDatas).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *OrderMgr) FetchByPrimaryKey(orderID int) (result models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`order_id` = ?", orderID).First(&result).Error

	return
}

// FetchIndexByIndOrderMemberid  获取多个内容
func (obj *OrderMgr) FetchIndexByIndOrderMemberid(memberID int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// FetchIndexByIndOrderState  获取多个内容
func (obj *OrderMgr) FetchIndexByIndOrderState(orderStatus string, payStatus string, shipStatus string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`order_status` = ? AND `pay_status` = ? AND `ship_status` = ?", orderStatus, payStatus, shipStatus).Find(&results).Error

	return
}

// FetchIndexByIndOrderTerm  获取多个内容
func (obj *OrderMgr) FetchIndexByIndOrderTerm(disabled int) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`disabled` = ?", disabled).Find(&results).Error

	return
}

// FetchIndexByIndOrderSn  获取多个内容
func (obj *OrderMgr) FetchIndexByIndOrderSn(sn string) (results []*models.EsOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrder{}).Where("`sn` = ?", sn).Find(&results).Error

	return
}
