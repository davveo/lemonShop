package models

// EsOrderComplain [...]
type EsOrderComplain struct {
	ComplainID        int     `gorm:"primaryKey;column:complain_id" json:"-"`              // 主键
	ComplainTopic     string  `gorm:"column:complain_topic" json:"complain_topic"`         // 投诉主题
	Content           string  `gorm:"column:content" json:"content"`                       // 投诉内容
	CreateTime        int64   `gorm:"column:create_time" json:"create_time"`               // 投诉时间
	Images            string  `gorm:"column:images" json:"images"`                         // 投诉凭证图片
	Status            string  `gorm:"column:status" json:"status"`                         // 状态，见ComplainStatusEnum.java
	AppealContent     string  `gorm:"column:appeal_content" json:"appeal_content"`         // 商家申诉内容
	AppealTime        int64   `gorm:"column:appeal_time" json:"appeal_time"`               // 商家申诉时间
	AppealImages      string  `gorm:"column:appeal_images" json:"appeal_images"`           // 商家申诉上传的图片
	OrderSn           string  `gorm:"column:order_sn" json:"order_sn"`                     // 订单号
	OrderTime         int64   `gorm:"column:order_time" json:"order_time"`                 // 下单时间
	GoodsName         string  `gorm:"column:goods_name" json:"goods_name"`                 // 商品名称
	GoodsID           int     `gorm:"column:goods_id" json:"goods_id"`                     // 商品id
	GoodsPrice        float64 `gorm:"column:goods_price" json:"goods_price"`               // 商品价格
	Num               int     `gorm:"column:num" json:"num"`                               // 购买的商品数量
	ShippingPrice     float64 `gorm:"column:shipping_price" json:"shipping_price"`         // 运费
	OrderPrice        float64 `gorm:"column:order_price" json:"order_price"`               // 订单金额
	ShipNo            string  `gorm:"column:ship_no" json:"ship_no"`                       // 物流单号
	SellerID          int     `gorm:"column:seller_id" json:"seller_id"`                   // 商家id
	SellerName        string  `gorm:"column:seller_name" json:"seller_name"`               // 商家名称
	MemberID          int     `gorm:"column:member_id" json:"member_id"`                   // 会员id
	MemberName        string  `gorm:"column:member_name" json:"member_name"`               // 会员名称
	ShipName          string  `gorm:"column:ship_name" json:"ship_name"`                   // 收货人
	ShipAddr          string  `gorm:"column:ship_addr" json:"ship_addr"`                   // 收货地址
	ShipMobile        string  `gorm:"column:ship_mobile" json:"ship_mobile"`               // 收货人手机
	ArbitrationResult string  `gorm:"column:arbitration_result" json:"arbitration_result"` // 仲裁结果
	SkuID             int     `gorm:"column:sku_id" json:"sku_id"`                         // sku 主键
	GoodsImage        string  `gorm:"column:goods_image" json:"goods_image"`               // 商品图片
}

// TableName get sql table name.获取数据库表名
func (m *EsOrderComplain) TableName() string {
	return "es_order_complain"
}
