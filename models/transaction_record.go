package models

// EsTransactionRecord 交易记录表(es_transaction_record)
type EsTransactionRecord struct {
	RecordID int     `gorm:"primaryKey;column:record_id" json:"-"` // 主键ID
	OrderSn  string  `gorm:"column:order_sn" json:"order_sn"`      // 订单编号
	GoodsID  int     `gorm:"column:goods_id" json:"goods_id"`      // 商品ID
	GoodsNum int     `gorm:"column:goods_num" json:"goods_num"`    // 商品数量
	RogTime  int64   `gorm:"column:rog_time" json:"rog_time"`      // 确认收货时间
	Uname    string  `gorm:"column:uname" json:"uname"`            // 用户名
	Price    float64 `gorm:"column:price" json:"price"`            // 交易价格
	MemberID int     `gorm:"column:member_id" json:"member_id"`    // 会员ID
}

// TableName get sql table name.获取数据库表名
func (m *EsTransactionRecord) TableName() string {
	return "es_transaction_record"
}