package models

// EsShopCat 店铺分组(es_shop_cat)
type EsShopCat struct {
	ShopCatID   int    `gorm:"primaryKey;column:shop_cat_id" json:"-"`    // 店铺分组id
	ShopCatPid  int    `gorm:"column:shop_cat_pid" json:"shop_cat_pid"`   // 店铺分组父ID
	ShopID      int    `gorm:"column:shop_id" json:"shop_id"`             // 店铺id
	ShopCatName string `gorm:"column:shop_cat_name" json:"shop_cat_name"` // 店铺分组名称
	Disable     int    `gorm:"column:disable" json:"disable"`             // 店铺分组显示状态:1显示 0不显示
	Sort        int    `gorm:"column:sort" json:"sort"`                   // 排序
	CatPath     string `gorm:"column:cat_path" json:"cat_path"`           // 分组路径
}

// TableName get sql table name.获取数据库表名
func (m *EsShopCat) TableName() string {
	return "es_shop_cat"
}