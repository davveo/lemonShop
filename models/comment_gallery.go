package models

// EsCommentGallery 评论图片(es_comment_gallery)
type EsCommentGallery struct {
	ImgID     int    `gorm:"primaryKey;column:img_id" json:"-"`   // 主键
	CommentID int    `gorm:"column:comment_id" json:"comment_id"` // 主键
	Original  string `gorm:"column:original" json:"original"`     // 图片路径
	Sort      int    `gorm:"column:sort" json:"sort"`             // 排序
	ImgBelong int    `gorm:"column:img_belong" json:"img_belong"` // 图片所属 0：初评，1：追评
}

// TableName get sql table name.获取数据库表名
func (m *EsCommentGallery) TableName() string {
	return "es_comment_gallery"
}
