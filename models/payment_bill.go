package models

// EsPaymentBill 支付帐单(es_payment_bill)
type EsPaymentBill struct {
	BillID          int     `gorm:"primaryKey;column:bill_id" json:"-"`                // 主键id
	Sn              string  `gorm:"column:sn" json:"sn"`                               // 单号
	OutTradeNo      string  `gorm:"column:out_trade_no" json:"out_trade_no"`           // 提交给第三方平台单号
	ReturnTradeNo   string  `gorm:"column:return_trade_no" json:"return_trade_no"`     // 第三方平台返回交易号
	IsPay           int     `gorm:"column:is_pay" json:"is_pay"`                       // 是否已支付
	TradeType       string  `gorm:"column:trade_type" json:"trade_type"`               // 交易类型
	PaymentName     string  `gorm:"column:payment_name" json:"payment_name"`           // 支付方式名称
	PayConfig       string  `gorm:"column:pay_config" json:"pay_config"`               // 支付参数
	TradePrice      float64 `gorm:"column:trade_price" json:"trade_price"`             // 交易金额
	PaymentPluginID string  `gorm:"column:payment_plugin_id" json:"payment_plugin_id"` // 支付插件
}

// TableName get sql table name.获取数据库表名
func (m *EsPaymentBill) TableName() string {
	return "es_payment_bill"
}
