package models

// EsSssOrderData2017 [...]
type EsSssOrderData2017 struct {
	ID             int     `gorm:"primaryKey;column:id" json:"-"`                   // id
	Sn             string  `gorm:"column:sn" json:"sn"`                             // 订单编号
	BuyerID        int     `gorm:"column:buyer_id" json:"buyer_id"`                 // 会员id
	BuyerName      string  `gorm:"column:buyer_name" json:"buyer_name"`             // 商家名称
	SellerID       int     `gorm:"column:seller_id" json:"seller_id"`               // 商家id
	SellerName     string  `gorm:"column:seller_name" json:"seller_name"`           // 商家名称
	OrderStatus    string  `gorm:"column:order_status" json:"order_status"`         // 订单状态
	PayStatus      string  `gorm:"column:pay_status" json:"pay_status"`             // 付款状态
	OrderPrice     float64 `gorm:"column:order_price" json:"order_price"`           // 订单金额
	GoodsNum       int     `gorm:"column:goods_num" json:"goods_num"`               // 订单商品数量
	ShipProvinceID int     `gorm:"column:ship_province_id" json:"ship_province_id"` // 省id
	ShipCityID     int     `gorm:"column:ship_city_id" json:"ship_city_id"`         // 市id
	CreateTime     int     `gorm:"column:create_time" json:"create_time"`           // 订单创建时间
}

// TableName get sql table name.获取数据库表名
func (m *EsSssOrderData2017) TableName() string {
	return "es_sss_order_data_2017"
}