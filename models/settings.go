package models

// EsSettings 系统设置(es_settings)
type EsSettings struct {
	ID       int    `gorm:"primaryKey;column:id" json:"-"`     // 系统设置id
	CfgValue string `gorm:"column:cfg_value" json:"cfg_value"` // 系统配置信息
	CfgGroup string `gorm:"column:cfg_group" json:"cfg_group"` // 业务设置标识
}

// TableName get sql table name.获取数据库表名
func (m *EsSettings) TableName() string {
	return "es_settings"
}