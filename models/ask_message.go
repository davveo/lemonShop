package models

// EsAskMessage [...]
type EsAskMessage struct {
	ID             int    `gorm:"primaryKey;column:id" json:"-"`                 // 主键ID
	MemberID       int    `gorm:"column:member_id" json:"member_id"`             // 会员ID
	GoodsID        int    `gorm:"column:goods_id" json:"goods_id"`               // 商品ID
	GoodsName      string `gorm:"column:goods_name" json:"goods_name"`           // 商品名称
	GoodsImg       string `gorm:"column:goods_img" json:"goods_img"`             // 商品图片
	AskID          int    `gorm:"column:ask_id" json:"ask_id"`                   // 会员咨询id
	Ask            string `gorm:"column:ask" json:"ask"`                         // 咨询内容
	AskMember      string `gorm:"column:ask_member" json:"ask_member"`           // 咨询会员
	ReplyID        int    `gorm:"column:reply_id" json:"reply_id"`               // 会员回复咨询id
	Reply          string `gorm:"column:reply" json:"reply"`                     // 回复内容
	ReplyMember    string `gorm:"column:reply_member" json:"reply_member"`       // 回复会员
	SendTime       int64  `gorm:"column:send_time" json:"send_time"`             // 消息发送时间
	IsDel          string `gorm:"column:is_del" json:"is_del"`                   // 删除状态 DELETED：已删除 NORMAL：正常
	IsRead         string `gorm:"column:is_read" json:"is_read"`                 // 是否已读 YES：是 NO：否
	ReceiveTime    int64  `gorm:"column:receive_time" json:"receive_time"`       // 消息接收时间
	MsgType        string `gorm:"column:msg_type" json:"msg_type"`               // 咨询消息类型 ASK：提问消息 REPLY：回复消息
	AskAnonymous   string `gorm:"column:ask_anonymous" json:"ask_anonymous"`     // 咨询人是否匿名 YES:是，NO:否
	ReplyAnonymous string `gorm:"column:reply_anonymous" json:"reply_anonymous"` // 回复咨询人是否匿名 YES:是，NO:否
}

// TableName get sql table name.获取数据库表名
func (m *EsAskMessage) TableName() string {
	return "es_ask_message"
}
