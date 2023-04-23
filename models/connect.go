package models

// EsConnect 信任登录(es_connect)
type EsConnect struct {
	ID          int    `gorm:"primaryKey;column:id" json:"-"`           // id
	MemberID    int    `gorm:"column:member_id" json:"member_id"`       // 会员id
	UnionID     string `gorm:"column:union_id" json:"union_id"`         // 第三方唯一标示
	UnionType   string `gorm:"column:union_type" json:"union_type"`     // 信任登录类型
	UnboundTime int64  `gorm:"column:unbound_time" json:"unbound_time"` // 解绑时间
}

// TableName get sql table name.获取数据库表名
func (m *EsConnect) TableName() string {
	return "es_connect"
}
