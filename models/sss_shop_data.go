package models

// EsSssShopData 商品数据表(es_sss_shop_data)
type EsSssShopData struct {
	ID          int    `gorm:"primaryKey;column:id" json:"-"`           // 主键id
	SellerID    int    `gorm:"column:seller_id" json:"seller_id"`       // 店铺id
	SellerName  string `gorm:"column:seller_name" json:"seller_name"`   // 店铺名称
	FavoriteNum int    `gorm:"column:favorite_num" json:"favorite_num"` // 店铺被收藏数量
	ShopDisable string `gorm:"column:shop_disable" json:"shop_disable"` // 店铺是否开启
}

// TableName get sql table name.获取数据库表名
func (m *EsSssShopData) TableName() string {
	return "es_sss_shop_data"
}