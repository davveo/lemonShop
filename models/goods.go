package models

// EsGoods 商品(es_goods)
type EsGoods struct {
	GoodsID             int     `gorm:"primaryKey;column:goods_id" json:"-"`                       // 主键
	GoodsName           string  `gorm:"column:goods_name" json:"goods_name"`                       // 商品名称
	Sn                  string  `gorm:"column:sn" json:"sn"`                                       // 商品编号
	BrandID             int     `gorm:"column:brand_id" json:"brand_id"`                           // 品牌id
	CategoryID          int     `gorm:"column:category_id" json:"category_id"`                     // 分类id
	GoodsType           string  `gorm:"column:goods_type" json:"goods_type"`                       // 商品类型normal普通point积分
	Weight              float64 `gorm:"column:weight" json:"weight"`                               // 重量
	MarketEnable        int     `gorm:"column:market_enable" json:"market_enable"`                 // 上架状态 1上架  0下架
	Intro               string  `gorm:"column:intro" json:"intro"`                                 // 详情
	Price               float64 `gorm:"column:price" json:"price"`                                 // 商品价格
	Cost                float64 `gorm:"column:cost" json:"cost"`                                   // 成本价格
	Mktprice            float64 `gorm:"column:mktprice" json:"mktprice"`                           // 市场价格
	HaveSpec            int     `gorm:"column:have_spec" json:"have_spec"`                         // 是否有规格0没有 1有
	CreateTime          int64   `gorm:"column:create_time" json:"create_time"`                     // 创建时间
	LastModify          int64   `gorm:"column:last_modify" json:"last_modify"`                     // 最后修改时间
	ViewCount           int     `gorm:"column:view_count" json:"view_count"`                       // 浏览数量
	BuyCount            int     `gorm:"column:buy_count" json:"buy_count"`                         // 购买数量
	Disabled            int     `gorm:"column:disabled" json:"disabled"`                           // 是否被删除0 删除 1未删除
	Quantity            int     `gorm:"column:quantity" json:"quantity"`                           // 库存
	EnableQuantity      int     `gorm:"column:enable_quantity" json:"enable_quantity"`             // 可用库存
	Point               int     `gorm:"column:point" json:"point"`                                 // 如果是积分商品需要使用的积分
	PageTitle           string  `gorm:"column:page_title" json:"page_title"`                       // seo标题
	MetaKeywords        string  `gorm:"column:meta_keywords" json:"meta_keywords"`                 // seo关键字
	MetaDescription     string  `gorm:"column:meta_description" json:"meta_description"`           // seo描述
	Grade               float64 `gorm:"column:grade" json:"grade"`                                 // 商品好评率
	Thumbnail           string  `gorm:"column:thumbnail" json:"thumbnail"`                         // 缩略图路径
	Big                 string  `gorm:"column:big" json:"big"`                                     // 大图路径
	Small               string  `gorm:"column:small" json:"small"`                                 // 小图路径
	Original            string  `gorm:"column:original" json:"original"`                           // 原图路径
	SellerID            int     `gorm:"column:seller_id" json:"seller_id"`                         // 卖家id
	ShopCatID           int     `gorm:"column:shop_cat_id" json:"shop_cat_id"`                     // 店铺分类id
	CommentNum          int     `gorm:"column:comment_num" json:"comment_num"`                     // 评论数量
	TemplateID          int     `gorm:"column:template_id" json:"template_id"`                     // 运费模板id
	GoodsTransfeeCharge int     `gorm:"column:goods_transfee_charge" json:"goods_transfee_charge"` // 谁承担运费0：买家承担，1：卖家承担
	SellerName          string  `gorm:"column:seller_name" json:"seller_name"`                     // 卖家名字
	IsAuth              int     `gorm:"column:is_auth" json:"is_auth"`                             // 0 需要审核 并且待审核，1 不需要审核 2需要审核 且审核通过 3 需要审核 且审核未通过
	AuthMessage         string  `gorm:"column:auth_message" json:"auth_message"`                   // 审核信息
	SelfOperated        int     `gorm:"column:self_operated" json:"self_operated"`                 // 是否是自营商品 0 不是 1是
	UnderMessage        string  `gorm:"column:under_message" json:"under_message"`                 // 下架原因
	Priority            int     `gorm:"column:priority" json:"priority"`                           // 优先级:高(3)、中(2)、低(1)
}

// TableName get sql table name.获取数据库表名
func (m *EsGoods) TableName() string {
	return "es_goods"
}
