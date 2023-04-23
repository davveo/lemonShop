package models

// EsCommissionTpl 模版(es_commission_tpl)
type EsCommissionTpl struct {
	ID          int64   `gorm:"primaryKey;column:id" json:"-"`           // id
	TplName     string  `gorm:"column:tpl_name" json:"tpl_name"`         // 名称
	TplDescribe string  `gorm:"column:tpl_describe" json:"tpl_describe"` // 描述
	Cycle       int     `gorm:"column:cycle" json:"cycle"`               // 周期
	Grade1      float64 `gorm:"column:grade1" json:"grade1"`             // 1级返利
	Grade2      float64 `gorm:"column:grade2" json:"grade2"`             // 2级返利
	IsDefault   int16   `gorm:"column:is_default" json:"is_default"`     // 是否默认
}

// TableName get sql table name.获取数据库表名
func (m *EsCommissionTpl) TableName() string {
	return "es_commission_tpl"
}
