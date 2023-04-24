package models

// EsGroupbuyArea 团购地区(es_groupbuy_area)
type EsGroupbuyArea struct {
	AreaID    int    `gorm:"primaryKey;column:area_id" json:"-"`  // 地区ID
	AreaName  string `gorm:"column:area_name" json:"area_name"`   // 地区名称
	AreaOrder int    `gorm:"column:area_order" json:"area_order"` // 地区排序
	AreaPath  string `gorm:"column:area_path" json:"area_path"`   // 地区路径结构
	ParentID  int    `gorm:"column:parent_id" json:"parent_id"`   // 地区父节点
}

// TableName get sql table name.获取数据库表名
func (m *EsGroupbuyArea) TableName() string {
	return "es_groupbuy_area"
}
