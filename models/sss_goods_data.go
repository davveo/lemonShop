package models

// EsSssGoodsData 统计库商品数据(es_sss_goods_data)
type EsSssGoodsData struct {
	ID           int     `gorm:"primaryKey;column:id" json:"-"`             // id
	GoodsID      int     `gorm:"column:goods_id" json:"goods_id"`           // 商品id
	GoodsName    string  `gorm:"column:goods_name" json:"goods_name"`       // 商品名称
	BrandID      int     `gorm:"column:brand_id" json:"brand_id"`           // 品牌id
	CategoryID   int     `gorm:"column:category_id" json:"category_id"`     // 分类id
	CategoryPath string  `gorm:"column:category_path" json:"category_path"` // 分类路径
	SellerID     int     `gorm:"column:seller_id" json:"seller_id"`         // 商家id
	ShopCatID    int     `gorm:"column:shop_cat_id" json:"shop_cat_id"`     // 商家分类id
	Price        float64 `gorm:"column:price" json:"price"`                 // 商品价格
	FavoriteNum  int     `gorm:"column:favorite_num" json:"favorite_num"`   // 收藏数量
	MarketEnable int16   `gorm:"column:market_enable" json:"market_enable"` // 是否上架
}

// TableName get sql table name.获取数据库表名
func (m *EsSssGoodsData) TableName() string {
	return "es_sss_goods_data"
}