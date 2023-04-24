package models

// EsShipTemplate 运费模版(es_ship_template)
type EsShipTemplate struct {
	ID       int    `gorm:"primaryKey;column:id" json:"-"`     // 主键
	SellerID int    `gorm:"column:seller_id" json:"seller_id"` // 卖家id
	Name     string `gorm:"column:name" json:"name"`           // 模板名称
	Type     int16  `gorm:"column:type" json:"type"`           // 1 重量算运费 2 计件算运费
}

// TableName get sql table name.获取数据库表名
func (m *EsShipTemplate) TableName() string {
	return "es_ship_template"
}