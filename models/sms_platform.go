package models

// EsSmsPlatform 短信网关(es_sms_platform)
type EsSmsPlatform struct {
	ID     int    `gorm:"primaryKey;column:id" json:"-"` // 主键ID
	Name   string `gorm:"column:name" json:"name"`       // 平台名称
	Open   int    `gorm:"column:open" json:"open"`       // 是否开启
	Config string `gorm:"column:config" json:"config"`   // 配置
	Bean   string `gorm:"column:bean" json:"bean"`       // beanid
}

// TableName get sql table name.获取数据库表名
func (m *EsSmsPlatform) TableName() string {
	return "es_sms_platform"
}