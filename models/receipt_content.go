package models

// EsReceiptContent 发票内容(es_receipt_content)
type EsReceiptContent struct {
	ID      int    `gorm:"primaryKey;column:id" json:"-"` // 发票内容id
	Content string `gorm:"column:content" json:"content"` // 发票内容
}

// TableName get sql table name.获取数据库表名
func (m *EsReceiptContent) TableName() string {
	return "es_receipt_content"
}