package models

// EsShopNoticeLog 店铺站内消息(es_shop_notice_log)
type EsShopNoticeLog struct {
	ID            int    `gorm:"primaryKey;column:id" json:"-"`               // id
	ShopID        int    `gorm:"column:shop_id" json:"shop_id"`               // 店铺ID
	NoticeContent string `gorm:"column:notice_content" json:"notice_content"` // 站内信内容
	SendTime      int64  `gorm:"column:send_time" json:"send_time"`           // 发送时间
	IsDelete      int    `gorm:"column:is_delete" json:"is_delete"`           // 是否删除 ：1 删除   0  未删除
	IsRead        int    `gorm:"column:is_read" json:"is_read"`               // 是否已读 ：1已读   0 未读
	Type          string `gorm:"column:type" json:"type"`                     // 消息类型
}

// TableName get sql table name.获取数据库表名
func (m *EsShopNoticeLog) TableName() string {
	return "es_shop_notice_log"
}