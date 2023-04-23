package models


// EsBrand 品牌(es_brand)
type EsBrand struct {
	BrandID  int    `gorm:"primaryKey;column:brand_id" json:"-"` // 主键
	Name     string `gorm:"column:name" json:"name"`             // 品牌名称
	Logo     string `gorm:"column:logo" json:"logo"`             // 品牌图标
	Disabled int    `gorm:"column:disabled" json:"disabled"`     // 是否删除，0删除1未删除
}

// TableName get sql table name.获取数据库表名
func (m *EsBrand) TableName() string {
	return "es_brand"
}

