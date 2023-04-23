package models


// EsCategoryBrand 分类品牌关联表(es_category_brand)
type EsCategoryBrand struct {
	CategoryID int `gorm:"column:category_id" json:"category_id"` // 分类id
	BrandID    int `gorm:"column:brand_id" json:"brand_id"`       // 品牌id
}

// TableName get sql table name.获取数据库表名
func (m *EsCategoryBrand) TableName() string {
	return "es_category_brand"
}
