package models

// EsReceiptFile 会员电子发票附件表(es_receipt_file)
type EsReceiptFile struct {
	ID        int    `gorm:"primaryKey;column:id" json:"-"`
	HistoryID int    `gorm:"column:history_id" json:"history_id"`
	ElecFile  string `gorm:"column:elec_file" json:"elec_file"`
}

// TableName get sql table name.获取数据库表名
func (m *EsReceiptFile) TableName() string {
	return "es_receipt_file"
}