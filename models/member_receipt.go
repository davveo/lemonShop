package models

// EsMemberReceipt 会员发票信息缓存表(es_member_receipt)
type EsMemberReceipt struct {
	ReceiptID      int    `gorm:"primaryKey;column:receipt_id" json:"-"`         // 主键ID
	MemberID       int    `gorm:"column:member_id" json:"member_id"`             // 会员ID
	ReceiptType    string `gorm:"column:receipt_type" json:"receipt_type"`       // 发票类型 ELECTRO：电子普通发票，VATORDINARY：增值税普通发票
	ReceiptTitle   string `gorm:"column:receipt_title" json:"receipt_title"`     // 发票抬头
	ReceiptContent string `gorm:"column:receipt_content" json:"receipt_content"` // 发票内容
	TaxNo          string `gorm:"column:tax_no" json:"tax_no"`                   // 纳税人识别号
	MemberMobile   string `gorm:"column:member_mobile" json:"member_mobile"`     // 收票人手机号
	MemberEmail    string `gorm:"column:member_email" json:"member_email"`       // 收票人邮箱
	IsDefault      int16  `gorm:"column:is_default" json:"is_default"`           // 是否为默认选项 0：否，1：是
}

// TableName get sql table name.获取数据库表名
func (m *EsMemberReceipt) TableName() string {
	return "es_member_receipt"
}
