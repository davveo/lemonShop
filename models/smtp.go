package models

// EsSMTP 邮件(es_smtp)
type EsSMTP struct {
	ID           int    `gorm:"primaryKey;column:id" json:"-"`               // 主键ID
	Host         string `gorm:"column:host" json:"host"`                     // 主机
	Username     string `gorm:"column:username" json:"username"`             // 用户名
	Password     string `gorm:"column:password" json:"password"`             // 密码
	LastSendTime int64  `gorm:"column:last_send_time" json:"last_send_time"` // 最后发信时间
	SendCount    int    `gorm:"column:send_count" json:"send_count"`         // 已发数
	MaxCount     int    `gorm:"column:max_count" json:"max_count"`           // 最大发信数
	MailFrom     string `gorm:"column:mail_from" json:"mail_from"`           // 发信邮箱
	Port         int    `gorm:"column:port" json:"port"`                     // 端口
	OpenSsl      int    `gorm:"column:open_ssl" json:"open_ssl"`             // ssl是否开启
}

// TableName get sql table name.获取数据库表名
func (m *EsSMTP) TableName() string {
	return "es_smtp"
}