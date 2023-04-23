package models

// EsBill 结算单(es_bill)
type EsBill struct {
	BillID                   int     `gorm:"primaryKey;column:bill_id" json:"-"`                                  // 主键ID
	BillSn                   string  `gorm:"column:bill_sn" json:"bill_sn"`                                       // 结算单编号
	StartTime                int64   `gorm:"column:start_time" json:"start_time"`                                 // 结算开始时间
	EndTime                  int64   `gorm:"column:end_time" json:"end_time"`                                     // 结算结束时间
	Price                    float64 `gorm:"column:price" json:"price"`                                           // 结算总金额
	CommiPrice               float64 `gorm:"column:commi_price" json:"commi_price"`                               // 佣金
	DiscountPrice            float64 `gorm:"column:discount_price" json:"discount_price"`                         // 优惠金额
	Status                   string  `gorm:"column:status" json:"status"`                                         // 状态
	BillType                 int     `gorm:"column:bill_type" json:"bill_type"`                                   // 账单类型
	SellerID                 int     `gorm:"column:seller_id" json:"seller_id"`                                   // 店铺id
	PayTime                  int64   `gorm:"column:pay_time" json:"pay_time"`                                     // 付款时间
	CreateTime               int64   `gorm:"column:create_time" json:"create_time"`                               // 出账日期
	BillPrice                float64 `gorm:"column:bill_price" json:"bill_price"`                                 // 结算金额
	RefundPrice              float64 `gorm:"column:refund_price" json:"refund_price"`                             // 在线支付退款金额
	RefundCommiPrice         float64 `gorm:"column:refund_commi_price" json:"refund_commi_price"`                 // 退还佣金金额
	Sn                       string  `gorm:"column:sn" json:"sn"`                                                 // 账单号
	ShopName                 string  `gorm:"column:shop_name" json:"shop_name"`                                   // 店铺名称
	BankAccountName          string  `gorm:"column:bank_account_name" json:"bank_account_name"`                   // 银行开户名
	BankAccountNumber        string  `gorm:"column:bank_account_number" json:"bank_account_number"`               // 公司银行账号
	BankName                 string  `gorm:"column:bank_name" json:"bank_name"`                                   // 开户银行支行名称
	BankCode                 string  `gorm:"column:bank_code" json:"bank_code"`                                   // 支行联行号
	BankAddress              string  `gorm:"column:bank_address" json:"bank_address"`                             // 开户银行地址
	CodPrice                 float64 `gorm:"column:cod_price" json:"cod_price"`                                   // 货到付款金额
	CodRefundPrice           float64 `gorm:"column:cod_refund_price" json:"cod_refund_price"`                     // 货到付款退款金额
	DistributionRebate       float64 `gorm:"column:distribution_rebate" json:"distribution_rebate"`               // 分销返现支出
	DistributionReturnRebate float64 `gorm:"column:distribution_return_rebate" json:"distribution_return_rebate"` // 分销返现支出返还
	SiteCouponCommi          float64 `gorm:"column:site_coupon_commi" json:"site_coupon_commi"`                   // 平台优惠券佣金
	OrderTotalPrice          float64 `gorm:"column:order_total_price" json:"order_total_price"`                   // 结算周期内订单付款总金额
	RefundTotalPrice         float64 `gorm:"column:refund_total_price" json:"refund_total_price"`                 // 结算周期内订单退款总金额
}

// TableName get sql table name.获取数据库表名
func (m *EsBill) TableName() string {
	return "es_bill"
}
