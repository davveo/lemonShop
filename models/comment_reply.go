package models

// EsCommentReply 评论回复(es_comment_reply)
type EsCommentReply struct {
	ReplyID    int    `gorm:"primaryKey;column:reply_id" json:"-"`   // 主键
	CommentID  int    `gorm:"column:comment_id" json:"comment_id"`   // 评论id
	ParentID   int    `gorm:"column:parent_id" json:"parent_id"`     // 回复父id
	Content    string `gorm:"column:content" json:"content"`         // 回复内容
	CreateTime int64  `gorm:"column:create_time" json:"create_time"` // 创建时间
	Role       string `gorm:"column:role" json:"role"`               // 商家或者买家
	Path       string `gorm:"column:path" json:"path"`               // 父子路径0|10|
	ReplyType  string `gorm:"column:reply_type" json:"reply_type"`   // INITIAL
}

// TableName get sql table name.获取数据库表名
func (m *EsCommentReply) TableName() string {
	return "es_comment_reply"
}
