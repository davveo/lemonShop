package models


// EsShopLogisticsSetting [...]
type EsShopLogisticsSetting struct {
	ID          int    `gorm:"primaryKey;column:id" json:"-"`
	ShopID      int    `gorm:"column:shop_id" json:"shop_id"`
	LogisticsID int    `gorm:"column:logistics_id" json:"logistics_id"`
	Config      string `gorm:"column:config" json:"config"`
}

// TableName get sql table name.获取数据库表名
func (m *EsShopLogisticsSetting) TableName() string {
	return "es_shop_logistics_setting"
}