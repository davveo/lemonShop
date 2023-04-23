package models

// EsAskReply [...]
type EsAskReply struct {
	ID          int    `gorm:"primaryKey;column:id" json:"-"`           // 主键ID
	AskID       int    `gorm:"column:ask_id" json:"ask_id"`             // 会员咨询id
	MemberID    int    `gorm:"column:member_id" json:"member_id"`       // 会员id
	MemberName  string `gorm:"column:member_name" json:"member_name"`   // 会员名称
	Content     string `gorm:"column:content" json:"content"`           // 回复内容
	ReplyTime   int64  `gorm:"column:reply_time" json:"reply_time"`     // 回复时间
	Anonymous   string `gorm:"column:anonymous" json:"anonymous"`       // 是否匿名 YES:是，NO:否
	AuthStatus  string `gorm:"column:auth_status" json:"auth_status"`   // 审核状态 WAIT_AUDIT:待审核,PASS_AUDIT:审核通过,REFUSE_AUDIT:审核未通过
	IsDel       string `gorm:"column:is_del" json:"is_del"`             // 删除状态 DELETED：已删除 NORMAL：正常
	ReplyStatus string `gorm:"column:reply_status" json:"reply_status"` // 是否已回复 YES:是，NO:否
	CreateTime  int64  `gorm:"column:create_time" json:"create_time"`   // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *EsAskReply) TableName() string {
	return "es_ask_reply"
}
