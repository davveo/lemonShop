package models

// EsWechatMsgTemplate [...]
type EsWechatMsgTemplate struct {
	ID         int    `gorm:"primaryKey;column:id" json:"-"`           // 主键
	MsgTmpName string `gorm:"column:msg_tmp_name" json:"msg_tmp_name"` // 模板名称
	MsgTmpSn   string `gorm:"column:msg_tmp_sn" json:"msg_tmp_sn"`     // 消息编号
	TemplateID string `gorm:"column:template_id" json:"template_id"`   // 模板ID
	MsgFirst   string `gorm:"column:msg_first" json:"msg_first"`       // 消息开头文字
	MsgRemark  string `gorm:"column:msg_remark" json:"msg_remark"`     // 消息结尾备注文字
	IsOpen     int16  `gorm:"column:is_open" json:"is_open"`           // 是否开启
	TmpType    string `gorm:"column:tmp_type" json:"tmp_type"`         // 模板类型，枚举
}

// TableName get sql table name.获取数据库表名
func (m *EsWechatMsgTemplate) TableName() string {
	return "es_wechat_msg_template"
}