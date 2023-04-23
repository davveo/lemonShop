package models

// EsExpressPlatform 快递平台(es_express_platform)
type EsExpressPlatform struct {
	ID     int    `gorm:"primaryKey;column:id" json:"-"` // 快递平台id
	Name   string `gorm:"column:name" json:"name"`       // 快递平台名称
	Open   int    `gorm:"column:open" json:"open"`       // 是否开启快递平台,1开启，0未开启
	Config string `gorm:"column:config" json:"config"`   // 快递平台配置
	Bean   string `gorm:"column:bean" json:"bean"`       // 快递平台beanid
}

// TableName get sql table name.获取数据库表名
func (m *EsExpressPlatform) TableName() string {
	return "es_express_platform"
}
