package models

// EsWithdrawApply 提现申请(es_withdraw_apply)
type EsWithdrawApply struct {
	ID             int     `gorm:"primaryKey;column:id" json:"-"`                 // id
	ApplyMoney     float64 `gorm:"column:apply_money" json:"apply_money"`         // 提现金额
	MemberID       int     `gorm:"column:member_id" json:"member_id"`             // 会员id
	MemberName     string  `gorm:"column:member_name" json:"member_name"`         // 会员名称
	ApplyRemark    string  `gorm:"column:apply_remark" json:"apply_remark"`       // 申请备注
	InspectRemark  string  `gorm:"column:inspect_remark" json:"inspect_remark"`   // 审核备注
	TransferRemark string  `gorm:"column:transfer_remark" json:"transfer_remark"` // 转账备注
	ApplyTime      int     `gorm:"column:apply_time" json:"apply_time"`           // 申请时间
	InspectTime    int     `gorm:"column:inspect_time" json:"inspect_time"`       // 审核时间
	TransferTime   int     `gorm:"column:transfer_time" json:"transfer_time"`     // 转账时间
	Status         string  `gorm:"column:status" json:"status"`                   // 状态
	Sn             string  `gorm:"column:sn" json:"sn"`                           // 编号
	IP             string  `gorm:"column:ip" json:"ip"`                           // ip地址
}

// TableName get sql table name.获取数据库表名
func (m *EsWithdrawApply) TableName() string {
	return "es_withdraw_apply"
}