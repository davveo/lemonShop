package models

// EsEmail 邮件记录(es_email)
type EsEmail struct {
	ID       int    `gorm:"primaryKey;column:id" json:"-"`     // 邮件记录id
	Title    string `gorm:"column:title" json:"title"`         // 邮件标题
	Type     string `gorm:"column:type" json:"type"`           // 邮件类型
	Success  int    `gorm:"column:success" json:"success"`     // 是否成功
	Email    string `gorm:"column:email" json:"email"`         // 邮件接收者
	Content  string `gorm:"column:content" json:"content"`     // 邮件内容
	ErrorNum int    `gorm:"column:error_num" json:"error_num"` // 错误次数
	LastSend int64  `gorm:"column:last_send" json:"last_send"` // 最后发送时间
}

// TableName get sql table name.获取数据库表名
func (m *EsEmail) TableName() string {
	return "es_email"
}
