package models

// EsArticleCategory 文章分类(es_article_category)
type EsArticleCategory struct {
	ID          int    `gorm:"primaryKey;column:id" json:"-"`           // 主键
	Name        string `gorm:"column:name" json:"name"`                 // 分类名称
	ParentID    int    `gorm:"column:parent_id" json:"parent_id"`       // 父分类id
	Path        string `gorm:"column:path" json:"path"`                 // 父子路径0|10|
	AllowDelete int16  `gorm:"column:allow_delete" json:"allow_delete"` // 是否允许删除1允许 0不允许
	Type        string `gorm:"column:type" json:"type"`                 // 分类类型，枚举值
	Sort        int    `gorm:"column:sort" json:"sort"`                 // 排序，正序123
}

// TableName get sql table name.获取数据库表名
func (m *EsArticleCategory) TableName() string {
	return "es_article_category"
}
