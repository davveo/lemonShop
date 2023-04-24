package models

// EsSiteNavigation 导航栏(es_site_navigation)
type EsSiteNavigation struct {
	NavigationID   int    `gorm:"primaryKey;column:navigation_id" json:"-"`      // 主键
	NavigationName string `gorm:"column:navigation_name" json:"navigation_name"` // 导航名称
	URL            string `gorm:"column:url" json:"url"`                         // 导航地址
	ClientType     string `gorm:"column:client_type" json:"client_type"`         // 客户端类型
	Image          string `gorm:"column:image" json:"image"`                     // 图片
	Sort           int    `gorm:"column:sort" json:"sort"`                       // 排序
}

// TableName get sql table name.获取数据库表名
func (m *EsSiteNavigation) TableName() string {
	return "es_site_navigation"
}