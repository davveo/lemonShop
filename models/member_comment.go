package models

// EsMemberComment 评论(es_member_comment)
type EsMemberComment struct {
	CommentID           int    `gorm:"primaryKey;column:comment_id" json:"-"`                     // 评论主键
	GoodsID             int    `gorm:"column:goods_id" json:"goods_id"`                           // 商品id
	SkuID               int    `gorm:"column:sku_id" json:"sku_id"`                               // skuid
	MemberID            int    `gorm:"column:member_id" json:"member_id"`                         // 会员id
	SellerID            int    `gorm:"column:seller_id" json:"seller_id"`                         // 卖家id
	MemberName          string `gorm:"column:member_name" json:"member_name"`                     // 会员名称
	MemberFace          string `gorm:"column:member_face" json:"member_face"`                     // 会员头像
	GoodsName           string `gorm:"column:goods_name" json:"goods_name"`                       // 商品名称
	GoodsImg            string `gorm:"column:goods_img" json:"goods_img"`                         // 商品图片
	Content             string `gorm:"column:content" json:"content"`                             // 评论内容
	CreateTime          int64  `gorm:"column:create_time" json:"create_time"`                     // 评论时间
	HaveImage           int16  `gorm:"column:have_image" json:"have_image"`                       // 是否有图片 1 有 0 没有
	Status              int16  `gorm:"column:status" json:"status"`                               // 状态  1 正常 0 删除
	Grade               string `gorm:"column:grade" json:"grade"`                                 // 好中差评
	OrderSn             string `gorm:"column:order_sn" json:"order_sn"`                           // 订单编号
	ReplyStatus         int16  `gorm:"column:reply_status" json:"reply_status"`                   // 是否回复 1 是 0 否
	AuditStatus         string `gorm:"column:audit_status" json:"audit_status"`                   // 初评审核:待审核(WAIT_AUDIT),审核通过(PASS_AUDIT),审核拒绝(REFUSE_AUDIT)
	CommentsType        string `gorm:"column:comments_type" json:"comments_type"`                 // 评论类型：初评(INITIAL),追评(ADDITIONAL)
	ParentID            int    `gorm:"column:parent_id" json:"parent_id"`                         // 初评id，默认为0
	AdditionalStatus    int    `gorm:"column:additional_status" json:"additional_status"`         // 追加评论状态 0：未追加，1：已追加
	AdditionalContent   string `gorm:"column:additional_content" json:"additional_content"`       // 追加的评论内容
	AdditionalTime      int64  `gorm:"column:additional_time" json:"additional_time"`             // 追加评论时间
	AdditionalHaveImage int    `gorm:"column:additional_have_image" json:"additional_have_image"` // 追加的评论是否上传了图片 0：未上传，1：已上传
}

// TableName get sql table name.获取数据库表名
func (m *EsMemberComment) TableName() string {
	return "es_member_comment"
}
