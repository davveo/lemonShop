package models

// EsRefund 售后服务退款单(es_refund)
type EsRefund struct {
	ID                int     `gorm:"primaryKey;column:id" json:"-"`                         // 主键id
	Sn                string  `gorm:"column:sn" json:"sn"`                                   // 售后服务单号
	MemberID          int     `gorm:"column:member_id" json:"member_id"`                     // 会员id
	MemberName        string  `gorm:"column:member_name" json:"member_name"`                 // 会员名称
	SellerID          int     `gorm:"column:seller_id" json:"seller_id"`                     // 卖家id
	SellerName        string  `gorm:"column:seller_name" json:"seller_name"`                 // 卖家姓名
	OrderSn           string  `gorm:"column:order_sn" json:"order_sn"`                       // 订单编号
	RefundStatus      string  `gorm:"column:refund_status" json:"refund_status"`             // 退款单状态 APPLY：新创建，REFUNDING：退款中，REFUNDFAIL：退款失败，COMPLETE：完成
	CreateTime        int64   `gorm:"column:create_time" json:"create_time"`                 // 创建时间
	RefundPrice       float64 `gorm:"column:refund_price" json:"refund_price"`               // 申请退款金额
	RefundWay         string  `gorm:"column:refund_way" json:"refund_way"`                   // 退款方式 ORIGINAL：原路退回，OFFLINE：线下退款，ACCOUNT：账户退款
	AccountType       string  `gorm:"column:account_type" json:"account_type"`               // 退款账户类型
	ReturnAccount     string  `gorm:"column:return_account" json:"return_account"`           // 退款账户
	BankName          string  `gorm:"column:bank_name" json:"bank_name"`                     // 银行名称
	BankAccountNumber string  `gorm:"column:bank_account_number" json:"bank_account_number"` // 银行账号
	BankAccountName   string  `gorm:"column:bank_account_name" json:"bank_account_name"`     // 银行开户名
	BankDepositName   string  `gorm:"column:bank_deposit_name" json:"bank_deposit_name"`     // 银行开户行
	PayOrderNo        string  `gorm:"column:pay_order_no" json:"pay_order_no"`               // 支付结果交易号
	PaymentType       string  `gorm:"column:payment_type" json:"payment_type"`               // 订单付款类型 ONLINE：在线支付，COD：货到付款
	RefundFailReason  string  `gorm:"column:refund_fail_reason" json:"refund_fail_reason"`   // 退款失败原因
	RefundTime        int64   `gorm:"column:refund_time" json:"refund_time"`                 // 退款时间
	Mobile            string  `gorm:"column:mobile" json:"mobile"`                           // 手机号
	AgreePrice        float64 `gorm:"column:agree_price" json:"agree_price"`                 // 商家同意的退款金额
	ActualPrice       float64 `gorm:"column:actual_price" json:"actual_price"`               // 实际退款金额
	GoodsJSON         string  `gorm:"column:goods_json" json:"goods_json"`                   // 售后商品信息json
	Disabled          string  `gorm:"column:disabled" json:"disabled"`                       // 删除状态 DELETED：已删除 NORMAL：正常
	CreateChannel     string  `gorm:"column:create_channel" json:"create_channel"`           // 售后服务单创建渠道 NORMAL：正常渠道创建，PINTUAN：拼团失败自动创建
}

// TableName get sql table name.获取数据库表名
func (m *EsRefund) TableName() string {
	return "es_refund"
}
