package models

// EsParameters 参数(es_parameters)
type EsParameters struct {
	ParamID    int    `gorm:"primaryKey;column:param_id" json:"-"`   // 主键
	ParamName  string `gorm:"column:param_name" json:"param_name"`   // 参数名称
	ParamType  int    `gorm:"column:param_type" json:"param_type"`   // 参数类型1 输入项   2 选择项
	Options    string `gorm:"column:options" json:"options"`         // 选择值，当参数类型是选择项2时，必填，逗号分隔
	IsIndex    int    `gorm:"column:is_index" json:"is_index"`       // 是否可索引，0 不显示 1 显示
	Required   int    `gorm:"column:required" json:"required"`       // 是否必填是  1    否   0
	GroupID    int    `gorm:"column:group_id" json:"group_id"`       // 参数分组id
	CategoryID int    `gorm:"column:category_id" json:"category_id"` // 分类id
	Sort       int    `gorm:"column:sort" json:"sort"`               // 排序
}

// TableName get sql table name.获取数据库表名
func (m *EsParameters) TableName() string {
	return "es_parameters"
}
