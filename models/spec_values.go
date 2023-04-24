package models

// EsSpecValues 规格值(es_spec_values)
type EsSpecValues struct {
	SpecValueID int    `gorm:"primaryKey;column:spec_value_id" json:"-"` // 主键
	SpecID      int    `gorm:"column:spec_id" json:"spec_id"`            // 规格项id
	SpecValue   string `gorm:"column:spec_value" json:"spec_value"`      // 规格值名字
	SellerID    int    `gorm:"column:seller_id" json:"seller_id"`        // 所属卖家
	SpecName    string `gorm:"column:spec_name" json:"spec_name"`        // 规格名字
}

// TableName get sql table name.获取数据库表名
func (m *EsSpecValues) TableName() string {
	return "es_spec_values"
}