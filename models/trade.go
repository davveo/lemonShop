package models

// EsTrade 交易表(es_trade)
type EsTrade struct {
	TradeID             int64   `gorm:"primaryKey;column:trade_id" json:"-"`                       // trade_id
	TradeSn             string  `gorm:"column:trade_sn" json:"trade_sn"`                           // 交易编号
	MemberID            int     `gorm:"column:member_id" json:"member_id"`                         // 买家id
	MemberName          string  `gorm:"column:member_name" json:"member_name"`                     // 买家用户名
	PaymentMethodID     string  `gorm:"column:payment_method_id" json:"payment_method_id"`         // 支付方式id
	PaymentPluginID     string  `gorm:"column:payment_plugin_id" json:"payment_plugin_id"`         // 支付插件id
	PaymentMethodName   string  `gorm:"column:payment_method_name" json:"payment_method_name"`     // 支付方式名称
	PaymentType         string  `gorm:"column:payment_type" json:"payment_type"`                   // 支付方式类型
	TotalPrice          float64 `gorm:"column:total_price" json:"total_price"`                     // 总价格
	GoodsPrice          float64 `gorm:"column:goods_price" json:"goods_price"`                     // 商品价格
	FreightPrice        float64 `gorm:"column:freight_price" json:"freight_price"`                 // 运费
	DiscountPrice       float64 `gorm:"column:discount_price" json:"discount_price"`               // 优惠的金额
	ConsigneeID         int     `gorm:"column:consignee_id" json:"consignee_id"`                   // 收货人id
	ConsigneeName       string  `gorm:"column:consignee_name" json:"consignee_name"`               // 收货人姓名
	ConsigneeCountry    string  `gorm:"column:consignee_country" json:"consignee_country"`         // 收货国家
	ConsigneeCountryID  int     `gorm:"column:consignee_country_id" json:"consignee_country_id"`   // 收货国家id
	ConsigneeProvince   string  `gorm:"column:consignee_province" json:"consignee_province"`       // 收货省
	ConsigneeProvinceID int     `gorm:"column:consignee_province_id" json:"consignee_province_id"` // 收货省id
	ConsigneeCity       string  `gorm:"column:consignee_city" json:"consignee_city"`               // 收货市
	ConsigneeCityID     int     `gorm:"column:consignee_city_id" json:"consignee_city_id"`         // 收货市id
	ConsigneeCounty     string  `gorm:"column:consignee_county" json:"consignee_county"`           // 收货区
	ConsigneeCountyID   int     `gorm:"column:consignee_county_id" json:"consignee_county_id"`     // 收货区id
	ConsigneeTown       string  `gorm:"column:consignee_town" json:"consignee_town"`               // 收货镇
	ConsigneeTownID     int     `gorm:"column:consignee_town_id" json:"consignee_town_id"`         // 收货镇id
	ConsigneeAddress    string  `gorm:"column:consignee_address" json:"consignee_address"`         // 收货详细地址
	ConsigneeMobile     string  `gorm:"column:consignee_mobile" json:"consignee_mobile"`           // 收货人手机号
	ConsigneeTelephone  string  `gorm:"column:consignee_telephone" json:"consignee_telephone"`     // 收货人电话
	CreateTime          int64   `gorm:"column:create_time" json:"create_time"`                     // 交易创建时间
	OrderJSON           string  `gorm:"column:order_json" json:"order_json"`                       // 订单json(预留，7.0可能废弃)
	TradeStatus         string  `gorm:"column:trade_status" json:"trade_status"`                   // 订单状态
}

// TableName get sql table name.获取数据库表名
func (m *EsTrade) TableName() string {
	return "es_trade"
}
