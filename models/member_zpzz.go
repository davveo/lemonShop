package models

// EsMemberZpzz 会员增票资质表(es_member_zpzz)
type EsMemberZpzz struct {
	ID              int    `gorm:"primaryKey;column:id" json:"-"`                   // 主键ID
	MemberID        int    `gorm:"column:member_id" json:"member_id"`               // 会员ID
	Uname           string `gorm:"column:uname" json:"uname"`                       // 会员用户名
	Status          string `gorm:"column:status" json:"status"`                     // 状态 NEW_APPLY：新申请，AUDIT_PASS：审核通过，AUDIT_REFUSE：审核未通过
	CompanyName     string `gorm:"column:company_name" json:"company_name"`         // 单位名称
	TaxpayerCode    string `gorm:"column:taxpayer_code" json:"taxpayer_code"`       // 纳税人识别码
	RegisterAddress string `gorm:"column:register_address" json:"register_address"` // 公司注册地址
	RegisterTel     string `gorm:"column:register_tel" json:"register_tel"`         // 公司注册电话
	BankName        string `gorm:"column:bank_name" json:"bank_name"`               // 开户银行
	BankAccount     string `gorm:"column:bank_account" json:"bank_account"`         // 银行账户
	AuditRemark     string `gorm:"column:audit_remark" json:"audit_remark"`         // 平台审核备注
	ApplyTime       int64  `gorm:"column:apply_time" json:"apply_time"`             // 申请时间
}

// TableName get sql table name.获取数据库表名
func (m *EsMemberZpzz) TableName() string {
	return "es_member_zpzz"
}
