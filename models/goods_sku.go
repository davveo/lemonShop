package models

// EsGoodsSku 商品sku(es_goods_sku)
type EsGoodsSku struct {
	SkuID          int     `gorm:"primaryKey;column:sku_id" json:"-"`             // 主键
	GoodsID        int     `gorm:"column:goods_id" json:"goods_id"`               // 商品id
	GoodsName      string  `gorm:"column:goods_name" json:"goods_name"`           // 商品名称
	Sn             string  `gorm:"column:sn" json:"sn"`                           // 商品编号
	Quantity       int     `gorm:"column:quantity" json:"quantity"`               // 库存
	EnableQuantity int     `gorm:"column:enable_quantity" json:"enable_quantity"` // 可用库存
	Price          float64 `gorm:"column:price" json:"price"`                     // 商品价格
	Specs          string  `gorm:"column:specs" json:"specs"`                     // 规格信息json
	Cost           float64 `gorm:"column:cost" json:"cost"`                       // 成本价格
	Weight         float64 `gorm:"column:weight" json:"weight"`                   // 重量
	SellerID       int     `gorm:"column:seller_id" json:"seller_id"`             // 卖家id
	SellerName     string  `gorm:"column:seller_name" json:"seller_name"`         // 卖家名称
	CategoryID     int     `gorm:"column:category_id" json:"category_id"`         // 分类id
	Thumbnail      string  `gorm:"column:thumbnail" json:"thumbnail"`             // 缩略图
	HashCode       int     `gorm:"column:hash_code" json:"hash_code"`             // 标识规格唯一性
}

// TableName get sql table name.获取数据库表名
func (m *EsGoodsSku) TableName() string {
	return "es_goods_sku"
}
