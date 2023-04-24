package models

// EsPintuanGoods [...]
type EsPintuanGoods struct {
	ID             int     `gorm:"primaryKey;column:id" json:"-"`                 // id
	SkuID          int     `gorm:"column:sku_id" json:"sku_id"`                   // sku_id
	GoodsName      string  `gorm:"column:goods_name" json:"goods_name"`           // 商品名称
	OriginPrice    float64 `gorm:"column:origin_price" json:"origin_price"`       // 原价
	SalesPrice     float64 `gorm:"column:sales_price" json:"sales_price"`         // 活动价
	Sn             string  `gorm:"column:sn" json:"sn"`                           // sn
	SoldQuantity   int     `gorm:"column:sold_quantity" json:"sold_quantity"`     // 已售数量
	LockedQuantity int     `gorm:"column:locked_quantity" json:"locked_quantity"` // 待发货数量
	PintuanID      int     `gorm:"column:pintuan_id" json:"pintuan_id"`           // 拼团活动id
	GoodsID        int     `gorm:"column:goods_id" json:"goods_id"`               // goods_id
	Specs          string  `gorm:"column:specs" json:"specs"`                     // 规格
	SellerID       int     `gorm:"column:seller_id" json:"seller_id"`             // 卖家id
	SellerName     string  `gorm:"column:seller_name" json:"seller_name"`         // 卖家名称
	Thumbnail      string  `gorm:"column:thumbnail" json:"thumbnail"`             // 商品缩略图
}

// TableName get sql table name.获取数据库表名
func (m *EsPintuanGoods) TableName() string {
	return "es_pintuan_goods"
}
