package models

// EsPage 楼层(es_page)
type EsPage struct {
	PageID     int    `gorm:"primaryKey;column:page_id" json:"-"`    // 主键id
	PageName   string `gorm:"column:page_name" json:"page_name"`     // 楼层名称
	PageData   string `gorm:"column:page_data" json:"page_data"`     // 楼层数据
	PageType   string `gorm:"column:page_type" json:"page_type"`     // 页面类型
	ClientType string `gorm:"column:client_type" json:"client_type"` // 客户端类型
}

// TableName get sql table name.获取数据库表名
func (m *EsPage) TableName() string {
	return "es_page"
}
