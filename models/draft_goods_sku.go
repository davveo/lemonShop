package models

// EsDraftGoodsSku 草稿商品sku(es_draft_goods_sku)
type EsDraftGoodsSku struct {
	DraftSkuID   int     `gorm:"primaryKey;column:draft_sku_id" json:"-"`     // 主键ID
	DraftGoodsID int     `gorm:"column:draft_goods_id" json:"draft_goods_id"` // 草稿id
	Sn           string  `gorm:"column:sn" json:"sn"`                         // 货号
	Quantity     int     `gorm:"column:quantity" json:"quantity"`             // 总库存
	Price        float64 `gorm:"column:price" json:"price"`                   // 价格
	Specs        string  `gorm:"column:specs" json:"specs"`                   // 规格
	Cost         float64 `gorm:"column:cost" json:"cost"`                     // 成本
	Weight       float64 `gorm:"column:weight" json:"weight"`                 // 重量
}

// TableName get sql table name.获取数据库表名
func (m *EsDraftGoodsSku) TableName() string {
	return "es_draft_goods_sku"
}
