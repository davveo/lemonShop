package models

// EsMessage 站内消息(es_message)
type EsMessage struct {
	ID        int    `gorm:"primaryKey;column:id" json:"-"`       // 站内消息主键
	Title     string `gorm:"column:title" json:"title"`           // 标题
	Content   string `gorm:"column:content" json:"content"`       // 消息内容
	MemberIDs string `gorm:"column:member_ids" json:"member_ids"` // 会员id
	AdminID   int    `gorm:"column:admin_id" json:"admin_id"`     // 管理员id
	AdminName string `gorm:"column:admin_name" json:"admin_name"` // 管理员名称
	SendTime  int64  `gorm:"column:send_time" json:"send_time"`   // 发送时间
	SendType  int    `gorm:"column:send_type" json:"send_type"`   // 发送类型
	Disabled  int    `gorm:"column:disabled" json:"disabled"`     // 是否删除 0：否，1：是
}

// TableName get sql table name.获取数据库表名
func (m *EsMessage) TableName() string {
	return "es_message"
}
