package models

// EsArticle 文章(es_article)
type EsArticle struct {
	ArticleID    int    `gorm:"primaryKey;column:article_id" json:"-"`     // 主键
	ArticleName  string `gorm:"column:article_name" json:"article_name"`   // 文章名称
	CategoryID   int    `gorm:"column:category_id" json:"category_id"`     // 分类id
	Sort         int    `gorm:"column:sort" json:"sort"`                   // 文章排序
	OutsideURL   string `gorm:"column:outside_url" json:"outside_url"`     // 外链url
	Content      string `gorm:"column:content" json:"content"`             // 文章内容
	ShowPosition string `gorm:"column:show_position" json:"show_position"` // 显示位置
	CreateTime   int64  `gorm:"column:create_time" json:"create_time"`     // 添加时间
	ModifyTime   int64  `gorm:"column:modify_time" json:"modify_time"`     // 修改时间
}

// TableName get sql table name.获取数据库表名
func (m *EsArticle) TableName() string {
	return "es_article"
}
