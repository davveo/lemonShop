package models

// EsShopSilde 店铺幻灯片(es_shop_silde)
type EsShopSilde struct {
	SildeID  int    `gorm:"primaryKey;column:silde_id" json:"-"` // 幻灯片Id
	ShopID   int    `gorm:"column:shop_id" json:"shop_id"`       // 店铺Id
	SildeURL string `gorm:"column:silde_url" json:"silde_url"`   // 幻灯片URL
	Img      string `gorm:"column:img" json:"img"`               // 图片
}

// TableName get sql table name.获取数据库表名
func (m *EsShopSilde) TableName() string {
	return "es_shop_silde"
}