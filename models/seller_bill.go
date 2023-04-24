package models

// EsSellerBill 店铺结算(es_seller_bill)
type EsSellerBill struct {
	ID                int     `gorm:"primaryKey;column:id" json:"-"`                       // id
	SellerID          int     `gorm:"column:seller_id" json:"seller_id"`                   // 商家id
	CreateTime        int64   `gorm:"column:create_time" json:"create_time"`               // 创建时间
	OrderSn           string  `gorm:"column:order_sn" json:"order_sn"`                     // 订单编号
	Expenditure       float64 `gorm:"column:expenditure" json:"expenditure"`               // 商品反现支出
	ReturnExpenditure float64 `gorm:"column:return_expenditure" json:"return_expenditure"` // 返还商品反现支出
}

// TableName get sql table name.获取数据库表名
func (m *EsSellerBill) TableName() string {
	return "es_seller_bill"
}