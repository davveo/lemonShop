package models

// EsParameterGroup 参数组(es_parameter_group)
type EsParameterGroup struct {
	GroupID    int    `gorm:"primaryKey;column:group_id" json:"-"`   // 主键
	GroupName  string `gorm:"column:group_name" json:"group_name"`   // 参数组名称
	CategoryID int    `gorm:"column:category_id" json:"category_id"` // 关联分类id
	Sort       int    `gorm:"column:sort" json:"sort"`               // 排序
}

// TableName get sql table name.获取数据库表名
func (m *EsParameterGroup) TableName() string {
	return "es_parameter_group"
}
