package models

// EsWaybill 电子面单(es_waybill)
type EsWaybill struct {
	ID     int    `gorm:"primaryKey;column:id" json:"-"` // 电子面单id
	Name   string `gorm:"column:name" json:"name"`       // 名称
	Open   int    `gorm:"column:open" json:"open"`       // 是否开启
	Config string `gorm:"column:config" json:"config"`   // 电子面单配置
	Bean   string `gorm:"column:bean" json:"bean"`       // beanid
}

// TableName get sql table name.获取数据库表名
func (m *EsWaybill) TableName() string {
	return "es_waybill"
}