package models

// EsShopThemes 店铺模版(es_shop_themes)
type EsShopThemes struct {
	ID        int    `gorm:"primaryKey;column:id" json:"-"`       // 模版id
	Name      string `gorm:"column:name" json:"name"`             // 模版名称
	ImagePath string `gorm:"column:image_path" json:"image_path"` // 模版图片路径
	IsDefault int    `gorm:"column:is_default" json:"is_default"` // 是否为默认
	Type      string `gorm:"column:type" json:"type"`             // 模版类型
	Mark      string `gorm:"column:mark" json:"mark"`             // 模版标识
}

// TableName get sql table name.获取数据库表名
func (m *EsShopThemes) TableName() string {
	return "es_shop_themes"
}