package models

// EsHotKeyword 热门关键字(es_hot_keyword)
type EsHotKeyword struct {
	ID      int    `gorm:"primaryKey;column:id" json:"-"`   // 主键ID
	HotName string `gorm:"column:hot_name" json:"hot_name"` // 文字内容
	Sort    int    `gorm:"column:sort" json:"sort"`         // 排序
}

// TableName get sql table name.获取数据库表名
func (m *EsHotKeyword) TableName() string {
	return "es_hot_keyword"
}
