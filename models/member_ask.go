package models

// EsMemberAsk 咨询(es_member_ask)
type EsMemberAsk struct {
	AskID       int    `gorm:"primaryKey;column:ask_id" json:"-"`       // 主键
	GoodsID     int    `gorm:"column:goods_id" json:"goods_id"`         // 商品id
	MemberID    int    `gorm:"column:member_id" json:"member_id"`       // 会员id
	Content     string `gorm:"column:content" json:"content"`           // 咨询内容
	CreateTime  int64  `gorm:"column:create_time" json:"create_time"`   // 咨询时间
	SellerID    int    `gorm:"column:seller_id" json:"seller_id"`       // 卖家id
	Reply       string `gorm:"column:reply" json:"reply"`               // 卖家回复内容
	ReplyTime   int64  `gorm:"column:reply_time" json:"reply_time"`     // 卖家回复时间
	ReplyStatus string `gorm:"column:reply_status" json:"reply_status"` // 商家回复状态 YES:已回复,NO:未回复
	Status      string `gorm:"column:status" json:"status"`             // 删除状态 DELETED:已删除,NORMAL:正常
	MemberName  string `gorm:"column:member_name" json:"member_name"`   // 会员名称
	GoodsName   string `gorm:"column:goods_name" json:"goods_name"`     // 商品名称
	MemberFace  string `gorm:"column:member_face" json:"member_face"`   // 会员头像
	AuthStatus  string `gorm:"column:auth_status" json:"auth_status"`   // 审核状态 WAIT_AUDIT:待审核,PASS_AUDIT:审核通过,REFUSE_AUDIT:审核未通过
	GoodsImg    string `gorm:"column:goods_img" json:"goods_img"`       // 商品图片
	Anonymous   string `gorm:"column:anonymous" json:"anonymous"`       // 是否匿名 YES:是，NO:否
	ReplyNum    int    `gorm:"column:reply_num" json:"reply_num"`       // 会员问题咨询回复数量
}

// TableName get sql table name.获取数据库表名
func (m *EsMemberAsk) TableName() string {
	return "es_member_ask"
}
