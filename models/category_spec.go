package models

// EsCategorySpec 分类规格关联表(es_category_spec)
type EsCategorySpec struct {
	CategoryID int `gorm:"column:category_id" json:"category_id"` // 分类id
	SpecID     int `gorm:"column:spec_id" json:"spec_id"`         // 规格id
}

// TableName get sql table name.获取数据库表名
func (m *EsCategorySpec) TableName() string {
	return "es_category_spec"
}
