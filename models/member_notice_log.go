package models

// EsMemberNoticeLog 会员站内消息历史(es_member_notice_log)
type EsMemberNoticeLog struct {
	ID          int    `gorm:"primaryKey;column:id" json:"-"`           // 会员历史消息id
	MemberID    int    `gorm:"column:member_id" json:"member_id"`       // 会员id
	Content     string `gorm:"column:content" json:"content"`           // 站内信内容
	SendTime    int64  `gorm:"column:send_time" json:"send_time"`       // 发送时间
	IsDel       int    `gorm:"column:is_del" json:"is_del"`             // 是否删除，0删除，1没有删除
	IsRead      int    `gorm:"column:is_read" json:"is_read"`           // 是否已读，0未读，1已读
	ReceiveTime int64  `gorm:"column:receive_time" json:"receive_time"` // 接收时间
	Title       string `gorm:"column:title" json:"title"`               // 标题
}

// TableName get sql table name.获取数据库表名
func (m *EsMemberNoticeLog) TableName() string {
	return "es_member_notice_log"
}
