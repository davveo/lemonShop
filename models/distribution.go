package models

// EsDistribution 分销商(es_distribution)
type EsDistribution struct {
	ID                  int64   `gorm:"primaryKey;column:id" json:"-"`                             // id
	MemberID            int     `gorm:"column:member_id" json:"member_id"`                         // 会员id
	MemberName          string  `gorm:"column:member_name" json:"member_name"`                     // 会员名称
	Path                string  `gorm:"column:path" json:"path"`                                   // 关系路径
	MemberIDLv2         int64   `gorm:"column:member_id_lv2" json:"member_id_lv2"`                 // 上上级
	MemberIDLv1         int64   `gorm:"column:member_id_lv1" json:"member_id_lv1"`                 // 上级
	Downline            int64   `gorm:"column:downline" json:"downline"`                           // 下线人数
	OrderNum            int64   `gorm:"column:order_num" json:"order_num"`                         // 订单数
	RebateTotal         float64 `gorm:"column:rebate_total" json:"rebate_total"`                   // 返利总额
	TurnoverPrice       float64 `gorm:"column:turnover_price" json:"turnover_price"`               // 营业额总额
	CanRebate           float64 `gorm:"column:can_rebate" json:"can_rebate"`                       // 可提现金额
	CommissionFrozen    float64 `gorm:"column:commission_frozen" json:"commission_frozen"`         // 返利金额冻结
	CurrentTplName      string  `gorm:"column:current_tpl_name" json:"current_tpl_name"`           // 使用模版名称
	CurrentTplID        int     `gorm:"column:current_tpl_id" json:"current_tpl_id"`               // 使用模版
	WithdrawFrozenPrice float64 `gorm:"column:withdraw_frozen_price" json:"withdraw_frozen_price"` // 提现冻结
	CreateTime          int64   `gorm:"column:create_time" json:"create_time"`                     // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *EsDistribution) TableName() string {
	return "es_distribution"
}
