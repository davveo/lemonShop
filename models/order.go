package models

// EsOrder 订单表(es_order)
type EsOrder struct {
	OrderID           int     `gorm:"primaryKey;column:order_id" json:"-"`                   // 主键ID
	TradeSn           string  `gorm:"column:trade_sn" json:"trade_sn"`                       // 交易编号
	SellerID          int     `gorm:"column:seller_id" json:"seller_id"`                     // 店铺ID
	SellerName        string  `gorm:"column:seller_name" json:"seller_name"`                 // 店铺名称
	MemberID          int     `gorm:"column:member_id" json:"member_id"`                     // 会员ID
	MemberName        string  `gorm:"column:member_name" json:"member_name"`                 // 买家账号
	OrderStatus       string  `gorm:"column:order_status" json:"order_status"`               // 订单状态
	PayStatus         string  `gorm:"column:pay_status" json:"pay_status"`                   // 付款状态
	ShipStatus        string  `gorm:"column:ship_status" json:"ship_status"`                 // 货运状态
	ShippingID        int     `gorm:"column:shipping_id" json:"shipping_id"`                 // 配送方式ID
	CommentStatus     string  `gorm:"column:comment_status" json:"comment_status"`           // 评论是否完成
	ShippingType      string  `gorm:"column:shipping_type" json:"shipping_type"`             // 配送方式
	PaymentMethodID   string  `gorm:"column:payment_method_id" json:"payment_method_id"`     // 支付方式id
	PaymentPluginID   string  `gorm:"column:payment_plugin_id" json:"payment_plugin_id"`     // 支付插件id
	PaymentMethodName string  `gorm:"column:payment_method_name" json:"payment_method_name"` // 支付方式名称
	PaymentType       string  `gorm:"column:payment_type" json:"payment_type"`               // 支付方式类型
	PaymentTime       int64   `gorm:"column:payment_time" json:"payment_time"`               // 支付时间
	PayMoney          float64 `gorm:"column:pay_money" json:"pay_money"`                     // 已支付金额
	ShipName          string  `gorm:"column:ship_name" json:"ship_name"`                     // 收货人姓名
	ShipAddr          string  `gorm:"column:ship_addr" json:"ship_addr"`                     // 收货地址
	ShipZip           string  `gorm:"column:ship_zip" json:"ship_zip"`                       // 收货人邮编
	ShipMobile        string  `gorm:"column:ship_mobile" json:"ship_mobile"`                 // 收货人手机
	ShipTel           string  `gorm:"column:ship_tel" json:"ship_tel"`                       // 收货人电话
	ReceiveTime       string  `gorm:"column:receive_time" json:"receive_time"`               // 收货时间
	ShipProvinceID    int     `gorm:"column:ship_province_id" json:"ship_province_id"`       // 配送地区-省份ID
	ShipCityID        int     `gorm:"column:ship_city_id" json:"ship_city_id"`               // 配送地区-城市ID
	ShipCountyID      int     `gorm:"column:ship_county_id" json:"ship_county_id"`           // 配送地区-区(县)ID
	ShipTownID        int     `gorm:"column:ship_town_id" json:"ship_town_id"`               // 配送街道id
	ShipProvince      string  `gorm:"column:ship_province" json:"ship_province"`             // 配送地区-省份
	ShipCity          string  `gorm:"column:ship_city" json:"ship_city"`                     // 配送地区-城市
	ShipCounty        string  `gorm:"column:ship_county" json:"ship_county"`                 // 配送地区-区(县)
	ShipTown          string  `gorm:"column:ship_town" json:"ship_town"`                     // 配送街道
	OrderPrice        float64 `gorm:"column:order_price" json:"order_price"`                 // 订单总额
	GoodsPrice        float64 `gorm:"column:goods_price" json:"goods_price"`                 // 商品总额
	ShippingPrice     float64 `gorm:"column:shipping_price" json:"shipping_price"`           // 配送费用
	DiscountPrice     float64 `gorm:"column:discount_price" json:"discount_price"`           // 优惠金额
	Disabled          int     `gorm:"column:disabled" json:"disabled"`                       // 是否被删除
	Weight            float64 `gorm:"column:weight" json:"weight"`                           // 订单商品总重量
	GoodsNum          int     `gorm:"column:goods_num" json:"goods_num"`                     // 商品数量
	Remark            string  `gorm:"column:remark" json:"remark"`                           // 订单备注
	CancelReason      string  `gorm:"column:cancel_reason" json:"cancel_reason"`             // 订单取消原因
	TheSign           string  `gorm:"column:the_sign" json:"the_sign"`                       // 签收人
	ItemsJSON         string  `gorm:"column:items_json" json:"items_json"`                   // 货物列表json
	WarehouseID       int     `gorm:"column:warehouse_id" json:"warehouse_id"`               // 发货仓库ID
	NeedPayMoney      float64 `gorm:"column:need_pay_money" json:"need_pay_money"`           // 应付金额
	ShipNo            string  `gorm:"column:ship_no" json:"ship_no"`                         // 发货单号
	AddressID         int     `gorm:"column:address_id" json:"address_id"`                   // 收货地址ID
	AdminRemark       int     `gorm:"column:admin_remark" json:"admin_remark"`               // 管理员备注
	LogiID            int     `gorm:"column:logi_id" json:"logi_id"`                         // 物流公司ID
	LogiName          string  `gorm:"column:logi_name" json:"logi_name"`                     // 物流公司名称
	CompleteTime      int64   `gorm:"column:complete_time" json:"complete_time"`             // 完成时间
	CreateTime        int64   `gorm:"column:create_time" json:"create_time"`                 // 订单创建时间
	SigningTime       int64   `gorm:"column:signing_time" json:"signing_time"`               // 签收时间
	ShipTime          int64   `gorm:"column:ship_time" json:"ship_time"`                     // 送货时间
	PayOrderNo        string  `gorm:"column:pay_order_no" json:"pay_order_no"`               // 支付方式返回的交易号
	ServiceStatus     string  `gorm:"column:service_status" json:"service_status"`           // 售后状态
	BillStatus        int     `gorm:"column:bill_status" json:"bill_status"`                 // 结算状态
	BillSn            string  `gorm:"column:bill_sn" json:"bill_sn"`                         // 结算单号
	ClientType        string  `gorm:"column:client_type" json:"client_type"`                 // 订单来源
	Sn                string  `gorm:"column:sn" json:"sn"`                                   // 订单编号
	NeedReceipt       int     `gorm:"column:need_receipt" json:"need_receipt"`               // 是否需要发票
	OrderType         string  `gorm:"column:order_type" json:"order_type"`                   // 订单类型
	OrderData         string  `gorm:"column:order_data" json:"order_data"`                   // 订单数据
}

// TableName get sql table name.获取数据库表名
func (m *EsOrder) TableName() string {
	return "es_order"
}
