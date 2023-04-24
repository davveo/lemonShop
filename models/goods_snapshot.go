package models

// EsGoodsSnapshot 交易快照(es_goods_snapshot)
type EsGoodsSnapshot struct {
	SnapshotID    int     `gorm:"primaryKey;column:snapshot_id" json:"-"`      // 主键
	GoodsID       int     `gorm:"column:goods_id" json:"goods_id"`             // 商品id
	Name          string  `gorm:"column:name" json:"name"`                     // 商品名称
	Sn            string  `gorm:"column:sn" json:"sn"`                         // 商品编号
	BrandName     string  `gorm:"column:brand_name" json:"brand_name"`         // 品牌名称
	CategoryName  string  `gorm:"column:category_name" json:"category_name"`   // 分类名称
	GoodsType     string  `gorm:"column:goods_type" json:"goods_type"`         // 商品类型
	Weight        float64 `gorm:"column:weight" json:"weight"`                 // 重量
	Intro         string  `gorm:"column:intro" json:"intro"`                   // 商品详情
	Price         float64 `gorm:"column:price" json:"price"`                   // 商品价格
	Cost          float64 `gorm:"column:cost" json:"cost"`                     // 商品成本价
	Mktprice      float64 `gorm:"column:mktprice" json:"mktprice"`             // 商品市场价
	HaveSpec      int16   `gorm:"column:have_spec" json:"have_spec"`           // 商品是否开启规格1 开启 0 未开启
	ParamsJSON    string  `gorm:"column:params_json" json:"params_json"`       // 参数json
	ImgJSON       string  `gorm:"column:img_json" json:"img_json"`             // 图片json
	CreateTime    int64   `gorm:"column:create_time" json:"create_time"`       // 快照时间
	Point         int     `gorm:"column:point" json:"point"`                   // 商品使用积分
	SellerID      int     `gorm:"column:seller_id" json:"seller_id"`           // 所属卖家
	PromotionJSON string  `gorm:"column:promotion_json" json:"promotion_json"` // 促销json值
	CouponJSON    string  `gorm:"column:coupon_json" json:"coupon_json"`       // 优惠券json值
	MemberID      int     `gorm:"column:member_id" json:"member_id"`           // 会员id
}

// TableName get sql table name.获取数据库表名
func (m *EsGoodsSnapshot) TableName() string {
	return "es_goods_snapshot"
}
