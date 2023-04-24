package models

// EsPintuanChildOrder [...]
type EsPintuanChildOrder struct {
	ChildOrderID int     `gorm:"primaryKey;column:child_order_id" json:"-"` // 子订单id
	OrderSn      string  `gorm:"column:order_sn" json:"order_sn"`           // 订单编号
	MemberID     int     `gorm:"column:member_id" json:"member_id"`         // 会员id
	SkuID        int     `gorm:"column:sku_id" json:"sku_id"`               // 会员id
	PintuanID    int     `gorm:"column:pintuan_id" json:"pintuan_id"`       // 拼团活动id
	OrderStatus  string  `gorm:"column:order_status" json:"order_status"`   // 订单状态
	OrderID      int     `gorm:"column:order_id" json:"order_id"`           // 主订单id
	MemberName   string  `gorm:"column:member_name" json:"member_name"`     // 买家名称
	OriginPrice  float64 `gorm:"column:origin_price" json:"origin_price"`   // 原价
	SalesPrice   float64 `gorm:"column:sales_price" json:"sales_price"`     // 拼团价
}

// TableName get sql table name.获取数据库表名
func (m *EsPintuanChildOrder) TableName() string {
	return "es_pintuan_child_order"
}
