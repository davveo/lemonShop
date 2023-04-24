package models

// EsPintuanOrder [...]
type EsPintuanOrder struct {
	OrderID        int    `gorm:"primaryKey;column:order_id" json:"-"`           // 订单id
	PintuanID      int    `gorm:"column:pintuan_id" json:"pintuan_id"`           // 拼团id
	EndTime        int64  `gorm:"column:end_time" json:"end_time"`               // 结束时间
	SkuID          int    `gorm:"column:sku_id" json:"sku_id"`                   // sku_id
	RequiredNum    int    `gorm:"column:required_num" json:"required_num"`       // 成团人数
	OfferedNum     int    `gorm:"column:offered_num" json:"offered_num"`         // 已参团人数
	OfferedPersons string `gorm:"column:offered_persons" json:"offered_persons"` // 参团人
	OrderStatus    string `gorm:"column:order_status" json:"order_status"`       // 订单状态
	GoodsID        int    `gorm:"column:goods_id" json:"goods_id"`               // 商品id
	Thumbnail      string `gorm:"column:thumbnail" json:"thumbnail"`             // 商品缩略图
	GoodsName      string `gorm:"column:goods_name" json:"goods_name"`           // 商品名称
}

// TableName get sql table name.获取数据库表名
func (m *EsPintuanOrder) TableName() string {
	return "es_pintuan_order"
}
