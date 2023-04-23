package models

// EsAsRefund 售后服务退款相关信息表(es_as_refund)
type EsAsRefund struct {
	ID                int     `gorm:"primaryKey;column:id" json:"-"`                         // 主键ID
	ServiceSn         string  `gorm:"column:service_sn" json:"service_sn"`                   // 售后服务单号
	RefundPrice       float64 `gorm:"column:refund_price" json:"refund_price"`               // 申请的退款金额
	AgreePrice        float64 `gorm:"column:agree_price" json:"agree_price"`                 // 商家同意的退款金额
	ActualPrice       float64 `gorm:"column:actual_price" json:"actual_price"`               // 实际退款金额
	RefundWay         string  `gorm:"column:refund_way" json:"refund_way"`                   // 退款方式 ORIGINAL：原路退回，OFFLINE：线下支付
	AccountType       string  `gorm:"column:account_type" json:"account_type"`               // 账号类型
	ReturnAccount     string  `gorm:"column:return_account" json:"return_account"`           // 退款账号
	BankName          string  `gorm:"column:bank_name" json:"bank_name"`                     // 银行名称
	BankAccountNumber string  `gorm:"column:bank_account_number" json:"bank_account_number"` // 银行账户
	BankAccountName   string  `gorm:"column:bank_account_name" json:"bank_account_name"`     // 银行开户名
	BankDepositName   string  `gorm:"column:bank_deposit_name" json:"bank_deposit_name"`     // 开户行
	PayOrderNo        string  `gorm:"column:pay_order_no" json:"pay_order_no"`               // 订单支付方式返回的交易号
	RefundTime        int64   `gorm:"column:refund_time" json:"refund_time"`                 // 退款时间
}

// TableName get sql table name.获取数据库表名
func (m *EsAsRefund) TableName() string {
	return "es_as_refund"
}
