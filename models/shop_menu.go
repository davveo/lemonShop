package models

// EsShopMenu 菜单管理店铺(es_shop_menu)
type EsShopMenu struct {
	ID          int    `gorm:"primaryKey;column:id" json:"-"`           // 店铺菜单id
	ParentID    int    `gorm:"column:parent_id" json:"parent_id"`       // 菜单父id
	Title       string `gorm:"column:title" json:"title"`               // 菜单标题
	IDentifier  string `gorm:"column:identifier" json:"identifier"`     // 菜单唯一标识
	AuthRegular string `gorm:"column:auth_regular" json:"auth_regular"` // 权限表达式
	DeleteFlag  int    `gorm:"column:delete_flag" json:"delete_flag"`   // 删除标记
	Path        string `gorm:"column:path" json:"path"`                 // 菜单级别标识
	Grade       int    `gorm:"column:grade" json:"grade"`               // 菜单级别
}

// TableName get sql table name.获取数据库表名
func (m *EsShopMenu) TableName() string {
	return "es_shop_menu"
}