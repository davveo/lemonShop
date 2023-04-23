package models

// EsConnectSetting 信任登录参数(es_connect_setting)
type EsConnectSetting struct {
	ID     int    `gorm:"primaryKey;column:id" json:"-"` // id
	Type   string `gorm:"column:type" json:"type"`       // 授权类型
	Config string `gorm:"column:config" json:"config"`   // 参数组
	Name   string `gorm:"column:name" json:"name"`       // 标题
}

// TableName get sql table name.获取数据库表名
func (m *EsConnectSetting) TableName() string {
	return "es_connect_setting"
}
