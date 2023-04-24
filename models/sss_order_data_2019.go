package models

// EsSssOrderData2019 [...]
type EsSssOrderData2019 struct {
	ID             int     `gorm:"primaryKey;column:id" json:"-"`
	Sn             string  `gorm:"column:sn" json:"sn"`
	BuyerID        int     `gorm:"column:buyer_id" json:"buyer_id"`
	BuyerName      string  `gorm:"column:buyer_name" json:"buyer_name"`
	SellerID       int     `gorm:"column:seller_id" json:"seller_id"`
	SellerName     string  `gorm:"column:seller_name" json:"seller_name"`
	OrderStatus    string  `gorm:"column:order_status" json:"order_status"`
	PayStatus      string  `gorm:"column:pay_status" json:"pay_status"`
	OrderPrice     float64 `gorm:"column:order_price" json:"order_price"`
	GoodsNum       int     `gorm:"column:goods_num" json:"goods_num"`
	ShipProvinceID int     `gorm:"column:ship_province_id" json:"ship_province_id"`
	ShipCityID     int     `gorm:"column:ship_city_id" json:"ship_city_id"`
	CreateTime     int     `gorm:"column:create_time" json:"create_time"`
}

// TableName get sql table name.获取数据库表名
func (m *EsSssOrderData2019) TableName() string {
	return "es_sss_order_data_2019"
}
