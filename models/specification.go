package models

// EsSpecification 规格项(es_specification)
type EsSpecification struct {
	SpecID   int    `gorm:"primaryKey;column:spec_id" json:"spec_id"` // 主键
	SpecName string `gorm:"column:spec_name" json:"spec_name"`        // 规格项名称
	Disabled int    `gorm:"column:disabled" json:"disabled"`          // 是否被删除0 删除   1  没有删除
	SpecMemo string `gorm:"column:spec_memo" json:"spec_memo"`        // 规格描述
	SellerID int    `gorm:"column:seller_id" json:"seller_id"`        // 所属卖家
}

// TableName get sql table name.获取数据库表名
func (m *EsSpecification) TableName() string {
	return "es_specification"
}
