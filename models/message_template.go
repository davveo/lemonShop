package models

// EsMessageTemplate 消息模版(es_message_template)
type EsMessageTemplate struct {
	ID           int    `gorm:"primaryKey;column:id" json:"-"`             // 模版ID
	TplCode      string `gorm:"column:tpl_code" json:"tpl_code"`           // 模版编号
	TplName      string `gorm:"column:tpl_name" json:"tpl_name"`           // 模板名称
	Type         string `gorm:"column:type" json:"type"`                   // 类型(会员 ,店铺 ,其他)
	NoticeState  string `gorm:"column:notice_state" json:"notice_state"`   // 站内信提醒是否开启
	SmsState     string `gorm:"column:sms_state" json:"sms_state"`         // 短信提醒是否开启
	EmailState   string `gorm:"column:email_state" json:"email_state"`     // 邮件提醒是否开启
	Content      string `gorm:"column:content" json:"content"`             // 站内信内容
	SmsContent   string `gorm:"column:sms_content" json:"sms_content"`     // 短信内容
	EmailContent string `gorm:"column:email_content" json:"email_content"` // 邮件内容
	EmailTitle   string `gorm:"column:email_title" json:"email_title"`     // 邮件标题
}

// TableName get sql table name.获取数据库表名
func (m *EsMessageTemplate) TableName() string {
	return "es_message_template"
}
