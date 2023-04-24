package models

// EsNavigation 店铺导航管理(es_navigation)
type EsNavigation struct {
	ID       int    `gorm:"primaryKey;column:id" json:"-"`   // 导航id
	Name     string `gorm:"column:name" json:"name"`         // 名称
	Disable  int    `gorm:"column:disable" json:"disable"`   // 是否显示
	Sort     int    `gorm:"column:sort" json:"sort"`         // 排序
	NavURL   string `gorm:"column:nav_url" json:"nav_url"`   // URL
	Target   int    `gorm:"column:target" json:"target"`     // 新窗口打开
	ShopID   int    `gorm:"column:shop_id" json:"shop_id"`   // 店铺id
	Contents string `gorm:"column:contents" json:"contents"` // 内容
}

// TableName get sql table name.获取数据库表名
func (m *EsNavigation) TableName() string {
	return "es_navigation"
}
