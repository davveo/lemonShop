package models

// EsSssRefundData 退货数据(es_sss_refund_data)
type EsSssRefundData struct {
	ID          int     `gorm:"primaryKey;column:id" json:"-"`           // id
	SellerID    int     `gorm:"column:seller_id" json:"seller_id"`       // 商家id
	OrderSn     string  `gorm:"column:order_sn" json:"order_sn"`         // 订单sn
	RefundSn    string  `gorm:"column:refund_sn" json:"refund_sn"`       // 售后订单sn
	RefundPrice float64 `gorm:"column:refund_price" json:"refund_price"` // 退还金额
	CreateTime  int     `gorm:"column:create_time" json:"create_time"`   // 创建日期
}

// TableName get sql table name.获取数据库表名
func (m *EsSssRefundData) TableName() string {
	return "es_sss_refund_data"
}